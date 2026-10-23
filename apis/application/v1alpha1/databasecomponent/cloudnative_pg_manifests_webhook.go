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

package databasecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	applicationv1alpha1 "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/application/v1alpha1/databasecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateMutatingWebhookCnpgMutatingWebhookConfiguration creates the MutatingWebhookConfiguration resource with name cnpg-mutating-webhook-configuration.
func CreateMutatingWebhookCnpgMutatingWebhookConfiguration(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Provider != "cloudnative-pg" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=provider,value=cloudnative-pg,include
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "cnpg-mutating-webhook-configuration",
				"labels": map[string]interface{}{
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/mutate-postgresql-cnpg-io-v1-backup",
						},
					},
					"failurePolicy": "Fail",
					"name":          "mbackup.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"backups",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/mutate-postgresql-cnpg-io-v1-cluster",
						},
					},
					"failurePolicy": "Fail",
					"name":          "mcluster.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"clusters",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/mutate-postgresql-cnpg-io-v1-database",
						},
					},
					"failurePolicy": "Fail",
					"name":          "mdatabase.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"databases",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/mutate-postgresql-cnpg-io-v1-scheduledbackup",
						},
					},
					"failurePolicy": "Fail",
					"name":          "mscheduledbackup.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"scheduledbackups",
							},
						},
					},
					"sideEffects": "None",
				},
			},
		},
	}

	return mutate.MutateMutatingWebhookCnpgMutatingWebhookConfiguration(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingWebhookCnpgValidatingWebhookConfiguration creates the ValidatingWebhookConfiguration resource with name cnpg-validating-webhook-configuration.
func CreateValidatingWebhookCnpgValidatingWebhookConfiguration(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Provider != "cloudnative-pg" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=provider,value=cloudnative-pg,include
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "cnpg-validating-webhook-configuration",
				"labels": map[string]interface{}{
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-postgresql-cnpg-io-v1-backup",
						},
					},
					"failurePolicy": "Fail",
					"name":          "vbackup.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"backups",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-postgresql-cnpg-io-v1-cluster",
						},
					},
					"failurePolicy": "Fail",
					"name":          "vcluster.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"clusters",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-postgresql-cnpg-io-v1-database",
						},
					},
					"failurePolicy": "Fail",
					"name":          "vdatabase.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"databases",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-postgresql-cnpg-io-v1-pooler",
						},
					},
					"failurePolicy": "Fail",
					"name":          "vpooler.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"poolers",
							},
						},
					},
					"sideEffects": "None",
				},
				map[string]interface{}{
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cnpg-webhook-service",
							"namespace": parent.Spec.Namespace, //  controlled by field: namespace
							"path":      "/validate-postgresql-cnpg-io-v1-scheduledbackup",
						},
					},
					"failurePolicy": "Fail",
					"name":          "vscheduledbackup.cnpg.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"postgresql.cnpg.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"scheduledbackups",
							},
						},
					},
					"sideEffects": "None",
				},
			},
		},
	}

	return mutate.MutateValidatingWebhookCnpgValidatingWebhookConfiguration(resourceObj, parent, collection, reconciler, req)
}
