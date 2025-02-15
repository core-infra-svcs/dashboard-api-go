/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberRadiusRadsec The current settings for RADIUS RADSec
type NetworksNetworkIdWirelessSsidsNumberRadiusRadsec struct {
	TlsTunnel *NetworksNetworkIdWirelessSsidsNumberRadiusRadsecTlsTunnel `json:"tlsTunnel,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberRadiusRadsec instantiates a new NetworksNetworkIdWirelessSsidsNumberRadiusRadsec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberRadiusRadsec() *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec {
	this := NetworksNetworkIdWirelessSsidsNumberRadiusRadsec{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberRadiusRadsecWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberRadiusRadsec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberRadiusRadsecWithDefaults() *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec {
	this := NetworksNetworkIdWirelessSsidsNumberRadiusRadsec{}
	return &this
}

// GetTlsTunnel returns the TlsTunnel field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) GetTlsTunnel() NetworksNetworkIdWirelessSsidsNumberRadiusRadsecTlsTunnel {
	if o == nil || isNil(o.TlsTunnel) {
		var ret NetworksNetworkIdWirelessSsidsNumberRadiusRadsecTlsTunnel
		return ret
	}
	return *o.TlsTunnel
}

// GetTlsTunnelOk returns a tuple with the TlsTunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) GetTlsTunnelOk() (*NetworksNetworkIdWirelessSsidsNumberRadiusRadsecTlsTunnel, bool) {
	if o == nil || isNil(o.TlsTunnel) {
    return nil, false
	}
	return o.TlsTunnel, true
}

// HasTlsTunnel returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) HasTlsTunnel() bool {
	if o != nil && !isNil(o.TlsTunnel) {
		return true
	}

	return false
}

// SetTlsTunnel gets a reference to the given NetworksNetworkIdWirelessSsidsNumberRadiusRadsecTlsTunnel and assigns it to the TlsTunnel field.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) SetTlsTunnel(v NetworksNetworkIdWirelessSsidsNumberRadiusRadsecTlsTunnel) {
	o.TlsTunnel = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TlsTunnel) {
		toSerialize["tlsTunnel"] = o.TlsTunnel
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec struct {
	value *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec) Get() *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec) Set(val *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec(val *NetworksNetworkIdWirelessSsidsNumberRadiusRadsec) *NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec {
	return &NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberRadiusRadsec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


