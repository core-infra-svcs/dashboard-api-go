/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion The version to be updated to for switch catalyst devices
type NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion struct {
	// The version ID
	Id string `json:"id"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion(id string) *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersionWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersionWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) SetId(v string) {
	o.Id = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


