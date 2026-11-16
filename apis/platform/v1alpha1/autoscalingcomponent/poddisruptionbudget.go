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

// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=get;list;watch;create;update;patch;delete

// CreatePDBNamespaceKedaOperator creates the PodDisruptionBudget resource with name keda-operator.
func CreatePDBNamespaceKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "policy/v1",
			"kind":       "PodDisruptionBudget",
			"metadata": map[string]interface{}{
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"name":      "keda-operator",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"minAvailable": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-operator",
					},
				},
				"unhealthyPodEvictionPolicy": "AlwaysAllow",
			},
		},
	}

	return mutate.MutatePDBNamespaceKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=get;list;watch;create;update;patch;delete

// CreatePDBNamespaceKedaOperatorMetricsApiserver creates the PodDisruptionBudget resource with name keda-operator-metrics-apiserver.
func CreatePDBNamespaceKedaOperatorMetricsApiserver(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "policy/v1",
			"kind":       "PodDisruptionBudget",
			"metadata": map[string]interface{}{
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"name":      "keda-operator-metrics-apiserver",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator-metrics-apiserver",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"minAvailable": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-operator-metrics-apiserver",
					},
				},
				"unhealthyPodEvictionPolicy": "AlwaysAllow",
			},
		},
	}

	return mutate.MutatePDBNamespaceKedaOperatorMetricsApiserver(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=get;list;watch;create;update;patch;delete

// CreatePDBNamespaceKedaAdmissionWebhooks creates the PodDisruptionBudget resource with name keda-admission-webhooks.
func CreatePDBNamespaceKedaAdmissionWebhooks(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "policy/v1",
			"kind":       "PodDisruptionBudget",
			"metadata": map[string]interface{}{
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"name":      "keda-admission-webhooks",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-admission-webhooks",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"minAvailable": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-admission-webhooks",
					},
				},
				"unhealthyPodEvictionPolicy": "AlwaysAllow",
			},
		},
	}

	return mutate.MutatePDBNamespaceKedaAdmissionWebhooks(resourceObj, parent, collection, reconciler, req)
}
