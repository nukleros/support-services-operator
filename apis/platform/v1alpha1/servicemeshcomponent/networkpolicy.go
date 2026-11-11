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

package servicemeshcomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/servicemeshcomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceIstioCni creates the NetworkPolicy resource with name istio-cni.
func CreateNetworkPolicyNamespaceIstioCni(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name":      "istio-cni",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"k8s-app":      "istio-cni-node",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"k8s-app": "istio-cni-node",
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
								"port":     15014,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     8000,
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

	return mutate.MutateNetworkPolicyNamespaceIstioCni(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceIstiod creates the NetworkPolicy resource with name istiod.
func CreateNetworkPolicyNamespaceIstiod(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name":      "istiod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":          "istiod",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Pilot",
					"istio":                                     "pilot",
					"app.kubernetes.io/name":                    "istiod",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app":          "istiod",
						"istio.io/rev": "default",
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
								"port":     15017,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15010,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15011,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15012,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     8080,
							},
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15014,
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

	return mutate.MutateNetworkPolicyNamespaceIstiod(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceZtunnel creates the NetworkPolicy resource with name ztunnel.
func CreateNetworkPolicyNamespaceZtunnel(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name":      "ztunnel",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                    "ztunnel",
					"app.kubernetes.io/name": "ztunnel",
					"istio.io/rev":           "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Ztunnel",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "ztunnel",
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
								"port":     15021,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15020,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15000,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15008,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15001,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15006,
							},
						},
					},
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"protocol": "TCP",
								"port":     15053,
							},
							map[string]interface{}{
								"protocol": "UDP",
								"port":     15053,
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

	return mutate.MutateNetworkPolicyNamespaceZtunnel(resourceObj, parent, collection, reconciler, req)
}
