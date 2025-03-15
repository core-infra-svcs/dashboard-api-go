/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200282 struct for InlineResponse200282
type InlineResponse200282 struct {
	// Database ID for the new entity entry.
	ImportId *string `json:"importId,omitempty"`
	Device *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice `json:"device,omitempty"`
}

// NewInlineResponse200282 instantiates a new InlineResponse200282 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200282() *InlineResponse200282 {
	this := InlineResponse200282{}
	return &this
}

// NewInlineResponse200282WithDefaults instantiates a new InlineResponse200282 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200282WithDefaults() *InlineResponse200282 {
	this := InlineResponse200282{}
	return &this
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *InlineResponse200282) GetImportId() string {
	if o == nil || isNil(o.ImportId) {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetImportIdOk() (*string, bool) {
	if o == nil || isNil(o.ImportId) {
    return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *InlineResponse200282) HasImportId() bool {
	if o != nil && !isNil(o.ImportId) {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *InlineResponse200282) SetImportId(v string) {
	o.ImportId = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200282) GetDevice() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice {
	if o == nil || isNil(o.Device) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200282) GetDeviceOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200282) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice and assigns it to the Device field.
func (o *InlineResponse200282) SetDevice(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevice) {
	o.Device = &v
}

func (o InlineResponse200282) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ImportId) {
		toSerialize["importId"] = o.ImportId
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200282 struct {
	value *InlineResponse200282
	isSet bool
}

func (v NullableInlineResponse200282) Get() *InlineResponse200282 {
	return v.value
}

func (v *NullableInlineResponse200282) Set(val *InlineResponse200282) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200282) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200282) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200282(val *InlineResponse200282) *NullableInlineResponse200282 {
	return &NullableInlineResponse200282{value: val, isSet: true}
}

func (v NullableInlineResponse200282) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200282) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


