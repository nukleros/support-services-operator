/*
Copyright 2026.

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

package autoscalingcomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/autoscalingcomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceKedaOperator creates the NetworkPolicy resource with name keda-operator.
func CreateNetworkPolicyNamespaceKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name":      "keda-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "keda-operator",
					"name":                          "keda-operator",
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-operator",
					},
				},
				"policyTypes": []interface{}{
					"Ingress",
					"Egress",
				},
				"ingress": []interface{}{
					map[string]interface{}{
						"from": []interface{}{
							map[string]interface{}{
								"podSelector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app": "keda-operator-metrics-apiserver",
									},
								},
							},
						},
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     9666,
							},
						},
					},
					map[string]interface{}{
						"from": []interface{}{
							map[string]interface{}{
								"namespaceSelector": map[string]interface{}{},
							},
						},
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     8080,
							},
						},
					},
				},
				"egress": []interface{}{
					map[string]interface{}{},
				},
			},
		},
	}

	return mutate.MutateNetworkPolicyNamespaceKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceKedaOperatorMetricsApiserver creates the NetworkPolicy resource with name keda-operator-metrics-apiserver.
func CreateNetworkPolicyNamespaceKedaOperatorMetricsApiserver(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name":      "keda-operator-metrics-apiserver",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "keda-operator-metrics-apiserver",
					"app.kubernetes.io/name":        "keda-operator-metrics-apiserver",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-operator-metrics-apiserver",
					},
				},
				"policyTypes": []interface{}{
					"Ingress",
					"Egress",
				},
				"ingress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     6443,
							},
						},
					},
					map[string]interface{}{
						"from": []interface{}{
							map[string]interface{}{
								"namespaceSelector": map[string]interface{}{},
							},
						},
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     8080,
							},
						},
					},
				},
				"egress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "UDP",
								"port":     53,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     53,
							},
							map[string]interface{}{
								"protocol": "UDP",
								"port":     5353,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     5353,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     443,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     6443,
							},
						},
					},
					map[string]interface{}{
						"to": []interface{}{
							map[string]interface{}{
								"podSelector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app": "keda-operator",
									},
								},
							},
						},
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     9666,
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateNetworkPolicyNamespaceKedaOperatorMetricsApiserver(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceKedaAdmissionWebhooks creates the NetworkPolicy resource with name keda-admission-webhooks.
func CreateNetworkPolicyNamespaceKedaAdmissionWebhooks(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name":      "keda-admission-webhooks",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "keda-admission-webhooks",
					"name":                          "keda-admission-webhooks",
					"app.kubernetes.io/name":        "keda-admission-webhooks",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-admission-webhooks",
					},
				},
				"policyTypes": []interface{}{
					"Ingress",
					"Egress",
				},
				"ingress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     9443,
							},
						},
					},
					map[string]interface{}{
						"from": []interface{}{
							map[string]interface{}{
								"namespaceSelector": map[string]interface{}{},
							},
						},
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     8080,
							},
						},
					},
				},
				"egress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "UDP",
								"port":     53,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     53,
							},
							map[string]interface{}{
								"protocol": "UDP",
								"port":     5353,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     5353,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     443,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     6443,
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateNetworkPolicyNamespaceKedaAdmissionWebhooks(resourceObj, parent, collection, reconciler, req)
}
