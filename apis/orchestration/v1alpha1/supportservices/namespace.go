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

package supportservices

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1/supportservices/mutate"
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// CreateNamespaceParentName creates the Namespace resource with name parent.Name.
func CreateNamespaceParentName(
	parent *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// controlled by field: tier
			// controlled by field: defaultIngressController
			//  +kubebuilder:validation:Enum=development;staging;production
			//  The tier of cluster being used.  One of: development | staging | production.
			//  +kubebuilder:validation:Enum=kong;nginx
			//  The default ingress for setting TLS certs.  One of: kong | nginx.
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": parent.Name, //  controlled by field:
			},
		},
	}

	return mutate.MutateNamespaceParentName(resourceObj, parent, reconciler, req)
}
