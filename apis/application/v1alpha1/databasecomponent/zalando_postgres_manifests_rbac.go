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

package databasecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	applicationv1alpha1 "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

const ServiceAccountNamespacePostgresOperator = "postgres-operator"

// CreateServiceAccountNamespacePostgresOperator creates the postgres-operator ServiceAccount resource.
func CreateServiceAccountNamespacePostgresOperator(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "postgres-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=acid.zalan.do,resources=postgresqls,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=acid.zalan.do,resources=postgresqls/status,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=acid.zalan.do,resources=operatorconfigurations,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=acid.zalan.do,resources=postgresteams,verbs=get;list;watch
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=create;get;patch;update
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;delete;get;update
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=delete;get;list;patch;update
// +kubebuilder:rbac:groups=core,resources=persistentvolumes,verbs=get;list;update
// +kubebuilder:rbac:groups=core,resources=pods,verbs=delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=pods/exec,verbs=create
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;delete;get;patch;update
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=create;delete;get;list;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;patch
// +kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=create;delete;get;list;patch;update
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get
// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=create;delete;get
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;create
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;create

const ClusterRolePostgresOperator = "postgres-operator"

// CreateClusterRolePostgresOperator creates the postgres-operator ClusterRole resource.
func CreateClusterRolePostgresOperator(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "postgres-operator",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acid.zalan.do",
					},
					"resources": []interface{}{
						"postgresqls",
						"postgresqls/status",
						"operatorconfigurations",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"acid.zalan.do",
					},
					"resources": []interface{}{
						"postgresteams",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apiextensions.k8s.io",
					},
					"resources": []interface{}{
						"customresourcedefinitions",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
					},
					"verbs": []interface{}{
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"events",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"endpoints",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"update",
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
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"persistentvolumeclaims",
					},
					"verbs": []interface{}{
						"delete",
						"get",
						"list",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"persistentvolumes",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
					},
					"verbs": []interface{}{
						"delete",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods/exec",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"services",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"statefulsets",
						"deployments",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"batch",
					},
					"resources": []interface{}{
						"cronjobs",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"namespaces",
					},
					"verbs": []interface{}{
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"policy",
					},
					"resources": []interface{}{
						"poddisruptionbudgets",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"serviceaccounts",
					},
					"verbs": []interface{}{
						"get",
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"rbac.authorization.k8s.io",
					},
					"resources": []interface{}{
						"rolebindings",
					},
					"verbs": []interface{}{
						"get",
						"create",
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

const ClusterRoleBindingPostgresOperator = "postgres-operator"

// CreateClusterRoleBindingPostgresOperator creates the postgres-operator ClusterRoleBinding resource.
func CreateClusterRoleBindingPostgresOperator(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "postgres-operator",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "postgres-operator",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "postgres-operator",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=create;delete;deletecollection;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=create

const ClusterRolePostgresPod = "postgres-pod"

// CreateClusterRolePostgresPod creates the postgres-pod ClusterRole resource.
func CreateClusterRolePostgresPod(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "postgres-pod",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"endpoints",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"services",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
