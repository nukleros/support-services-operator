/*
Copyright 2024.

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

package mutate

import (
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// MutateSecretNamespaceExternalDnsGoogle mutates the Secret resource with name external-dns-google.
func MutateSecretNamespaceExternalDnsGoogle(
	original client.Object,
	parent *gatewayv1alpha1.ExternalDNS, collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	secret, ok := original.(*unstructured.Unstructured)
	if !ok {
		return []client.Object{original}, fmt.Errorf("expected *unstructured.Unstructured, got %T", original)
	}

	if parent.Spec.GcpProject == "" {
		return []client.Object{original}, errors.New("gcpProject is required when provider is \"google\"")
	}

	if err := unstructured.SetNestedField(secret.Object, parent.Spec.GcpProject, "stringData", "EXTERNAL_DNS_GOOGLE_PROJECT"); err != nil {
		return []client.Object{original}, fmt.Errorf("failed to set GCP project: %w", err)
	}

	if parent.Spec.DomainName == "" {
		return []client.Object{original}, errors.New("domainName is required")
	}

	if err := unstructured.SetNestedField(secret.Object, parent.Spec.DomainName, "stringData", "EXTERNAL_DNS_DOMAIN_FILTER"); err != nil {
		return []client.Object{original}, fmt.Errorf("failed to set domain filter: %w", err)
	}

	if parent.Spec.ZoneType != "" {
		if err := unstructured.SetNestedField(secret.Object, parent.Spec.ZoneType, "stringData", "EXTERNAL_DNS_GOOGLE_ZONE_VISIBILITY"); err != nil {
			return []client.Object{original}, fmt.Errorf("failed to set zone visibility: %w", err)
		}
	}

	return []client.Object{secret}, nil
}
