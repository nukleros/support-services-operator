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

package storagecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/storagecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateMutatingWebhookSeaweedfsOperatorMutatingWebhookConfiguration creates the MutatingWebhookConfiguration resource with name seaweedfs-operator-mutating-webhook-configuration.
func CreateMutatingWebhookSeaweedfsOperatorMutatingWebhookConfiguration(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "seaweedfs-operator-mutating-webhook-configuration",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
				"annotations": map[string]interface{}{
					// controlled by field: namespace
					"cert-manager.io/inject-ca-from": "" + parent.Spec.Namespace + "/seaweedfs-operator-webhook-cert",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "seaweedfs-operator-webhook",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"port":      443,
							"path":      "/mutate-seaweed-seaweedfs-com-v1-seaweed",
						},
					},
					"name":          "mseaweed.kb.io",
					"failurePolicy": "Fail",
					"sideEffects":   "None",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"seaweed.seaweedfs.com",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"seaweeds",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateMutatingWebhookSeaweedfsOperatorMutatingWebhookConfiguration(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookSeaweedfsOperatorValidatingWebhookConfiguration creates the ValidatingWebhookConfiguration resource with name seaweedfs-operator-validating-webhook-configuration.
func CreateValidatingWebhookSeaweedfsOperatorValidatingWebhookConfiguration(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "seaweedfs-operator-validating-webhook-configuration",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
				"annotations": map[string]interface{}{
					// controlled by field: namespace
					"cert-manager.io/inject-ca-from": "" + parent.Spec.Namespace + "/seaweedfs-operator-webhook-cert",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "seaweedfs-operator-webhook",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"port":      443,
							"path":      "/validate-seaweed-seaweedfs-com-v1-seaweed",
						},
					},
					"name":          "vseaweed.kb.io",
					"failurePolicy": "Fail",
					"sideEffects":   "None",
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"seaweed.seaweedfs.com",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"seaweeds",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookSeaweedfsOperatorValidatingWebhookConfiguration(resourceObj, parent, collection, reconciler, req)
}
