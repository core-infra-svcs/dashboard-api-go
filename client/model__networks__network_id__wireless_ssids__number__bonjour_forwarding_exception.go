/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException Bonjour forwarding exception
type NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException struct {
	// If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException instantiates a new NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException() *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException {
	this := NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberBonjourForwardingExceptionWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberBonjourForwardingExceptionWithDefaults() *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException {
	this := NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException struct {
	value *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) Get() *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) Set(val *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException(val *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) *NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException {
	return &NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


