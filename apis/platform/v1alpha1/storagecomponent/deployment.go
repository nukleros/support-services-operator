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

package storagecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/storagecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceSeaweedfsOperator creates the Deployment resource with name seaweedfs-operator.
func CreateDeploymentNamespaceSeaweedfsOperator(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "seaweedfs-operator",
					"app.kubernetes.io/instance":    "seaweedfs-operator",
					"app":                           "seaweedfs-operator",
					"platform.nukleros.io/category": "storage",
					"platform.nukleros.io/project":  "seaweedfs-operator",
					"app.kubernetes.io/component":   "seaweedfs-operator",
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: seaweedfsOperator.replicas
				//  Number of replicas to use for the seaweedfs-operator deployment.
				"replicas": parent.Spec.SeaweedfsOperator.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":      "seaweedfs-operator",
						"app.kubernetes.io/instance":  "seaweedfs-operator",
						"app.kubernetes.io/component": "seaweedfs-operator",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":        "seaweedfs-operator",
							"app.kubernetes.io/instance":    "seaweedfs-operator",
							"platform.nukleros.io/category": "storage",
							"platform.nukleros.io/project":  "seaweedfs-operator",
							"app.kubernetes.io/component":   "seaweedfs-operator",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "seaweedfs-operator",
						"securityContext": map[string]interface{}{
							"fsGroup":      65532,
							"runAsNonRoot": true,
							"runAsUser":    65532,
							"seccompProfile": map[string]interface{}{
								"type": "RuntimeDefault",
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "seaweedfs-operator",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"runAsNonRoot":           true,
								},
								// controlled by field: seaweedfsOperator.image
								// controlled by field: seaweedfsOperator.version
								//  Image repo and name to use for seaweedfs-operator.
								//  Version of seaweedfs-operator to use.
								"image":           "" + parent.Spec.SeaweedfsOperator.Image + ":" + parent.Spec.SeaweedfsOperator.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"command": []interface{}{
									"/manager",
								},
								"args": []interface{}{
									"--leader-elect",
									"--metrics-bind-address=:8080",
									"--health-probe-bind-address=:8081",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"name":          "http",
										"containerPort": 8080,
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"name":          "https",
										"containerPort": 9443,
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"name":          "health",
										"containerPort": 8081,
										"protocol":      "TCP",
									},
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/healthz",
										"port": "health",
									},
									"initialDelaySeconds": 15,
									"periodSeconds":       20,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/readyz",
										"port": "health",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "50m",
										"memory": "64Mi",
									},
									"limits": map[string]interface{}{
										"memory": "256Mi",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/tmp/k8s-webhook-server/serving-certs",
										"name":      "cert",
										"readOnly":  true,
									},
								},
							},
						},
						"terminationGracePeriodSeconds": 10,
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "cert",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "seaweedfs-operator-webhook-server-cert",
									"items": []interface{}{
										map[string]interface{}{
											"key":  "tls.crt",
											"path": "tls.crt",
										},
										map[string]interface{}{
											"key":  "tls.key",
											"path": "tls.key",
										},
									},
								},
							},
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"weight": 100,
										"podAffinityTerm": map[string]interface{}{
											"topologyKey": "kubernetes.io/hostname",
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"seaweedfs-operator",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceSeaweedfsOperator(resourceObj, parent, collection, reconciler, req)
}
