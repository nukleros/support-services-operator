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
)

// TestBoolJSONKeepsFalse marshals a false bool and requires the field to stay.
func TestBoolJSONKeepsFalse(t *testing.T) {
	// marshal the include field and one status field
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
			name:  "ingress created",
			value: IngressComponentStatus{},
			field: "created",
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
