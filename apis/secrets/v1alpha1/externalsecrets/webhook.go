/*
Copyright 2024.

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

package externalsecrets

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
	secretsv1alpha1 "github.com/nukleros/support-services-operator/apis/secrets/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/secrets/v1alpha1/externalsecrets/mutate"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookSecretstoreValidate creates the ValidatingWebhookConfiguration resource with name secretstore-validate.
func CreateValidatingWebhookSecretstoreValidate(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "secretstore-validate",
				"labels": map[string]interface{}{
					"external-secrets.io/component": "webhook",
					"platform.nukleros.io/group":    "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "validate.secretstore.external-secrets.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"external-secrets.io",
							},
							"apiVersions": []interface{}{
								"v1beta1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"resources": []interface{}{
								"secretstores",
							},
							"scope": "Namespaced",
						},
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"namespace": "nukleros-secrets-system",
							"name":      "external-secrets-webhook",
							"path":      "/validate-external-secrets-io-v1beta1-secretstore",
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
						"v1beta1",
					},
					"sideEffects":    "None",
					"timeoutSeconds": 5,
				},
				map[string]interface{}{
					"name": "validate.clustersecretstore.external-secrets.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"external-secrets.io",
							},
							"apiVersions": []interface{}{
								"v1beta1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"resources": []interface{}{
								"clustersecretstores",
							},
							"scope": "Cluster",
						},
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"namespace": "nukleros-secrets-system",
							"name":      "external-secrets-webhook",
							"path":      "/validate-external-secrets-io-v1beta1-clustersecretstore",
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
						"v1beta1",
					},
					"sideEffects":    "None",
					"timeoutSeconds": 5,
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookSecretstoreValidate(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookExternalsecretValidate creates the ValidatingWebhookConfiguration resource with name externalsecret-validate.
func CreateValidatingWebhookExternalsecretValidate(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "externalsecret-validate",
				"labels": map[string]interface{}{
					"external-secrets.io/component": "webhook",
					"platform.nukleros.io/group":    "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "validate.externalsecret.external-secrets.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"external-secrets.io",
							},
							"apiVersions": []interface{}{
								"v1beta1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
								"DELETE",
							},
							"resources": []interface{}{
								"externalsecrets",
							},
							"scope": "Namespaced",
						},
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"namespace": "nukleros-secrets-system",
							"name":      "external-secrets-webhook",
							"path":      "/validate-external-secrets-io-v1beta1-externalsecret",
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
						"v1beta1",
					},
					"sideEffects":    "None",
					"timeoutSeconds": 5,
					"failurePolicy":  "Fail",
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookExternalsecretValidate(resourceObj, parent, collection, reconciler, req)
}
