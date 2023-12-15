/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel TLS Related Parameters
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel struct {
	// Name of the configured TLS certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Name of the configured TLS tunnel
	Name *string `json:"name,omitempty"`
	// Number of the configured Loopback Interface used for TLS overlay
	LoopbackNumber *int32 `json:"loopbackNumber,omitempty"`
	// Number of the vlan expected to be used to connect to the cloud
	LocalInterface *int32 `json:"localInterface,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnelWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnelWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel{}
	return &this
}

// GetCertificateName returns the CertificateName field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetCertificateName() string {
	if o == nil || isNil(o.CertificateName) {
		var ret string
		return ret
	}
	return *o.CertificateName
}

// GetCertificateNameOk returns a tuple with the CertificateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetCertificateNameOk() (*string, bool) {
	if o == nil || isNil(o.CertificateName) {
    return nil, false
	}
	return o.CertificateName, true
}

// HasCertificateName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) HasCertificateName() bool {
	if o != nil && !isNil(o.CertificateName) {
		return true
	}

	return false
}

// SetCertificateName gets a reference to the given string and assigns it to the CertificateName field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) SetCertificateName(v string) {
	o.CertificateName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) SetName(v string) {
	o.Name = &v
}

// GetLoopbackNumber returns the LoopbackNumber field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetLoopbackNumber() int32 {
	if o == nil || isNil(o.LoopbackNumber) {
		var ret int32
		return ret
	}
	return *o.LoopbackNumber
}

// GetLoopbackNumberOk returns a tuple with the LoopbackNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetLoopbackNumberOk() (*int32, bool) {
	if o == nil || isNil(o.LoopbackNumber) {
    return nil, false
	}
	return o.LoopbackNumber, true
}

// HasLoopbackNumber returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) HasLoopbackNumber() bool {
	if o != nil && !isNil(o.LoopbackNumber) {
		return true
	}

	return false
}

// SetLoopbackNumber gets a reference to the given int32 and assigns it to the LoopbackNumber field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) SetLoopbackNumber(v int32) {
	o.LoopbackNumber = &v
}

// GetLocalInterface returns the LocalInterface field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetLocalInterface() int32 {
	if o == nil || isNil(o.LocalInterface) {
		var ret int32
		return ret
	}
	return *o.LocalInterface
}

// GetLocalInterfaceOk returns a tuple with the LocalInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) GetLocalInterfaceOk() (*int32, bool) {
	if o == nil || isNil(o.LocalInterface) {
    return nil, false
	}
	return o.LocalInterface, true
}

// HasLocalInterface returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) HasLocalInterface() bool {
	if o != nil && !isNil(o.LocalInterface) {
		return true
	}

	return false
}

// SetLocalInterface gets a reference to the given int32 and assigns it to the LocalInterface field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) SetLocalInterface(v int32) {
	o.LocalInterface = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CertificateName) {
		toSerialize["certificateName"] = o.CertificateName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.LoopbackNumber) {
		toSerialize["loopbackNumber"] = o.LoopbackNumber
	}
	if !isNil(o.LocalInterface) {
		toSerialize["localInterface"] = o.LocalInterface
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


