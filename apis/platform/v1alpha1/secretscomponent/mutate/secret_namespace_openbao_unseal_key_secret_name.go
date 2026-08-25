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
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"regexp"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// unsealKeyDataKey is the key that must exist in the secret's data, whether
// managed by us or provided externally (unmanaged).
const unsealKeyDataKey = "unseal-key.key"

// unsealKeyShapePattern matches the shape of `openssl rand -base64 32`'s
// output: 32 random bytes, standard base64 encoded (43 characters + 1 '=' pad).
var unsealKeyShapePattern = regexp.MustCompile(`^[A-Za-z0-9+/]{43}=$`)

// MutateSecretNamespaceOpenbaoUnsealKeySecretName mutates the Secret resource with name parent.Spec.Openbao.UnsealKey.Secret.Name.
func MutateSecretNamespaceOpenbaoUnsealKeySecretName(
	original client.Object,
	parent *platformv1alpha1.SecretsComponent, collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler, req *workload.Request,
) ([]client.Object, error) {
	// if either the reconciler or request are found to be nil, return the base object.
	if reconciler == nil || req == nil {
		return []client.Object{original}, nil
	}

	secretSpec := parent.Spec.Openbao.UnsealKey.Secret

	// unmanaged: the secret is expected to already exist somewhere else -
	// validate it rather than create/own anything.
	if secretSpec.Type == "unmanaged" {
		if secretSpec.Name == "" || secretSpec.Namespace == "" {
			return nil, fmt.Errorf(
				"openbao.unsealKey.secret.name and openbao.unsealKey.secret.namespace are both required when openbao.unsealKey.secret.type is 'unmanaged'",
			)
		}

		existing := &unstructured.Unstructured{}
		existing.SetAPIVersion("v1")
		existing.SetKind("Secret")

		if err := reconciler.Get(req.Context, client.ObjectKey{Namespace: secretSpec.Namespace, Name: secretSpec.Name}, existing); err != nil {
			if apierrors.IsNotFound(err) {
				return nil, fmt.Errorf("unmanaged secret %s/%s (openbao.unsealKey.secret) was not found", secretSpec.Namespace, secretSpec.Name)
			}

			return nil, fmt.Errorf("unable to look up unmanaged secret %s/%s: %w", secretSpec.Namespace, secretSpec.Name, err)
		}

		encodedKey, found, err := unstructured.NestedString(existing.Object, "data", unsealKeyDataKey)
		if err != nil || !found || encodedKey == "" {
			return nil, fmt.Errorf(
				"unmanaged secret %s/%s (openbao.unsealKey.secret) must contain a non-empty '%s' key in its data",
				secretSpec.Namespace, secretSpec.Name, unsealKeyDataKey,
			)
		}

		// best-effort shape check only: OpenBao's static seal also accepts hex
		// and other encodings, so a mismatch here is a warning, not a failure.
		if decoded, decodeErr := base64.StdEncoding.DecodeString(encodedKey); decodeErr != nil || len(decoded) != 32 || !unsealKeyShapePattern.MatchString(encodedKey) {
			req.Log.Info(
				"unseal-key.key does not look like base64(openssl rand 32); continuing anyway since other valid key encodings exist",
				"namespace", secretSpec.Namespace, "name", secretSpec.Name,
			)
		}

		// nothing for us to manage - the referenced secret already exists and is valid.
		return []client.Object{}, nil
	}

	// managed (the default): generate the key once, ourselves, the first time
	// this secret doesn't already exist with data set.
	u, ok := original.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected *unstructured.Unstructured, got %T", original)
	}

	existing := &unstructured.Unstructured{}
	existing.SetAPIVersion("v1")
	existing.SetKind("Secret")

	err := reconciler.Get(req.Context, client.ObjectKey{Namespace: u.GetNamespace(), Name: u.GetName()}, existing)

	switch {
	case err == nil:
		if encodedKey, found, ferr := unstructured.NestedString(existing.Object, "data", unsealKeyDataKey); ferr == nil && found && encodedKey != "" {
			// already generated - leave the live secret untouched. We must never
			// regenerate/rotate this key out from under an already-initialized
			// OpenBao raft cluster, or it becomes permanently unable to unseal.
			return []client.Object{}, nil
		}
	case !apierrors.IsNotFound(err):
		return nil, fmt.Errorf("unable to look up managed secret %s/%s: %w", u.GetNamespace(), u.GetName(), err)
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("unable to generate unseal key: %w", err)
	}

	unstructured.RemoveNestedField(u.Object, "stringData")

	if err := unstructured.SetNestedField(u.Object, base64.StdEncoding.EncodeToString(key), "data", unsealKeyDataKey); err != nil {
		return nil, fmt.Errorf("unable to set generated unseal key: %w", err)
	}

	return []client.Object{u}, nil
}
