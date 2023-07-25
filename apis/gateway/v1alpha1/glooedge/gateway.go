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

// +kubebuilder:rbac:groups=gateway.solo.io,resources=gateways,verbs=get;list;watch;create;update;patch;delete

// CreateGatewayNamespaceGatewayProxy creates the Gateway resource with name gateway-proxy.
func CreateGatewayNamespaceGatewayProxy(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "gateway.solo.io/v1",
			"kind":       "Gateway",
			"metadata": map[string]interface{}{
				"name":      "gateway-proxy",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app": "gloo",
				},
			},
			"spec": map[string]interface{}{
				"bindAddress":   "::",
				"bindPort":      8080,
				"httpGateway":   map[string]interface{}{},
				"useProxyProto": false,
				"ssl":           false,
				"options": map[string]interface{}{
					"accessLoggingService": map[string]interface{}{
						"accessLog": []interface{}{
							map[string]interface{}{
								"fileSink": map[string]interface{}{
									"path":         "/dev/stdout",
									"stringFormat": "",
								},
							},
						},
					},
				},
				"proxyNames": []interface{}{
					"gateway-proxy",
				},
			},
		},
	}

	return mutate.MutateGatewayNamespaceGatewayProxy(resourceObj, parent, collection, reconciler, req)
}
