/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	"errors"

	internal "github.com/zoetrope/ac-deepcopy/applyconfigurations/internal"
	v1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
	imagepolicyv1alpha1 "k8s.io/api/imagepolicy/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ImageReviewApplyConfiguration represents an declarative configuration of the ImageReview type for use
// with apply.
type ImageReviewApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ImageReviewSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *ImageReviewStatusApplyConfiguration `json:"status,omitempty"`
}

// ImageReview constructs an declarative configuration of the ImageReview type for use with
// apply.
func ImageReview(name string) *ImageReviewApplyConfiguration {
	b := &ImageReviewApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ImageReview")
	b.WithAPIVersion("imagepolicy.k8s.io/v1alpha1")
	return b
}

// ExtractImageReview extracts the applied configuration owned by fieldManager from
// imageReview. If no managedFields are found in imageReview for fieldManager, a
// ImageReviewApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// imageReview must be a unmodified ImageReview API object that was retrieved from the Kubernetes API.
// ExtractImageReview provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractImageReview(imageReview *imagepolicyv1alpha1.ImageReview, fieldManager string) (*ImageReviewApplyConfiguration, error) {
	return extractImageReview(imageReview, fieldManager, "")
}

// ExtractImageReviewStatus is the same as ExtractImageReview except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractImageReviewStatus(imageReview *imagepolicyv1alpha1.ImageReview, fieldManager string) (*ImageReviewApplyConfiguration, error) {
	return extractImageReview(imageReview, fieldManager, "status")
}

func extractImageReview(imageReview *imagepolicyv1alpha1.ImageReview, fieldManager string, subresource string) (*ImageReviewApplyConfiguration, error) {
	b := &ImageReviewApplyConfiguration{}
	err := managedfields.ExtractInto(imageReview, internal.Parser().Type("io.k8s.api.imagepolicy.v1alpha1.ImageReview"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(imageReview.Name)

	b.WithKind("ImageReview")
	b.WithAPIVersion("imagepolicy.k8s.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithKind(value string) *ImageReviewApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithAPIVersion(value string) *ImageReviewApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithName(value string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithGenerateName(value string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithNamespace(value string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithSelfLink(value string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithUID(value types.UID) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithResourceVersion(value string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithGeneration(value int64) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ImageReviewApplyConfiguration) WithLabels(entries map[string]string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ImageReviewApplyConfiguration) WithAnnotations(entries map[string]string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ImageReviewApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ImageReviewApplyConfiguration) WithFinalizers(values ...string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithClusterName(value string) *ImageReviewApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *ImageReviewApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithSpec(value *ImageReviewSpecApplyConfiguration) *ImageReviewApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ImageReviewApplyConfiguration) WithStatus(value *ImageReviewStatusApplyConfiguration) *ImageReviewApplyConfiguration {
	b.Status = value
	return b
}
func (b *ImageReviewApplyConfiguration) Original() client.Object {
	return &imagepolicyv1alpha1.ImageReview{}
}

func (b *ImageReviewApplyConfiguration) Extract(obj client.Object, fieldManager string, subresource string) (*ImageReviewApplyConfiguration, error) {
	return extractImageReview(obj.(*imagepolicyv1alpha1.ImageReview), fieldManager, subresource)
}
func (b *ImageReviewApplyConfiguration) ObjectKey() (client.ObjectKey, error) {
	if b.Name == nil {
		return client.ObjectKey{}, errors.New("The ImageReviewApplyConfiguration name should not be empty.")
	}
	return client.ObjectKey{
		Name: *b.Name,
	}, nil
}
