package mutate

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// nginxWorkload builds an nginx workload whose container args are the ones passed in.
func nginxWorkload(args ...string) *unstructured.Unstructured {
	// put the args and an int port on the nginx container
	rawArgs := make([]interface{}, len(args))
	for i, arg := range args {
		rawArgs[i] = arg
	}
	return &unstructured.Unstructured{Object: map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{
							"name": nginxContainerName,
							"args": rawArgs,
							"ports": []interface{}{
								map[string]interface{}{"containerPort": 80},
							},
						},
					},
				},
			},
		},
	}}
}

// TestSetGlobalConfigurationArg checks an empty name, a set name, a replaced name, and a dropped arg.
func TestSetGlobalConfigurationArg(t *testing.T) {
	// an empty name leaves the existing args
	empty := nginxWorkload("-nginx-configmaps=$(POD_NAMESPACE)/nginx-config")
	if err := setGlobalConfigurationArg(empty, ""); err != nil {
		t.Fatal(err)
	}
	if got := containerArgs(nginxContainer(t, empty)); len(got) != 1 || got[0] != "-nginx-configmaps=$(POD_NAMESPACE)/nginx-config" {
		t.Fatalf("empty name args: %v", got)
	}

	// a name sets the namespace and the name
	named := nginxWorkload("-nginx-configmaps=$(POD_NAMESPACE)/nginx-config")
	if err := setGlobalConfigurationArg(named, "nginx-ingress"); err != nil {
		t.Fatal(err)
	}
	got := containerArgs(nginxContainer(t, named))
	if len(got) != 2 || got[0] != "-global-configuration=$(POD_NAMESPACE)/nginx-ingress" {
		t.Fatalf("named args: %v", got)
	}

	// a second name keeps one arg
	if err := setGlobalConfigurationArg(named, "other"); err != nil {
		t.Fatal(err)
	}
	got = containerArgs(nginxContainer(t, named))
	if len(got) != 2 || got[0] != "-global-configuration=$(POD_NAMESPACE)/other" {
		t.Fatalf("second name args: %v", got)
	}

	// an empty name drops a global configuration arg
	dropped := nginxWorkload(
		"-nginx-configmaps=$(POD_NAMESPACE)/nginx-config",
		"-global-configuration=$(POD_NAMESPACE)/nginx-ingress",
	)
	if err := setGlobalConfigurationArg(dropped, ""); err != nil {
		t.Fatal(err)
	}
	if got := containerArgs(nginxContainer(t, dropped)); len(got) != 1 || got[0] != "-nginx-configmaps=$(POD_NAMESPACE)/nginx-config" {
		t.Fatalf("empty name drops global configuration arg: %v", got)
	}
}

// TestSetGlobalConfigurationArgMissingContainers rejects a workload that has no containers.
func TestSetGlobalConfigurationArgMissingContainers(t *testing.T) {
	// build a workload with no containers
	workload := &unstructured.Unstructured{Object: map[string]interface{}{}}

	// reject the missing containers
	err := setGlobalConfigurationArg(workload, "nginx-ingress")
	if err == nil || err.Error() != "failed to find containers in unstructured object" {
		t.Fatalf("missing containers: %v", err)
	}
}

// nginxContainer returns the nginx container from the workload.
func nginxContainer(t *testing.T, workload *unstructured.Unstructured) map[string]interface{} {
	t.Helper()
	// read the single container in place
	raw, found, err := unstructured.NestedFieldNoCopy(workload.Object, "spec", "template", "spec", "containers")
	containers, ok := raw.([]interface{})
	if err != nil || !found || !ok || len(containers) != 1 {
		t.Fatalf("containers: found=%v err=%v", found, err)
	}
	container, ok := containers[0].(map[string]interface{})
	if !ok {
		t.Fatalf("container type %T", containers[0])
	}
	return container
}
