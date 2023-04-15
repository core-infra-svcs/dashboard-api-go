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

// NetworksNetworkIdWirelessSsidsNumberSpeedBurst The SpeedBurst setting for this SSID'
type NetworksNetworkIdWirelessSsidsNumberSpeedBurst struct {
	// Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSpeedBurst instantiates a new NetworksNetworkIdWirelessSsidsNumberSpeedBurst object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSpeedBurst() *NetworksNetworkIdWirelessSsidsNumberSpeedBurst {
	this := NetworksNetworkIdWirelessSsidsNumberSpeedBurst{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSpeedBurstWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSpeedBurst object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSpeedBurstWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSpeedBurst {
	this := NetworksNetworkIdWirelessSsidsNumberSpeedBurst{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSpeedBurst) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSpeedBurst) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSpeedBurst) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberSpeedBurst) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSpeedBurst) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst struct {
	value *NetworksNetworkIdWirelessSsidsNumberSpeedBurst
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst) Get() *NetworksNetworkIdWirelessSsidsNumberSpeedBurst {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst) Set(val *NetworksNetworkIdWirelessSsidsNumberSpeedBurst) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst(val *NetworksNetworkIdWirelessSsidsNumberSpeedBurst) *NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSpeedBurst) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


