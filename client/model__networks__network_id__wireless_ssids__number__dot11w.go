/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberDot11w The current setting for Protected Management Frames (802.11w).
type NetworksNetworkIdWirelessSsidsNumberDot11w struct {
	// Whether 802.11w is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// (Optional) Whether 802.11w is required or not.
	Required *bool `json:"required,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberDot11w instantiates a new NetworksNetworkIdWirelessSsidsNumberDot11w object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberDot11w() *NetworksNetworkIdWirelessSsidsNumberDot11w {
	this := NetworksNetworkIdWirelessSsidsNumberDot11w{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberDot11wWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberDot11w object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberDot11wWithDefaults() *NetworksNetworkIdWirelessSsidsNumberDot11w {
	this := NetworksNetworkIdWirelessSsidsNumberDot11w{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
    return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11w) SetRequired(v bool) {
	o.Required = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberDot11w) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberDot11w struct {
	value *NetworksNetworkIdWirelessSsidsNumberDot11w
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDot11w) Get() *NetworksNetworkIdWirelessSsidsNumberDot11w {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDot11w) Set(val *NetworksNetworkIdWirelessSsidsNumberDot11w) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDot11w) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDot11w) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberDot11w(val *NetworksNetworkIdWirelessSsidsNumberDot11w) *NullableNetworksNetworkIdWirelessSsidsNumberDot11w {
	return &NullableNetworksNetworkIdWirelessSsidsNumberDot11w{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDot11w) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDot11w) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


