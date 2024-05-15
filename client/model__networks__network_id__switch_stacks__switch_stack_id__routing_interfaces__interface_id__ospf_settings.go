/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings The OSPF routing settings of the interface.
type NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings struct {
	// The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an existing OSPF area.
	Area *string `json:"area,omitempty"`
	// The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.
	Cost *int32 `json:"cost,omitempty"`
	// When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings instantiates a new NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings {
	this := NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings{}
	return &this
}

// NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettingsWithDefaults instantiates a new NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettingsWithDefaults() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings {
	this := NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) GetArea() string {
	if o == nil || isNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) GetAreaOk() (*string, bool) {
	if o == nil || isNil(o.Area) {
    return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) GetCost() int32 {
	if o == nil || isNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) GetCostOk() (*int32, bool) {
	if o == nil || isNil(o.Cost) {
    return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) HasCost() bool {
	if o != nil && !isNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) GetIsPassiveEnabled() bool {
	if o == nil || isNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsPassiveEnabled) {
    return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) HasIsPassiveEnabled() bool {
	if o != nil && !isNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings struct {
	value *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) Get() *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) Set(val *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings(val *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings {
	return &NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesInterfaceIdOspfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


