//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020-2023 Red Hat, Inc.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDevfileRegistriesList) DeepCopyInto(out *ClusterDevfileRegistriesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDevfileRegistriesList.
func (in *ClusterDevfileRegistriesList) DeepCopy() *ClusterDevfileRegistriesList {
	if in == nil {
		return nil
	}
	out := new(ClusterDevfileRegistriesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDevfileRegistriesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDevfileRegistriesListList) DeepCopyInto(out *ClusterDevfileRegistriesListList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterDevfileRegistriesList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDevfileRegistriesListList.
func (in *ClusterDevfileRegistriesListList) DeepCopy() *ClusterDevfileRegistriesListList {
	if in == nil {
		return nil
	}
	out := new(ClusterDevfileRegistriesListList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDevfileRegistriesListList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistriesList) DeepCopyInto(out *DevfileRegistriesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistriesList.
func (in *DevfileRegistriesList) DeepCopy() *DevfileRegistriesList {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistriesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevfileRegistriesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistriesListList) DeepCopyInto(out *DevfileRegistriesListList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevfileRegistriesList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistriesListList.
func (in *DevfileRegistriesListList) DeepCopy() *DevfileRegistriesListList {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistriesListList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevfileRegistriesListList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistriesListSpec) DeepCopyInto(out *DevfileRegistriesListSpec) {
	*out = *in
	if in.DevfileRegistries != nil {
		in, out := &in.DevfileRegistries, &out.DevfileRegistries
		*out = make([]DevfileRegistryService, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistriesListSpec.
func (in *DevfileRegistriesListSpec) DeepCopy() *DevfileRegistriesListSpec {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistriesListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistriesListStatus) DeepCopyInto(out *DevfileRegistriesListStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistriesListStatus.
func (in *DevfileRegistriesListStatus) DeepCopy() *DevfileRegistriesListStatus {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistriesListStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistry) DeepCopyInto(out *DevfileRegistry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistry.
func (in *DevfileRegistry) DeepCopy() *DevfileRegistry {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevfileRegistry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistryList) DeepCopyInto(out *DevfileRegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevfileRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistryList.
func (in *DevfileRegistryList) DeepCopy() *DevfileRegistryList {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevfileRegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistryService) DeepCopyInto(out *DevfileRegistryService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistryService.
func (in *DevfileRegistryService) DeepCopy() *DevfileRegistryService {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistrySpec) DeepCopyInto(out *DevfileRegistrySpec) {
	*out = *in
	out.DevfileIndex = in.DevfileIndex
	out.OciRegistry = in.OciRegistry
	out.RegistryViewer = in.RegistryViewer
	in.Storage.DeepCopyInto(&out.Storage)
	in.TLS.DeepCopyInto(&out.TLS)
	out.K8s = in.K8s
	out.Telemetry = in.Telemetry
	if in.Headless != nil {
		in, out := &in.Headless, &out.Headless
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistrySpec.
func (in *DevfileRegistrySpec) DeepCopy() *DevfileRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistrySpecContainer) DeepCopyInto(out *DevfileRegistrySpecContainer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistrySpecContainer.
func (in *DevfileRegistrySpecContainer) DeepCopy() *DevfileRegistrySpecContainer {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistrySpecContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistrySpecK8sOnly) DeepCopyInto(out *DevfileRegistrySpecK8sOnly) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistrySpecK8sOnly.
func (in *DevfileRegistrySpecK8sOnly) DeepCopy() *DevfileRegistrySpecK8sOnly {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistrySpecK8sOnly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistrySpecStorage) DeepCopyInto(out *DevfileRegistrySpecStorage) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistrySpecStorage.
func (in *DevfileRegistrySpecStorage) DeepCopy() *DevfileRegistrySpecStorage {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistrySpecStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistrySpecTLS) DeepCopyInto(out *DevfileRegistrySpecTLS) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistrySpecTLS.
func (in *DevfileRegistrySpecTLS) DeepCopy() *DevfileRegistrySpecTLS {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistrySpecTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistrySpecTelemetry) DeepCopyInto(out *DevfileRegistrySpecTelemetry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistrySpecTelemetry.
func (in *DevfileRegistrySpecTelemetry) DeepCopy() *DevfileRegistrySpecTelemetry {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistrySpecTelemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevfileRegistryStatus) DeepCopyInto(out *DevfileRegistryStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevfileRegistryStatus.
func (in *DevfileRegistryStatus) DeepCopy() *DevfileRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(DevfileRegistryStatus)
	in.DeepCopyInto(out)
	return out
}
