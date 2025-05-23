/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200270Devices struct for InlineResponse200270Devices
type InlineResponse200270Devices struct {
	Device *InlineResponse200270Device `json:"device,omitempty"`
}

// NewInlineResponse200270Devices instantiates a new InlineResponse200270Devices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200270Devices() *InlineResponse200270Devices {
	this := InlineResponse200270Devices{}
	return &this
}

// NewInlineResponse200270DevicesWithDefaults instantiates a new InlineResponse200270Devices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200270DevicesWithDefaults() *InlineResponse200270Devices {
	this := InlineResponse200270Devices{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InlineResponse200270Devices) GetDevice() InlineResponse200270Device {
	if o == nil || isNil(o.Device) {
		var ret InlineResponse200270Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Devices) GetDeviceOk() (*InlineResponse200270Device, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InlineResponse200270Devices) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given InlineResponse200270Device and assigns it to the Device field.
func (o *InlineResponse200270Devices) SetDevice(v InlineResponse200270Device) {
	o.Device = &v
}

func (o InlineResponse200270Devices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200270Devices struct {
	value *InlineResponse200270Devices
	isSet bool
}

func (v NullableInlineResponse200270Devices) Get() *InlineResponse200270Devices {
	return v.value
}

func (v *NullableInlineResponse200270Devices) Set(val *InlineResponse200270Devices) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200270Devices) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200270Devices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200270Devices(val *InlineResponse200270Devices) *NullableInlineResponse200270Devices {
	return &NullableInlineResponse200270Devices{value: val, isSet: true}
}

func (v NullableInlineResponse200270Devices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200270Devices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


