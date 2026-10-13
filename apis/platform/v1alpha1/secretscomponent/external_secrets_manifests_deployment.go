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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceExternalSecretsCertController creates the Deployment resource with name external-secrets-cert-controller.
func CreateDeploymentNamespaceExternalSecretsCertController(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-cert-controller",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-cert-controller",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/version": parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
						},
					},
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: externalSecrets.certController.replicas
				//  Number of replicas to use for the external-secrets cert-controller deployment.
				"replicas":             parent.Spec.ExternalSecrets.CertController.Replicas,
				"revisionHistoryLimit": 10,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":     "external-secrets-cert-controller",
						"app.kubernetes.io/instance": "external-secrets",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":        "external-secrets-cert-controller",
							"app.kubernetes.io/instance":    "external-secrets",
							"app.kubernetes.io/version":     "v2.12.0",
							"platform.nukleros.io/category": "secrets",
							"platform.nukleros.io/project":  "external-secrets",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName":           "external-secrets-cert-controller",
						"automountServiceAccountToken": true,
						"hostNetwork":                  false,
						"containers": []interface{}{
							map[string]interface{}{
								"name": "cert-controller",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"runAsNonRoot":             true,
									"runAsUser":                65534,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								// controlled by field: externalSecrets.image
								// controlled by field: externalSecrets.version
								//  Image repo and name to use for external-secrets.
								//  Version of external-secrets to use.
								"image":           "" + parent.Spec.ExternalSecrets.Image + ":" + parent.Spec.ExternalSecrets.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"certcontroller",
									"--crd-requeue-interval=5m",
									"--service-name=external-secrets-webhook",
									"--service-namespace=nukleros-secrets-system",
									"--secret-name=external-secrets-webhook",
									"--secret-namespace=nukleros-secrets-system",
									"--metrics-addr=:8080",
									"--healthz-addr=:8081",
									"--loglevel=info",
									"--zap-time-encoding=epoch",
									"--enable-partial-cache=true",
									"--enable-leader-election=true",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"protocol":      "TCP",
										"name":          "metrics",
									},
									map[string]interface{}{
										"containerPort": 8081,
										"protocol":      "TCP",
										"name":          "ready",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port": "ready",
										"path": "/readyz",
									},
									"initialDelaySeconds": 20,
									"periodSeconds":       5,
									"timeoutSeconds":      5,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "50m",
										"memory": "32Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "100m",
										"memory": "64Mi",
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
															"external-secrets-cert-controller",
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

	return mutate.MutateDeploymentNamespaceExternalSecretsCertController(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceExternalSecrets creates the Deployment resource with name external-secrets.
func CreateDeploymentNamespaceExternalSecrets(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "external-secrets",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/version": parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
						},
					},
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: externalSecrets.controller.replicas
				//  Number of replicas to use for the external-secrets controller deployment.
				"replicas":             parent.Spec.ExternalSecrets.Controller.Replicas,
				"revisionHistoryLimit": 10,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":     "external-secrets",
						"app.kubernetes.io/instance": "external-secrets",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":        "external-secrets",
							"app.kubernetes.io/instance":    "external-secrets",
							"app.kubernetes.io/version":     "v2.12.0",
							"platform.nukleros.io/category": "secrets",
							"platform.nukleros.io/project":  "external-secrets",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName":           "external-secrets",
						"automountServiceAccountToken": true,
						"hostNetwork":                  false,
						"containers": []interface{}{
							map[string]interface{}{
								"name": "external-secrets",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"runAsNonRoot":             true,
									"runAsUser":                65534,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								// controlled by field: externalSecrets.image
								// controlled by field: externalSecrets.version
								//  Image repo and name to use for external-secrets.
								//  Version of external-secrets to use.
								"image":           "" + parent.Spec.ExternalSecrets.Image + ":" + parent.Spec.ExternalSecrets.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"--enable-leader-election=true",
									"--concurrent=2",
									"--metrics-addr=:8080",
									"--loglevel=info",
									"--zap-time-encoding=epoch",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"protocol":      "TCP",
										"name":          "metrics",
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "50m",
										"memory": "128Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "100m",
										"memory": "256Mi",
									},
								},
							},
						},
						"dnsPolicy": "ClusterFirst",
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
															"external-secrets",
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

	return mutate.MutateDeploymentNamespaceExternalSecrets(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceExternalSecretsWebhook creates the Deployment resource with name external-secrets-webhook.
func CreateDeploymentNamespaceExternalSecretsWebhook(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-webhook",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/version": parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
						},
					},
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: externalSecrets.webhook.replicas
				//  Number of replicas to use for the external-secrets webhook deployment.
				"replicas":             parent.Spec.ExternalSecrets.Webhook.Replicas,
				"revisionHistoryLimit": 10,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":     "external-secrets-webhook",
						"app.kubernetes.io/instance": "external-secrets",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":        "external-secrets-webhook",
							"app.kubernetes.io/instance":    "external-secrets",
							"app.kubernetes.io/version":     "v2.12.0",
							"platform.nukleros.io/category": "secrets",
							"platform.nukleros.io/project":  "external-secrets",
						},
					},
					"spec": map[string]interface{}{
						"hostNetwork":                  false,
						"serviceAccountName":           "external-secrets-webhook",
						"automountServiceAccountToken": true,
						"containers": []interface{}{
							map[string]interface{}{
								"name": "webhook",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"runAsNonRoot":             true,
									"runAsUser":                65534,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								// controlled by field: externalSecrets.image
								// controlled by field: externalSecrets.version
								//  Image repo and name to use for external-secrets.
								//  Version of external-secrets to use.
								"image":           "" + parent.Spec.ExternalSecrets.Image + ":" + parent.Spec.ExternalSecrets.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"args": []interface{}{
									"webhook",
									"--port=10250",
									"--dns-name=external-secrets-webhook.nukleros-secrets-system.svc",
									"--cert-dir=/tmp/certs",
									"--check-interval=5m",
									"--metrics-addr=:8080",
									"--healthz-addr=:8081",
									"--loglevel=info",
									"--zap-time-encoding=epoch",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"protocol":      "TCP",
										"name":          "metrics",
									},
									map[string]interface{}{
										"containerPort": 10250,
										"protocol":      "TCP",
										"name":          "webhook",
									},
									map[string]interface{}{
										"containerPort": 8081,
										"protocol":      "TCP",
										"name":          "ready",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port": "ready",
										"path": "/readyz",
									},
									"initialDelaySeconds": 20,
									"periodSeconds":       5,
									"timeoutSeconds":      5,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "certs",
										"mountPath": "/tmp/certs",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "50m",
										"memory": "32Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "100m",
										"memory": "64Mi",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "certs",
								"secret": map[string]interface{}{
									"secretName": "external-secrets-webhook",
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
															"external-secrets-webhook",
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

	return mutate.MutateDeploymentNamespaceExternalSecretsWebhook(resourceObj, parent, collection, reconciler, req)
}
