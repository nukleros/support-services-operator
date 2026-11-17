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

// +kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=get;list;watch;create;update;patch;delete

// CreateAPIServiceV1beta1ExternalMetricsK8sIo creates the APIService resource with name v1beta1.external.metrics.k8s.io.
func CreateAPIServiceV1beta1ExternalMetricsK8sIo(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiregistration.k8s.io/v1",
			"kind":       "APIService",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "v1beta1.external.metrics.k8s.io",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "v1beta1.external.metrics.k8s.io",
			},
			"spec": map[string]interface{}{
				"service": map[string]interface{}{
					"name":      "keda-operator-metrics-apiserver",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"port":      443,
				},
				"group":                "external.metrics.k8s.io",
				"version":              "v1beta1",
				"groupPriorityMinimum": 100,
				"versionPriority":      100,
			},
		},
	}

	return mutate.MutateAPIServiceV1beta1ExternalMetricsK8sIo(resourceObj, parent, collection, reconciler, req)
}
