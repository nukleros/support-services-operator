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

package secretscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/secretscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// CreateMutatingWebhookOpenbaoAgentInjectorCfg creates the MutatingWebhookConfiguration resource with name openbao-agent-injector-cfg.
func CreateMutatingWebhookOpenbaoAgentInjectorCfg(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Openbao.Injector.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=openbao.injector.include,value=true,include
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "openbao-agent-injector-cfg",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao-agent-injector",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name":           "vault.hashicorp.com",
					"failurePolicy":  "Ignore",
					"matchPolicy":    "Exact",
					"sideEffects":    "None",
					"timeoutSeconds": 30,
					"admissionReviewVersions": []interface{}{
						"v1",
						"v1beta1",
					},
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "openbao-agent-injector-svc",
							"namespace": "nukleros-secrets-system",
							"path":      "/mutate",
						},
						"caBundle": "",
					},
					"rules": []interface{}{
						map[string]interface{}{
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
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
					"objectSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "app.kubernetes.io/name",
								"operator": "NotIn",
								"values": []interface{}{
									"openbao-agent-injector",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateMutatingWebhookOpenbaoAgentInjectorCfg(resourceObj, parent, collection, reconciler, req)
}
