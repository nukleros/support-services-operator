package mutate

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// nginxWorkload builds an nginx workload whose container args are the ones passed in.
func nginxWorkload(args ...string) *unstructured.Unstructured {
	// put the args on the nginx container
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
						},
					},
				},
			},
		},
	}}
}

// TestSetGlobalConfigurationArg checks an empty name, a set name, and a second name.
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
	if len(got) != 2 || got[1] != "-global-configuration=$(POD_NAMESPACE)/nginx-ingress" {
		t.Fatalf("named args: %v", got)
	}

	// a second name keeps one arg
	if err := setGlobalConfigurationArg(named, "other"); err != nil {
		t.Fatal(err)
	}
	got = containerArgs(nginxContainer(t, named))
	if len(got) != 2 || got[1] != "-global-configuration=$(POD_NAMESPACE)/other" {
		t.Fatalf("second name args: %v", got)
	}
}

// nginxContainer returns the nginx container from the workload.
func nginxContainer(t *testing.T, workload *unstructured.Unstructured) map[string]interface{} {
	t.Helper()
	// read the single container
	containers, found, err := unstructured.NestedSlice(workload.Object, "spec", "template", "spec", "containers")
	if err != nil || !found || len(containers) != 1 {
		t.Fatalf("containers: found=%v err=%v", found, err)
	}
	container, ok := containers[0].(map[string]interface{})
	if !ok {
		t.Fatalf("container type %T", containers[0])
	}
	return container
}
