/*
Copyright 2022.

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

package supportservicescollection

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

const NamespaceParentName = "parent.Name"

// CreateNamespaceParentName creates the parent.Name Namespace resource.
func CreateNamespaceParentName(
	parent *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": parent.Name, //  controlled by field:
				"labels": map[string]interface{}{
					// controlled by field: cloudProvider
					//  +kubebuilder:validation:Enum=aws;gcp;azure;none
					"cloudProvider": parent.Spec.CloudProvider,
					// controlled by field: certProvider
					//  +kubebuilder:validation:Enum=letsencrypt-staging;letsencrypt-production
					"certProvider": parent.Spec.CertProvider,
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
