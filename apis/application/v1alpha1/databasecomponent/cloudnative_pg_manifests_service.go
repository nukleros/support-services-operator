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

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceCnpgWebhookService creates the Service resource with name cnpg-webhook-service.
func CreateServiceNamespaceCnpgWebhookService(
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
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "cnpg-webhook-service",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port":       443,
						"targetPort": 9443,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name": "cloudnative-pg",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceCnpgWebhookService(resourceObj, parent, collection, reconciler, req)
}
