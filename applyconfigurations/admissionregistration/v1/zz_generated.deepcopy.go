//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookApplyConfiguration) DeepCopyInto(out *MutatingWebhookApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ClientConfig != nil {
		in, out := &in.ClientConfig, &out.ClientConfig
		*out = new(WebhookClientConfigApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RuleWithOperationsApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(admissionregistrationv1.FailurePolicyType)
		**out = **in
	}
	if in.MatchPolicy != nil {
		in, out := &in.MatchPolicy, &out.MatchPolicy
		*out = new(admissionregistrationv1.MatchPolicyType)
		**out = **in
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectSelector != nil {
		in, out := &in.ObjectSelector, &out.ObjectSelector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.SideEffects != nil {
		in, out := &in.SideEffects, &out.SideEffects
		*out = new(admissionregistrationv1.SideEffectClass)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.AdmissionReviewVersions != nil {
		in, out := &in.AdmissionReviewVersions, &out.AdmissionReviewVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReinvocationPolicy != nil {
		in, out := &in.ReinvocationPolicy, &out.ReinvocationPolicy
		*out = new(admissionregistrationv1.ReinvocationPolicyType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookApplyConfiguration.
func (in *MutatingWebhookApplyConfiguration) DeepCopy() *MutatingWebhookApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookConfigurationApplyConfiguration) DeepCopyInto(out *MutatingWebhookConfigurationApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]MutatingWebhookApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookConfigurationApplyConfiguration.
func (in *MutatingWebhookConfigurationApplyConfiguration) DeepCopy() *MutatingWebhookConfigurationApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookConfigurationApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleApplyConfiguration) DeepCopyInto(out *RuleApplyConfiguration) {
	*out = *in
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIVersions != nil {
		in, out := &in.APIVersions, &out.APIVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(admissionregistrationv1.ScopeType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleApplyConfiguration.
func (in *RuleApplyConfiguration) DeepCopy() *RuleApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(RuleApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWithOperationsApplyConfiguration) DeepCopyInto(out *RuleWithOperationsApplyConfiguration) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]admissionregistrationv1.OperationType, len(*in))
		copy(*out, *in)
	}
	in.RuleApplyConfiguration.DeepCopyInto(&out.RuleApplyConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleWithOperationsApplyConfiguration.
func (in *RuleWithOperationsApplyConfiguration) DeepCopy() *RuleWithOperationsApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(RuleWithOperationsApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReferenceApplyConfiguration) DeepCopyInto(out *ServiceReferenceApplyConfiguration) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceReferenceApplyConfiguration.
func (in *ServiceReferenceApplyConfiguration) DeepCopy() *ServiceReferenceApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceReferenceApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookApplyConfiguration) DeepCopyInto(out *ValidatingWebhookApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ClientConfig != nil {
		in, out := &in.ClientConfig, &out.ClientConfig
		*out = new(WebhookClientConfigApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RuleWithOperationsApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(admissionregistrationv1.FailurePolicyType)
		**out = **in
	}
	if in.MatchPolicy != nil {
		in, out := &in.MatchPolicy, &out.MatchPolicy
		*out = new(admissionregistrationv1.MatchPolicyType)
		**out = **in
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectSelector != nil {
		in, out := &in.ObjectSelector, &out.ObjectSelector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.SideEffects != nil {
		in, out := &in.SideEffects, &out.SideEffects
		*out = new(admissionregistrationv1.SideEffectClass)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.AdmissionReviewVersions != nil {
		in, out := &in.AdmissionReviewVersions, &out.AdmissionReviewVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookApplyConfiguration.
func (in *ValidatingWebhookApplyConfiguration) DeepCopy() *ValidatingWebhookApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookConfigurationApplyConfiguration) DeepCopyInto(out *ValidatingWebhookConfigurationApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]ValidatingWebhookApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookConfigurationApplyConfiguration.
func (in *ValidatingWebhookConfigurationApplyConfiguration) DeepCopy() *ValidatingWebhookConfigurationApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookConfigurationApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookClientConfigApplyConfiguration) DeepCopyInto(out *WebhookClientConfigApplyConfiguration) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceReferenceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookClientConfigApplyConfiguration.
func (in *WebhookClientConfigApplyConfiguration) DeepCopy() *WebhookClientConfigApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebhookClientConfigApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}
