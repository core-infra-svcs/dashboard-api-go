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

// NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch Version information for the switch network being upgraded
type NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) GetNextUpgrade() NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) GetNextUpgradeOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade and assigns it to the NextUpgrade field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) SetNextUpgrade(v NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) {
	o.NextUpgrade = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


