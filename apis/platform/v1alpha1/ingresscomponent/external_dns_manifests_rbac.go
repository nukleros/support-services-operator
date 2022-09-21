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

package ingresscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

const ClusterRoleBindingExternalDnsViewer = "external-dns-viewer"

// CreateClusterRoleBindingExternalDnsViewer creates the external-dns-viewer ClusterRoleBinding resource.
func CreateClusterRoleBindingExternalDnsViewer(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "external-dns-viewer",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "external-dns",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "external-dns",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;watch;list
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;watch;list
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;watch;list
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;watch;list
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;watch;list
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=watch;list

const ClusterRoleNamespaceExternalDns = "external-dns"

// CreateClusterRoleNamespaceExternalDns creates the external-dns ClusterRole resource.
func CreateClusterRoleNamespaceExternalDns(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "external-dns",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"endpoints",
						"pods",
						"services",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"extensions",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"nodes",
					},
					"verbs": []interface{}{
						"watch",
						"list",
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

const ServiceAccountNamespaceExternalDns = "external-dns"

// CreateServiceAccountNamespaceExternalDns creates the external-dns ServiceAccount resource.
func CreateServiceAccountNamespaceExternalDns(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "external-dns",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
