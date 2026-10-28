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

package storagecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/storagecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceSeaweedfsOperator creates the ServiceAccount resource with name seaweedfs-operator.
func CreateServiceAccountNamespaceSeaweedfsOperator(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"automountServiceAccountToken": true,
		},
	}

	return mutate.MutateServiceAccountNamespaceSeaweedfsOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=batch,resources=cronjobs,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=monitoring.coreos.com,resources=servicemonitors,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=adminscripts,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=bucketlifecyclepolicies,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=bucketnotifications,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=buckets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3credentials,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3identities,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3oidcproviders,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policies,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policybindings,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweedbackups,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweedrestores,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=adminscripts/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=bucketlifecyclepolicies/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=bucketnotifications/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=buckets/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3credentials/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3identities/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3oidcproviders/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policies/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policybindings/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweedbackups/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweedrestores/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds/finalizers,verbs=update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=adminscripts/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=bucketlifecyclepolicies/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=bucketnotifications/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=buckets/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3credentials/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3identities/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3oidcproviders/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policies/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policybindings/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweedbackups/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweedrestores/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=resourcereferencegrants,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=storage.k8s.io,resources=storageclasses,verbs=get;list;watch

// CreateClusterRoleSeaweedfsOperatorManagerRole creates the ClusterRole resource with name seaweedfs-operator-manager-role.
func CreateClusterRoleSeaweedfsOperatorManagerRole(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "seaweedfs-operator-manager-role",
				"labels": map[string]interface{}{
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
						"secrets",
						"services",
					},
					"verbs": []interface{}{
						"create",
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
						"events",
					},
					"verbs": []interface{}{
						"create",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"namespaces",
						"pods",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"daemonsets",
						"deployments",
						"statefulsets",
					},
					"verbs": []interface{}{
						"create",
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
						"batch",
					},
					"resources": []interface{}{
						"cronjobs",
						"jobs",
					},
					"verbs": []interface{}{
						"create",
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
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
						"issuers",
					},
					"verbs": []interface{}{
						"create",
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
						"extensions",
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"create",
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
						"monitoring.coreos.com",
					},
					"resources": []interface{}{
						"servicemonitors",
					},
					"verbs": []interface{}{
						"create",
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
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"adminscripts",
						"bucketlifecyclepolicies",
						"bucketnotifications",
						"buckets",
						"s3credentials",
						"s3identities",
						"s3oidcproviders",
						"s3policies",
						"s3policybindings",
						"seaweedbackups",
						"seaweedrestores",
						"seaweeds",
					},
					"verbs": []interface{}{
						"create",
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
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"adminscripts/finalizers",
						"bucketlifecyclepolicies/finalizers",
						"bucketnotifications/finalizers",
						"buckets/finalizers",
						"s3credentials/finalizers",
						"s3identities/finalizers",
						"s3oidcproviders/finalizers",
						"s3policies/finalizers",
						"s3policybindings/finalizers",
						"seaweedbackups/finalizers",
						"seaweedrestores/finalizers",
						"seaweeds/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"adminscripts/status",
						"bucketlifecyclepolicies/status",
						"bucketnotifications/status",
						"buckets/status",
						"s3credentials/status",
						"s3identities/status",
						"s3oidcproviders/status",
						"s3policies/status",
						"s3policybindings/status",
						"seaweedbackups/status",
						"seaweedrestores/status",
						"seaweeds/status",
					},
					"verbs": []interface{}{
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"resourcereferencegrants",
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
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"storage.k8s.io",
					},
					"resources": []interface{}{
						"storageclasses",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleSeaweedfsOperatorManagerRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds/status,verbs=get

// CreateClusterRoleSeaweedEditorRole creates the ClusterRole resource with name seaweed-editor-role.
func CreateClusterRoleSeaweedEditorRole(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "seaweed-editor-role",
				"labels": map[string]interface{}{
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"seaweeds",
					},
					"verbs": []interface{}{
						"create",
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
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"seaweeds/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleSeaweedEditorRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds,verbs=get;list;watch
// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds/status,verbs=get

// CreateClusterRoleSeaweedViewerRole creates the ClusterRole resource with name seaweed-viewer-role.
func CreateClusterRoleSeaweedViewerRole(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "seaweed-viewer-role",
				"labels": map[string]interface{}{
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"seaweeds",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"seaweed.seaweedfs.com",
					},
					"resources": []interface{}{
						"seaweeds/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleSeaweedViewerRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingSeaweedfsOperatorManagerRolebinding creates the ClusterRoleBinding resource with name seaweedfs-operator-manager-rolebinding.
func CreateClusterRoleBindingSeaweedfsOperatorManagerRolebinding(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "seaweedfs-operator-manager-rolebinding",
				"labels": map[string]interface{}{
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "seaweedfs-operator-manager-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "seaweedfs-operator",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingSeaweedfsOperatorManagerRolebinding(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;list;update

// CreateRoleNamespaceSeaweedfsOperatorLeaderElectionRole creates the Role resource with name seaweedfs-operator-leader-election-role.
func CreateRoleNamespaceSeaweedfsOperatorLeaderElectionRole(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-operator-leader-election-role",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"update",
						"patch",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps/status",
					},
					"verbs": []interface{}{
						"get",
						"update",
						"patch",
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
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"list",
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceSeaweedfsOperatorLeaderElectionRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceSeaweedfsOperatorLeaderElectionRolebinding creates the RoleBinding resource with name seaweedfs-operator-leader-election-rolebinding.
func CreateRoleBindingNamespaceSeaweedfsOperatorLeaderElectionRolebinding(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-operator-leader-election-rolebinding",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "seaweedfs-operator-leader-election-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "seaweedfs-operator",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceSeaweedfsOperatorLeaderElectionRolebinding(resourceObj, parent, collection, reconciler, req)
}
