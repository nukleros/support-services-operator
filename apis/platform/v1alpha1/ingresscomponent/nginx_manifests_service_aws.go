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

package ingresscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/ingresscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceNginxIngressAws creates the Service resource with name nginx-ingress-aws.
func CreateServiceNamespaceNginxIngressAws(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "nginx-ingress-aws",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"annotations": map[string]interface{}{
					"service.beta.kubernetes.io/aws-load-balancer-backend-protocol": "tcp",
					"service.beta.kubernetes.io/aws-load-balancer-proxy-protocol":   "*",
				},
				"labels": map[string]interface{}{
					"platform.nukleros.io/group":   "ingress",
					"platform.nukleros.io/project": "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"type": "LoadBalancer",
				"ports": []interface{}{
					map[string]interface{}{
						"port":       80,
						"targetPort": 80,
						"protocol":   "TCP",
						"name":       "http",
					},
					map[string]interface{}{
						"port":       443,
						"targetPort": 443,
						"protocol":   "TCP",
						"name":       "https",
					},
				},
				"selector": map[string]interface{}{
					"app": "nginx-ingress",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceNginxIngressAws(resourceObj, parent, collection, reconciler, req)
}
