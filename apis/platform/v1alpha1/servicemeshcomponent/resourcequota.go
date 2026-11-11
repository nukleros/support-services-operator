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

// +kubebuilder:rbac:groups=core,resources=resourcequotas,verbs=get;list;watch;create;update;patch;delete

// CreateResourceQuotaNamespaceIstioCniResourceQuota creates the ResourceQuota resource with name istio-cni-resource-quota.
func CreateResourceQuotaNamespaceIstioCniResourceQuota(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ResourceQuota",
			"metadata": map[string]interface{}{
				"name":      "istio-cni-resource-quota",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "istio-cni",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
			},
			"spec": map[string]interface{}{
				"hard": map[string]interface{}{
					"pods": "5k",
				},
				"scopeSelector": map[string]interface{}{
					"matchExpressions": []interface{}{
						map[string]interface{}{
							"operator":  "In",
							"scopeName": "PriorityClass",
							"values": []interface{}{
								"system-node-critical",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateResourceQuotaNamespaceIstioCniResourceQuota(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=resourcequotas,verbs=get;list;watch;create;update;patch;delete

// CreateResourceQuotaNamespaceZtunnel creates the ResourceQuota resource with name ztunnel.
func CreateResourceQuotaNamespaceZtunnel(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ResourceQuota",
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
			},
			"spec": map[string]interface{}{
				"hard": map[string]interface{}{
					"pods": "5k",
				},
				"scopeSelector": map[string]interface{}{
					"matchExpressions": []interface{}{
						map[string]interface{}{
							"operator":  "In",
							"scopeName": "PriorityClass",
							"values": []interface{}{
								"system-node-critical",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateResourceQuotaNamespaceZtunnel(resourceObj, parent, collection, reconciler, req)
}
