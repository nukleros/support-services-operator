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

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

// CreateSecretNamespaceOpenbaoUnsealKeySecretName creates the Secret resource with name parent.Spec.Openbao.UnsealKey.Secret.Name.
func CreateSecretNamespaceOpenbaoUnsealKeySecretName(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"stringData": map[string]interface{}{
				"unseal-key.key": "",
			},
			"kind": "Secret",
			"metadata": map[string]interface{}{
				// controlled by field: openbao.unsealKey.secret.type

				// controlled by field: openbao.unsealKey.secret.namespace

				// controlled by field: openbao.unsealKey.secret.name
				//  Type of secret.  One of 'managed' or 'unmanaged'.

				//  +kubebuilder:validation:Enum=managed;unmanaged
				//  Namespace of the secret which contains the unseal key at openbao.unsealKey.secret.name.
				//
				//  Only relevant when openbao.unsealKey.secret.type is not managed.
				//  Name of the secret which contains the unseal key.  The secret must contain
				//  the key 'unseal-key.key'.
				//
				//  Only relevant when openbao.unsealKey.secret.type is not managed.
				"name":      parent.Spec.Openbao.UnsealKey.Secret.Name,
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"type": "Opaque",
		},
	}

	return mutate.MutateSecretNamespaceOpenbaoUnsealKeySecretName(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

// CreateSecretNamespaceOpenbaoInjectorCerts creates the Secret resource with name openbao-injector-certs.
func CreateSecretNamespaceOpenbaoInjectorCerts(
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
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "openbao-injector-certs",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao-agent-injector",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
		},
	}

	return mutate.MutateSecretNamespaceOpenbaoInjectorCerts(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapNamespaceOpenbaoConfig creates the ConfigMap resource with name openbao-config.
func CreateConfigMapNamespaceOpenbaoConfig(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "openbao-config",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao",
					"app.kubernetes.io/instance":    "openbao",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"data": map[string]interface{}{
				"extraconfig-from-values.hcl": `ui = true

listener "tcp" {
  tls_disable = 1
  address = "[::]:8200"
  cluster_address = "[::]:8201"
  # Enable unauthenticated metrics access (necessary for Prometheus Operator)
  #telemetry {
  #  unauthenticated_metrics_access = "true"
  #}
}

storage "raft" {
  path = "/openbao/data"

  # each peer tries the others in turn so the 3 replicas form a single
  # raft cluster automatically, without a manual "operator raft join" per pod
  retry_join {
    leader_api_addr = "http://openbao-0.openbao-internal:8200"
  }
  retry_join {
    leader_api_addr = "http://openbao-1.openbao-internal:8200"
  }
  retry_join {
    leader_api_addr = "http://openbao-2.openbao-internal:8200"
  }
}

service_registration "kubernetes" {}

# Auto-unseal via a locally-mounted static key (no cloud provider, no HSM,
# no second Vault/OpenBao instance) - suited to a local, airgapped install.
# The key file comes from the openbao-unseal-key Secret mounted via
# server.extraVolumes above. Generate and load the key out-of-band, e.g.:
#   kubectl create secret generic openbao-unseal-key \
#     --namespace nukleros-secrets-system \
#     --from-literal=unseal-key.key=$(openssl rand -base64 32)
# (base64, not raw bytes: OpenBao's static seal accepts base64/hex key
# material, and raw random bytes can contain a NUL byte that a shell
# command substitution silently strips, corrupting the key)
seal "static" {
  current_key_id = "initial"
  current_key    = "file:///openbao/userconfig/openbao-unseal-key/unseal-key.key"
}`,
			},
		},
	}

	return mutate.MutateConfigMapNamespaceOpenbaoConfig(resourceObj, parent, collection, reconciler, req)
}
