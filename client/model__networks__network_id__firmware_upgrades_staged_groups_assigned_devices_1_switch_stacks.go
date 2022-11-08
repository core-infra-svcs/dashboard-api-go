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

// NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks struct for NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks
type NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks struct {
	// ID of the Switch Stack
	Id string `json:"id"`
	// Name of the Switch Stack
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks(id string) *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacksWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacksWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) Get() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) Set(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1SwitchStacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


