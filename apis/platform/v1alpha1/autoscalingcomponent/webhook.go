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

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookKedaAdmission creates the ValidatingWebhookConfiguration resource with name keda-admission.
func CreateValidatingWebhookKedaAdmission(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-admission-webhooks",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-admission",
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "keda-admission-webhooks",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-keda-sh-v1alpha1-scaledobject",
						},
					},
					"failurePolicy":     "Ignore",
					"matchPolicy":       "Equivalent",
					"name":              "vscaledobject.kb.io",
					"namespaceSelector": map[string]interface{}{},
					"objectSelector":    map[string]interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"keda.sh",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"scaledobjects",
							},
						},
					},
					"sideEffects":    "None",
					"timeoutSeconds": 10,
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "keda-admission-webhooks",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-keda-sh-v1alpha1-scaledjob",
						},
					},
					"failurePolicy":     "Ignore",
					"matchPolicy":       "Equivalent",
					"name":              "vscaledjob.kb.io",
					"namespaceSelector": map[string]interface{}{},
					"objectSelector":    map[string]interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"keda.sh",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"scaledjobs",
							},
						},
					},
					"sideEffects":    "None",
					"timeoutSeconds": 10,
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "keda-admission-webhooks",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-keda-sh-v1alpha1-triggerauthentication",
						},
					},
					"failurePolicy":     "Ignore",
					"matchPolicy":       "Equivalent",
					"name":              "vstriggerauthentication.kb.io",
					"namespaceSelector": map[string]interface{}{},
					"objectSelector":    map[string]interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"keda.sh",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"triggerauthentications",
							},
						},
					},
					"sideEffects":    "None",
					"timeoutSeconds": 10,
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "keda-admission-webhooks",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-keda-sh-v1alpha1-clustertriggerauthentication",
						},
					},
					"failurePolicy":     "Ignore",
					"matchPolicy":       "Equivalent",
					"name":              "vsclustertriggerauthentication.kb.io",
					"namespaceSelector": map[string]interface{}{},
					"objectSelector":    map[string]interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"keda.sh",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"clustertriggerauthentications",
							},
						},
					},
					"sideEffects":    "None",
					"timeoutSeconds": 10,
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "keda-admission-webhooks",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-eventing-keda-sh-v1alpha1-cloudeventsource",
						},
					},
					"failurePolicy":     "Ignore",
					"matchPolicy":       "Equivalent",
					"name":              "vcloudeventsource.kb.io",
					"namespaceSelector": map[string]interface{}{},
					"objectSelector":    map[string]interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"eventing.keda.sh",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"cloudeventsources",
							},
						},
					},
					"sideEffects":    "None",
					"timeoutSeconds": 10,
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "keda-admission-webhooks",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-eventing-keda-sh-v1alpha1-clustercloudeventsource",
						},
					},
					"failurePolicy":     "Ignore",
					"matchPolicy":       "Equivalent",
					"name":              "vclustercloudeventsource.kb.io",
					"namespaceSelector": map[string]interface{}{},
					"objectSelector":    map[string]interface{}{},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"eventing.keda.sh",
							},
							"apiVersions": []interface{}{
								"v1alpha1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"clustercloudeventsources",
							},
						},
					},
					"sideEffects":    "None",
					"timeoutSeconds": 10,
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookKedaAdmission(resourceObj, parent, collection, reconciler, req)
}
