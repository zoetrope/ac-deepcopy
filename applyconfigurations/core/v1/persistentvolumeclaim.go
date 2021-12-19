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

package v1

import (
	"errors"

	internal "github.com/zoetrope/ssa-helper/applyconfigurations/internal"
	v1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PersistentVolumeClaimApplyConfiguration represents an declarative configuration of the PersistentVolumeClaim type for use
// with apply.
type PersistentVolumeClaimApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *PersistentVolumeClaimSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *PersistentVolumeClaimStatusApplyConfiguration `json:"status,omitempty"`
}

// PersistentVolumeClaim constructs an declarative configuration of the PersistentVolumeClaim type for use with
// apply.
func PersistentVolumeClaim(name, namespace string) *PersistentVolumeClaimApplyConfiguration {
	b := &PersistentVolumeClaimApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("PersistentVolumeClaim")
	b.WithAPIVersion("v1")
	return b
}

// ExtractPersistentVolumeClaim extracts the applied configuration owned by fieldManager from
// persistentVolumeClaim. If no managedFields are found in persistentVolumeClaim for fieldManager, a
// PersistentVolumeClaimApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// persistentVolumeClaim must be a unmodified PersistentVolumeClaim API object that was retrieved from the Kubernetes API.
// ExtractPersistentVolumeClaim provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractPersistentVolumeClaim(persistentVolumeClaim *apicorev1.PersistentVolumeClaim, fieldManager string) (*PersistentVolumeClaimApplyConfiguration, error) {
	return extractPersistentVolumeClaim(persistentVolumeClaim, fieldManager, "")
}

// ExtractPersistentVolumeClaimStatus is the same as ExtractPersistentVolumeClaim except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractPersistentVolumeClaimStatus(persistentVolumeClaim *apicorev1.PersistentVolumeClaim, fieldManager string) (*PersistentVolumeClaimApplyConfiguration, error) {
	return extractPersistentVolumeClaim(persistentVolumeClaim, fieldManager, "status")
}

func extractPersistentVolumeClaim(persistentVolumeClaim *apicorev1.PersistentVolumeClaim, fieldManager string, subresource string) (*PersistentVolumeClaimApplyConfiguration, error) {
	b := &PersistentVolumeClaimApplyConfiguration{}
	err := managedfields.ExtractInto(persistentVolumeClaim, internal.Parser().Type("io.k8s.api.core.v1.PersistentVolumeClaim"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(persistentVolumeClaim.Name)
	b.WithNamespace(persistentVolumeClaim.Namespace)

	b.WithKind("PersistentVolumeClaim")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithKind(value string) *PersistentVolumeClaimApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithAPIVersion(value string) *PersistentVolumeClaimApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithName(value string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithGenerateName(value string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithNamespace(value string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithSelfLink(value string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithUID(value types.UID) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithResourceVersion(value string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithGeneration(value int64) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithCreationTimestamp(value metav1.Time) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *PersistentVolumeClaimApplyConfiguration) WithLabels(entries map[string]string) *PersistentVolumeClaimApplyConfiguration {
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
func (b *PersistentVolumeClaimApplyConfiguration) WithAnnotations(entries map[string]string) *PersistentVolumeClaimApplyConfiguration {
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
func (b *PersistentVolumeClaimApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *PersistentVolumeClaimApplyConfiguration {
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
func (b *PersistentVolumeClaimApplyConfiguration) WithFinalizers(values ...string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithClusterName(value string) *PersistentVolumeClaimApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *PersistentVolumeClaimApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithSpec(value *PersistentVolumeClaimSpecApplyConfiguration) *PersistentVolumeClaimApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PersistentVolumeClaimApplyConfiguration) WithStatus(value *PersistentVolumeClaimStatusApplyConfiguration) *PersistentVolumeClaimApplyConfiguration {
	b.Status = value
	return b
}
func (b *PersistentVolumeClaimApplyConfiguration) Original() client.Object {
	return &apicorev1.PersistentVolumeClaim{}
}

func (b *PersistentVolumeClaimApplyConfiguration) Extract(obj client.Object, fieldManager string, subresource string) (*PersistentVolumeClaimApplyConfiguration, error) {
	return extractPersistentVolumeClaim(obj.(*apicorev1.PersistentVolumeClaim), fieldManager, subresource)
}
func (b *PersistentVolumeClaimApplyConfiguration) ObjectKey() (client.ObjectKey, error) {
	if b.Namespace == nil {
		return client.ObjectKey{}, errors.New("The PersistentVolumeClaimApplyConfiguration namespace should not be empty.")
	}
	if b.Name == nil {
		return client.ObjectKey{}, errors.New("The PersistentVolumeClaimApplyConfiguration name should not be empty.")
	}
	return client.ObjectKey{
		Name:      *b.Name,
		Namespace: *b.Namespace,
	}, nil
}
