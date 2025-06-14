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

// NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade The next upgrade version for the switch network
type NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade struct {
	ToVersion *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade{}
	return &this
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) GetToVersion() NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) SetToVersion(v NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchCatalystNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


