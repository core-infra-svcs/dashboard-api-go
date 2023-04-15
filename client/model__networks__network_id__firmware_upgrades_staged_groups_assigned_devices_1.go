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

// NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 The devices and Switch Stacks assigned to the Group
type NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 struct {
	// Data Array of Devices containing the name and serial
	Devices []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices `json:"devices,omitempty"`
	// Data Array of Switch Stacks containing the name and id
	SwitchStacks []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks `json:"switchStacks,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1WithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1WithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) GetDevices() []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices {
	if o == nil || isNil(o.Devices) {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) GetDevicesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices and assigns it to the Devices field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) SetDevices(v []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) {
	o.Devices = v
}

// GetSwitchStacks returns the SwitchStacks field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) GetSwitchStacks() []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks {
	if o == nil || isNil(o.SwitchStacks) {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks
		return ret
	}
	return o.SwitchStacks
}

// GetSwitchStacksOk returns a tuple with the SwitchStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) GetSwitchStacksOk() ([]NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks, bool) {
	if o == nil || isNil(o.SwitchStacks) {
    return nil, false
	}
	return o.SwitchStacks, true
}

// HasSwitchStacks returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) HasSwitchStacks() bool {
	if o != nil && !isNil(o.SwitchStacks) {
		return true
	}

	return false
}

// SetSwitchStacks gets a reference to the given []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks and assigns it to the SwitchStacks field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) SetSwitchStacks(v []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) {
	o.SwitchStacks = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.SwitchStacks) {
		toSerialize["switchStacks"] = o.SwitchStacks
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) Get() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) Set(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


