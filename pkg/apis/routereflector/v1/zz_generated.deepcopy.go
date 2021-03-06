// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteReflector) DeepCopyInto(out *RouteReflector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteReflector.
func (in *RouteReflector) DeepCopy() *RouteReflector {
	if in == nil {
		return nil
	}
	out := new(RouteReflector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteReflector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteReflectorList) DeepCopyInto(out *RouteReflectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteReflector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteReflectorList.
func (in *RouteReflectorList) DeepCopy() *RouteReflectorList {
	if in == nil {
		return nil
	}
	out := new(RouteReflectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteReflectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteReflectorSpec) DeepCopyInto(out *RouteReflectorSpec) {
	*out = *in
	if in.IgnoredPools != nil {
		in, out := &in.IgnoredPools, &out.IgnoredPools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteReflectorSpec.
func (in *RouteReflectorSpec) DeepCopy() *RouteReflectorSpec {
	if in == nil {
		return nil
	}
	out := new(RouteReflectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteReflectorStatus) DeepCopyInto(out *RouteReflectorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteReflectorStatus.
func (in *RouteReflectorStatus) DeepCopy() *RouteReflectorStatus {
	if in == nil {
		return nil
	}
	out := new(RouteReflectorStatus)
	in.DeepCopyInto(out)
	return out
}
