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

// MutateGatewayNamespaceGatewayProxy mutates the Gateway resource with name gateway-proxy.
func MutateGatewayNamespaceGatewayProxy(
	original client.Object,
	parent *gatewayv1alpha1.GlooEdge, collection *orchestrationv1alpha1.SupportServices,
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

	var mutatedGateways []client.Object

	// create a gateway object for each port requested
	for _, portSpec := range parent.Spec.Ports {
		target := unstructuredObj.DeepCopy()

		spec, found, err := unstructured.NestedMap(target.Object, "spec")
		if err != nil {
			return mutatedGateways, fmt.Errorf("failed to retrieve spec field for gateway: %w", err)
		}
		if !found {
			return mutatedGateways, fmt.Errorf("spec field not found in gateway object: %w", err)
		}

		bindPort, err := getUnprivilegedPort(portSpec.Port, parent.Spec.Ports)
		if err != nil {
			return mutatedGateways, fmt.Errorf("failed to get unprivileged port: %w", err)
		}

		spec["bindPort"] = bindPort
		spec["ssl"] = portSpec.SSL

		if err := unstructured.SetNestedMap(target.Object, spec, "spec"); err != nil {
			return mutatedGateways, fmt.Errorf("failed to set spec on gateway: %w", err)
		}

		target.SetName(portSpec.Name)

		mutatedGateways = append(mutatedGateways, target)
	}

	return mutatedGateways, nil
}

// getUnprivilegedPort returns an unprivileged port if the port is privileged.
func getUnprivilegedPort(port int64, portSpec []gatewayv1alpha1.PortSpec) (int64, error) {
	if port < 1 {
		return port, fmt.Errorf("port %d is invalid, must be greater than 0", port)
	}
	if port < 1024 {
		unprivilegedPort := port + 8000
		if !isPortInUse(unprivilegedPort, portSpec) {
			return unprivilegedPort, nil
		}
		return port, fmt.Errorf("port %d is already in use", unprivilegedPort)
	}
	return port, nil
}

// isPortInUse returns true if the requested port is already in use.
func isPortInUse(port int64, portSpec []gatewayv1alpha1.PortSpec) bool {
	for _, portSpec := range portSpec {
		if portSpec.Port == int64(port) {
			return true
		}
	}
	return false
}
