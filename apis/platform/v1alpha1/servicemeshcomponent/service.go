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

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceIstiod creates the Service resource with name istiod.
func CreateServiceNamespaceIstiod(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "istiod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Pilot",
					"app":                                       "istiod",
					"istio":                                     "pilot",
					"app.kubernetes.io/name":                    "istiod",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port":     15010,
						"name":     "grpc-xds",
						"protocol": "TCP",
					},
					map[string]interface{}{
						"port":     15012,
						"name":     "https-dns",
						"protocol": "TCP",
					},
					map[string]interface{}{
						"port":       443,
						"name":       "https-webhook",
						"targetPort": 15017,
						"protocol":   "TCP",
					},
					map[string]interface{}{
						"port":     15014,
						"name":     "http-monitoring",
						"protocol": "TCP",
					},
				},
				"selector": map[string]interface{}{
					"app":   "istiod",
					"istio": "pilot",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceIstiod(resourceObj, parent, collection, reconciler, req)
}
