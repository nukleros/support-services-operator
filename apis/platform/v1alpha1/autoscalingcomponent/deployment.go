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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceKedaOperator creates the Deployment resource with name keda-operator.
func CreateDeploymentNamespaceKedaOperator(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "keda-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "keda-operator",
					"name":                          "keda-operator",
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"revisionHistoryLimit": 10,
				// controlled by field: keda.operator.replicas
				//  Number of replicas to use for the KEDA operator deployment.  With 1 replica, its pod
				//  disruption budget (minAvailable: 1) blocks node drains.
				"replicas": parent.Spec.Keda.Operator.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-operator",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                           "keda-operator",
							"name":                          "keda-operator",
							"app.kubernetes.io/name":        "keda-operator",
							"app.kubernetes.io/component":   "operator",
							"app.kubernetes.io/part-of":     "keda-operator",
							"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
							"app.kubernetes.io/instance":    "keda",
							"platform.nukleros.io/category": "autoscaling",
							"platform.nukleros.io/project":  "keda",
						},
					},
					"spec": map[string]interface{}{
						"enableServiceLinks":           false,
						"serviceAccountName":           "keda-operator",
						"automountServiceAccountToken": true,
						"securityContext": map[string]interface{}{
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "keda-operator",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
								// controlled by field: keda.operator.image
								// controlled by field: keda.version
								//  Image repo and name to use for the KEDA operator.
								//  Version of KEDA to use.  Applies to the operator, metrics apiserver and admission webhooks,
								//  which are released together.
								"image": "" + parent.Spec.Keda.Operator.Image + ":" + parent.Spec.Keda.Version + "",
								"command": []interface{}{
									"/keda",
								},
								"args": []interface{}{
									"--leader-elect",
									"--disable-compression=true",
									"--zap-log-level=info",
									"--zap-encoder=json",
									"--zap-time-encoding=rfc3339",
									"--enable-webhook-patching=true",
									"--cert-dir=/certs",
									"--enable-cert-rotation=true",
									"--cert-secret-name=kedaorg-certs",
									"--operator-service-name=keda-operator",
									"--metrics-service-bind-address=:9666",
									"--metrics-server-service-name=keda-operator-metrics-apiserver",
									"--webhooks-service-name=keda-admission-webhooks",
									"--k8s-cluster-name=kubernetes-default",
									"--k8s-cluster-domain=cluster.local",
									"--enable-prometheus-metrics=true",
									"--metrics-bind-address=:8080",
									"--http-max-idle-conns=0",
									"--http-max-idle-conns-per-host=1000",
									"--http-idle-conn-timeout=90s",
									"--service-account-token-mode=enforce-audience",
									"--vault-kubernetes-auth-token-file=/var/run/secrets/keda-vault/token",
								},
								"imagePullPolicy": "IfNotPresent",
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/healthz",
										"port": 8081,
									},
									"initialDelaySeconds": 25,
									"periodSeconds":       10,
									"timeoutSeconds":      1,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/readyz",
										"port": 8081,
									},
									"initialDelaySeconds": 20,
									"periodSeconds":       3,
									"timeoutSeconds":      1,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"name":          "metrics",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 9666,
										"name":          "metricsservice",
										"protocol":      "TCP",
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name":  "KEDA_SERVICE_ACCOUNT_TOKEN_AUDIENCES",
										"value": "[{\"audience\":\"vault\"}]",
									},
									map[string]interface{}{
										"name":  "WATCH_NAMESPACE",
										"value": "",
									},
									map[string]interface{}{
										"name":  "WATCH_LABEL_SELECTOR",
										"value": "",
									},
									map[string]interface{}{
										"name":  "WATCH_LABEL_SELECTOR_FOR_TRIGGERAUTH",
										"value": "",
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
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "OPERATOR_NAME",
										"value": "keda-operator",
									},
									map[string]interface{}{
										"name":  "KEDA_HTTP_DEFAULT_TIMEOUT",
										"value": "3000",
									},
									map[string]interface{}{
										"name":  "KEDA_KUBERNETES_API_TIMEOUT",
										"value": "5s",
									},
									map[string]interface{}{
										"name":  "KEDA_HTTP_MIN_TLS_VERSION",
										"value": "TLS12",
									},
									map[string]interface{}{
										"name":  "KEDA_HTTP_TLS_CIPHER_LIST",
										"value": "",
									},
									map[string]interface{}{
										"name":  "KEDA_SERVICE_MIN_TLS_VERSION",
										"value": "TLS13",
									},
									map[string]interface{}{
										"name":  "KEDA_SERVICE_TLS_CIPHER_LIST",
										"value": "",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "keda-vault-e6f0a1fbb43c8919",
										"mountPath": "/var/run/secrets/keda-vault",
										"readOnly":  true,
									},
									map[string]interface{}{
										"mountPath": "/certs",
										"name":      "certificates",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "1000Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "100Mi",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "keda-vault-e6f0a1fbb43c8919",
								"projected": map[string]interface{}{
									"sources": []interface{}{
										map[string]interface{}{
											"serviceAccountToken": map[string]interface{}{
												"path":              "token",
												"audience":          "vault",
												"expirationSeconds": 3600,
											},
										},
									},
								},
							},
							map[string]interface{}{
								"name": "certificates",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "kedaorg-certs",
									"optional":    true,
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"podAffinityTerm": map[string]interface{}{
											"labelSelector": map[string]interface{}{
												"matchLabels": map[string]interface{}{
													"app": "keda-operator",
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"topologySpreadConstraints": []interface{}{
							map[string]interface{}{
								"labelSelector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app": "keda-operator",
									},
								},
								"maxSkew":           1,
								"topologyKey":       "topology.kubernetes.io/zone",
								"whenUnsatisfiable": "ScheduleAnyway",
							},
						},
						"dnsPolicy":   "ClusterFirst",
						"hostNetwork": false,
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceKedaOperator(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceKedaOperatorMetricsApiserver creates the Deployment resource with name keda-operator-metrics-apiserver.
func CreateDeploymentNamespaceKedaOperatorMetricsApiserver(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "keda-operator-metrics-apiserver",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "keda-operator-metrics-apiserver",
					"app.kubernetes.io/name":        "keda-operator-metrics-apiserver",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"revisionHistoryLimit": 10,
				// controlled by field: keda.metricsApiserver.replicas
				//  Number of replicas to use for the KEDA metrics apiserver deployment.  With 1 replica, its
				//  pod disruption budget (minAvailable: 1) blocks node drains.
				"replicas": parent.Spec.Keda.MetricsApiserver.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-operator-metrics-apiserver",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                           "keda-operator-metrics-apiserver",
							"app.kubernetes.io/name":        "keda-operator-metrics-apiserver",
							"app.kubernetes.io/component":   "operator",
							"app.kubernetes.io/part-of":     "keda-operator",
							"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
							"app.kubernetes.io/instance":    "keda",
							"platform.nukleros.io/category": "autoscaling",
							"platform.nukleros.io/project":  "keda",
						},
						"annotations": map[string]interface{}{
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "8080",
						},
					},
					"spec": map[string]interface{}{
						"enableServiceLinks":           false,
						"serviceAccountName":           "keda-metrics-server",
						"automountServiceAccountToken": true,
						"securityContext": map[string]interface{}{
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "keda-operator-metrics-apiserver",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
								// controlled by field: keda.metricsApiserver.image
								// controlled by field: keda.version
								//  Image repo and name to use for the KEDA metrics apiserver.
								"image":           "" + parent.Spec.Keda.MetricsApiserver.Image + ":" + parent.Spec.Keda.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/healthz",
										"port":   6443,
										"scheme": "HTTPS",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
									"timeoutSeconds":      1,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/readyz",
										"port":   6443,
										"scheme": "HTTPS",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       3,
									"timeoutSeconds":      1,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"env": []interface{}{
									map[string]interface{}{
										"name":  "WATCH_NAMESPACE",
										"value": "",
									},
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "KEDA_HTTP_DEFAULT_TIMEOUT",
										"value": "3000",
									},
									map[string]interface{}{
										"name":  "KEDA_HTTP_MIN_TLS_VERSION",
										"value": "TLS12",
									},
									map[string]interface{}{
										"name":  "KEDA_HTTP_TLS_CIPHER_LIST",
										"value": "",
									},
									map[string]interface{}{
										"name":  "KEDA_SERVICE_MIN_TLS_VERSION",
										"value": "TLS13",
									},
									map[string]interface{}{
										"name":  "KEDA_SERVICE_TLS_CIPHER_LIST",
										"value": "",
									},
								},
								"command": []interface{}{
									"/keda-adapter",
								},
								"args": []interface{}{
									"--port=8080",
									"--secure-port=6443",
									"--logtostderr=true",
									"--stderrthreshold=ERROR",
									"--disable-compression=true",
									"--metrics-service-address=keda-operator.$(POD_NAMESPACE).svc.cluster.local:9666",
									"--client-ca-file=/certs/ca.crt",
									"--tls-cert-file=/certs/tls.crt",
									"--tls-private-key-file=/certs/tls.key",
									"--cert-dir=/certs",
									"--v=0",
									"--zap-log-level=info",
									"--zap-encoder=json",
									"--zap-time-encoding=rfc3339",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 6443,
										"name":          "https",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 8080,
										"name":          "metrics",
										"protocol":      "TCP",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/certs",
										"name":      "certificates",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "1000Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "100Mi",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "certificates",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "kedaorg-certs",
								},
							},
						},
						"dnsPolicy":   "ClusterFirst",
						"hostNetwork": false,
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"podAffinityTerm": map[string]interface{}{
											"labelSelector": map[string]interface{}{
												"matchLabels": map[string]interface{}{
													"app": "keda-operator-metrics-apiserver",
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"topologySpreadConstraints": []interface{}{
							map[string]interface{}{
								"labelSelector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app": "keda-operator-metrics-apiserver",
									},
								},
								"maxSkew":           1,
								"topologyKey":       "topology.kubernetes.io/zone",
								"whenUnsatisfiable": "ScheduleAnyway",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceKedaOperatorMetricsApiserver(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceKedaAdmissionWebhooks creates the Deployment resource with name keda-admission-webhooks.
func CreateDeploymentNamespaceKedaAdmissionWebhooks(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "keda-admission-webhooks",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "keda-admission-webhooks",
					"name":                          "keda-admission-webhooks",
					"app.kubernetes.io/name":        "keda-admission-webhooks",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"app.kubernetes.io/instance":    "keda",
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
			},
			"spec": map[string]interface{}{
				"revisionHistoryLimit": 10,
				// controlled by field: keda.admissionWebhooks.replicas
				//  Number of replicas to use for the KEDA admission webhooks deployment.  With 1 replica, its
				//  pod disruption budget (minAvailable: 1) blocks node drains.
				"replicas": parent.Spec.Keda.AdmissionWebhooks.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "keda-admission-webhooks",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                           "keda-admission-webhooks",
							"name":                          "keda-admission-webhooks",
							"app.kubernetes.io/name":        "keda-admission-webhooks",
							"app.kubernetes.io/component":   "operator",
							"app.kubernetes.io/part-of":     "keda-operator",
							"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
							"app.kubernetes.io/instance":    "keda",
							"platform.nukleros.io/category": "autoscaling",
							"platform.nukleros.io/project":  "keda",
						},
					},
					"spec": map[string]interface{}{
						"enableServiceLinks":           false,
						"serviceAccountName":           "keda-webhook",
						"automountServiceAccountToken": true,
						"securityContext": map[string]interface{}{
							"runAsNonRoot": true,
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "keda-admission-webhooks",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"readOnlyRootFilesystem": true,
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
								// controlled by field: keda.admissionWebhooks.image
								// controlled by field: keda.version
								//  Image repo and name to use for the KEDA admission webhooks.
								"image": "" + parent.Spec.Keda.AdmissionWebhooks.Image + ":" + parent.Spec.Keda.Version + "",
								"command": []interface{}{
									"/keda-admission-webhooks",
								},
								"args": []interface{}{
									"--zap-log-level=info",
									"--zap-encoder=json",
									"--zap-time-encoding=rfc3339",
									"--cert-dir=/certs",
									"--health-probe-bind-address=:8081",
									"--metrics-bind-address=:8080",
								},
								"imagePullPolicy": "IfNotPresent",
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/healthz",
										"port": 8081,
									},
									"initialDelaySeconds": 25,
									"periodSeconds":       10,
									"timeoutSeconds":      1,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/readyz",
										"port": 8081,
									},
									"initialDelaySeconds": 20,
									"periodSeconds":       3,
									"timeoutSeconds":      1,
									"failureThreshold":    3,
									"successThreshold":    1,
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 9443,
										"name":          "http",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 8080,
										"name":          "metrics",
										"protocol":      "TCP",
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name":  "WATCH_NAMESPACE",
										"value": "",
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
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/certs",
										"name":      "certificates",
										"readOnly":  true,
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "1000Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "100Mi",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "certificates",
								"secret": map[string]interface{}{
									"defaultMode": 420,
									"secretName":  "kedaorg-certs",
								},
							},
						},
						"hostNetwork": false,
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"podAffinityTerm": map[string]interface{}{
											"labelSelector": map[string]interface{}{
												"matchLabels": map[string]interface{}{
													"app": "keda-admission-webhooks",
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"topologySpreadConstraints": []interface{}{
							map[string]interface{}{
								"labelSelector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app": "keda-admission-webhooks",
									},
								},
								"maxSkew":           1,
								"topologyKey":       "topology.kubernetes.io/zone",
								"whenUnsatisfiable": "ScheduleAnyway",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceKedaAdmissionWebhooks(resourceObj, parent, collection, reconciler, req)
}
