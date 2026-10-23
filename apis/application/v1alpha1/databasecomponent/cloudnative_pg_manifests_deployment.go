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

package databasecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	applicationv1alpha1 "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/application/v1alpha1/databasecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceCnpgControllerManager creates the Deployment resource with name cnpg-controller-manager.
func CreateDeploymentNamespaceCnpgControllerManager(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Provider != "cloudnative-pg" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=provider,value=cloudnative-pg,include
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":           "cloudnative-pg",
					"application.nukleros.io/category": "database",
					"application.nukleros.io/project":  "cloudnative-pg",
					"app.kubernetes.io/component":      "database",
				},
				"name":      "cnpg-controller-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				// controlled by field: cloudNativePg.replicas
				//  Number of replicas to use for the cloudnative-pg controller deployment.
				"replicas": parent.Spec.CloudNativePg.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name": "cloudnative-pg",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":           "cloudnative-pg",
							"application.nukleros.io/category": "database",
							"application.nukleros.io/project":  "cloudnative-pg",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"args": []interface{}{
									"controller",
									"--leader-elect",
									"--max-concurrent-reconciles=10",
									"--config-map-name=cnpg-controller-manager-config",
									"--secret-name=cnpg-controller-manager-config",
									"--webhook-port=9443",
								},
								"command": []interface{}{
									"/manager",
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "OPERATOR_IMAGE_NAME",
										// controlled by field: cloudNativePg.image
										// controlled by field: cloudNativePg.version
										"value": "" + parent.Spec.CloudNativePg.Image + ":" + parent.Spec.CloudNativePg.Version + "",
									},
									map[string]interface{}{
										"name": "OPERATOR_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "MONITORING_QUERIES_CONFIGMAP",
										"value": "cnpg-default-monitoring",
									},
								},
								// controlled by field: cloudNativePg.image
								// controlled by field: cloudNativePg.version
								//  Image repo and name to use for cloudnative-pg.
								//  Version of cloudnative-pg to use.
								"image":           "" + parent.Spec.CloudNativePg.Image + ":" + parent.Spec.CloudNativePg.Version + "",
								"imagePullPolicy": "Always",
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/readyz",
										"port":   9443,
										"scheme": "HTTPS",
									},
								},
								"name": "manager",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"name":          "metrics",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9443,
										"name":          "webhook-server",
										"protocol":      "TCP",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/readyz",
										"port":   9443,
										"scheme": "HTTPS",
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "100Mi",
									},
									"limits": map[string]interface{}{
										"memory": "200Mi",
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
									"runAsGroup":             10001,
									"runAsUser":              10001,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
								"startupProbe": map[string]interface{}{
									"failureThreshold": 6,
									"httpGet": map[string]interface{}{
										"path":   "/readyz",
										"port":   9443,
										"scheme": "HTTPS",
									},
									"periodSeconds": 5,
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/controller",
										"name":      "scratch-data",
									},
									map[string]interface{}{
										"mountPath": "/run/secrets/cnpg.io/webhook",
										"name":      "webhook-certificates",
									},
								},
							},
						},
						"securityContext": map[string]interface{}{
							"runAsNonRoot": true,
							"seccompProfile": map[string]interface{}{
								"type": "RuntimeDefault",
							},
						},
						"serviceAccountName":            "cnpg-manager",
						"terminationGracePeriodSeconds": 10,
						"volumes": []interface{}{
							map[string]interface{}{
								"emptyDir": map[string]interface{}{},
								"name":     "scratch-data",
							},
							map[string]interface{}{
								"name": "webhook-certificates",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"optional":    true,
									"secretName":  "cnpg-webhook-cert",
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
															"cloudnative-pg",
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

	return mutate.MutateDeploymentNamespaceCnpgControllerManager(resourceObj, parent, collection, reconciler, req)
}
