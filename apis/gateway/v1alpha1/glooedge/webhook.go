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

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookGlooGatewayValidationWebhookDefault creates the ValidatingWebhookConfiguration resource with name gloo-gateway-validation-webhook-default.
func CreateValidatingWebhookGlooGatewayValidationWebhookDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "gloo-gateway-validation-webhook-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "gateway",
				},
				"annotations": map[string]interface{}{
					"cert-manager.io/inject-ca-from-secret": "nukleros-gateway-system/certificate-authority",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "gloo.default.svc", //  must be a domain with at least three segments separated by dots
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "gloo",
							"namespace": "nukleros-gateway-system",
							"path":      "/validation",
						},
					},
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"apiGroups": []interface{}{
								"gateway.solo.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"virtualservices",
							},
						},
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"apiGroups": []interface{}{
								"gateway.solo.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"routetables",
							},
						},
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"apiGroups": []interface{}{
								"gateway.solo.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"gateways",
							},
						},
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"apiGroups": []interface{}{
								"gloo.solo.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"upstreams",
							},
						},
						map[string]interface{}{
							"operations": []interface{}{
								"DELETE",
							},
							"apiGroups": []interface{}{
								"gloo.solo.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"secrets",
							},
						},
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"apiGroups": []interface{}{
								"ratelimit.solo.io",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"resources": []interface{}{
								"ratelimitconfigs",
							},
						},
					},
					"sideEffects": "None",
					"matchPolicy": "Exact",
					"admissionReviewVersions": []interface{}{
						"v1beta1", //  v1beta1 still live in 1.22 https://github.com/kubernetes/api/blob/release-1.22/admission/v1beta1/types.go#L33
					},
					"failurePolicy": "Ignore",
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookGlooGatewayValidationWebhookDefault(resourceObj, parent, collection, reconciler, req)
}
