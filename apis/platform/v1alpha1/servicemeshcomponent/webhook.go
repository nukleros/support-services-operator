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

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateMutatingWebhookIstioSidecarInjectorNamespace creates the MutatingWebhookConfiguration resource with name istio-sidecar-injector- + parent.spec.namespace.
func CreateMutatingWebhookIstioSidecarInjectorNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istio-sidecar-injector-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Pilot",
					"app":                                       "sidecar-injector",
					"app.kubernetes.io/name":                    "istiod",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "rev.namespace.sidecar-injector.istio.io",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "istiod",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/inject",
							"port":      443,
						},
					},
					"sideEffects": "None",
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
							},
							"apiGroups": []interface{}{
								"",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"pods",
							},
						},
					},
					"failurePolicy":      "Fail",
					"reinvocationPolicy": "Never",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"namespaceSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "istio.io/rev",
								"operator": "In",
								"values": []interface{}{
									"default",
								},
							},
							map[string]interface{}{
								"key":      "istio-injection",
								"operator": "DoesNotExist",
							},
						},
					},
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "sidecar.istio.io/inject",
								"operator": "NotIn",
								"values": []interface{}{
									"false",
								},
							},
						},
					},
				},
				map[string]interface{}{
					"name": "rev.object.sidecar-injector.istio.io",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "istiod",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/inject",
							"port":      443,
						},
					},
					"sideEffects": "None",
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
							},
							"apiGroups": []interface{}{
								"",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"pods",
							},
						},
					},
					"failurePolicy":      "Fail",
					"reinvocationPolicy": "Never",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"namespaceSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "istio.io/rev",
								"operator": "DoesNotExist",
							},
							map[string]interface{}{
								"key":      "istio-injection",
								"operator": "DoesNotExist",
							},
						},
					},
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "sidecar.istio.io/inject",
								"operator": "NotIn",
								"values": []interface{}{
									"false",
								},
							},
							map[string]interface{}{
								"key":      "istio.io/rev",
								"operator": "In",
								"values": []interface{}{
									"default",
								},
							},
						},
					},
				},
				map[string]interface{}{
					"name": "namespace.sidecar-injector.istio.io",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "istiod",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/inject",
							"port":      443,
						},
					},
					"sideEffects": "None",
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
							},
							"apiGroups": []interface{}{
								"",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"pods",
							},
						},
					},
					"failurePolicy":      "Fail",
					"reinvocationPolicy": "Never",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"namespaceSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "istio-injection",
								"operator": "In",
								"values": []interface{}{
									"enabled",
								},
							},
						},
					},
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "sidecar.istio.io/inject",
								"operator": "NotIn",
								"values": []interface{}{
									"false",
								},
							},
						},
					},
				},
				map[string]interface{}{
					"name": "object.sidecar-injector.istio.io",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "istiod",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/inject",
							"port":      443,
						},
					},
					"sideEffects": "None",
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
							},
							"apiGroups": []interface{}{
								"",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"pods",
							},
						},
					},
					"failurePolicy":      "Fail",
					"reinvocationPolicy": "Never",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"namespaceSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "istio-injection",
								"operator": "DoesNotExist",
							},
							map[string]interface{}{
								"key":      "istio.io/rev",
								"operator": "DoesNotExist",
							},
						},
					},
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "sidecar.istio.io/inject",
								"operator": "In",
								"values": []interface{}{
									"true",
								},
							},
							map[string]interface{}{
								"key":      "istio.io/rev",
								"operator": "DoesNotExist",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateMutatingWebhookIstioSidecarInjectorNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookIstiodDefaultValidator creates the ValidatingWebhookConfiguration resource with name istiod-default-validator.
func CreateValidatingWebhookIstiodDefaultValidator(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "istiod-default-validator",
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"istio":                         "istiod",
					"istio.io/rev":                  "default",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "validation.istio.io",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "istiod",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate",
						},
					},
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"apiGroups": []interface{}{
								"security.istio.io",
								"networking.istio.io",
								"telemetry.istio.io",
								"extensions.istio.io",
							},
							"apiVersions": []interface{}{
								"*",
							},
							"resources": []interface{}{
								"*",
							},
						},
					},
					"sideEffects": "None",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookIstiodDefaultValidator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookIstioValidatorNamespace creates the ValidatingWebhookConfiguration resource with name istio-validator- + parent.spec.namespace.
func CreateValidatingWebhookIstioValidatorNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istio-validator-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"istio":                         "istiod",
					"istio.io/rev":                  "default",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "rev.validation.istio.io",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "istiod",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate",
						},
					},
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"apiGroups": []interface{}{
								"security.istio.io",
								"networking.istio.io",
								"telemetry.istio.io",
								"extensions.istio.io",
							},
							"apiVersions": []interface{}{
								"*",
							},
							"resources": []interface{}{
								"*",
							},
						},
					},
					"sideEffects": "None",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "istio.io/rev",
								"operator": "In",
								"values": []interface{}{
									"default",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookIstioValidatorNamespace(resourceObj, parent, collection, reconciler, req)
}
