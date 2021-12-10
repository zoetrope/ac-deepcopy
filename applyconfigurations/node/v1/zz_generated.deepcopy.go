//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	applyconfigurationscorev1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
	metav1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverheadApplyConfiguration) DeepCopyInto(out *OverheadApplyConfiguration) {
	*out = *in
	if in.PodFixed != nil {
		in, out := &in.PodFixed, &out.PodFixed
		*out = new(corev1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[corev1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverheadApplyConfiguration.
func (in *OverheadApplyConfiguration) DeepCopy() *OverheadApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(OverheadApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeClassApplyConfiguration) DeepCopyInto(out *RuntimeClassApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.Overhead != nil {
		in, out := &in.Overhead, &out.Overhead
		*out = new(OverheadApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Scheduling != nil {
		in, out := &in.Scheduling, &out.Scheduling
		*out = new(SchedulingApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeClassApplyConfiguration.
func (in *RuntimeClassApplyConfiguration) DeepCopy() *RuntimeClassApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(RuntimeClassApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingApplyConfiguration) DeepCopyInto(out *SchedulingApplyConfiguration) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]applyconfigurationscorev1.TolerationApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingApplyConfiguration.
func (in *SchedulingApplyConfiguration) DeepCopy() *SchedulingApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(SchedulingApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}
