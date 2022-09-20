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

package application

import (
	v1alpha1application "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	//+kubebuilder:scaffold:operator-builder:imports

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DatabaseComponentGroupVersions returns all group version objects associated with this kind.
func DatabaseComponentGroupVersions() []schema.GroupVersion {
	return []schema.GroupVersion{
		v1alpha1application.GroupVersion,
		//+kubebuilder:scaffold:operator-builder:groupversions
	}
}
