/*
Copyright 2022.

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

package platform

import (
	v1alpha1platform "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	v1alpha1secretscomponent "github.com/nukleros/support-services-operator/apis/platform/v1alpha1/secretscomponent"
)

// Code generated by operator-builder. DO NOT EDIT.

// SecretsComponentLatestGroupVersion returns the latest group version object associated with this
// particular kind.
var SecretsComponentLatestGroupVersion = v1alpha1platform.GroupVersion

// SecretsComponentLatestSample returns the latest sample manifest associated with this
// particular kind.
var SecretsComponentLatestSample = v1alpha1secretscomponent.Sample(false)
