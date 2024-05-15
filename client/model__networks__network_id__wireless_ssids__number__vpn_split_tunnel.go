/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel The VPN split tunnel settings for this SSID.
type NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel struct {
	// If true, VPN split tunnel is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of VPN split tunnel rules.
	Rules []NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules `json:"rules,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel() *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel {
	this := NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelWithDefaults() *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel {
	this := NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) GetRules() []NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) GetRulesOk() ([]NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules and assigns it to the Rules field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) SetRules(v []NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnelRules) {
	o.Rules = v
}

func (o NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel struct {
	value *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) Get() *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) Set(val *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel(val *NetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel {
	return &NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnSplitTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


