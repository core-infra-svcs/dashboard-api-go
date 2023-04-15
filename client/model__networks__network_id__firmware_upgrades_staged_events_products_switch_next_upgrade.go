/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade The next upgrade version for the switch network
type NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade{}
	return &this
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) GetToVersion() NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) SetToVersion(v NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


