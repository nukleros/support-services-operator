package mutate

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	nginxContainerName            = "nginx-ingress"
	globalConfigurationArgPrefix  = "-global-configuration="
	globalConfigurationArgPattern = "-global-configuration=$(POD_NAMESPACE)/%s"

	// globalConfigurationAnnotation records the name on the pod template.
	// An added container arg alone compares equal, so the update skips the deployment.
	globalConfigurationAnnotation = "addons.nukleros.io/nginx-global-configuration"
)

// setGlobalConfigurationArg sets the nginx container's global configuration arg.
// An empty name omits the arg.
func setGlobalConfigurationArg(obj client.Object, name string) error {
	workloadObject, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return fmt.Errorf("nginx workload must be *unstructured.Unstructured, got %T", obj)
	}

	// read the pod containers in place, since a deep copy panics on int ports
	rawContainers, found, err := unstructured.NestedFieldNoCopy(
		workloadObject.Object, "spec", "template", "spec", "containers",
	)
	if err != nil {
		return fmt.Errorf("failed to get containers from unstructured object: %w", err)
	}
	if !found {
		return fmt.Errorf("failed to find containers in unstructured object")
	}
	containers, ok := rawContainers.([]interface{})
	if !ok {
		return fmt.Errorf("nginx containers must be []interface{}, got %T", rawContainers)
	}

	// set the arg on the nginx container
	for _, item := range containers {
		container, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		containerName, _, _ := unstructured.NestedString(container, "name")
		if containerName != nginxContainerName {
			continue
		}
		args := containerArgs(container)
		kept := make([]interface{}, 0, len(args)+1)
		for _, arg := range args {
			if strings.HasPrefix(arg, globalConfigurationArgPrefix) {
				continue
			}
			kept = append(kept, arg)
		}
		// put the arg first so the equality check sees a changed element
		if name != "" {
			kept = append([]interface{}{fmt.Sprintf(globalConfigurationArgPattern, name)}, kept...)
		}
		container["args"] = kept
	}

	// record the name where the update comparison can see it
	annotations, _, _ := unstructured.NestedStringMap(
		workloadObject.Object, "spec", "template", "metadata", "annotations",
	)
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[globalConfigurationAnnotation] = name
	if err := unstructured.SetNestedStringMap(
		workloadObject.Object, annotations, "spec", "template", "metadata", "annotations",
	); err != nil {
		return fmt.Errorf("failed to set nginx global configuration annotation: %w", err)
	}
	return nil
}

// containerArgs returns the container arg strings.
func containerArgs(container map[string]interface{}) []string {
	// collect the string args
	raw, _, _ := unstructured.NestedSlice(container, "args")
	args := make([]string, 0, len(raw))
	for _, item := range raw {
		arg, ok := item.(string)
		if ok {
			args = append(args, arg)
		}
	}
	return args
}
