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

// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateNetworkPolicyNamespaceOpenbao creates the NetworkPolicy resource with name openbao.
func CreateNetworkPolicyNamespaceOpenbao(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1",
			"kind":       "NetworkPolicy",
			"metadata": map[string]interface{}{
				"name": "openbao",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"podSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":     "openbao",
						"app.kubernetes.io/instance": "openbao",
					},
				},
				"ingress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"port":     8200,
								"protocol": "TCP",
							},
							map[string]interface{}{
								"port":     8201,
								"protocol": "TCP",
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

	return mutate.MutateNetworkPolicyNamespaceOpenbao(resourceObj, parent, collection, reconciler, req)
}
