// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVirtualSwitch) DeepCopyInto(out *HostVirtualSwitch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVirtualSwitch.
func (in *HostVirtualSwitch) DeepCopy() *HostVirtualSwitch {
	if in == nil {
		return nil
	}
	out := new(HostVirtualSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostVirtualSwitch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVirtualSwitchList) DeepCopyInto(out *HostVirtualSwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostVirtualSwitch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVirtualSwitchList.
func (in *HostVirtualSwitchList) DeepCopy() *HostVirtualSwitchList {
	if in == nil {
		return nil
	}
	out := new(HostVirtualSwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostVirtualSwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVirtualSwitchObservation) DeepCopyInto(out *HostVirtualSwitchObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVirtualSwitchObservation.
func (in *HostVirtualSwitchObservation) DeepCopy() *HostVirtualSwitchObservation {
	if in == nil {
		return nil
	}
	out := new(HostVirtualSwitchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVirtualSwitchParameters) DeepCopyInto(out *HostVirtualSwitchParameters) {
	*out = *in
	if in.StandbyNics != nil {
		in, out := &in.StandbyNics, &out.StandbyNics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetworkAdapters != nil {
		in, out := &in.NetworkAdapters, &out.NetworkAdapters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActiveNics != nil {
		in, out := &in.ActiveNics, &out.ActiveNics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVirtualSwitchParameters.
func (in *HostVirtualSwitchParameters) DeepCopy() *HostVirtualSwitchParameters {
	if in == nil {
		return nil
	}
	out := new(HostVirtualSwitchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVirtualSwitchSpec) DeepCopyInto(out *HostVirtualSwitchSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVirtualSwitchSpec.
func (in *HostVirtualSwitchSpec) DeepCopy() *HostVirtualSwitchSpec {
	if in == nil {
		return nil
	}
	out := new(HostVirtualSwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostVirtualSwitchStatus) DeepCopyInto(out *HostVirtualSwitchStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostVirtualSwitchStatus.
func (in *HostVirtualSwitchStatus) DeepCopy() *HostVirtualSwitchStatus {
	if in == nil {
		return nil
	}
	out := new(HostVirtualSwitchStatus)
	in.DeepCopyInto(out)
	return out
}