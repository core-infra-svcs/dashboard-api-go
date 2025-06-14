/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberVpnConcentrator The VPN concentrator settings for this SSID.
type NetworksNetworkIdWirelessSsidsNumberVpnConcentrator struct {
	// The NAT ID of the concentrator that should be set.
	NetworkId *string `json:"networkId,omitempty"`
	// The VLAN that should be tagged for the concentrator.
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnConcentrator instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnConcentrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberVpnConcentrator() *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator {
	this := NetworksNetworkIdWirelessSsidsNumberVpnConcentrator{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnConcentratorWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnConcentrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberVpnConcentratorWithDefaults() *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator {
	this := NetworksNetworkIdWirelessSsidsNumberVpnConcentrator{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator struct {
	value *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator) Get() *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator) Set(val *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator(val *NetworksNetworkIdWirelessSsidsNumberVpnConcentrator) *NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator {
	return &NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnConcentrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


