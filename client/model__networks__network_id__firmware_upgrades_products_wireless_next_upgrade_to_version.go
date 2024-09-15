/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion The version to be updated to
type NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion struct {
	// The version ID
	Id *string `json:"id,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion instantiates a new NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion() *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion {
	this := NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersionWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersionWithDefaults() *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion {
	this := NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

func (o NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion struct {
	value *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) Get() *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) Set(val *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion(val *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion {
	return &NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


