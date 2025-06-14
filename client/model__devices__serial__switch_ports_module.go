/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsModule Expansion module
type DevicesSerialSwitchPortsModule struct {
	// The model of the expansion module.
	Model *string `json:"model,omitempty"`
}

// NewDevicesSerialSwitchPortsModule instantiates a new DevicesSerialSwitchPortsModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsModule() *DevicesSerialSwitchPortsModule {
	this := DevicesSerialSwitchPortsModule{}
	return &this
}

// NewDevicesSerialSwitchPortsModuleWithDefaults instantiates a new DevicesSerialSwitchPortsModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsModuleWithDefaults() *DevicesSerialSwitchPortsModule {
	this := DevicesSerialSwitchPortsModule{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsModule) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsModule) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsModule) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DevicesSerialSwitchPortsModule) SetModel(v string) {
	o.Model = &v
}

func (o DevicesSerialSwitchPortsModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsModule struct {
	value *DevicesSerialSwitchPortsModule
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsModule) Get() *DevicesSerialSwitchPortsModule {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsModule) Set(val *DevicesSerialSwitchPortsModule) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsModule) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsModule(val *DevicesSerialSwitchPortsModule) *NullableDevicesSerialSwitchPortsModule {
	return &NullableDevicesSerialSwitchPortsModule{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


