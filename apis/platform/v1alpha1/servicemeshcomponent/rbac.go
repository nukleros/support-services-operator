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

package servicemeshcomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/servicemeshcomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceIstioReaderServiceAccount creates the ServiceAccount resource with name istio-reader-service-account.
func CreateServiceAccountNamespaceIstioReaderServiceAccount(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "istio-reader-service-account",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "istio-reader",
					"app.kubernetes.io/name":        "istio-reader",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceIstioReaderServiceAccount(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceIstioCni creates the ServiceAccount resource with name istio-cni.
func CreateServiceAccountNamespaceIstioCni(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "istio-cni",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":          "istio-cni",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceIstioCni(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceIstiod creates the ServiceAccount resource with name istiod.
func CreateServiceAccountNamespaceIstiod(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "istiod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceIstiod(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceZtunnel creates the ServiceAccount resource with name ztunnel.
func CreateServiceAccountNamespaceZtunnel(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "ztunnel",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "ztunnel",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
				"annotations": map[string]interface{}{},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceZtunnel(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch

// CreateClusterRoleIstioCni creates the ClusterRole resource with name istio-cni.
func CreateClusterRoleIstioCni(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "istio-cni",
				"labels": map[string]interface{}{
					"app":          "istio-cni",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
						"nodes",
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

	return mutate.MutateClusterRoleIstioCni(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=watch;get;list

// CreateClusterRoleIstioCniRepairRole creates the ClusterRole resource with name istio-cni-repair-role.
func CreateClusterRoleIstioCniRepairRole(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "istio-cni-repair-role",
				"labels": map[string]interface{}{
					"app":          "istio-cni",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"rules": []interface{}{
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
						"pods",
					},
					"verbs": []interface{}{
						"watch",
						"get",
						"list",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIstioCniRepairRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=patch;update
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get

// CreateClusterRoleIstioCniAmbient creates the ClusterRole resource with name istio-cni-ambient.
func CreateClusterRoleIstioCniAmbient(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "istio-cni-ambient",
				"labels": map[string]interface{}{
					"app":          "istio-cni",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods/status",
					},
					"verbs": []interface{}{
						"patch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"daemonsets",
					},
					"resourceNames": []interface{}{
						"istio-cni-node",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIstioCniAmbient(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=config.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=security.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=networking.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=authentication.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=rbac.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=telemetry.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=extensions.istio.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=networking.istio.io,resources=workloadentries,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=networking.istio.io,resources=workloadentries/status,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=networking.istio.io,resources=serviceentries/status,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=security.istio.io,resources=authorizationpolicies/status,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=core,resources=services/status,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch
// +kubebuilder:rbac:groups=discovery.k8s.io,resources=endpointslices,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingressclasses,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=*
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=create;get;list;watch;update
// +kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenreviews,verbs=create
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=gateway.networking.x-k8s.io,resources=*,verbs=get;watch;list
// +kubebuilder:rbac:groups=gateway.networking.x-k8s.io,resources=xbackendtrafficpolicies/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=backendtlspolicies/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gatewayclasses/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gateways/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=grpcroutes/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=httproutes/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=referencegrants/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=tcproutes/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=tlsroutes/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=udproutes/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=listenersets/status,verbs=update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gatewayclasses,verbs=create;update;patch;delete
// +kubebuilder:rbac:groups=inference.networking.k8s.io,resources=inferencepools,verbs=get;watch;list
// +kubebuilder:rbac:groups=inference.networking.k8s.io,resources=inferencepools/status,verbs=update;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;watch;list
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceexports,verbs=get;watch;list;create;delete
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceimports,verbs=get;watch;list

// CreateClusterRoleIstiodClusterroleNamespace creates the ClusterRole resource with name istiod-clusterrole- + parent.spec.namespace.
func CreateClusterRoleIstiodClusterroleNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-clusterrole-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"mutatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"config.istio.io",
						"security.istio.io",
						"networking.istio.io",
						"authentication.istio.io",
						"rbac.istio.io",
						"telemetry.istio.io",
						"extensions.istio.io",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
					"resources": []interface{}{
						"*",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.istio.io",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"workloadentries",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.istio.io",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"workloadentries/status",
						"serviceentries/status",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"security.istio.io",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"authorizationpolicies/status",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"services/status",
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
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
						"nodes",
						"services",
						"namespaces",
						"endpoints",
					},
					"verbs": []interface{}{
						"get",
						"list",
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
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
						"ingressclasses",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses/status",
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
						"create",
						"get",
						"list",
						"watch",
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authentication.k8s.io",
					},
					"resources": []interface{}{
						"tokenreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authorization.k8s.io",
					},
					"resources": []interface{}{
						"subjectaccessreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
						"gateway.networking.x-k8s.io",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.x-k8s.io",
					},
					"resources": []interface{}{
						"xbackendtrafficpolicies/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"backendtlspolicies/status",
						"gatewayclasses/status",
						"gateways/status",
						"grpcroutes/status",
						"httproutes/status",
						"referencegrants/status",
						"tcproutes/status",
						"tlsroutes/status",
						"udproutes/status",
						"listenersets/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"gatewayclasses",
					},
					"verbs": []interface{}{
						"create",
						"update",
						"patch",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"inference.networking.k8s.io",
					},
					"resources": []interface{}{
						"inferencepools",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"inference.networking.k8s.io",
					},
					"resources": []interface{}{
						"inferencepools/status",
					},
					"verbs": []interface{}{
						"update",
						"patch",
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
						"get",
						"watch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"multicluster.x-k8s.io",
					},
					"resources": []interface{}{
						"serviceexports",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"create",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"multicluster.x-k8s.io",
					},
					"resources": []interface{}{
						"serviceimports",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIstiodClusterroleNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;watch;list;update;patch;create;delete

// CreateClusterRoleIstiodGatewayControllerNamespace creates the ClusterRole resource with name istiod-gateway-controller- + parent.spec.namespace.
func CreateClusterRoleIstiodGatewayControllerNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-gateway-controller-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"deployments",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"autoscaling",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"horizontalpodautoscalers",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"policy",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"poddisruptionbudgets",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"services",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
						"update",
						"patch",
						"create",
						"delete",
					},
					"resources": []interface{}{
						"serviceaccounts",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIstiodGatewayControllerNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=security.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=authentication.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=rbac.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=telemetry.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=extensions.istio.io,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=replicationcontrollers,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.istio.io,resources=workloadentries,verbs=get;watch;list
// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=gateways,verbs=get;watch;list
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gateways,verbs=get;watch;list
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch
// +kubebuilder:rbac:groups=discovery.k8s.io,resources=endpointslices,verbs=get;list;watch
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceexports,verbs=get;list;watch;create;delete
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceimports,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=replicasets,verbs=get;list;watch
// +kubebuilder:rbac:groups=authentication.k8s.io,resources=tokenreviews,verbs=create
// +kubebuilder:rbac:groups=authorization.k8s.io,resources=subjectaccessreviews,verbs=create

// CreateClusterRoleIstioReaderClusterroleNamespace creates the ClusterRole resource with name istio-reader-clusterrole- + parent.spec.namespace.
func CreateClusterRoleIstioReaderClusterroleNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istio-reader-clusterrole-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istio-reader",
					"app.kubernetes.io/name":        "istio-reader",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"config.istio.io",
						"security.istio.io",
						"networking.istio.io",
						"authentication.istio.io",
						"rbac.istio.io",
						"telemetry.istio.io",
						"extensions.istio.io",
					},
					"resources": []interface{}{
						"*",
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
						"endpoints",
						"pods",
						"services",
						"nodes",
						"replicationcontrollers",
						"namespaces",
						"secrets",
						"configmaps",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.istio.io",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
					"resources": []interface{}{
						"workloadentries",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.x-k8s.io",
						"gateway.networking.k8s.io",
					},
					"resources": []interface{}{
						"gateways",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
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
						"multicluster.x-k8s.io",
					},
					"resources": []interface{}{
						"serviceexports",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"create",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"multicluster.x-k8s.io",
					},
					"resources": []interface{}{
						"serviceimports",
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
						"replicasets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authentication.k8s.io",
					},
					"resources": []interface{}{
						"tokenreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"authorization.k8s.io",
					},
					"resources": []interface{}{
						"subjectaccessreviews",
					},
					"verbs": []interface{}{
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIstioReaderClusterroleNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioCni creates the ClusterRoleBinding resource with name istio-cni.
func CreateClusterRoleBindingIstioCni(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "istio-cni",
				"labels": map[string]interface{}{
					"app":          "istio-cni",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "istio-cni",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istio-cni",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioCni(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioCniRepairRolebinding creates the ClusterRoleBinding resource with name istio-cni-repair-rolebinding.
func CreateClusterRoleBindingIstioCniRepairRolebinding(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "istio-cni-repair-rolebinding",
				"labels": map[string]interface{}{
					"k8s-app":      "istio-cni-repair",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istio-cni",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "istio-cni-repair-role",
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioCniRepairRolebinding(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioCniAmbient creates the ClusterRoleBinding resource with name istio-cni-ambient.
func CreateClusterRoleBindingIstioCniAmbient(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "istio-cni-ambient",
				"labels": map[string]interface{}{
					"k8s-app":      "istio-cni-repair",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istio-cni",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "istio-cni-ambient",
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioCniAmbient(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstiodClusterroleNamespace creates the ClusterRoleBinding resource with name istiod-clusterrole- + parent.spec.namespace.
func CreateClusterRoleBindingIstiodClusterroleNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-clusterrole-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				// controlled by field: namespace
				"name": "istiod-clusterrole-" + parent.Spec.Namespace + "",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istiod",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstiodClusterroleNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstiodGatewayControllerNamespace creates the ClusterRoleBinding resource with name istiod-gateway-controller- + parent.spec.namespace.
func CreateClusterRoleBindingIstiodGatewayControllerNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-gateway-controller-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				// controlled by field: namespace
				"name": "istiod-gateway-controller-" + parent.Spec.Namespace + "",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istiod",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstiodGatewayControllerNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioReaderClusterroleNamespace creates the ClusterRoleBinding resource with name istio-reader-clusterrole- + parent.spec.namespace.
func CreateClusterRoleBindingIstioReaderClusterroleNamespace(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istio-reader-clusterrole-" + parent.Spec.Namespace + "",
				"labels": map[string]interface{}{
					"app":                           "istio-reader",
					"app.kubernetes.io/name":        "istio-reader",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				// controlled by field: namespace
				"name": "istio-reader-clusterrole-" + parent.Spec.Namespace + "",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istio-reader-service-account",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioReaderClusterroleNamespace(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.istio.io,resources=gateways,verbs=create
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;get;watch;list;update;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;update;patch;create

// CreateRoleNamespaceIstiod creates the Role resource with name istiod.
func CreateRoleNamespaceIstiod(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "istiod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.istio.io",
					},
					"verbs": []interface{}{
						"create",
					},
					"resources": []interface{}{
						"gateways",
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
						"get",
						"watch",
						"list",
						"update",
						"delete",
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
						"delete",
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
						"get",
						"update",
						"patch",
						"create",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceIstiod(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceIstiod creates the RoleBinding resource with name istiod.
func CreateRoleBindingNamespaceIstiod(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "istiod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "istiod",
					"app.kubernetes.io/name":        "istiod",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "istiod",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istiod",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceIstiod(resourceObj, parent, collection, reconciler, req)
}
