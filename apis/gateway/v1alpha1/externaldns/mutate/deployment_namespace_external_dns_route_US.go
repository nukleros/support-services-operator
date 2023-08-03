/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mutate

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/resources"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// MutateDeploymentNamespaceExternalDnsRoute53 mutates the Deployment resource with name external-dns-route53.
func MutateDeploymentNamespaceExternalDnsRoute53(
	original client.Object,
	parent *gatewayv1alpha1.ExternalDNS, collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	// convert object to unstructured
	unstructuredObj, err := resources.ToUnstructured(original)
	if err != nil {
		return []client.Object{original}, fmt.Errorf("failed to convert client.Object to unstructured object: %w", err)
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
		return []client.Object{original}, fmt.Errorf("failed to get containers from unstructured object: %w", err)
	}
	if !found {
		return []client.Object{original}, fmt.Errorf("failed to find containers in unstructured object")
	}

	// unwrap external dns container
	container := containers[0].(map[string]interface{})

	// unwrap args
	args, found, err := unstructured.NestedSlice(container, "args")
	if err != nil {
		return []client.Object{original}, fmt.Errorf("failed to get args from unstructured object: %w", err)
	}
	if !found {
		return []client.Object{original}, fmt.Errorf("failed to find args in unstructured object")
	}

	// append extra args
	for _, arg := range parent.Spec.ExtraArgs {
		args = append(args, arg)
	}

	// set args
	err = unstructured.SetNestedSlice(container, args, "args")
	if err != nil {
		return []client.Object{original}, fmt.Errorf("failed to set args in unstructured object: %w", err)
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
		return []client.Object{original}, fmt.Errorf("failed to set containers in unstructured object: %w", err)
	}

	return []client.Object{original}, nil
}
