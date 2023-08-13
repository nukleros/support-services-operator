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

// MutateServiceNamespaceGatewayProxy mutates the Service resource with name gateway-proxy.
func MutateServiceNamespaceGatewayProxy(
	original client.Object,
	parent *gatewayv1alpha1.GlooEdge, collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	var mutatedService []client.Object

	// convert object to unstructured
	service, err := resources.ToUnstructured(original)
	if err != nil {
		return []client.Object{original}, fmt.Errorf("failed to convert client.Object to unstructured object: %w", err)
	}

	// target := unstructuredObj.DeepCopy()
	existingPorts, found, err := unstructured.NestedSlice(service.Object, "spec", "ports")
	if err != nil {
		return mutatedService, fmt.Errorf("failed to retrieve spec field for gateway: %w", err)
	}
	if !found {
		return mutatedService, fmt.Errorf("spec field not found in gateway object: %w", err)
	}

	newPorts := []interface{}{}
	// create a port object for each port requested
	for _, portSpec := range parent.Spec.Ports {
		port := map[string]interface{}{}

		existingPorts = removeDuplicatePorts(portSpec.Port, existingPorts)

		targetPort, err := getUnprivilegedPort(portSpec.Port, parent.Spec.Ports)
		if err != nil {
			return mutatedService, err
		}

		port["port"] = portSpec.SSL
		port["targetPort"] = targetPort
		port["protocol"] = "TCP"
		port["name"] = portSpec.Name

		newPorts = append(newPorts, port)
	}

	newPorts = append(newPorts, existingPorts)

	err = unstructured.SetNestedSlice(service.Object, newPorts, "spec", "ports")
	if err != nil {
		return mutatedService, fmt.Errorf("failed to set ports field for service: %w", err)
	}

	return mutatedService, nil
}

// removeDuplicatePorts removes any ports that have the same port number as the one passed in.
func removeDuplicatePorts(port int64, ports []interface{}) []interface{} {
	deduplicated := []interface{}{}
	for _, p := range ports {
		if pMap, ok := p.(map[string]interface{}); ok {
			if pMap["port"].(int64) == port {
				continue
			}
		}
		deduplicated = append(deduplicated, p)
	}
	return deduplicated
}
