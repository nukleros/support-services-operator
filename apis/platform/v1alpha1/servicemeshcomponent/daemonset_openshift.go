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

// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete

// CreateDaemonSetNamespaceIstioCniNodeOpenshift creates the DaemonSet resource with name istio-cni-node-openshift.
func CreateDaemonSetNamespaceIstioCniNodeOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": "DaemonSet",
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "apps/v1",
			"metadata": map[string]interface{}{
				"name":      "istio-cni-node-openshift",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"k8s-app":      "istio-cni-node",
					"istio.io/rev": "default",
					"install.operator.istio.io/owning-resource": "unknown",
					"operator.istio.io/component":               "Cni",
					"app.kubernetes.io/name":                    "istio-cni",
					"app.kubernetes.io/instance":                "istio",
					"app.kubernetes.io/part-of":                 "istio",
					"app.kubernetes.io/version":                 parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category":             "service-mesh",
					"platform.nukleros.io/project":              "istio",
				},
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"k8s-app": "istio-cni-node",
					},
				},
				"updateStrategy": map[string]interface{}{
					"rollingUpdate": map[string]interface{}{
						"maxUnavailable": 1,
					},
					"type": "RollingUpdate",
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"k8s-app":                       "istio-cni-node",
							"sidecar.istio.io/inject":       "false",
							"istio.io/dataplane-mode":       "none",
							"app.kubernetes.io/name":        "istio-cni",
							"app.kubernetes.io/instance":    "istio",
							"app.kubernetes.io/part-of":     "istio",
							"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
							"platform.nukleros.io/category": "service-mesh",
							"platform.nukleros.io/project":  "istio",
						},
						"annotations": map[string]interface{}{
							"sidecar.istio.io/inject": "false",
							"prometheus.io/scrape":    "true",
							"prometheus.io/port":      "15014",
							"prometheus.io/path":      "/metrics",
							"container.apparmor.security.beta.kubernetes.io/install-cni": "unconfined",
						},
					},
					"spec": map[string]interface{}{
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"tolerations": []interface{}{
							map[string]interface{}{
								"effect":   "NoSchedule",
								"operator": "Exists",
							},
							map[string]interface{}{
								"key":      "CriticalAddonsOnly",
								"operator": "Exists",
							},
							map[string]interface{}{
								"effect":   "NoExecute",
								"operator": "Exists",
							},
						},
						"priorityClassName":             "system-node-critical",
						"serviceAccountName":            "istio-cni-openshift",
						"terminationGracePeriodSeconds": 30,
						"containers": []interface{}{
							map[string]interface{}{
								"name": "install-cni",
								// controlled by field: istio.images.cni
								// controlled by field: istio.version
								"image": "" + parent.Spec.Istio.Images.Cni + ":" + parent.Spec.Istio.Version + "-distroless",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 15014,
										"name":          "metrics",
										"protocol":      "TCP",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/readyz",
										"port": 8000,
									},
								},
								"securityContext": map[string]interface{}{
									"privileged":   false,
									"runAsGroup":   0,
									"runAsUser":    0,
									"runAsNonRoot": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
										"add": []interface{}{
											"NET_ADMIN",
											"NET_RAW",
											"SYS_PTRACE",
											"SYS_ADMIN",
											"DAC_OVERRIDE",
										},
									},
									"seLinuxOptions": map[string]interface{}{
										"type": "spc_t",
									},
									"seccompProfile": map[string]interface{}{
										"type": "RuntimeDefault",
									},
								},
								"command": []interface{}{
									"install-cni",
								},
								"args": []interface{}{
									"--log_output_level=info",
									"--log_as_json",
								},
								"envFrom": []interface{}{
									map[string]interface{}{
										"configMapRef": map[string]interface{}{
											"name": "istio-cni-config-openshift",
										},
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "REPAIR_NODE_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "spec.nodeName",
											},
										},
									},
									map[string]interface{}{
										"name":  "REPAIR_RUN_AS_DAEMON",
										"value": "true",
									},
									map[string]interface{}{
										"name":  "REPAIR_SIDECAR_ANNOTATION",
										"value": "sidecar.istio.io/status",
									},
									map[string]interface{}{
										"name":  "ALLOW_SWITCH_TO_HOST_NS",
										"value": "true",
									},
									map[string]interface{}{
										"name": "NODE_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"apiVersion": "v1",
												"fieldPath":  "spec.nodeName",
											},
										},
									},
									map[string]interface{}{
										"name": "GOMEMLIMIT",
										"valueFrom": map[string]interface{}{
											"resourceFieldRef": map[string]interface{}{
												"resource": "limits.memory",
												"divisor":  "1",
											},
										},
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
										"mountPath": "/host/opt/cni/bin",
										"name":      "cni-bin-dir",
									},
									map[string]interface{}{
										"mountPath": "/host/proc",
										"name":      "cni-host-procfs",
										"readOnly":  true,
									},
									map[string]interface{}{
										"mountPath": "/host/etc/cni/net.d",
										"name":      "cni-net-dir",
									},
									map[string]interface{}{
										"mountPath": "/var/run/istio-cni",
										"name":      "cni-socket-dir",
									},
									map[string]interface{}{
										"mountPath":        "/host/var/run/netns",
										"mountPropagation": "HostToContainer",
										"name":             "cni-netns-dir",
									},
									map[string]interface{}{
										"mountPath": "/var/run/ztunnel",
										"name":      "cni-ztunnel-sock-dir",
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "256Mi",
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
								"name": "cni-bin-dir",
								"hostPath": map[string]interface{}{
									"path": "/var/lib/cni/bin",
								},
							},
							map[string]interface{}{
								"name": "cni-host-procfs",
								"hostPath": map[string]interface{}{
									"path": "/proc",
									"type": "Directory",
								},
							},
							map[string]interface{}{
								"name": "cni-ztunnel-sock-dir",
								"hostPath": map[string]interface{}{
									"path": "/var/run/ztunnel",
									"type": "DirectoryOrCreate",
								},
							},
							map[string]interface{}{
								"name": "cni-net-dir",
								"hostPath": map[string]interface{}{
									"path": "/etc/cni/multus/net.d",
								},
							},
							map[string]interface{}{
								"name": "cni-socket-dir",
								"hostPath": map[string]interface{}{
									"path": "/var/run/istio-cni",
								},
							},
							map[string]interface{}{
								"name": "cni-netns-dir",
								"hostPath": map[string]interface{}{
									"path": "/var/run/netns",
									"type": "DirectoryOrCreate",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDaemonSetNamespaceIstioCniNodeOpenshift(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete

// CreateDaemonSetNamespaceZtunnelOpenshift creates the DaemonSet resource with name ztunnel-openshift.
func CreateDaemonSetNamespaceZtunnelOpenshift(
	parent *platformv1alpha1.ServiceMeshComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Platform != "openshift" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=platform,value="openshift",include
			"apiVersion": "apps/v1",
			"kind":       "DaemonSet",
			"metadata": map[string]interface{}{
				"name":      "ztunnel-openshift",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "ztunnel",
					"app.kubernetes.io/instance":    "istio",
					"app.kubernetes.io/part-of":     "istio",
					"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
					"platform.nukleros.io/category": "service-mesh",
					"platform.nukleros.io/project":  "istio",
				},
				"annotations": map[string]interface{}{},
			},
			"spec": map[string]interface{}{
				"updateStrategy": map[string]interface{}{
					"rollingUpdate": map[string]interface{}{
						"maxSurge":       1,
						"maxUnavailable": 0,
					},
					"type": "RollingUpdate",
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "ztunnel",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"sidecar.istio.io/inject":       "false",
							"istio.io/dataplane-mode":       "none",
							"app":                           "ztunnel",
							"app.kubernetes.io/name":        "ztunnel",
							"app.kubernetes.io/instance":    "istio",
							"app.kubernetes.io/part-of":     "istio",
							"app.kubernetes.io/version":     parent.Spec.Istio.Version, //  controlled by field: istio.version
							"platform.nukleros.io/category": "service-mesh",
							"platform.nukleros.io/project":  "istio",
						},
						"annotations": map[string]interface{}{
							"sidecar.istio.io/inject": "false",
							"prometheus.io/port":      "15020",
							"prometheus.io/scrape":    "true",
						},
					},
					"spec": map[string]interface{}{
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"serviceAccountName": "ztunnel-openshift",
						"tolerations": []interface{}{
							map[string]interface{}{
								"effect":   "NoSchedule",
								"operator": "Exists",
							},
							map[string]interface{}{
								"key":      "CriticalAddonsOnly",
								"operator": "Exists",
							},
							map[string]interface{}{
								"effect":   "NoExecute",
								"operator": "Exists",
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "istio-proxy",
								// controlled by field: istio.images.ztunnel
								// controlled by field: istio.version
								"image": "" + parent.Spec.Istio.Images.Ztunnel + ":" + parent.Spec.Istio.Version + "-distroless",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 15020,
										"name":          "ztunnel-stats",
										"protocol":      "TCP",
									},
								},
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "1Gi",
									},
									"requests": map[string]interface{}{
										"cpu":    "200m",
										"memory": "512Mi",
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": true,
									"privileged":               false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
										"add": []interface{}{
											"NET_ADMIN",
											"SYS_ADMIN",
											"NET_RAW",
										},
									},
									"readOnlyRootFilesystem": true,
									"runAsGroup":             1337,
									"runAsNonRoot":           false,
									"runAsUser":              0,
									"seLinuxOptions": map[string]interface{}{
										"type": "spc_t",
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"port": 15021,
										"path": "/healthz/ready",
									},
								},
								"args": []interface{}{
									"proxy",
									"ztunnel",
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "CA_ADDRESS",
										// controlled by field: namespace
										"value": "istiod." + parent.Spec.Namespace + ".svc:15012",
									},
									map[string]interface{}{
										"name": "XDS_ADDRESS",
										// controlled by field: namespace
										"value": "istiod." + parent.Spec.Namespace + ".svc:15012",
									},
									map[string]interface{}{
										"name":  "LOG_FORMAT",
										"value": "json",
									},
									map[string]interface{}{
										"name":  "RUST_LOG",
										"value": "info",
									},
									map[string]interface{}{
										"name":  "RUST_BACKTRACE",
										"value": "1",
									},
									map[string]interface{}{
										"name":  "ISTIO_META_CLUSTER_ID",
										"value": "Kubernetes",
									},
									map[string]interface{}{
										"name":  "INPOD_ENABLED",
										"value": "true",
									},
									map[string]interface{}{
										"name":  "TERMINATION_GRACE_PERIOD_SECONDS",
										"value": "30",
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
										"name": "NODE_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "spec.nodeName",
											},
										},
									},
									map[string]interface{}{
										"name": "INSTANCE_IP",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "status.podIP",
											},
										},
									},
									map[string]interface{}{
										"name": "SERVICE_ACCOUNT",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "spec.serviceAccountName",
											},
										},
									},
									map[string]interface{}{
										"name":  "ISTIO_META_ENABLE_HBONE",
										"value": "true",
									},
									map[string]interface{}{
										"name": "ZTUNNEL_CPU_LIMIT",
										"valueFrom": map[string]interface{}{
											"resourceFieldRef": map[string]interface{}{
												"resource": "limits.cpu",
												"divisor":  "1",
											},
										},
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/var/run/secrets/istio",
										"name":      "istiod-ca-cert",
									},
									map[string]interface{}{
										"mountPath": "/var/run/secrets/tokens",
										"name":      "istio-token",
									},
									map[string]interface{}{
										"mountPath": "/var/run/ztunnel",
										"name":      "cni-ztunnel-sock-dir",
									},
									map[string]interface{}{
										"mountPath": "/tmp",
										"name":      "tmp",
									},
								},
							},
						},
						"priorityClassName":             "system-node-critical",
						"terminationGracePeriodSeconds": 30,
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "istio-token",
								"projected": map[string]interface{}{
									"sources": []interface{}{
										map[string]interface{}{
											"serviceAccountToken": map[string]interface{}{
												"path":              "istio-token",
												"expirationSeconds": 43200,
												"audience":          "istio-ca",
											},
										},
									},
								},
							},
							map[string]interface{}{
								"name": "istiod-ca-cert",
								"configMap": map[string]interface{}{
									"name": "istio-ca-root-cert",
								},
							},
							map[string]interface{}{
								"name": "cni-ztunnel-sock-dir",
								"hostPath": map[string]interface{}{
									"path": "/var/run/ztunnel",
									"type": "DirectoryOrCreate",
								},
							},
							map[string]interface{}{
								"name":     "tmp",
								"emptyDir": map[string]interface{}{},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDaemonSetNamespaceZtunnelOpenshift(resourceObj, parent, collection, reconciler, req)
}
