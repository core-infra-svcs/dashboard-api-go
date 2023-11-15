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

// NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst Version information for the switch network being upgraded
type NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst struct {
	NextUpgrade *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) GetNextUpgrade() NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) GetNextUpgradeOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade and assigns it to the NextUpgrade field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) SetNextUpgrade(v NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) {
	o.NextUpgrade = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalyst) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


