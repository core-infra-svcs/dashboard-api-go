/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesProductsWireless The network device to be updated
type NetworksNetworkIdFirmwareUpgradesProductsWireless struct {
	NextUpgrade *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade `json:"nextUpgrade,omitempty"`
	// Whether or not the network wants beta firmware
	ParticipateInNextBetaRelease *bool `json:"participateInNextBetaRelease,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWireless instantiates a new NetworksNetworkIdFirmwareUpgradesProductsWireless object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesProductsWireless() *NetworksNetworkIdFirmwareUpgradesProductsWireless {
	this := NetworksNetworkIdFirmwareUpgradesProductsWireless{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWirelessWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesProductsWireless object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesProductsWirelessWithDefaults() *NetworksNetworkIdFirmwareUpgradesProductsWireless {
	this := NetworksNetworkIdFirmwareUpgradesProductsWireless{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) GetNextUpgrade() NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) GetNextUpgradeOk() (*NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade and assigns it to the NextUpgrade field.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) SetNextUpgrade(v NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) {
	o.NextUpgrade = &v
}

// GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) GetParticipateInNextBetaRelease() bool {
	if o == nil || isNil(o.ParticipateInNextBetaRelease) {
		var ret bool
		return ret
	}
	return *o.ParticipateInNextBetaRelease
}

// GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) GetParticipateInNextBetaReleaseOk() (*bool, bool) {
	if o == nil || isNil(o.ParticipateInNextBetaRelease) {
    return nil, false
	}
	return o.ParticipateInNextBetaRelease, true
}

// HasParticipateInNextBetaRelease returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) HasParticipateInNextBetaRelease() bool {
	if o != nil && !isNil(o.ParticipateInNextBetaRelease) {
		return true
	}

	return false
}

// SetParticipateInNextBetaRelease gets a reference to the given bool and assigns it to the ParticipateInNextBetaRelease field.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWireless) SetParticipateInNextBetaRelease(v bool) {
	o.ParticipateInNextBetaRelease = &v
}

func (o NetworksNetworkIdFirmwareUpgradesProductsWireless) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	if !isNil(o.ParticipateInNextBetaRelease) {
		toSerialize["participateInNextBetaRelease"] = o.ParticipateInNextBetaRelease
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesProductsWireless struct {
	value *NetworksNetworkIdFirmwareUpgradesProductsWireless
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWireless) Get() *NetworksNetworkIdFirmwareUpgradesProductsWireless {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWireless) Set(val *NetworksNetworkIdFirmwareUpgradesProductsWireless) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWireless) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWireless) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesProductsWireless(val *NetworksNetworkIdFirmwareUpgradesProductsWireless) *NullableNetworksNetworkIdFirmwareUpgradesProductsWireless {
	return &NullableNetworksNetworkIdFirmwareUpgradesProductsWireless{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWireless) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWireless) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


