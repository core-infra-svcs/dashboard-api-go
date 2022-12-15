/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialCellularSimsSimFailover SIM Failover settings.
type DevicesSerialCellularSimsSimFailover struct {
	// Failover to secondary SIM (optional)
	Enabled *bool `json:"enabled,omitempty"`
}

// NewDevicesSerialCellularSimsSimFailover instantiates a new DevicesSerialCellularSimsSimFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularSimsSimFailover() *DevicesSerialCellularSimsSimFailover {
	this := DevicesSerialCellularSimsSimFailover{}
	return &this
}

// NewDevicesSerialCellularSimsSimFailoverWithDefaults instantiates a new DevicesSerialCellularSimsSimFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularSimsSimFailoverWithDefaults() *DevicesSerialCellularSimsSimFailover {
	this := DevicesSerialCellularSimsSimFailover{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsSimFailover) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsSimFailover) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsSimFailover) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialCellularSimsSimFailover) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o DevicesSerialCellularSimsSimFailover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularSimsSimFailover struct {
	value *DevicesSerialCellularSimsSimFailover
	isSet bool
}

func (v NullableDevicesSerialCellularSimsSimFailover) Get() *DevicesSerialCellularSimsSimFailover {
	return v.value
}

func (v *NullableDevicesSerialCellularSimsSimFailover) Set(val *DevicesSerialCellularSimsSimFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularSimsSimFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularSimsSimFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularSimsSimFailover(val *DevicesSerialCellularSimsSimFailover) *NullableDevicesSerialCellularSimsSimFailover {
	return &NullableDevicesSerialCellularSimsSimFailover{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularSimsSimFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularSimsSimFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


