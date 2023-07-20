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

package glooedge

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/gateway/v1alpha1/glooedge/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemGlooResourceCleanup creates the ServiceAccount resource with name gloo-resource-cleanup.
func CreateServiceAccountNuklerosGatewaySystemGlooResourceCleanup(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
				"name":      "gloo-resource-cleanup",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemGlooResourceCleanup(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemGlooResourceMigration creates the ServiceAccount resource with name gloo-resource-migration.
func CreateServiceAccountNuklerosGatewaySystemGlooResourceMigration(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
				"name":      "gloo-resource-migration",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemGlooResourceMigration(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemGlooResourceRollout creates the ServiceAccount resource with name gloo-resource-rollout.
func CreateServiceAccountNuklerosGatewaySystemGlooResourceRollout(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
				"name":      "gloo-resource-rollout",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemGlooResourceRollout(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemCertgen creates the ServiceAccount resource with name certgen.
func CreateServiceAccountNuklerosGatewaySystemCertgen(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
				"name":      "certgen",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemCertgen(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemGloo creates the ServiceAccount resource with name gloo.
func CreateServiceAccountNuklerosGatewaySystemGloo(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "gloo",
				},
				"name":      "gloo",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemGloo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemDiscovery creates the ServiceAccount resource with name discovery.
func CreateServiceAccountNuklerosGatewaySystemDiscovery(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "discovery",
				},
				"name":      "discovery",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemDiscovery(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNuklerosGatewaySystemGatewayProxy creates the ServiceAccount resource with name gateway-proxy.
func CreateServiceAccountNuklerosGatewaySystemGatewayProxy(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "gateway-proxy",
				},
				"name":      "gateway-proxy",
				"namespace": "nukleros-gateway-system",
			},
		},
	}

	return mutate.MutateServiceAccountNuklerosGatewaySystemGatewayProxy(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch

// CreateClusterRoleKubeResourceWatcherDefault creates the ClusterRole resource with name kube-resource-watcher-default.
func CreateClusterRoleKubeResourceWatcherDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "kube-resource-watcher-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
						"services",
						"secrets",
						"endpoints",
						"configmaps",
						"namespaces",
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

	return mutate.MutateClusterRoleKubeResourceWatcherDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=*
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=*

// CreateClusterRoleKubeLeaderElectionDefault creates the ClusterRole resource with name kube-leader-election-default.
func CreateClusterRoleKubeLeaderElectionDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "kube-leader-election-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"*",
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
						"*",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleKubeLeaderElectionDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gloo.solo.io,resources=upstreams,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleGlooUpstreamMutatorDefault creates the ClusterRole resource with name gloo-upstream-mutator-default.
func CreateClusterRoleGlooUpstreamMutatorDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-upstream-mutator-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"upstreams",
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
			},
		},
	}

	return mutate.MutateClusterRoleGlooUpstreamMutatorDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gloo.solo.io,resources=upstreams,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gloo.solo.io,resources=upstreamgroups,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gloo.solo.io,resources=proxies,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=enterprise.gloo.solo.io,resources=authconfigs,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=ratelimit.solo.io,resources=ratelimitconfigs,verbs=get;list;watch;patch;update
// +kubebuilder:rbac:groups=ratelimit.solo.io,resources=ratelimitconfigs/status,verbs=get;list;watch;patch;update
// +kubebuilder:rbac:groups=graphql.gloo.solo.io,resources=graphqlapis,verbs=get;list;watch;patch;update
// +kubebuilder:rbac:groups=graphql.gloo.solo.io,resources=graphqlapis/status,verbs=get;list;watch;patch;update

// CreateClusterRoleGlooResourceReaderDefault creates the ClusterRole resource with name gloo-resource-reader-default.
func CreateClusterRoleGlooResourceReaderDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-reader-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"upstreams",
						"upstreamgroups",
						"proxies",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"patch", //  needed for status updates for skv1
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"enterprise.gloo.solo.io",
					},
					"resources": []interface{}{
						"authconfigs",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"patch", //  needed for status updates for skv1
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"ratelimit.solo.io",
					},
					"resources": []interface{}{
						"ratelimitconfigs",
						"ratelimitconfigs/status",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"patch",  //  needed for status updates for skv1
						"update", //  needed for status updates for skv2
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"graphql.gloo.solo.io",
					},
					"resources": []interface{}{
						"graphqlapis",
						"graphqlapis/status",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"patch",  //  needed for status updates for skv1
						"update", //  needed for status updates for skv2
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooResourceReaderDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gloo.solo.io,resources=settings,verbs=get;list;watch

// CreateClusterRoleSettingsUserDefault creates the ClusterRole resource with name settings-user-default.
func CreateClusterRoleSettingsUserDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "settings-user-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"settings",
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

	return mutate.MutateClusterRoleSettingsUserDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gloo.solo.io,resources=proxies,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleGlooResourceMutatorDefault creates the ClusterRole resource with name gloo-resource-mutator-default.
func CreateClusterRoleGlooResourceMutatorDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-mutator-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"proxies",
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
			},
		},
	}

	return mutate.MutateClusterRoleGlooResourceMutatorDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gateway.solo.io,resources=gateways,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.solo.io,resources=httpgateways,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.solo.io,resources=tcpgateways,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.solo.io,resources=virtualservices,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.solo.io,resources=routetables,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.solo.io,resources=virtualhostoptions,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.solo.io,resources=routeoptions,verbs=get;list;watch;patch

// CreateClusterRoleGatewayResourceReaderDefault creates the ClusterRole resource with name gateway-resource-reader-default.
func CreateClusterRoleGatewayResourceReaderDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gateway-resource-reader-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.solo.io",
					},
					"resources": []interface{}{
						"gateways",
						"httpgateways",
						"tcpgateways",
						"virtualservices",
						"routetables",
						"virtualhostoptions",
						"routeoptions",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"patch", //  needed for status updates
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGatewayResourceReaderDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=graphql.gloo.solo.io,resources=graphqlapis,verbs=get;list;watch;update;patch;create
// +kubebuilder:rbac:groups=graphql.gloo.solo.io,resources=graphqlapis/status,verbs=get;list;watch;update;patch;create

// CreateClusterRoleGlooGraphqlapiMutatorDefault creates the ClusterRole resource with name gloo-graphqlapi-mutator-default.
func CreateClusterRoleGlooGraphqlapiMutatorDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-graphqlapi-mutator-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"graphql.gloo.solo.io",
					},
					"resources": []interface{}{
						"graphqlapis",
						"graphqlapis/status",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooGraphqlapiMutatorDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKubeResourceWatcherBindingDefault creates the ClusterRoleBinding resource with name kube-resource-watcher-binding-default.
func CreateClusterRoleBindingKubeResourceWatcherBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "kube-resource-watcher-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",
					"namespace": "nukleros-gateway-system",
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "kube-resource-watcher-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingKubeResourceWatcherBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKubeLeaderElectionBindingDefault creates the ClusterRoleBinding resource with name kube-leader-election-binding-default.
func CreateClusterRoleBindingKubeLeaderElectionBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "kube-leader-election-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",
					"namespace": "nukleros-gateway-system",
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "kube-leader-election-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingKubeLeaderElectionBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooUpstreamMutatorBindingDefault creates the ClusterRoleBinding resource with name gloo-upstream-mutator-binding-default.
func CreateClusterRoleBindingGlooUpstreamMutatorBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-upstream-mutator-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-upstream-mutator-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooUpstreamMutatorBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooResourceReaderBindingDefault creates the ClusterRoleBinding resource with name gloo-resource-reader-binding-default.
func CreateClusterRoleBindingGlooResourceReaderBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-reader-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-resource-reader-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooResourceReaderBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingSettingsUserBindingDefault creates the ClusterRoleBinding resource with name settings-user-binding-default.
func CreateClusterRoleBindingSettingsUserBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "settings-user-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",
					"namespace": "nukleros-gateway-system",
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gateway",
					"namespace": "nukleros-gateway-system",
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "settings-user-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingSettingsUserBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooResourceMutatorBindingDefault creates the ClusterRoleBinding resource with name gloo-resource-mutator-binding-default.
func CreateClusterRoleBindingGlooResourceMutatorBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-mutator-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gateway",
					"namespace": "nukleros-gateway-system",
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-resource-mutator-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooResourceMutatorBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGatewayResourceReaderBindingDefault creates the ClusterRoleBinding resource with name gateway-resource-reader-binding-default.
func CreateClusterRoleBindingGatewayResourceReaderBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gateway-resource-reader-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gateway",
					"namespace": "nukleros-gateway-system",
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo", //  used to support gloo/gateway running in same pod
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gateway-resource-reader-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGatewayResourceReaderBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooGraphqlapiMutatorBindingDefault creates the ClusterRoleBinding resource with name gloo-graphqlapi-mutator-binding-default.
func CreateClusterRoleBindingGlooGraphqlapiMutatorBindingDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-graphqlapi-mutator-binding-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-graphqlapi-mutator-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooGraphqlapiMutatorBindingDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=delete
// +kubebuilder:rbac:groups=gateway.solo.io,resources=*,verbs=list;delete

// CreateClusterRoleGlooResourceCleanupDefault creates the ClusterRole resource with name gloo-resource-cleanup-default.
func CreateClusterRoleGlooResourceCleanupDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-cleanup-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.solo.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"list",
						"delete",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooResourceCleanupDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gateway.solo.io,resources=*,verbs=get;list;update;patch
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list

// CreateClusterRoleGlooResourceMigrationDefault creates the ClusterRole resource with name gloo-resource-migration-default.
func CreateClusterRoleGlooResourceMigrationDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-migration-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.solo.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"update",
						"patch",
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
						"get",
						"list",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooResourceMigrationDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gateway.solo.io,resources=*,verbs=get;list;create;update;patch

// CreateClusterRoleGlooResourceRolloutDefault creates the ClusterRole resource with name gloo-resource-rollout-default.
func CreateClusterRoleGlooResourceRolloutDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-rollout-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.solo.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"create",
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooResourceRolloutDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;update

// CreateClusterRoleGlooGatewayVwcUpdateDefault creates the ClusterRole resource with name gloo-gateway-vwc-update-default.
func CreateClusterRoleGlooGatewayVwcUpdateDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-gateway-vwc-update-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"get",
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooGatewayVwcUpdateDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;get;update

// CreateClusterRoleGlooGatewaySecretCreateDefault creates the ClusterRole resource with name gloo-gateway-secret-create-default.
func CreateClusterRoleGlooGatewaySecretCreateDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-gateway-secret-create-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"create",
						"get",
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleGlooGatewaySecretCreateDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooResourceCleanupDefault creates the ClusterRoleBinding resource with name gloo-resource-cleanup-default.
func CreateClusterRoleBindingGlooResourceCleanupDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-cleanup-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-resource-cleanup-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo-resource-cleanup",
					"namespace": "nukleros-gateway-system",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooResourceCleanupDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooResourceMigrationDefault creates the ClusterRoleBinding resource with name gloo-resource-migration-default.
func CreateClusterRoleBindingGlooResourceMigrationDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-migration-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-resource-migration-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo-resource-migration",
					"namespace": "nukleros-gateway-system",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooResourceMigrationDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooResourceRolloutDefault creates the ClusterRoleBinding resource with name gloo-resource-rollout-default.
func CreateClusterRoleBindingGlooResourceRolloutDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-resource-rollout-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-resource-rollout-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo-resource-rollout",
					"namespace": "nukleros-gateway-system",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooResourceRolloutDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooGatewayVwcUpdateDefault creates the ClusterRoleBinding resource with name gloo-gateway-vwc-update-default.
func CreateClusterRoleBindingGlooGatewayVwcUpdateDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-gateway-vwc-update-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "certgen",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-gateway-vwc-update-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooGatewayVwcUpdateDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingGlooGatewaySecretCreateDefault creates the ClusterRoleBinding resource with name gloo-gateway-secret-create-default.
func CreateClusterRoleBindingGlooGatewaySecretCreateDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "gloo-gateway-secret-create-default",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "certgen",
					"namespace": "nukleros-gateway-system",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "gloo-gateway-secret-create-default",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingGlooGatewaySecretCreateDefault(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gloo.solo.io,resources=*,verbs=list;delete

// CreateRoleNuklerosGatewaySystemGlooResourceCleanup creates the Role resource with name gloo-resource-cleanup.
func CreateRoleNuklerosGatewaySystemGlooResourceCleanup(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Role",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "gloo-resource-cleanup",
				"namespace": "nukleros-gateway-system",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"list",
						"delete",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNuklerosGatewaySystemGlooResourceCleanup(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=gloo.solo.io,resources=*,verbs=get;list;update;patch

// CreateRoleNuklerosGatewaySystemGlooResourceMigration creates the Role resource with name gloo-resource-migration.
func CreateRoleNuklerosGatewaySystemGlooResourceMigration(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Role",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "gloo-resource-migration",
				"namespace": "nukleros-gateway-system",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNuklerosGatewaySystemGlooResourceMigration(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch
// +kubebuilder:rbac:groups=gloo.solo.io,resources=*,verbs=get;list;create;update;patch

// CreateRoleNuklerosGatewaySystemGlooResourceRollout creates the Role resource with name gloo-resource-rollout.
func CreateRoleNuklerosGatewaySystemGlooResourceRollout(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Role",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "gloo-resource-rollout",
				"namespace": "nukleros-gateway-system",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"deployments",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gloo.solo.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"create",
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNuklerosGatewaySystemGlooResourceRollout(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNuklerosGatewaySystemGlooResourceCleanup creates the RoleBinding resource with name gloo-resource-cleanup.
func CreateRoleBindingNuklerosGatewaySystemGlooResourceCleanup(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "RoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "gloo-resource-cleanup",
				"namespace": "nukleros-gateway-system",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "Role",
				"name":     "gloo-resource-cleanup",
				"apiGroup": "rbac.authorization.k8s.io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo-resource-cleanup",
					"namespace": "nukleros-gateway-system",
				},
			},
		},
	}

	return mutate.MutateRoleBindingNuklerosGatewaySystemGlooResourceCleanup(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNuklerosGatewaySystemGlooResourceMigration creates the RoleBinding resource with name gloo-resource-migration.
func CreateRoleBindingNuklerosGatewaySystemGlooResourceMigration(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "RoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "gloo-resource-migration",
				"namespace": "nukleros-gateway-system",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "Role",
				"name":     "gloo-resource-migration",
				"apiGroup": "rbac.authorization.k8s.io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo-resource-migration",
					"namespace": "nukleros-gateway-system",
				},
			},
		},
	}

	return mutate.MutateRoleBindingNuklerosGatewaySystemGlooResourceMigration(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNuklerosGatewaySystemGlooResourceRollout creates the RoleBinding resource with name gloo-resource-rollout.
func CreateRoleBindingNuklerosGatewaySystemGlooResourceRollout(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "RoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name":      "gloo-resource-rollout",
				"namespace": "nukleros-gateway-system",
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "rbac",
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "Role",
				"name":     "gloo-resource-rollout",
				"apiGroup": "rbac.authorization.k8s.io",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo-resource-rollout",
					"namespace": "nukleros-gateway-system",
				},
			},
		},
	}

	return mutate.MutateRoleBindingNuklerosGatewaySystemGlooResourceRollout(resourceObj, parent, collection, reconciler, req)
}
