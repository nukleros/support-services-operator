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

package orchestration

import (
	v1alpha1orchestration "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
	v1alpha1supportservices "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1/supportservices"
)

// Code generated by operator-builder. DO NOT EDIT.

// SupportServicesLatestGroupVersion returns the latest group version object associated with this
// particular kind.
var SupportServicesLatestGroupVersion = v1alpha1orchestration.GroupVersion

// SupportServicesLatestSample returns the latest sample manifest associated with this
// particular kind.
var SupportServicesLatestSample = v1alpha1supportservices.Sample(false)
