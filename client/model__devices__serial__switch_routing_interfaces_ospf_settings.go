/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchRoutingInterfacesOspfSettings IPv4 OSPF Settings
type DevicesSerialSwitchRoutingInterfacesOspfSettings struct {
	// Area id
	Area *string `json:"area,omitempty"`
	// OSPF Cost
	Cost *int32 `json:"cost,omitempty"`
	// Disable sending Hello packets on this interface's IPv4 area
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewDevicesSerialSwitchRoutingInterfacesOspfSettings instantiates a new DevicesSerialSwitchRoutingInterfacesOspfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchRoutingInterfacesOspfSettings() *DevicesSerialSwitchRoutingInterfacesOspfSettings {
	this := DevicesSerialSwitchRoutingInterfacesOspfSettings{}
	return &this
}

// NewDevicesSerialSwitchRoutingInterfacesOspfSettingsWithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesOspfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchRoutingInterfacesOspfSettingsWithDefaults() *DevicesSerialSwitchRoutingInterfacesOspfSettings {
	this := DevicesSerialSwitchRoutingInterfacesOspfSettings{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) GetArea() string {
	if o == nil || isNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) GetAreaOk() (*string, bool) {
	if o == nil || isNil(o.Area) {
    return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) GetCost() int32 {
	if o == nil || isNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) GetCostOk() (*int32, bool) {
	if o == nil || isNil(o.Cost) {
    return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) HasCost() bool {
	if o != nil && !isNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) GetIsPassiveEnabled() bool {
	if o == nil || isNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsPassiveEnabled) {
    return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) HasIsPassiveEnabled() bool {
	if o != nil && !isNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o DevicesSerialSwitchRoutingInterfacesOspfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !isNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !isNil(o.IsPassiveEnabled) {
		toSerialize["isPassiveEnabled"] = o.IsPassiveEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchRoutingInterfacesOspfSettings struct {
	value *DevicesSerialSwitchRoutingInterfacesOspfSettings
	isSet bool
}

func (v NullableDevicesSerialSwitchRoutingInterfacesOspfSettings) Get() *DevicesSerialSwitchRoutingInterfacesOspfSettings {
	return v.value
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings) Set(val *DevicesSerialSwitchRoutingInterfacesOspfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchRoutingInterfacesOspfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchRoutingInterfacesOspfSettings(val *DevicesSerialSwitchRoutingInterfacesOspfSettings) *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings {
	return &NullableDevicesSerialSwitchRoutingInterfacesOspfSettings{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchRoutingInterfacesOspfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


