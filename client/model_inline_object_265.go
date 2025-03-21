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

// InlineObject265 struct for InlineObject265
type InlineObject265 struct {
	// The serial number of the device to assign this license to. Set this to  null to unassign the license. If a different license is already active on the device, this parameter will control queueing/dequeuing this license.
	DeviceSerial *string `json:"deviceSerial,omitempty"`
}

// NewInlineObject265 instantiates a new InlineObject265 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject265() *InlineObject265 {
	this := InlineObject265{}
	return &this
}

// NewInlineObject265WithDefaults instantiates a new InlineObject265 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject265WithDefaults() *InlineObject265 {
	this := InlineObject265{}
	return &this
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *InlineObject265) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject265) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *InlineObject265) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *InlineObject265) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

func (o InlineObject265) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject265 struct {
	value *InlineObject265
	isSet bool
}

func (v NullableInlineObject265) Get() *InlineObject265 {
	return v.value
}

func (v *NullableInlineObject265) Set(val *InlineObject265) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject265) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject265) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject265(val *InlineObject265) *NullableInlineObject265 {
	return &NullableInlineObject265{value: val, isSet: true}
}

func (v NullableInlineObject265) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject265) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


