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

// NetworksNetworkIdFirmwareUpgradesStagedStagesJson struct for NetworksNetworkIdFirmwareUpgradesStagedStagesJson
type NetworksNetworkIdFirmwareUpgradesStagedStagesJson struct {
	Group *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 `json:"group,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedStagesJson instantiates a new NetworksNetworkIdFirmwareUpgradesStagedStagesJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedStagesJson() *NetworksNetworkIdFirmwareUpgradesStagedStagesJson {
	this := NetworksNetworkIdFirmwareUpgradesStagedStagesJson{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedStagesJsonWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedStagesJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedStagesJsonWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedStagesJson {
	this := NetworksNetworkIdFirmwareUpgradesStagedStagesJson{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesJson) GetGroup() NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 {
	if o == nil || isNil(o.Group) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesJson) GetGroupOk() (*NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesJson) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 and assigns it to the Group field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesJson) SetGroup(v NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) {
	o.Group = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedStagesJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedStagesJson
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson) Get() *NetworksNetworkIdFirmwareUpgradesStagedStagesJson {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson) Set(val *NetworksNetworkIdFirmwareUpgradesStagedStagesJson) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson(val *NetworksNetworkIdFirmwareUpgradesStagedStagesJson) *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


