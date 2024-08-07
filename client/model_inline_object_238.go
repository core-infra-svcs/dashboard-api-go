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

// InlineObject238 struct for InlineObject238
type InlineObject238 struct {
	// A set of devices to import (or update)
	Devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices `json:"devices"`
}

// NewInlineObject238 instantiates a new InlineObject238 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject238(devices []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) *InlineObject238 {
	this := InlineObject238{}
	this.Devices = devices
	return &this
}

// NewInlineObject238WithDefaults instantiates a new InlineObject238 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject238WithDefaults() *InlineObject238 {
	this := InlineObject238{}
	return &this
}

// GetDevices returns the Devices field value
func (o *InlineObject238) GetDevices() []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices {
	if o == nil {
		var ret []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *InlineObject238) GetDevicesOk() ([]OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices, bool) {
	if o == nil {
    return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *InlineObject238) SetDevices(v []OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices) {
	o.Devices = v
}

func (o InlineObject238) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject238 struct {
	value *InlineObject238
	isSet bool
}

func (v NullableInlineObject238) Get() *InlineObject238 {
	return v.value
}

func (v *NullableInlineObject238) Set(val *InlineObject238) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject238) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject238) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject238(val *InlineObject238) *NullableInlineObject238 {
	return &NullableInlineObject238{value: val, isSet: true}
}

func (v NullableInlineObject238) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject238) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


