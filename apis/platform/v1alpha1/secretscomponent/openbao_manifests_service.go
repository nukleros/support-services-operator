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

package secretscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/secretscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceOpenbaoAgentInjectorSvc creates the Service resource with name openbao-agent-injector-svc.
func CreateServiceNamespaceOpenbaoAgentInjectorSvc(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Openbao.Injector.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=openbao.injector.include,value=true,include
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "openbao-agent-injector-svc",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao-agent-injector",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "https",
						"port":       443,
						"targetPort": 8080,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "openbao-agent-injector",
					"app.kubernetes.io/instance": "openbao",
					"component":                  "webhook",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceOpenbaoAgentInjectorSvc(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceOpenbaoActive creates the Service resource with name openbao-active.
func CreateServiceNamespaceOpenbaoActive(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "openbao-active",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao",
					"app.kubernetes.io/instance":    "openbao",
					"openbao-active":                "true",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				"publishNotReadyAddresses": true,
				"ports": []interface{}{
					map[string]interface{}{
						"name":        "http",
						"port":        8200,
						"targetPort":  8200,
						"appProtocol": "HTTP",
					},
					map[string]interface{}{
						"name":       "https-internal",
						"port":       8201,
						"targetPort": 8201,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "openbao",
					"app.kubernetes.io/instance": "openbao",
					"component":                  "server",
					"openbao-active":             "true",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceOpenbaoActive(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceOpenbaoStandby creates the Service resource with name openbao-standby.
func CreateServiceNamespaceOpenbaoStandby(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "openbao-standby",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				"publishNotReadyAddresses": true,
				"ports": []interface{}{
					map[string]interface{}{
						"name":        "http",
						"port":        8200,
						"targetPort":  8200,
						"appProtocol": "HTTP",
					},
					map[string]interface{}{
						"name":       "https-internal",
						"port":       8201,
						"targetPort": 8201,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "openbao",
					"app.kubernetes.io/instance": "openbao",
					"component":                  "server",
					"openbao-active":             "false",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceOpenbaoStandby(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceOpenbaoInternal creates the Service resource with name openbao-internal.
func CreateServiceNamespaceOpenbaoInternal(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "openbao-internal",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao",
					"app.kubernetes.io/instance":    "openbao",
					"openbao-internal":              "true",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				"clusterIP":                "None",
				"publishNotReadyAddresses": true,
				"ports": []interface{}{
					map[string]interface{}{
						"name":        "http",
						"port":        8200,
						"targetPort":  8200,
						"appProtocol": "HTTP",
					},
					map[string]interface{}{
						"name":       "https-internal",
						"port":       8201,
						"targetPort": 8201,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "openbao",
					"app.kubernetes.io/instance": "openbao",
					"component":                  "server",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceOpenbaoInternal(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceOpenbao creates the Service resource with name openbao.
func CreateServiceNamespaceOpenbao(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "openbao",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				"publishNotReadyAddresses": true,
				"ports": []interface{}{
					map[string]interface{}{
						"name":        "http",
						"port":        8200,
						"targetPort":  8200,
						"appProtocol": "HTTP",
					},
					map[string]interface{}{
						"name":       "https-internal",
						"port":       8201,
						"targetPort": 8201,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "openbao",
					"app.kubernetes.io/instance": "openbao",
					"component":                  "server",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceOpenbao(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceOpenbaoUi creates the Service resource with name openbao-ui.
func CreateServiceNamespaceOpenbaoUi(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "openbao-ui",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao-ui",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "openbao",
					"app.kubernetes.io/instance": "openbao",
					"component":                  "server",
				},
				"publishNotReadyAddresses": true,
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "http",
						"port":       8200,
						"targetPort": 8200,
					},
				},
				"type": "ClusterIP",
			},
		},
	}

	return mutate.MutateServiceNamespaceOpenbaoUi(resourceObj, parent, collection, reconciler, req)
}
