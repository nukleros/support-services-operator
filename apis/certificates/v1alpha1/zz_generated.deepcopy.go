//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/nukleros/operator-builder-tools/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManager) DeepCopyInto(out *CertManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManager.
func (in *CertManager) DeepCopy() *CertManager {
	if in == nil {
		return nil
	}
	out := new(CertManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerCollectionSpec) DeepCopyInto(out *CertManagerCollectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerCollectionSpec.
func (in *CertManagerCollectionSpec) DeepCopy() *CertManagerCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(CertManagerCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerList) DeepCopyInto(out *CertManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerList.
func (in *CertManagerList) DeepCopy() *CertManagerList {
	if in == nil {
		return nil
	}
	out := new(CertManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerSpec) DeepCopyInto(out *CertManagerSpec) {
	*out = *in
	out.Collection = in.Collection
	out.Cainjector = in.Cainjector
	out.Controller = in.Controller
	out.Webhook = in.Webhook
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerSpec.
func (in *CertManagerSpec) DeepCopy() *CertManagerSpec {
	if in == nil {
		return nil
	}
	out := new(CertManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerSpecCainjector) DeepCopyInto(out *CertManagerSpecCainjector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerSpecCainjector.
func (in *CertManagerSpecCainjector) DeepCopy() *CertManagerSpecCainjector {
	if in == nil {
		return nil
	}
	out := new(CertManagerSpecCainjector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerSpecController) DeepCopyInto(out *CertManagerSpecController) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerSpecController.
func (in *CertManagerSpecController) DeepCopy() *CertManagerSpecController {
	if in == nil {
		return nil
	}
	out := new(CertManagerSpecController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerSpecWebhook) DeepCopyInto(out *CertManagerSpecWebhook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerSpecWebhook.
func (in *CertManagerSpecWebhook) DeepCopy() *CertManagerSpecWebhook {
	if in == nil {
		return nil
	}
	out := new(CertManagerSpecWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerStatus) DeepCopyInto(out *CertManagerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerStatus.
func (in *CertManagerStatus) DeepCopy() *CertManagerStatus {
	if in == nil {
		return nil
	}
	out := new(CertManagerStatus)
	in.DeepCopyInto(out)
	return out
}
