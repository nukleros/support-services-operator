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

// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;list;watch;create;update;patch;delete

// CreateHorizontalPodAutoscalerNamespaceIstiodGke creates the HorizontalPodAutoscaler resource with name istiod-gke.
func CreateHorizontalPodAutoscalerNamespaceIstiodGke(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "gke" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="gke",include
			"apiVersion": "autoscaling/v2",
			"kind":       "HorizontalPodAutoscaler",
			"metadata": map[string]interface{}{
				"name":      "istiod-gke",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":          "istiod",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Pilot",
					"app.kubernetes.io/name":                    "istiod",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: istio.maxReplicas
				"maxReplicas": parent.Spec.Istio.MaxReplicas,
				// controlled by field: istio.replicas
				"minReplicas": parent.Spec.Istio.Replicas,
				"scaleTargetRef": map[string]interface{}{
					"apiVersion": "apps/v1",
					"kind":       "Deployment",
					"name":       "istiod-gke",
				},
				"metrics": []interface{}{
					map[string]interface{}{
						"type": "Resource",
						"resource": map[string]interface{}{
							"name": "cpu",
							"target": map[string]interface{}{
								"type":               "Utilization",
								"averageUtilization": 80,
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateHorizontalPodAutoscalerNamespaceIstiodGke(resourceObj, parent, collection, reconciler, req)
}
