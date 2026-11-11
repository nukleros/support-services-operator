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

// CreateServiceAccountNamespaceIstioReaderServiceAccountOpenshift creates the ServiceAccount resource with name istio-reader-service-account-openshift.
func CreateServiceAccountNamespaceIstioReaderServiceAccountOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "istio-reader-service-account-openshift",
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

	return mutate.MutateServiceAccountNamespaceIstioReaderServiceAccountOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceIstioCniOpenshift creates the ServiceAccount resource with name istio-cni-openshift.
func CreateServiceAccountNamespaceIstioCniOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "istio-cni-openshift",
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

	return mutate.MutateServiceAccountNamespaceIstioCniOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceIstiodOpenshift creates the ServiceAccount resource with name istiod-openshift.
func CreateServiceAccountNamespaceIstiodOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "istiod-openshift",
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

	return mutate.MutateServiceAccountNamespaceIstiodOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceZtunnelOpenshift creates the ServiceAccount resource with name ztunnel-openshift.
func CreateServiceAccountNamespaceZtunnelOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "ztunnel-openshift",
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

	return mutate.MutateServiceAccountNamespaceZtunnelOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=security.openshift.io,resources=securitycontextconstraints,verbs=use

// CreateClusterRoleIstioCniOpenshift creates the ClusterRole resource with name istio-cni-openshift.
func CreateClusterRoleIstioCniOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "istio-cni-openshift",
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
				map[string]interface{}{
					"apiGroups": []interface{}{
						"security.openshift.io",
					},
					"resources": []interface{}{
						"securitycontextconstraints",
					},
					"resourceNames": []interface{}{
						"privileged",
					},
					"verbs": []interface{}{
						"use",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleIstioCniOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=watch;get;list

// CreateClusterRoleIstioCniRepairRoleOpenshift creates the ClusterRole resource with name istio-cni-repair-role-openshift.
func CreateClusterRoleIstioCniRepairRoleOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "istio-cni-repair-role-openshift",
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

	return mutate.MutateClusterRoleIstioCniRepairRoleOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=patch;update
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get

// CreateClusterRoleIstioCniAmbientOpenshift creates the ClusterRole resource with name istio-cni-ambient-openshift.
func CreateClusterRoleIstioCniAmbientOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "istio-cni-ambient-openshift",
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

	return mutate.MutateClusterRoleIstioCniAmbientOpenshift(resourceObj, parent, collection, reconciler, req)
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

// CreateClusterRoleIstiodClusterroleNamespaceOpenshift creates the ClusterRole resource with name istiod-clusterrole- + parent.spec.namespace + -openshift.
func CreateClusterRoleIstiodClusterroleNamespaceOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-clusterrole-" + parent.Spec.Namespace + "-openshift",
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

	return mutate.MutateClusterRoleIstiodClusterroleNamespaceOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;watch;list;update;patch;create;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;watch;list;update;patch;create;delete

// CreateClusterRoleIstiodGatewayControllerNamespaceOpenshift creates the ClusterRole resource with name istiod-gateway-controller- + parent.spec.namespace + -openshift.
func CreateClusterRoleIstiodGatewayControllerNamespaceOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-gateway-controller-" + parent.Spec.Namespace + "-openshift",
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

	return mutate.MutateClusterRoleIstiodGatewayControllerNamespaceOpenshift(resourceObj, parent, collection, reconciler, req)
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

// CreateClusterRoleIstioReaderClusterroleNamespaceOpenshift creates the ClusterRole resource with name istio-reader-clusterrole- + parent.spec.namespace + -openshift.
func CreateClusterRoleIstioReaderClusterroleNamespaceOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istio-reader-clusterrole-" + parent.Spec.Namespace + "-openshift",
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

	return mutate.MutateClusterRoleIstioReaderClusterroleNamespaceOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=security.openshift.io,resources=securitycontextconstraints,verbs=use

// CreateClusterRoleZtunnelOpenshift creates the ClusterRole resource with name ztunnel-openshift.
func CreateClusterRoleZtunnelOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "ztunnel-openshift",
				"labels": map[string]interface{}{
					"app": "ztunnel",
					"install.operator.istio.io/owning-resource": "unknown",
					"app.kubernetes.io/name":                    "ztunnel",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
				"annotations": map[string]interface{}{},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"security.openshift.io",
					},
					"resources": []interface{}{
						"securitycontextconstraints",
					},
					"resourceNames": []interface{}{
						"privileged",
					},
					"verbs": []interface{}{
						"use",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleZtunnelOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioCniOpenshift creates the ClusterRoleBinding resource with name istio-cni-openshift.
func CreateClusterRoleBindingIstioCniOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "istio-cni-openshift",
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
				"name":     "istio-cni-openshift",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istio-cni-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioCniOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioCniRepairRolebindingOpenshift creates the ClusterRoleBinding resource with name istio-cni-repair-rolebinding-openshift.
func CreateClusterRoleBindingIstioCniRepairRolebindingOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "istio-cni-repair-rolebinding-openshift",
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
					"name":      "istio-cni-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "istio-cni-repair-role-openshift",
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioCniRepairRolebindingOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioCniAmbientOpenshift creates the ClusterRoleBinding resource with name istio-cni-ambient-openshift.
func CreateClusterRoleBindingIstioCniAmbientOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "istio-cni-ambient-openshift",
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
					"name":      "istio-cni-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "istio-cni-ambient-openshift",
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioCniAmbientOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstiodClusterroleNamespaceOpenshift creates the ClusterRoleBinding resource with name istiod-clusterrole- + parent.spec.namespace + -openshift.
func CreateClusterRoleBindingIstiodClusterroleNamespaceOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-clusterrole-" + parent.Spec.Namespace + "-openshift",
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
				"name": "istiod-clusterrole-" + parent.Spec.Namespace + "-openshift",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istiod-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstiodClusterroleNamespaceOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstiodGatewayControllerNamespaceOpenshift creates the ClusterRoleBinding resource with name istiod-gateway-controller- + parent.spec.namespace + -openshift.
func CreateClusterRoleBindingIstiodGatewayControllerNamespaceOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istiod-gateway-controller-" + parent.Spec.Namespace + "-openshift",
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
				"name": "istiod-gateway-controller-" + parent.Spec.Namespace + "-openshift",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istiod-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstiodGatewayControllerNamespaceOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingIstioReaderClusterroleNamespaceOpenshift creates the ClusterRoleBinding resource with name istio-reader-clusterrole- + parent.spec.namespace + -openshift.
func CreateClusterRoleBindingIstioReaderClusterroleNamespaceOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				"name": "istio-reader-clusterrole-" + parent.Spec.Namespace + "-openshift",
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
				"name": "istio-reader-clusterrole-" + parent.Spec.Namespace + "-openshift",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istio-reader-service-account-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingIstioReaderClusterroleNamespaceOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingZtunnelOpenshift creates the ClusterRoleBinding resource with name ztunnel-openshift.
func CreateClusterRoleBindingZtunnelOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "ztunnel-openshift",
				"labels": map[string]interface{}{
					"app": "ztunnel",
					"install.operator.istio.io/owning-resource": "unknown",
					"app.kubernetes.io/name":                    "ztunnel",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
				"annotations": map[string]interface{}{},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "ztunnel-openshift",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "ztunnel-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingZtunnelOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.istio.io,resources=gateways,verbs=create
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=create;get;watch;list;update;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;update;patch;create

// CreateRoleNamespaceIstiodOpenshift creates the Role resource with name istiod-openshift.
func CreateRoleNamespaceIstiodOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "istiod-openshift",
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

	return mutate.MutateRoleNamespaceIstiodOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceIstiodOpenshift creates the RoleBinding resource with name istiod-openshift.
func CreateRoleBindingNamespaceIstiodOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "istiod-openshift",
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
				"name":     "istiod-openshift",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "istiod-openshift",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceIstiodOpenshift(resourceObj, parent, collection, reconciler, req)
}
