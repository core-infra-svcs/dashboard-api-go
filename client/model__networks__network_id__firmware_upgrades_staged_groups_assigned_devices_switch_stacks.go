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

// NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks struct for NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks
type NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks struct {
	// ID of the Switch Stack
	Id *string `json:"id,omitempty"`
	// Name of the Switch Stack
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacksWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacksWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) Get() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) Set(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesSwitchStacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


