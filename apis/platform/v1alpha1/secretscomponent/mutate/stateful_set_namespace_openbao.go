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

package mutate

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// MutateStatefulSetNamespaceOpenbao mutates the StatefulSet resource with name openbao.
func MutateStatefulSetNamespaceOpenbao(
	original client.Object,
	parent *platformv1alpha1.SecretsComponent, collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	u, ok := original.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected *unstructured.Unstructured, got %T", original)
	}

	// An empty storageClassName explicitly disables dynamic provisioning in
	// Kubernetes - that's different from omitting the field, which lets the
	// cluster's default StorageClass apply. Remove the key entirely when the
	// user hasn't set a class, so an unset field behaves like an unset field.
	classByVolumeName := map[string]string{
		"data":  parent.Spec.Openbao.Storage.Data.Class,
		"audit": parent.Spec.Openbao.Storage.Audit.Class,
	}

	templates, found, err := unstructured.NestedFieldNoCopy(u.Object, "spec", "volumeClaimTemplates")
	if err != nil {
		return nil, fmt.Errorf("unable to read spec.volumeClaimTemplates: %w", err)
	}

	if found {
		templateSlice, ok := templates.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected spec.volumeClaimTemplates to be a slice, got %T", templates)
		}

		for _, template := range templateSlice {
			templateMap, ok := template.(map[string]interface{})
			if !ok {
				continue
			}

			name, _, _ := unstructured.NestedString(templateMap, "metadata", "name")

			class, tracked := classByVolumeName[name]
			if !tracked || class != "" {
				continue
			}

			specMap, ok := templateMap["spec"].(map[string]interface{})
			if !ok {
				continue
			}

			delete(specMap, "storageClassName")
		}
	}

	return []client.Object{u}, nil
}
