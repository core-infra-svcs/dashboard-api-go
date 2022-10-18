/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedEventsProducts Contains information about the network to update
type NetworksNetworkIdFirmwareUpgradesStagedEventsProducts struct {
	Switch *NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch `json:"switch,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProducts instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProducts() *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProducts{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsProductsWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsProducts{}
	return &this
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) GetSwitch() NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch {
	if o == nil || o.Switch == nil {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) GetSwitchOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch, bool) {
	if o == nil || o.Switch == nil {
		return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) HasSwitch() bool {
	if o != nil && o.Switch != nil {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch and assigns it to the Switch field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) SetSwitch(v NetworksNetworkIdFirmwareUpgradesStagedEventsProductsSwitch) {
	o.Switch = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Switch != nil {
		toSerialize["switch"] = o.Switch
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts(val *NetworksNetworkIdFirmwareUpgradesStagedEventsProducts) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


