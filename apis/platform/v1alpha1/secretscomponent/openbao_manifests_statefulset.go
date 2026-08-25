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

// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete

// CreateStatefulSetNamespaceOpenbao creates the StatefulSet resource with name openbao.
func CreateStatefulSetNamespaceOpenbao(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
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
				"serviceName":         "openbao-internal",
				"podManagementPolicy": "OrderedReady",
				// controlled by field: openbao.replicas
				//  Number of replicas to use for the OpenBao server statefulset.
				"replicas": parent.Spec.Openbao.Replicas,
				"updateStrategy": map[string]interface{}{
					"type": "OnDelete",
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":     "openbao",
						"app.kubernetes.io/instance": "openbao",
						"component":                  "server",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"helm.sh/chart":                 "openbao-0.29.2",
							"app.kubernetes.io/name":        "openbao",
							"app.kubernetes.io/instance":    "openbao",
							"component":                     "server",
							"platform.nukleros.io/category": "secrets",
							"platform.nukleros.io/project":  "openbao",
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
													"app.kubernetes.io/instance": "openbao",
													"app.kubernetes.io/name":     "openbao",
													"component":                  "server",
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"terminationGracePeriodSeconds": 10,
						"serviceAccountName":            "openbao",
						"securityContext": map[string]interface{}{
							"seccompProfile": map[string]interface{}{
								"type": "RuntimeDefault",
							},
							"runAsNonRoot": true,
							"runAsGroup":   1000,
							"runAsUser":    100,
							"fsGroup":      1000,
						},
						"hostNetwork": false,
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "config",
								"configMap": map[string]interface{}{
									"name": "openbao-config",
								},
							},
							map[string]interface{}{
								"name": "userconfig-openbao-unseal-key",
								"secret": map[string]interface{}{
									"secretName":  "openbao-unseal-key",
									"defaultMode": 256,
								},
							},
							map[string]interface{}{
								"name":     "home",
								"emptyDir": map[string]interface{}{},
							},
						},
						"initContainers": []interface{}{
							map[string]interface{}{
								"command": []interface{}{
									"/bin/sh",
									"-c",
									`[ "$(hostname)" = "openbao-0" ] || exit 0

cp /openbao/config/extraconfig-from-values.hcl /tmp/storageconfig.hcl
bao server -config=/tmp/storageconfig.hcl &
BAO_PID=$!

n=0
while ! bao status 2>/dev/null | grep -q '^Initialized'; do
  n=$((n+1))
  if [ "$n" -ge 30 ]; then
    kill -TERM "$BAO_PID"
    wait "$BAO_PID" 2>/dev/null
    exit 0
  fi
  sleep 2
done

if ! bao status 2>/dev/null | grep -Eq '^Initialized +true'; then
  INIT_OUTPUT="$(bao operator init -format=json 2>&1)"
  echo "$INIT_OUTPUT"

  # The root token + recovery keys above are only ever printed this
  # once. Persist them to a Secret so a later pod restart (which
  # skips straight past this block, since it's already initialized)
  # doesn't leave them unrecoverable in a log that's since rotated
  # away. Only ever creates this one, specific secret name - see the
  # RBAC note on the Role granting this.
  #
  # --no-check-certificate: this image's wget has no way to point at
  # a custom CA bundle (no --ca-certificate flag, and neither
  # SSL_CERT_FILE nor appending to the system bundle made it trust a
  # test cert when checked) so it can't validate the cluster's own
  # API server certificate. This call stays inside the pod network
  # and is bearer-token authenticated regardless, so skipping server
  # cert validation is an accepted tradeoff for this one-time
  # bootstrap call rather than a real MITM exposure.
  TOKEN_FILE=/var/run/secrets/kubernetes.io/serviceaccount/token
  if [ -f "$TOKEN_FILE" ]; then
    RECOVERY_KEYS=$(printf '%s' "$INIT_OUTPUT" | sed -n '/"recovery_keys_b64"/,/\]/p' | grep -oE '"[A-Za-z0-9+/=]{20,}"' | tr -d '"')
    B64_INIT=$(printf '%s' "$INIT_OUTPUT" | base64 | tr -d '\n')
    B64_RECOVERY=$(printf '%s' "$RECOVERY_KEYS" | base64 | tr -d '\n')
    printf '{"apiVersion":"v1","kind":"Secret","metadata":{"name":"openbao-token"},"type":"Opaque","data":{"init-output.json":"%s","recovery-keys.txt":"%s"}}' "$B64_INIT" "$B64_RECOVERY" > /tmp/secret.json
    echo "=== creating openbao-token secret ==="
    wget -q -O- --no-check-certificate \
      --header="Authorization: Bearer $(cat "$TOKEN_FILE")" \
      --header="Content-Type: application/json" \
      --post-file=/tmp/secret.json \
      "https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}/api/v1/namespaces/${BAO_K8S_NAMESPACE}/secrets" 2>&1
    echo ""
    echo "=== secret creation attempt complete (see result above) ==="
  fi
  sleep 2
fi

kill -TERM "$BAO_PID"
wait "$BAO_PID" 2>/dev/null
exit 0
`,
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "HOST_IP",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "status.hostIP",
											},
										},
									},
									map[string]interface{}{
										"name": "POD_IP",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "status.podIP",
											},
										},
									},
									map[string]interface{}{
										"name": "BAO_K8S_POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name": "BAO_K8S_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name": "HOSTNAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name":  "BAO_ADDR",
										"value": "http://127.0.0.1:8200",
									},
									map[string]interface{}{
										"name":  "BAO_API_ADDR",
										"value": "http://$(POD_IP):8200",
									},
									map[string]interface{}{
										"name":  "BAO_CLUSTER_ADDR",
										"value": "https://$(HOSTNAME).openbao-internal:8201",
									},
									map[string]interface{}{
										"name":  "HOME",
										"value": "/home/openbao",
									},
									map[string]interface{}{
										"name":  "BAO_LOG_FORMAT",
										"value": "json",
									},
								},
								"image": "quay.io/openbao/openbao:2.6.2",
								"name":  "bao-init",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/openbao/config",
										"name":      "config",
									},
									map[string]interface{}{
										"mountPath": "/openbao/data",
										"name":      "data",
									},
									map[string]interface{}{
										"mountPath": "/openbao/userconfig/openbao-unseal-key",
										"name":      "userconfig-openbao-unseal-key",
										"readOnly":  true,
									},
									map[string]interface{}{
										"mountPath": "/home/openbao",
										"name":      "home",
									},
								},
							},
						},
						"containers": []interface{}{
							map[string]interface{}{
								"name": "openbao",
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "256Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "250m",
										"memory": "256Mi",
									},
								},
								// controlled by field: openbao.image
								// controlled by field: openbao.version
								//  Image repo and name to use for OpenBao.
								//  Version of OpenBao to use.
								"image":           "" + parent.Spec.Openbao.Image + ":" + parent.Spec.Openbao.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"command": []interface{}{
									"/bin/sh",
									"-ec",
								},
								"args": []interface{}{
									`cp /openbao/config/extraconfig-from-values.hcl /tmp/storageconfig.hcl;
[ -n "${HOST_IP}" ] && sed -Ei "s|HOST_IP|${HOST_IP?}|g" /tmp/storageconfig.hcl;
[ -n "${POD_IP}" ] && sed -Ei "s|POD_IP|${POD_IP?}|g" /tmp/storageconfig.hcl;
[ -n "${HOSTNAME}" ] && sed -Ei "s|HOSTNAME|${HOSTNAME?}|g" /tmp/storageconfig.hcl;
[ -n "${API_ADDR}" ] && sed -Ei "s|API_ADDR|${API_ADDR?}|g" /tmp/storageconfig.hcl;
[ -n "${TRANSIT_ADDR}" ] && sed -Ei "s|TRANSIT_ADDR|${TRANSIT_ADDR?}|g" /tmp/storageconfig.hcl;
[ -n "${RAFT_ADDR}" ] && sed -Ei "s|RAFT_ADDR|${RAFT_ADDR?}|g" /tmp/storageconfig.hcl;
/usr/local/bin/docker-entrypoint.sh bao server -config=/tmp/storageconfig.hcl 
`,
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "HOST_IP",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "status.hostIP",
											},
										},
									},
									map[string]interface{}{
										"name": "POD_IP",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "status.podIP",
											},
										},
									},
									map[string]interface{}{
										"name": "BAO_K8S_POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name": "BAO_K8S_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "BAO_ADDR",
										"value": "http://127.0.0.1:8200",
									},
									map[string]interface{}{
										"name":  "BAO_API_ADDR",
										"value": "http://$(POD_IP):8200",
									},
									map[string]interface{}{
										"name":  "SKIP_CHOWN",
										"value": "true",
									},
									map[string]interface{}{
										"name":  "SKIP_SETCAP",
										"value": "true",
									},
									map[string]interface{}{
										"name": "HOSTNAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name":  "BAO_CLUSTER_ADDR",
										"value": "https://$(HOSTNAME).openbao-internal:8201",
									},
									map[string]interface{}{
										"name":  "HOME",
										"value": "/home/openbao",
									},
									map[string]interface{}{
										"name":  "BAO_LOG_FORMAT",
										"value": "json",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "audit",
										"mountPath": "/openbao/audit",
									},
									map[string]interface{}{
										"name":      "data",
										"mountPath": "/openbao/data",
									},
									map[string]interface{}{
										"name":      "config",
										"mountPath": "/openbao/config",
									},
									map[string]interface{}{
										"name":      "userconfig-openbao-unseal-key",
										"readOnly":  true,
										"mountPath": "/openbao/userconfig/openbao-unseal-key",
									},
									map[string]interface{}{
										"name":      "home",
										"mountPath": "/home/openbao",
									},
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8200,
										"name":          "http",
									},
									map[string]interface{}{
										"containerPort": 8201,
										"name":          "https-internal",
									},
									map[string]interface{}{
										"containerPort": 8202,
										"name":          "http-rep",
									},
								},
								"readinessProbe": map[string]interface{}{
									"exec": map[string]interface{}{
										"command": []interface{}{
											"/bin/sh",
											"-ec",
											"bao status -tls-skip-verify",
										},
									},
									"failureThreshold":    2,
									"initialDelaySeconds": 5,
									"periodSeconds":       5,
									"successThreshold":    1,
									"timeoutSeconds":      3,
								},
								"lifecycle": map[string]interface{}{
									"preStop": map[string]interface{}{
										"exec": map[string]interface{}{
											"command": []interface{}{
												"/bin/sh",
												"-c",
												"sleep 5 && kill -SIGTERM $(pidof bao)",
											},
										},
									},
								},
							},
						},
					},
				},
				"volumeClaimTemplates": []interface{}{
					map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "PersistentVolumeClaim",
						"metadata": map[string]interface{}{
							"name": "data",
						},
						"spec": map[string]interface{}{
							"accessModes": []interface{}{
								"ReadWriteOnce",
							},
							"resources": map[string]interface{}{
								"requests": map[string]interface{}{
									// controlled by field: openbao.storage.data.size
									//  Size of the PVC used for OpenBao's raft data storage.
									"storage": parent.Spec.Openbao.Storage.Data.Size,
								},
							},
							// controlled by field: openbao.storage.data.class
							//  StorageClass to use for OpenBao's raft data volume.  Leave empty to
							//  use the cluster's default StorageClass.
							"storageClassName": parent.Spec.Openbao.Storage.Data.Class,
						},
					},
					map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "PersistentVolumeClaim",
						"metadata": map[string]interface{}{
							"name": "audit",
						},
						"spec": map[string]interface{}{
							"accessModes": []interface{}{
								"ReadWriteOnce",
							},
							"resources": map[string]interface{}{
								"requests": map[string]interface{}{
									// controlled by field: openbao.storage.audit.size
									//  Size of the PVC used for OpenBao's audit log storage.
									"storage": parent.Spec.Openbao.Storage.Audit.Size,
								},
							},
							// controlled by field: openbao.storage.audit.class
							//  StorageClass to use for OpenBao's audit log volume.  Leave empty to
							//  use the cluster's default StorageClass.
							"storageClassName": parent.Spec.Openbao.Storage.Audit.Class,
						},
					},
				},
			},
		},
	}

	return mutate.MutateStatefulSetNamespaceOpenbao(resourceObj, parent, collection, reconciler, req)
}
