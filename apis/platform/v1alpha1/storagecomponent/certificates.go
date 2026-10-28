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

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceSeaweedfsOperatorWebhookCert creates the Certificate resource with name seaweedfs-operator-webhook-cert.
func CreateCertNamespaceSeaweedfsOperatorWebhookCert(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-operator-webhook-cert",
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
				"dnsNames": []interface{}{
					"seaweedfs-operator-webhook",
					"seaweedfs-operator-webhook." + parent.Spec.Namespace + "",                   //  controlled by field: namespace
					"seaweedfs-operator-webhook." + parent.Spec.Namespace + ".svc",               //  controlled by field: namespace
					"seaweedfs-operator-webhook." + parent.Spec.Namespace + ".svc.cluster.local", //  controlled by field: namespace
				},
				"secretName": "seaweedfs-operator-webhook-server-cert",
				"issuerRef": map[string]interface{}{
					"name":  "seaweedfs-operator-webhook-selfsigned-issuer",
					"kind":  "Issuer",
					"group": "cert-manager.io",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceSeaweedfsOperatorWebhookCert(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch;create;update;patch;delete

// CreateIssuerNamespaceSeaweedfsOperatorWebhookSelfsignedIssuer creates the Issuer resource with name seaweedfs-operator-webhook-selfsigned-issuer.
func CreateIssuerNamespaceSeaweedfsOperatorWebhookSelfsignedIssuer(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Issuer",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-operator-webhook-selfsigned-issuer",
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
				"selfSigned": map[string]interface{}{},
			},
		},
	}

	return mutate.MutateIssuerNamespaceSeaweedfsOperatorWebhookSelfsignedIssuer(resourceObj, parent, collection, reconciler, req)
}
