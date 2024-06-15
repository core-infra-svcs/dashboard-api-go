/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion The version to be updated to for switch devices
type NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	// The version ID
	Id string `json:"id"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion(id string) *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersionWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersionWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) SetId(v string) {
	o.Id = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


