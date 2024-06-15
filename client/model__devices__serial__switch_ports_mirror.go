/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsMirror Port mirror
type DevicesSerialSwitchPortsMirror struct {
	// The port mirror mode. Can be one of ('Destination port', 'Source port' or 'Not mirroring traffic').
	Mode *string `json:"mode,omitempty"`
}

// NewDevicesSerialSwitchPortsMirror instantiates a new DevicesSerialSwitchPortsMirror object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsMirror() *DevicesSerialSwitchPortsMirror {
	this := DevicesSerialSwitchPortsMirror{}
	return &this
}

// NewDevicesSerialSwitchPortsMirrorWithDefaults instantiates a new DevicesSerialSwitchPortsMirror object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsMirrorWithDefaults() *DevicesSerialSwitchPortsMirror {
	this := DevicesSerialSwitchPortsMirror{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsMirror) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsMirror) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsMirror) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DevicesSerialSwitchPortsMirror) SetMode(v string) {
	o.Mode = &v
}

func (o DevicesSerialSwitchPortsMirror) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsMirror struct {
	value *DevicesSerialSwitchPortsMirror
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsMirror) Get() *DevicesSerialSwitchPortsMirror {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsMirror) Set(val *DevicesSerialSwitchPortsMirror) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsMirror) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsMirror) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsMirror(val *DevicesSerialSwitchPortsMirror) *NullableDevicesSerialSwitchPortsMirror {
	return &NullableDevicesSerialSwitchPortsMirror{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsMirror) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsMirror) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


