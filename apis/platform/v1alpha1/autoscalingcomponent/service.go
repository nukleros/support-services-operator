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

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceKedaOperator creates the Service resource with name keda-operator.
func CreateServiceNamespaceKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "8080",
					"prometheus.io/path":   "/metrics",
				},
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
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "metricsservice",
						"port":       9666,
						"targetPort": 9666,
					},
					map[string]interface{}{
						"name":       "metrics",
						"port":       8080,
						"targetPort": 8080,
					},
				},
				"selector": map[string]interface{}{
					"app": "keda-operator",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceKedaOperatorMetricsApiserver creates the Service resource with name keda-operator-metrics-apiserver.
func CreateServiceNamespaceKedaOperatorMetricsApiserver(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-metrics-apiserver",
					"app":                           "keda-operator-metrics-apiserver",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-operator-metrics-apiserver",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"annotations": map[string]interface{}{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "8080",
					"prometheus.io/path":   "/metrics",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"name":        "https",
						"port":        443,
						"targetPort":  6443,
						"protocol":    "TCP",
						"appProtocol": "https",
					},
					map[string]interface{}{
						"name":       "metrics",
						"port":       8080,
						"targetPort": 8080,
						"protocol":   "TCP",
					},
				},
				"selector": map[string]interface{}{
					"app": "keda-operator-metrics-apiserver",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceKedaOperatorMetricsApiserver(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceKedaAdmissionWebhooks creates the Service resource with name keda-admission-webhooks.
func CreateServiceNamespaceKedaAdmissionWebhooks(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "8080",
					"prometheus.io/path":   "/metrics",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-admission-webhooks",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name":      "keda-admission-webhooks",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":        "https",
						"port":        443,
						"protocol":    "TCP",
						"targetPort":  9443,
						"appProtocol": "https",
					},
					map[string]interface{}{
						"name":       "metrics",
						"port":       8080,
						"targetPort": 8080,
					},
				},
				"selector": map[string]interface{}{
					"app": "keda-admission-webhooks",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceKedaAdmissionWebhooks(resourceObj, parent, collection, reconciler, req)
}
