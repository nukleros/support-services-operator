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
)

// setGlobalConfigurationArg sets the nginx container's global configuration arg.
// An empty name omits the arg.
func setGlobalConfigurationArg(obj client.Object, name string) error {
	workloadObject, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return fmt.Errorf("nginx workload is %T", obj)
	}

	// read the pod containers
	containers, found, err := unstructured.NestedSlice(
		workloadObject.Object, "spec", "template", "spec", "containers",
	)
	if err != nil || !found {
		return fmt.Errorf("nginx workload has no containers: %w", err)
	}

	// set the arg on the nginx container
	for i, item := range containers {
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
		if name != "" {
			kept = append(kept, fmt.Sprintf(globalConfigurationArgPattern, name))
		}
		container["args"] = kept
		containers[i] = container
	}
	if err := unstructured.SetNestedSlice(
		workloadObject.Object, containers, "spec", "template", "spec", "containers",
	); err != nil {
		return fmt.Errorf("failed to set nginx container args: %w", err)
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
