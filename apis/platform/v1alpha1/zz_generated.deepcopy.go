//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/nukleros/operator-builder-tools/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponent) DeepCopyInto(out *CertificatesComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponent.
func (in *CertificatesComponent) DeepCopy() *CertificatesComponent {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificatesComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentCollectionSpec) DeepCopyInto(out *CertificatesComponentCollectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentCollectionSpec.
func (in *CertificatesComponentCollectionSpec) DeepCopy() *CertificatesComponentCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentList) DeepCopyInto(out *CertificatesComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificatesComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentList.
func (in *CertificatesComponentList) DeepCopy() *CertificatesComponentList {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificatesComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentSpec) DeepCopyInto(out *CertificatesComponentSpec) {
	*out = *in
	out.Collection = in.Collection
	out.CertManager = in.CertManager
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentSpec.
func (in *CertificatesComponentSpec) DeepCopy() *CertificatesComponentSpec {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentSpecCertManager) DeepCopyInto(out *CertificatesComponentSpecCertManager) {
	*out = *in
	out.Cainjector = in.Cainjector
	out.Controller = in.Controller
	out.Webhook = in.Webhook
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentSpecCertManager.
func (in *CertificatesComponentSpecCertManager) DeepCopy() *CertificatesComponentSpecCertManager {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentSpecCertManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentSpecCertManagerCainjector) DeepCopyInto(out *CertificatesComponentSpecCertManagerCainjector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentSpecCertManagerCainjector.
func (in *CertificatesComponentSpecCertManagerCainjector) DeepCopy() *CertificatesComponentSpecCertManagerCainjector {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentSpecCertManagerCainjector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentSpecCertManagerController) DeepCopyInto(out *CertificatesComponentSpecCertManagerController) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentSpecCertManagerController.
func (in *CertificatesComponentSpecCertManagerController) DeepCopy() *CertificatesComponentSpecCertManagerController {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentSpecCertManagerController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentSpecCertManagerWebhook) DeepCopyInto(out *CertificatesComponentSpecCertManagerWebhook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentSpecCertManagerWebhook.
func (in *CertificatesComponentSpecCertManagerWebhook) DeepCopy() *CertificatesComponentSpecCertManagerWebhook {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentSpecCertManagerWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesComponentStatus) DeepCopyInto(out *CertificatesComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesComponentStatus.
func (in *CertificatesComponentStatus) DeepCopy() *CertificatesComponentStatus {
	if in == nil {
		return nil
	}
	out := new(CertificatesComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressComponent) DeepCopyInto(out *IngressComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressComponent.
func (in *IngressComponent) DeepCopy() *IngressComponent {
	if in == nil {
		return nil
	}
	out := new(IngressComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressComponentCollectionSpec) DeepCopyInto(out *IngressComponentCollectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressComponentCollectionSpec.
func (in *IngressComponentCollectionSpec) DeepCopy() *IngressComponentCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(IngressComponentCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressComponentList) DeepCopyInto(out *IngressComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressComponentList.
func (in *IngressComponentList) DeepCopy() *IngressComponentList {
	if in == nil {
		return nil
	}
	out := new(IngressComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressComponentSpec) DeepCopyInto(out *IngressComponentSpec) {
	*out = *in
	out.Collection = in.Collection
	out.ExternalDNS = in.ExternalDNS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressComponentSpec.
func (in *IngressComponentSpec) DeepCopy() *IngressComponentSpec {
	if in == nil {
		return nil
	}
	out := new(IngressComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressComponentSpecExternalDNS) DeepCopyInto(out *IngressComponentSpecExternalDNS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressComponentSpecExternalDNS.
func (in *IngressComponentSpecExternalDNS) DeepCopy() *IngressComponentSpecExternalDNS {
	if in == nil {
		return nil
	}
	out := new(IngressComponentSpecExternalDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressComponentStatus) DeepCopyInto(out *IngressComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressComponentStatus.
func (in *IngressComponentStatus) DeepCopy() *IngressComponentStatus {
	if in == nil {
		return nil
	}
	out := new(IngressComponentStatus)
	in.DeepCopyInto(out)
	return out
}
