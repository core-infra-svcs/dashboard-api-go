/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject211 struct for InlineObject211
type InlineObject211 struct {
	// The serial number of the device to assign this license to. Set this to  null to unassign the license. If a different license is already active on the device, this parameter will control queueing/dequeuing this license.
	DeviceSerial *string `json:"deviceSerial,omitempty"`
}

// NewInlineObject211 instantiates a new InlineObject211 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject211() *InlineObject211 {
	this := InlineObject211{}
	return &this
}

// NewInlineObject211WithDefaults instantiates a new InlineObject211 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject211WithDefaults() *InlineObject211 {
	this := InlineObject211{}
	return &this
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *InlineObject211) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject211) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *InlineObject211) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *InlineObject211) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

func (o InlineObject211) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject211 struct {
	value *InlineObject211
	isSet bool
}

func (v NullableInlineObject211) Get() *InlineObject211 {
	return v.value
}

func (v *NullableInlineObject211) Set(val *InlineObject211) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject211) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject211) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject211(val *InlineObject211) *NullableInlineObject211 {
	return &NullableInlineObject211{value: val, isSet: true}
}

func (v NullableInlineObject211) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject211) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


