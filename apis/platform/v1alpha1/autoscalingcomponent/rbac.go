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

package autoscalingcomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/autoscalingcomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceKedaOperator creates the ServiceAccount resource with name keda-operator.
func CreateServiceAccountNamespaceKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"automountServiceAccountToken": true,
		},
	}

	return mutate.MutateServiceAccountNamespaceKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceKedaMetricsServer creates the ServiceAccount resource with name keda-metrics-server.
func CreateServiceAccountNamespaceKedaMetricsServer(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-metrics-server",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-metrics-server",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"automountServiceAccountToken": true,
		},
	}

	return mutate.MutateServiceAccountNamespaceKedaMetricsServer(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceKedaWebhook creates the ServiceAccount resource with name keda-webhook.
func CreateServiceAccountNamespaceKedaWebhook(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-webhook",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"automountServiceAccountToken": true,
		},
	}

	return mutate.MutateServiceAccountNamespaceKedaWebhook(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps/status,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=limitranges,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch
// +kubebuilder:rbac:groups=events.k8s.io,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=discovery.k8s.io,resources=endpointslices,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=list;watch
// +kubebuilder:rbac:groups=*,resources=*/scale,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=*,resources=*,verbs=get
// +kubebuilder:rbac:groups=apps,resources=deployments/scale,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=apps,resources=statefulsets/scale,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=replicasets,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch
// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=cloudeventsources,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=cloudeventsources/status,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=clustercloudeventsources,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=clustercloudeventsources/status,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledjobs,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledjobs/finalizers,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledjobs/status,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledobjects,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledobjects/finalizers,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledobjects/status,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=triggerauthentications,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=triggerauthentications/status,verbs=get;list;patch;update;watch

// CreateClusterRoleKedaOperator creates the ClusterRole resource with name keda-operator.
func CreateClusterRoleKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
						"configmaps/status",
						"limitranges",
						"pods",
						"services",
						"serviceaccounts",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"events.k8s.io",
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
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"*",
					},
					"resources": []interface{}{
						"*/scale",
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
						"*",
					},
					"resources": []interface{}{
						"*",
					},
					"verbs": []interface{}{
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"deployments/scale",
						"statefulsets/scale",
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
						"apps",
					},
					"resources": []interface{}{
						"deployments",
						"replicasets",
						"statefulsets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"autoscaling",
					},
					"resources": []interface{}{
						"horizontalpodautoscalers",
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
						"update",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"eventing.keda.sh",
					},
					"resources": []interface{}{
						"cloudeventsources",
						"cloudeventsources/status",
						"clustercloudeventsources",
						"clustercloudeventsources/status",
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
						"keda.sh",
					},
					"resources": []interface{}{
						"scaledjobs",
						"scaledjobs/finalizers",
						"scaledjobs/status",
						"scaledobjects",
						"scaledobjects/finalizers",
						"scaledobjects/status",
						"triggerauthentications",
						"triggerauthentications/status",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=keda.sh,resources=clustertriggerauthentications,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=keda.sh,resources=clustertriggerauthentications/status,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=cloudeventsources,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=cloudeventsources/status,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=clustercloudeventsources,verbs=get;list;patch;update;watch
// +kubebuilder:rbac:groups=eventing.keda.sh,resources=clustercloudeventsources/status,verbs=get;list;patch;update;watch

// CreateClusterRoleKedaOperatorMinimalClusterRole creates the ClusterRole resource with name keda-operator-minimal-cluster-role.
func CreateClusterRoleKedaOperatorMinimalClusterRole(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-minimal-cluster-role",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-minimal-cluster-role",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"keda.sh",
					},
					"resources": []interface{}{
						"clustertriggerauthentications",
						"clustertriggerauthentications/status",
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
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
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
						"apiregistration.k8s.io",
					},
					"resources": []interface{}{
						"apiservices",
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
						"eventing.keda.sh",
					},
					"resources": []interface{}{
						"cloudeventsources",
						"cloudeventsources/status",
						"clustercloudeventsources",
						"clustercloudeventsources/status",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"patch",
						"update",
						"watch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleKedaOperatorMinimalClusterRole(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=external.metrics.k8s.io,resources=externalmetrics,verbs=get

// CreateClusterRoleKedaOperatorExternalMetricsReader creates the ClusterRole resource with name keda-operator-external-metrics-reader.
func CreateClusterRoleKedaOperatorExternalMetricsReader(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-external-metrics-reader",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-external-metrics-reader",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"external.metrics.k8s.io",
					},
					"resources": []interface{}{
						"externalmetrics",
					},
					"verbs": []interface{}{
						"get",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleKedaOperatorExternalMetricsReader(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=list;watch
// +kubebuilder:rbac:groups=keda.sh,resources=scaledobjects,verbs=list;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=limitranges,verbs=list
// +kubebuilder:rbac:groups=events.k8s.io,resources=events,verbs=create;patch

// CreateClusterRoleKedaOperatorWebhook creates the ClusterRole resource with name keda-operator-webhook.
func CreateClusterRoleKedaOperatorWebhook(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-webhook",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-webhook",
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"autoscaling",
					},
					"resources": []interface{}{
						"horizontalpodautoscalers",
					},
					"verbs": []interface{}{
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"keda.sh",
					},
					"resources": []interface{}{
						"scaledobjects",
					},
					"verbs": []interface{}{
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"deployments",
						"statefulsets",
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
						"limitranges",
					},
					"verbs": []interface{}{
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"events.k8s.io",
					},
					"resources": []interface{}{
						"events",
					},
					"verbs": []interface{}{
						"create",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleKedaOperatorWebhook(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKedaOperator creates the ClusterRoleBinding resource with name keda-operator.
func CreateClusterRoleBindingKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "keda-operator",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "keda-operator",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKedaOperatorMinimal creates the ClusterRoleBinding resource with name keda-operator-minimal.
func CreateClusterRoleBindingKedaOperatorMinimal(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-minimal",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-minimal",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "keda-operator-minimal-cluster-role",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "keda-operator",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingKedaOperatorMinimal(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKedaOperatorSystemAuthDelegator creates the ClusterRoleBinding resource with name keda-operator-system-auth-delegator.
func CreateClusterRoleBindingKedaOperatorSystemAuthDelegator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-system-auth-delegator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-system-auth-delegator",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "system:auth-delegator",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "keda-metrics-server",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingKedaOperatorSystemAuthDelegator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKedaOperatorHpaControllerExternalMetrics creates the ClusterRoleBinding resource with name keda-operator-hpa-controller-external-metrics.
func CreateClusterRoleBindingKedaOperatorHpaControllerExternalMetrics(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-hpa-controller-external-metrics",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-hpa-controller-external-metrics",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "keda-operator-external-metrics-reader",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "horizontal-pod-autoscaler",
					"namespace": "kube-system",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingKedaOperatorHpaControllerExternalMetrics(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingKedaOperatorWebhook creates the ClusterRoleBinding resource with name keda-operator-webhook.
func CreateClusterRoleBindingKedaOperatorWebhook(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-webhook",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "keda-operator-webhook",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "keda-operator-webhook",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "keda-webhook",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingKedaOperatorWebhook(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;create;update

// CreateRoleNamespaceKedaOperatorCerts creates the Role resource with name keda-operator-certs.
func CreateRoleNamespaceKedaOperatorCerts(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-certs",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-operator-certs",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
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
						"secrets",
					},
					"verbs": []interface{}{
						"get",
					},
					"resourceNames": []interface{}{
						"kedaorg-certs",
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
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceKedaOperatorCerts(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceKedaOperatorCerts creates the RoleBinding resource with name keda-operator-certs.
func CreateRoleBindingNamespaceKedaOperatorCerts(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-certs",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-operator-certs",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "keda-operator-certs",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "keda-operator",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceKedaOperatorCerts(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingKubeSystemKedaOperatorAuthReader creates the RoleBinding resource with name keda-operator-auth-reader.
func CreateRoleBindingKubeSystemKedaOperatorAuthReader(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-auth-reader",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-operator-auth-reader",
				"namespace": "kube-system",
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "extension-apiserver-authentication-reader",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "keda-metrics-server",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingKubeSystemKedaOperatorAuthReader(resourceObj, parent, collection, reconciler, req)
}
