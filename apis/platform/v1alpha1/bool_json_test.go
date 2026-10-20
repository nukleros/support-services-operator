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

package v1alpha1

import (
	"encoding/json"
	"testing"

	applicationv1alpha1 "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// TestBoolJSONKeepsFalse marshals each boolean and requires false to stay in the object.
func TestBoolJSONKeepsFalse(t *testing.T) {
	// marshal every bool field and require false to remain
	cases := []struct {
		name  string
		value any
		field string
	}{
		{
			name:  "kong include",
			value: IngressComponentSpecKong{Include: false, ProxyServiceName: "kong-proxy"},
			field: "include",
		},
		{
			name:  "nginx include",
			value: IngressComponentSpecNginx{Include: false},
			field: "include",
		},
		{
			name:  "openbao injector include",
			value: SecretsComponentSpecOpenbaoInjector{Include: false},
			field: "include",
		},
		{
			name:  "ingress created",
			value: IngressComponentStatus{},
			field: "created",
		},
		{
			name:  "ingress dependenciesSatisfied",
			value: IngressComponentStatus{},
			field: "dependenciesSatisfied",
		},
		{
			name:  "certificates created",
			value: CertificatesComponentStatus{},
			field: "created",
		},
		{
			name:  "certificates dependenciesSatisfied",
			value: CertificatesComponentStatus{},
			field: "dependenciesSatisfied",
		},
		{
			name:  "secrets created",
			value: SecretsComponentStatus{},
			field: "created",
		},
		{
			name:  "secrets dependenciesSatisfied",
			value: SecretsComponentStatus{},
			field: "dependenciesSatisfied",
		},
		{
			name:  "database created",
			value: applicationv1alpha1.DatabaseComponentStatus{},
			field: "created",
		},
		{
			name:  "database dependenciesSatisfied",
			value: applicationv1alpha1.DatabaseComponentStatus{},
			field: "dependenciesSatisfied",
		},
		{
			name:  "support services created",
			value: setupv1alpha1.SupportServicesStatus{},
			field: "created",
		},
		{
			name:  "support services dependenciesSatisfied",
			value: setupv1alpha1.SupportServicesStatus{},
			field: "dependenciesSatisfied",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// require the marshaled object to contain the field set to false
			raw, err := json.Marshal(tc.value)
			if err != nil {
				t.Fatalf("marshal %s: %v", tc.name, err)
			}

			decoded := map[string]any{}
			if err := json.Unmarshal(raw, &decoded); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.name, err)
			}

			got, ok := decoded[tc.field]
			if !ok {
				t.Fatalf("%s: field %s missing from %s", tc.name, tc.field, raw)
			}
			if got != false {
				t.Fatalf("%s: field %s = %v, want false", tc.name, tc.field, got)
			}
		})
	}
}
