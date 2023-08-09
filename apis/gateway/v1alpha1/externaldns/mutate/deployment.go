package mutate

import (
	"fmt"

	"github.com/nukleros/operator-builder-tools/pkg/resources"
	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func appendExtraArgs(original client.Object, parent *gatewayv1alpha1.ExternalDNS) (client.Object, error) {

	// convert object to unstructured
	unstructuredObj, err := resources.ToUnstructured(original)
	if err != nil {
		return original, fmt.Errorf("failed to convert client.Object to unstructured object: %w", err)
	}

	// get the container from the unstructured object
	containers, found, err := unstructured.NestedSlice(
		unstructuredObj.Object,
		"spec",
		"template",
		"spec",
		"containers",
	)
	if err != nil {
		return original, fmt.Errorf("failed to get containers from unstructured object: %w", err)
	}
	if !found {
		return original, fmt.Errorf("failed to find containers in unstructured object")
	}

	// unwrap external dns container
	container := containers[0].(map[string]interface{})

	// unwrap args
	args, found, err := unstructured.NestedSlice(container, "args")
	if err != nil {
		return original, fmt.Errorf("failed to get args from unstructured object: %w", err)
	}
	if !found {
		return original, fmt.Errorf("failed to find args in unstructured object")
	}

	// append extra args
	for _, arg := range parent.Spec.ExtraArgs {
		args = append(args, arg)
	}

	// set args
	err = unstructured.SetNestedSlice(container, args, "args")
	if err != nil {
		return original, fmt.Errorf("failed to set args in unstructured object: %w", err)
	}

	// set container
	containers[0] = container

	// set containers
	err = unstructured.SetNestedSlice(
		unstructuredObj.Object,
		containers,
		"spec",
		"template",
		"spec",
		"containers",
	)
	if err != nil {
		return original, fmt.Errorf("failed to set containers in unstructured object: %w", err)
	}

	return unstructuredObj, nil
}
