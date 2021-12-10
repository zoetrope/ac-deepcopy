//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
	v1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIStorageCapacityApplyConfiguration) DeepCopyInto(out *CSIStorageCapacityApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeTopology != nil {
		in, out := &in.NodeTopology, &out.NodeTopology
		*out = new(v1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.MaximumVolumeSize != nil {
		in, out := &in.MaximumVolumeSize, &out.MaximumVolumeSize
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIStorageCapacityApplyConfiguration.
func (in *CSIStorageCapacityApplyConfiguration) DeepCopy() *CSIStorageCapacityApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CSIStorageCapacityApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentApplyConfiguration) DeepCopyInto(out *VolumeAttachmentApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(VolumeAttachmentSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VolumeAttachmentStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentApplyConfiguration.
func (in *VolumeAttachmentApplyConfiguration) DeepCopy() *VolumeAttachmentApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentSourceApplyConfiguration) DeepCopyInto(out *VolumeAttachmentSourceApplyConfiguration) {
	*out = *in
	if in.PersistentVolumeName != nil {
		in, out := &in.PersistentVolumeName, &out.PersistentVolumeName
		*out = new(string)
		**out = **in
	}
	if in.InlineVolumeSpec != nil {
		in, out := &in.InlineVolumeSpec, &out.InlineVolumeSpec
		*out = new(corev1.PersistentVolumeSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentSourceApplyConfiguration.
func (in *VolumeAttachmentSourceApplyConfiguration) DeepCopy() *VolumeAttachmentSourceApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentSourceApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentSpecApplyConfiguration) DeepCopyInto(out *VolumeAttachmentSpecApplyConfiguration) {
	*out = *in
	if in.Attacher != nil {
		in, out := &in.Attacher, &out.Attacher
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(VolumeAttachmentSourceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentSpecApplyConfiguration.
func (in *VolumeAttachmentSpecApplyConfiguration) DeepCopy() *VolumeAttachmentSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentStatusApplyConfiguration) DeepCopyInto(out *VolumeAttachmentStatusApplyConfiguration) {
	*out = *in
	if in.Attached != nil {
		in, out := &in.Attached, &out.Attached
		*out = new(bool)
		**out = **in
	}
	if in.AttachmentMetadata != nil {
		in, out := &in.AttachmentMetadata, &out.AttachmentMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AttachError != nil {
		in, out := &in.AttachError, &out.AttachError
		*out = new(VolumeErrorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DetachError != nil {
		in, out := &in.DetachError, &out.DetachError
		*out = new(VolumeErrorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentStatusApplyConfiguration.
func (in *VolumeAttachmentStatusApplyConfiguration) DeepCopy() *VolumeAttachmentStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeErrorApplyConfiguration) DeepCopyInto(out *VolumeErrorApplyConfiguration) {
	*out = *in
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = (*in).DeepCopy()
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeErrorApplyConfiguration.
func (in *VolumeErrorApplyConfiguration) DeepCopy() *VolumeErrorApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeErrorApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}
