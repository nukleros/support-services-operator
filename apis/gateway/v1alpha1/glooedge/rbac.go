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

// CreateServiceAccountNamespaceCertgen creates the ServiceAccount resource with name certgen.
func CreateServiceAccountNamespaceCertgen(
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
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceCertgen(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceGloo creates the ServiceAccount resource with name gloo.
func CreateServiceAccountNamespaceGloo(
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
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceGloo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceDiscovery creates the ServiceAccount resource with name discovery.
func CreateServiceAccountNamespaceDiscovery(
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
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceDiscovery(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceGatewayProxy creates the ServiceAccount resource with name gateway-proxy.
func CreateServiceAccountNamespaceGatewayProxy(
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
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceGatewayProxy(resourceObj, parent, collection, reconciler, req)
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gateway",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "discovery",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "gloo",                //  used to support gloo/gateway running in same pod
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
