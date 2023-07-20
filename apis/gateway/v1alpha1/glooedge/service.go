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

package glooedge

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/gateway/v1alpha1/glooedge/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNuklerosGatewaySystemGloo creates the Service resource with name gloo.
func CreateServiceNuklerosGatewaySystemGloo(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "gloo",
				},
				"name":      "gloo",
				"namespace": "nukleros-gateway-system",
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":     "grpc-xds",
						"port":     9977,
						"protocol": "TCP",
					},
					map[string]interface{}{
						"name":     "rest-xds",
						"port":     9976,
						"protocol": "TCP",
					},
					map[string]interface{}{
						"name":     "grpc-validation",
						"port":     9988,
						"protocol": "TCP",
					},
					map[string]interface{}{
						"name":     "grpc-proxydebug",
						"port":     9966,
						"protocol": "TCP",
					},
					map[string]interface{}{
						"name":     "wasm-cache",
						"port":     9979,
						"protocol": "TCP",
					},
					map[string]interface{}{
						"name":     "https",
						"port":     443,
						"protocol": "TCP",
						// this should map to projects/gateway/pkg/defaults.ValidationWebhookBindPort
						"targetPort": 8443,
					},
				},
				"selector": map[string]interface{}{
					"gloo": "gloo",
				},
			},
		},
	}

	return mutate.MutateServiceNuklerosGatewaySystemGloo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNuklerosGatewaySystemGatewayProxy creates the Service resource with name gateway-proxy.
func CreateServiceNuklerosGatewaySystemGatewayProxy(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":              "gloo",
					"gloo":             "gateway-proxy",
					"gateway-proxy-id": "gateway-proxy",
				},
				"name":      "gateway-proxy",
				"namespace": "nukleros-gateway-system",
			},
			"spec": map[string]interface{}{
				// port order matters due to this issue: https://github.com/solo-io/gloo/issues/2571
				"ports": []interface{}{
					map[string]interface{}{
						"port":       80,
						"targetPort": 8080,
						"protocol":   "TCP",
						"name":       "http",
					},
					map[string]interface{}{
						"port":       443,
						"targetPort": 8443,
						"protocol":   "TCP",
						"name":       "https",
					},
				},
				"selector": map[string]interface{}{
					"gateway-proxy-id": "gateway-proxy",
					"gateway-proxy":    "live",
				},
				"type": "LoadBalancer",
			},
		},
	}

	return mutate.MutateServiceNuklerosGatewaySystemGatewayProxy(resourceObj, parent, collection, reconciler, req)
}
