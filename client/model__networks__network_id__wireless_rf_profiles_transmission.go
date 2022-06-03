/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesTransmission Settings related to radio transmission.
type NetworksNetworkIdWirelessRfProfilesTransmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesTransmission instantiates a new NetworksNetworkIdWirelessRfProfilesTransmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesTransmission() *NetworksNetworkIdWirelessRfProfilesTransmission {
	this := NetworksNetworkIdWirelessRfProfilesTransmission{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesTransmissionWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesTransmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesTransmissionWithDefaults() *NetworksNetworkIdWirelessRfProfilesTransmission {
	this := NetworksNetworkIdWirelessRfProfilesTransmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesTransmission) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesTransmission) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesTransmission) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessRfProfilesTransmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdWirelessRfProfilesTransmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesTransmission struct {
	value *NetworksNetworkIdWirelessRfProfilesTransmission
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesTransmission) Get() *NetworksNetworkIdWirelessRfProfilesTransmission {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesTransmission) Set(val *NetworksNetworkIdWirelessRfProfilesTransmission) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesTransmission) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesTransmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesTransmission(val *NetworksNetworkIdWirelessRfProfilesTransmission) *NullableNetworksNetworkIdWirelessRfProfilesTransmission {
	return &NullableNetworksNetworkIdWirelessRfProfilesTransmission{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesTransmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesTransmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


