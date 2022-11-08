/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialCellularSimsSims struct for DevicesSerialCellularSimsSims
type DevicesSerialCellularSimsSims struct {
	// SIM slot being configured. Must be 'sim1' on single-sim devices.
	Slot *string `json:"slot,omitempty"`
	// If true, this SIM is used for boot. Must be true on single-sim devices.
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// APN configurations. If empty, the default APN will be used.
	Apns []DevicesSerialCellularSimsApns `json:"apns,omitempty"`
}

// NewDevicesSerialCellularSimsSims instantiates a new DevicesSerialCellularSimsSims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularSimsSims() *DevicesSerialCellularSimsSims {
	this := DevicesSerialCellularSimsSims{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewDevicesSerialCellularSimsSimsWithDefaults instantiates a new DevicesSerialCellularSimsSims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularSimsSimsWithDefaults() *DevicesSerialCellularSimsSims {
	this := DevicesSerialCellularSimsSims{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsSims) GetSlot() string {
	if o == nil || isNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsSims) GetSlotOk() (*string, bool) {
	if o == nil || isNil(o.Slot) {
    return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsSims) HasSlot() bool {
	if o != nil && !isNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *DevicesSerialCellularSimsSims) SetSlot(v string) {
	o.Slot = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsSims) GetIsPrimary() bool {
	if o == nil || isNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsSims) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrimary) {
    return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsSims) HasIsPrimary() bool {
	if o != nil && !isNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *DevicesSerialCellularSimsSims) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetApns returns the Apns field value if set, zero value otherwise.
func (o *DevicesSerialCellularSimsSims) GetApns() []DevicesSerialCellularSimsApns {
	if o == nil || isNil(o.Apns) {
		var ret []DevicesSerialCellularSimsApns
		return ret
	}
	return o.Apns
}

// GetApnsOk returns a tuple with the Apns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularSimsSims) GetApnsOk() ([]DevicesSerialCellularSimsApns, bool) {
	if o == nil || isNil(o.Apns) {
    return nil, false
	}
	return o.Apns, true
}

// HasApns returns a boolean if a field has been set.
func (o *DevicesSerialCellularSimsSims) HasApns() bool {
	if o != nil && !isNil(o.Apns) {
		return true
	}

	return false
}

// SetApns gets a reference to the given []DevicesSerialCellularSimsApns and assigns it to the Apns field.
func (o *DevicesSerialCellularSimsSims) SetApns(v []DevicesSerialCellularSimsApns) {
	o.Apns = v
}

func (o DevicesSerialCellularSimsSims) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !isNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !isNil(o.Apns) {
		toSerialize["apns"] = o.Apns
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularSimsSims struct {
	value *DevicesSerialCellularSimsSims
	isSet bool
}

func (v NullableDevicesSerialCellularSimsSims) Get() *DevicesSerialCellularSimsSims {
	return v.value
}

func (v *NullableDevicesSerialCellularSimsSims) Set(val *DevicesSerialCellularSimsSims) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularSimsSims) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularSimsSims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularSimsSims(val *DevicesSerialCellularSimsSims) *NullableDevicesSerialCellularSimsSims {
	return &NullableDevicesSerialCellularSimsSims{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularSimsSims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularSimsSims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

