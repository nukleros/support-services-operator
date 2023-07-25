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

package glooedge

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/gateway/v1alpha1/glooedge/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceGloo creates the Deployment resource with name gloo.
func CreateDeploymentNamespaceGloo(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "gloo",
				},
				"name":      "gloo",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"gloo": "gloo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"gloo": "gloo",
						},
						"annotations": map[string]interface{}{
							"prometheus.io/path":         "/metrics",
							"prometheus.io/port":         "9091",
							"prometheus.io/scrape":       "true",
							"gloo.solo.io/oss-image-tag": "1.14.9",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "gloo",
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "labels-volume",
								"downwardAPI": map[string]interface{}{
									"items": []interface{}{
										map[string]interface{}{
											"path": "labels",
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.labels",
											},
										},
									},
								},
							},
							map[string]interface{}{
								"name": "validation-certs",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "certificate-authority",
								},
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"image":           "quay.io/solo-io/gloo:1.14.9",
								"imagePullPolicy": "IfNotPresent",
								"name":            "gloo",
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "500m",
										"memory": "256Mi",
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"runAsNonRoot":           true,
									"runAsUser":              10101,
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 9977,
										"name":          "grpc-xds",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9976,
										"name":          "rest-xds",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9988,
										"name":          "grpc-validation",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9966,
										"name":          "grpc-proxydebug",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9979,
										"name":          "wasm-cache",
										"protocol":      "TCP",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/etc/gateway/validation-certs",
										"name":      "validation-certs",
									},
									map[string]interface{}{
										"name":      "labels-volume",
										"mountPath": "/etc/gloo",
										"readOnly":  true,
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "START_STATS_SERVER",
										"value": "true",
									},
									map[string]interface{}{
										"name":  "VALIDATION_MUST_START",
										"value": "true",
									},
								},
								"readinessProbe": map[string]interface{}{
									"tcpSocket": map[string]interface{}{
										"port": 9977,
									},
									"initialDelaySeconds": 3,
									"periodSeconds":       10,
									"failureThreshold":    3,
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceGloo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceDiscovery creates the Deployment resource with name discovery.
func CreateDeploymentNamespaceDiscovery(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "discovery",
				},
				"name":      "discovery",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"gloo": "discovery",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"gloo": "discovery",
						},
						"annotations": map[string]interface{}{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/port":   "9091",
							"prometheus.io/scrape": "true",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "discovery",
						"containers": []interface{}{
							map[string]interface{}{
								"image":           "quay.io/solo-io/discovery:1.14.9",
								"imagePullPolicy": "IfNotPresent",
								"name":            "discovery",
								// container security context
								"securityContext": map[string]interface{}{
									"readOnlyRootFilesystem":   true,
									"allowPrivilegeEscalation": false,
									"runAsNonRoot":             true,
									"runAsUser":                10101,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "START_STATS_SERVER",
										"value": "true",
									},
								},
							},
						},
						// Pod security context
						"securityContext": map[string]interface{}{
							"fsGroup": 10101,
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceDiscovery(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceGatewayProxy creates the Deployment resource with name gateway-proxy.
func CreateDeploymentNamespaceGatewayProxy(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":              "gloo",
					"gloo":             "gateway-proxy",
					"gateway-proxy-id": "gateway-proxy",
				},
				"name":      "gateway-proxy",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": 1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"gloo":             "gateway-proxy",
						"gateway-proxy-id": "gateway-proxy",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"gloo":             "gateway-proxy",
							"gateway-proxy-id": "gateway-proxy",
							"gateway-proxy":    "live",
						},
						"annotations": map[string]interface{}{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/port":   "8081",
							"prometheus.io/scrape": "true",
						},
					},
					"spec": map[string]interface{}{
						"securityContext": map[string]interface{}{
							"fsGroup":   10101,
							"runAsUser": 10101,
						},
						"serviceAccountName": "gateway-proxy",
						"containers": []interface{}{
							map[string]interface{}{
								"args": []interface{}{
									"--disable-hot-restart",
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name": "POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name":  "DISABLE_CORE_DUMPS",
										"value": "false",
									},
								},
								"image":           "quay.io/solo-io/gloo-envoy-wrapper:1.14.9",
								"imagePullPolicy": "IfNotPresent",
								"name":            "gateway-proxy",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"runAsNonRoot":           true,
									"runAsUser":              10101,
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"name":          "http",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 8443,
										"name":          "https",
										"protocol":      "TCP",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/etc/envoy",
										"name":      "envoy-config",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"configMap": map[string]interface{}{
									"name": "gateway-proxy-envoy-config",
								},
								"name": "envoy-config",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceGatewayProxy(resourceObj, parent, collection, reconciler, req)
}
