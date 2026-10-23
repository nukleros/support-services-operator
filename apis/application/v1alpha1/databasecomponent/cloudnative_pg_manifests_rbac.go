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

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceCnpgManager creates the ServiceAccount resource with name cnpg-manager.
func CreateServiceAccountNamespaceCnpgManager(
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
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "cnpg-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceCnpgManager(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databases,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databases/status,verbs=get

// CreateClusterRoleCnpgDatabaseEditorRole creates the ClusterRole resource with name cnpg-database-editor-role.
func CreateClusterRoleCnpgDatabaseEditorRole(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by":     "kustomize",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name": "cnpg-database-editor-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"databases",
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
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"databases/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgDatabaseEditorRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databases,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databases/status,verbs=get

// CreateClusterRoleCnpgDatabaseViewerRole creates the ClusterRole resource with name cnpg-database-viewer-role.
func CreateClusterRoleCnpgDatabaseViewerRole(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by":     "kustomize",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name": "cnpg-database-viewer-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"databases",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"databases/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgDatabaseViewerRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=configmaps/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=core,resources=secrets/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=create;delete;get;list;patch;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=create;delete;get;list;patch;watch
// +kubebuilder:rbac:groups=core,resources=pods/exec,verbs=create;delete;get;list;patch;watch
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;delete;get;list;patch;watch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;get;list;update;watch
// +kubebuilder:rbac:groups=discovery.k8s.io,resources=endpointslices,verbs=get;list;watch
// +kubebuilder:rbac:groups=monitoring.coreos.com,resources=podmonitors,verbs=create;delete;get;list;patch;watch
// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=backups,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=clusters,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databaseroles,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databases,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=poolers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=publications,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=scheduledbackups,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=subscriptions,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=backups/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databases/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=publications/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=scheduledbackups/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=subscriptions/status,verbs=get;patch;update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=clusterimagecatalogs,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=imagecatalogs,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=clusters/finalizers,verbs=update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databaseroles/finalizers,verbs=update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=poolers/finalizers,verbs=update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=clusters/status,verbs=get;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=databaseroles/status,verbs=get;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=failoverquorums/status,verbs=get;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=poolers/status,verbs=get;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=failoverquorums,verbs=create;delete;get;list;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=create;get;list;patch;update;watch
// +kubebuilder:rbac:groups=snapshot.storage.k8s.io,resources=volumesnapshots,verbs=create;get;list;patch;watch

// CreateClusterRoleCnpgManager creates the ClusterRole resource with name cnpg-manager.
func CreateClusterRoleCnpgManager(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "cnpg-manager",
				"labels": map[string]interface{}{
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
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
						"configmaps/status",
						"secrets/status",
					},
					"verbs": []interface{}{
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
						"pods",
						"pods/exec",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods/status",
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
						"serviceaccounts",
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
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"mutatingwebhookconfigurations",
						"validatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"get",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"deployments",
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
						"jobs",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"watch",
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
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"discovery.k8s.io",
					},
					"resources": []interface{}{
						"endpointslices",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"monitoring.coreos.com",
					},
					"resources": []interface{}{
						"podmonitors",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"patch",
						"watch",
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
						"list",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"backups",
						"clusters",
						"databaseroles",
						"databases",
						"poolers",
						"publications",
						"scheduledbackups",
						"subscriptions",
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
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"backups/status",
						"databases/status",
						"publications/status",
						"scheduledbackups/status",
						"subscriptions/status",
					},
					"verbs": []interface{}{
						"get",
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"clusterimagecatalogs",
						"imagecatalogs",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"clusters/finalizers",
						"databaseroles/finalizers",
						"poolers/finalizers",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"clusters/status",
						"databaseroles/status",
						"failoverquorums/status",
						"poolers/status",
					},
					"verbs": []interface{}{
						"get",
						"patch",
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"failoverquorums",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"rbac.authorization.k8s.io",
					},
					"resources": []interface{}{
						"rolebindings",
						"roles",
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
						"snapshot.storage.k8s.io",
					},
					"resources": []interface{}{
						"volumesnapshots",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"list",
						"patch",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgManager(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=publications,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=publications/status,verbs=get

// CreateClusterRoleCnpgPublicationEditorRole creates the ClusterRole resource with name cnpg-publication-editor-role.
func CreateClusterRoleCnpgPublicationEditorRole(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by":     "kustomize",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name": "cnpg-publication-editor-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"publications",
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
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"publications/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgPublicationEditorRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=publications,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=publications/status,verbs=get

// CreateClusterRoleCnpgPublicationViewerRole creates the ClusterRole resource with name cnpg-publication-viewer-role.
func CreateClusterRoleCnpgPublicationViewerRole(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by":     "kustomize",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name": "cnpg-publication-viewer-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"publications",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"publications/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgPublicationViewerRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=subscriptions,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=subscriptions/status,verbs=get

// CreateClusterRoleCnpgSubscriptionEditorRole creates the ClusterRole resource with name cnpg-subscription-editor-role.
func CreateClusterRoleCnpgSubscriptionEditorRole(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by":     "kustomize",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name": "cnpg-subscription-editor-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"subscriptions",
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
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"subscriptions/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgSubscriptionEditorRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=subscriptions,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=subscriptions/status,verbs=get

// CreateClusterRoleCnpgSubscriptionViewerRole creates the ClusterRole resource with name cnpg-subscription-viewer-role.
func CreateClusterRoleCnpgSubscriptionViewerRole(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by":     "kustomize",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name": "cnpg-subscription-viewer-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"subscriptions",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"postgresql.cnpg.io",
					},
					"resources": []interface{}{
						"subscriptions/status",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleCnpgSubscriptionViewerRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingCnpgManagerRolebinding creates the ClusterRoleBinding resource with name cnpg-manager-rolebinding.
func CreateClusterRoleBindingCnpgManagerRolebinding(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "cnpg-manager-rolebinding",
				"labels": map[string]interface{}{
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/name":           "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "cnpg-manager",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "cnpg-manager",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingCnpgManagerRolebinding(resourceObj, parent, collection, reconciler, req)
}
