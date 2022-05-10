/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSettingsSecureConnect A hash of SecureConnect options applied to the Network.
type NetworksNetworkIdSettingsSecureConnect struct {
	// Enables / disables SecureConnect on the network. Optional.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdSettingsSecureConnect instantiates a new NetworksNetworkIdSettingsSecureConnect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSettingsSecureConnect() *NetworksNetworkIdSettingsSecureConnect {
	this := NetworksNetworkIdSettingsSecureConnect{}
	return &this
}

// NewNetworksNetworkIdSettingsSecureConnectWithDefaults instantiates a new NetworksNetworkIdSettingsSecureConnect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSettingsSecureConnectWithDefaults() *NetworksNetworkIdSettingsSecureConnect {
	this := NetworksNetworkIdSettingsSecureConnect{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSettingsSecureConnect) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSettingsSecureConnect) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSettingsSecureConnect) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSettingsSecureConnect) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdSettingsSecureConnect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSettingsSecureConnect struct {
	value *NetworksNetworkIdSettingsSecureConnect
	isSet bool
}

func (v NullableNetworksNetworkIdSettingsSecureConnect) Get() *NetworksNetworkIdSettingsSecureConnect {
	return v.value
}

func (v *NullableNetworksNetworkIdSettingsSecureConnect) Set(val *NetworksNetworkIdSettingsSecureConnect) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSettingsSecureConnect) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSettingsSecureConnect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSettingsSecureConnect(val *NetworksNetworkIdSettingsSecureConnect) *NullableNetworksNetworkIdSettingsSecureConnect {
	return &NullableNetworksNetworkIdSettingsSecureConnect{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSettingsSecureConnect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSettingsSecureConnect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


