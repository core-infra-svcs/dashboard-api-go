/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject237 struct for InlineObject237
type InlineObject237 struct {
	// A set of device imports to commit
	Devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices `json:"devices"`
}

// NewInlineObject237 instantiates a new InlineObject237 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject237(devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices) *InlineObject237 {
	this := InlineObject237{}
	this.Devices = devices
	return &this
}

// NewInlineObject237WithDefaults instantiates a new InlineObject237 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject237WithDefaults() *InlineObject237 {
	this := InlineObject237{}
	return &this
}

// GetDevices returns the Devices field value
func (o *InlineObject237) GetDevices() []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices {
	if o == nil {
		var ret []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *InlineObject237) GetDevicesOk() ([]OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices, bool) {
	if o == nil {
    return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *InlineObject237) SetDevices(v []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices) {
	o.Devices = v
}

func (o InlineObject237) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject237 struct {
	value *InlineObject237
	isSet bool
}

func (v NullableInlineObject237) Get() *InlineObject237 {
	return v.value
}

func (v *NullableInlineObject237) Set(val *InlineObject237) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject237) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject237) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject237(val *InlineObject237) *NullableInlineObject237 {
	return &NullableInlineObject237{value: val, isSet: true}
}

func (v NullableInlineObject237) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject237) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


