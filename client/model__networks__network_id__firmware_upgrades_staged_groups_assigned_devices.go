/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices The devices and Switch Stacks assigned to the Group
type NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices struct {
	// Data Array of Devices containing the name and serial
	Devices []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices `json:"devices,omitempty"`
	// Data Array of Switch Stacks containing the name and id
	SwitchStacks []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks `json:"switchStacks,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) GetDevices() []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices {
	if o == nil || isNil(o.Devices) {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) GetDevicesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices and assigns it to the Devices field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) SetDevices(v []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) {
	o.Devices = v
}

// GetSwitchStacks returns the SwitchStacks field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) GetSwitchStacks() []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks {
	if o == nil || isNil(o.SwitchStacks) {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks
		return ret
	}
	return o.SwitchStacks
}

// GetSwitchStacksOk returns a tuple with the SwitchStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) GetSwitchStacksOk() ([]NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks, bool) {
	if o == nil || isNil(o.SwitchStacks) {
    return nil, false
	}
	return o.SwitchStacks, true
}

// HasSwitchStacks returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) HasSwitchStacks() bool {
	if o != nil && !isNil(o.SwitchStacks) {
		return true
	}

	return false
}

// SetSwitchStacks gets a reference to the given []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks and assigns it to the SwitchStacks field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) SetSwitchStacks(v []NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) {
	o.SwitchStacks = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.SwitchStacks) {
		toSerialize["switchStacks"] = o.SwitchStacks
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) Get() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) Set(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


