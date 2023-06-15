/*
Copyright 2023.

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

package externaldns

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	ingressv1alpha1 "github.com/nukleros/support-services-operator/apis/ingress/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/ingress/v1alpha1/externaldns/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceServiceAccountName creates the ServiceAccount resource with name parent.Spec.ServiceAccountName.
func CreateServiceAccountNamespaceServiceAccountName(
	parent *ingressv1alpha1.ExternalDNS,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				// controlled by field: serviceAccountName
				//  The name of the external-dns service account which is referenced in role policy doc for AWS.
				"name": parent.Spec.ServiceAccountName,
				"labels": map[string]interface{}{
					"platform.nukleros.io/group":   "ingress",
					"platform.nukleros.io/project": "external-dns",
				},
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"annotations": map[string]interface{}{
					// controlled by field: iamRoleArn
					//  On AWS, the IAM Role ARN that gives external-dns access to Route53
					"eks.amazonaws.com/role-arn": parent.Spec.IamRoleArn,
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceServiceAccountName(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;watch;list
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;watch;list
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;watch;list
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;watch;list
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;watch;list
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=watch;list

// CreateClusterRoleNamespaceExternalDns creates the ClusterRole resource with name external-dns.
func CreateClusterRoleNamespaceExternalDns(
	parent *ingressv1alpha1.ExternalDNS,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "external-dns",
				"labels": map[string]interface{}{
					"platform.nukleros.io/group":   "ingress",
					"platform.nukleros.io/project": "external-dns",
				},
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

	return mutate.MutateClusterRoleNamespaceExternalDns(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingExternalDnsViewer creates the ClusterRoleBinding resource with name external-dns-viewer.
func CreateClusterRoleBindingExternalDnsViewer(
	parent *ingressv1alpha1.ExternalDNS,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "external-dns-viewer",
				"labels": map[string]interface{}{
					"platform.nukleros.io/group":   "ingress",
					"platform.nukleros.io/project": "external-dns",
				},
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

	return mutate.MutateClusterRoleBindingExternalDnsViewer(resourceObj, parent, collection, reconciler, req)
}
