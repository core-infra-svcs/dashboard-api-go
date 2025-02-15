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

// DevicesSerialSwitchPortsDot3az dot3az settings for the port
type DevicesSerialSwitchPortsDot3az struct {
	// The Energy Efficient Ethernet status of the switch port.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewDevicesSerialSwitchPortsDot3az instantiates a new DevicesSerialSwitchPortsDot3az object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsDot3az() *DevicesSerialSwitchPortsDot3az {
	this := DevicesSerialSwitchPortsDot3az{}
	return &this
}

// NewDevicesSerialSwitchPortsDot3azWithDefaults instantiates a new DevicesSerialSwitchPortsDot3az object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsDot3azWithDefaults() *DevicesSerialSwitchPortsDot3az {
	this := DevicesSerialSwitchPortsDot3az{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsDot3az) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsDot3az) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsDot3az) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialSwitchPortsDot3az) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o DevicesSerialSwitchPortsDot3az) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsDot3az struct {
	value *DevicesSerialSwitchPortsDot3az
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsDot3az) Get() *DevicesSerialSwitchPortsDot3az {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsDot3az) Set(val *DevicesSerialSwitchPortsDot3az) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsDot3az) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsDot3az) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsDot3az(val *DevicesSerialSwitchPortsDot3az) *NullableDevicesSerialSwitchPortsDot3az {
	return &NullableDevicesSerialSwitchPortsDot3az{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsDot3az) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsDot3az) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


