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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceIstiod creates the Deployment resource with name istiod.
func CreateDeploymentNamespaceIstiod(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform == "gke" {
		return []client.Object{}, nil
	}

	if collection.Spec.Platform == "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="gke",include=false
			// +operator-builder:resource:collectionField=platform,value="openshift",include=false
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "istiod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":          "istiod",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Pilot",
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
				"strategy": map[string]interface{}{
					"rollingUpdate": map[string]interface{}{
						"maxSurge":       "100%",
						"maxUnavailable": "25%",
					},
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"istio": "pilot",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":          "istiod",
							"istio.io/rev": "default",
							"install.operator.istio.io/owning-resource": "unknown",
							"sidecar.istio.io/inject":                   "false",
							"operator.istio.io/component":               "Pilot",
							"istio":                                     "pilot",
							"istio.io/dataplane-mode":                   "none",
							"app.kubernetes.io/name":                    "istiod",
							"app.kubernetes.io/instance":                "istio",
							"app.kubernetes.io/part-of":                 "istio",
							"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
							"platform.nukleros.io/category":             "service-mesh",
							"platform.nukleros.io/project":              "istio",
						},
						"annotations": map[string]interface{}{
							"prometheus.io/port":      "15014",
							"prometheus.io/scrape":    "true",
							"sidecar.istio.io/inject": "false",
						},
					},
					"spec": map[string]interface{}{
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"podAffinityTerm": map[string]interface{}{
											"labelSelector": map[string]interface{}{
												"matchLabels": map[string]interface{}{
													"app": "istiod",
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"tolerations": []interface{}{
							map[string]interface{}{
								"key":      "cni.istio.io/not-ready",
								"operator": "Exists",
							},
						},
						"topologySpreadConstraints": []interface{}{
							map[string]interface{}{
								"labelSelector": map[string]interface{}{
									"matchLabels": map[string]interface{}{
										"app": "istiod",
									},
								},
								"maxSkew":           1,
								"topologyKey":       "topology.kubernetes.io/zone",
								"whenUnsatisfiable": "ScheduleAnyway",
							},
						},
						"serviceAccountName": "istiod",
						"containers": []interface{}{
							map[string]interface{}{
								"name": "discovery",
								// controlled by field: istio.images.istiod
								// controlled by field: istio.version
								//  Image repo and name to use for istiod.
								//  Version of istio to use.  This version is used for all istio images, including the proxy
								//  image used for waypoints.  Distroless images are used, so this should not include the
								//  '-distroless' suffix.
								"image": "" + parent.Spec.Istio.Images.Istiod + ":" + parent.Spec.Istio.Version + "-distroless",
								"args": []interface{}{
									"discovery",
									"--monitoringAddr=:15014",
									"--log_output_level=default:info",
									"--log_as_json",
									"--domain",
									"cluster.local",
									"--keepaliveMaxServerConnectionAge",
									"30m",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"protocol":      "TCP",
										"name":          "http-debug",
									},
									map[string]interface{}{
										"containerPort": 15010,
										"protocol":      "TCP",
										"name":          "grpc-xds",
									},
									map[string]interface{}{
										"containerPort": 15012,
										"protocol":      "TCP",
										"name":          "tls-xds",
									},
									map[string]interface{}{
										"containerPort": 15017,
										"protocol":      "TCP",
										"name":          "https-webhooks",
									},
									map[string]interface{}{
										"containerPort": 15014,
										"protocol":      "TCP",
										"name":          "http-monitoring",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/ready",
										"port": 8080,
									},
									"initialDelaySeconds": 1,
									"periodSeconds":       3,
									"timeoutSeconds":      5,
								},
								"env": []interface{}{
									map[string]interface{}{
										"name":  "REVISION",
										"value": "default",
									},
									map[string]interface{}{
										"name":  "PILOT_CERT_PROVIDER",
										"value": "istiod",
									},
									map[string]interface{}{
										"name": "POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"apiVersion": "v1",
												"fieldPath":  "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"apiVersion": "v1",
												"fieldPath":  "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name": "SERVICE_ACCOUNT",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"apiVersion": "v1",
												"fieldPath":  "spec.serviceAccountName",
											},
										},
									},
									map[string]interface{}{
										"name":  "KUBECONFIG",
										"value": "/var/run/secrets/remote/config",
									},
									map[string]interface{}{
										"name": "CA_TRUSTED_NODE_ACCOUNTS",
										// controlled by field: namespace
										"value": "" + parent.Spec.Namespace + "/ztunnel",
									},
									map[string]interface{}{
										"name":  "PILOT_ENABLE_AMBIENT",
										"value": "true",
									},
									map[string]interface{}{
										"name":  "PILOT_TRACE_SAMPLING",
										"value": "1",
									},
									map[string]interface{}{
										"name":  "PILOT_ENABLE_ANALYSIS",
										"value": "false",
									},
									map[string]interface{}{
										"name":  "CLUSTER_ID",
										"value": "Kubernetes",
									},
									map[string]interface{}{
										"name": "GOMAXPROCS",
										"valueFrom": map[string]interface{}{
											"resourceFieldRef": map[string]interface{}{
												"resource": "limits.cpu",
												"divisor":  "1",
											},
										},
									},
									map[string]interface{}{
										"name":  "PLATFORM",
										"value": "",
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "2Gi",
									},
									"requests": map[string]interface{}{
										"cpu":    "500m",
										"memory": "2Gi",
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"runAsNonRoot":             true,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "istio-token",
										"mountPath": "/var/run/secrets/tokens",
										"readOnly":  true,
									},
									map[string]interface{}{
										"name":      "local-certs",
										"mountPath": "/var/run/secrets/istio-dns",
									},
									map[string]interface{}{
										"name":      "cacerts",
										"mountPath": "/etc/cacerts",
										"readOnly":  true,
									},
									map[string]interface{}{
										"name":      "istio-kubeconfig",
										"mountPath": "/var/run/secrets/remote",
										"readOnly":  true,
									},
									map[string]interface{}{
										"name":      "istio-csr-dns-cert",
										"mountPath": "/var/run/secrets/istiod/tls",
										"readOnly":  true,
									},
									map[string]interface{}{
										"name":      "istio-csr-ca-configmap",
										"mountPath": "/var/run/secrets/istiod/ca",
										"readOnly":  true,
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"emptyDir": map[string]interface{}{
									"medium": "Memory",
								},
								"name": "local-certs",
							},
							map[string]interface{}{
								"name": "istio-token",
								"projected": map[string]interface{}{
									"sources": []interface{}{
										map[string]interface{}{
											"serviceAccountToken": map[string]interface{}{
												"audience":          "istio-ca",
												"expirationSeconds": 43200,
												"path":              "istio-token",
											},
										},
									},
								},
							},
							map[string]interface{}{
								"name": "cacerts",
								"secret": map[string]interface{}{
									"secretName": "cacerts",
									"optional":   true,
								},
							},
							map[string]interface{}{
								"name": "istio-kubeconfig",
								"secret": map[string]interface{}{
									"secretName": "istio-kubeconfig",
									"optional":   true,
								},
							},
							map[string]interface{}{
								"name": "istio-csr-dns-cert",
								"secret": map[string]interface{}{
									"secretName": "istiod-tls",
									"optional":   true,
								},
							},
							map[string]interface{}{
								"name": "istio-csr-ca-configmap",
								"configMap": map[string]interface{}{
									"name":        "istio-ca-root-cert",
									"defaultMode": 420,
									"optional":    true,
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

	return mutate.MutateDeploymentNamespaceIstiod(resourceObj, parent, collection, reconciler, req)
}
