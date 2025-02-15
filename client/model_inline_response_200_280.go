/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200280 struct for InlineResponse200280
type InlineResponse200280 struct {
	// Database ID for the new entity entry.
	ImportId *string `json:"importId,omitempty"`
	Device *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice `json:"device,omitempty"`
}

// NewInlineResponse200280 instantiates a new InlineResponse200280 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200280() *InlineResponse200280 {
	this := InlineResponse200280{}
	return &this
}

// NewInlineResponse200280WithDefaults instantiates a new InlineResponse200280 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200280WithDefaults() *InlineResponse200280 {
	this := InlineResponse200280{}
	return &this
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *InlineResponse200280) GetImportId() string {
	if o == nil || isNil(o.ImportId) {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280) GetImportIdOk() (*string, bool) {
	if o == nil || isNil(o.ImportId) {
    return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *InlineResponse200280) HasImportId() bool {
	if o != nil && !isNil(o.ImportId) {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *InlineResponse200280) SetImportId(v string) {
	o.ImportId = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200280) GetDevice() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280) GetDeviceOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200280) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice and assigns it to the Device field.
func (o *InlineResponse200280) SetDevice(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) {
	o.Device = &v
}

func (o InlineResponse200280) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ImportId) {
		toSerialize["importId"] = o.ImportId
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200280 struct {
	value *InlineResponse200280
	isSet bool
}

func (v NullableInlineResponse200280) Get() *InlineResponse200280 {
	return v.value
}

func (v *NullableInlineResponse200280) Set(val *InlineResponse200280) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200280) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200280) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200280(val *InlineResponse200280) *NullableInlineResponse200280 {
	return &NullableInlineResponse200280{value: val, isSet: true}
}

func (v NullableInlineResponse200280) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200280) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


