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

// NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 The Staged Upgrade Group
type NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 struct {
	// ID of the Staged Upgrade Group
	Id string `json:"id"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 instantiates a new NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1(id string) *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 {
	this := NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1WithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1WithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 {
	this := NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) SetId(v string) {
	o.Id = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) Get() *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) Set(val *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1(val *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1 {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


