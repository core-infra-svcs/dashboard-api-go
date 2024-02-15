/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchRoutingInterfacesOspfSettings1 The OSPF routing settings of the interface.
type DevicesSerialSwitchRoutingInterfacesOspfSettings1 struct {
	// The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPF area. Defaults to 'disabled'.
	Area *string `json:"area,omitempty"`
	// The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	Cost *int32 `json:"cost,omitempty"`
	// When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewDevicesSerialSwitchRoutingInterfacesOspfSettings1 instantiates a new DevicesSerialSwitchRoutingInterfacesOspfSettings1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchRoutingInterfacesOspfSettings1() *DevicesSerialSwitchRoutingInterfacesOspfSettings1 {
	this := DevicesSerialSwitchRoutingInterfacesOspfSettings1{}
	return &this
}

// NewDevicesSerialSwitchRoutingInterfacesOspfSettings1WithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesOspfSettings1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchRoutingInterfacesOspfSettings1WithDefaults() *DevicesSerialSwitchRoutingInterfacesOspfSettings1 {
	this := DevicesSerialSwitchRoutingInterfacesOspfSettings1{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) GetArea() string {
	if o == nil || isNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) GetAreaOk() (*string, bool) {
	if o == nil || isNil(o.Area) {
    return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) GetCost() int32 {
	if o == nil || isNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) GetCostOk() (*int32, bool) {
	if o == nil || isNil(o.Cost) {
    return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) HasCost() bool {
	if o != nil && !isNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) GetIsPassiveEnabled() bool {
	if o == nil || isNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsPassiveEnabled) {
    return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) HasIsPassiveEnabled() bool {
	if o != nil && !isNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *DevicesSerialSwitchRoutingInterfacesOspfSettings1) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o DevicesSerialSwitchRoutingInterfacesOspfSettings1) MarshalJSON() ([]byte, error) {
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

type NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1 struct {
	value *DevicesSerialSwitchRoutingInterfacesOspfSettings1
	isSet bool
}

func (v NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1) Get() *DevicesSerialSwitchRoutingInterfacesOspfSettings1 {
	return v.value
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1) Set(val *DevicesSerialSwitchRoutingInterfacesOspfSettings1) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchRoutingInterfacesOspfSettings1(val *DevicesSerialSwitchRoutingInterfacesOspfSettings1) *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1 {
	return &NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesOspfSettings1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


