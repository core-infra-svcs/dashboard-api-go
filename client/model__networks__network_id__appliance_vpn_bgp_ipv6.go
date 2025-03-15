/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVpnBgpIpv6 Information regarding IPv6 address of the neighbor, Required if `ip` is not present.
type NetworksNetworkIdApplianceVpnBgpIpv6 struct {
	// The IPv6 address of the neighbor.
	Address string `json:"address"`
}

// NewNetworksNetworkIdApplianceVpnBgpIpv6 instantiates a new NetworksNetworkIdApplianceVpnBgpIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVpnBgpIpv6(address string) *NetworksNetworkIdApplianceVpnBgpIpv6 {
	this := NetworksNetworkIdApplianceVpnBgpIpv6{}
	this.Address = address
	return &this
}

// NewNetworksNetworkIdApplianceVpnBgpIpv6WithDefaults instantiates a new NetworksNetworkIdApplianceVpnBgpIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVpnBgpIpv6WithDefaults() *NetworksNetworkIdApplianceVpnBgpIpv6 {
	this := NetworksNetworkIdApplianceVpnBgpIpv6{}
	return &this
}

// GetAddress returns the Address field value
func (o *NetworksNetworkIdApplianceVpnBgpIpv6) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpIpv6) GetAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NetworksNetworkIdApplianceVpnBgpIpv6) SetAddress(v string) {
	o.Address = v
}

func (o NetworksNetworkIdApplianceVpnBgpIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVpnBgpIpv6 struct {
	value *NetworksNetworkIdApplianceVpnBgpIpv6
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVpnBgpIpv6) Get() *NetworksNetworkIdApplianceVpnBgpIpv6 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpIpv6) Set(val *NetworksNetworkIdApplianceVpnBgpIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVpnBgpIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVpnBgpIpv6(val *NetworksNetworkIdApplianceVpnBgpIpv6) *NullableNetworksNetworkIdApplianceVpnBgpIpv6 {
	return &NullableNetworksNetworkIdApplianceVpnBgpIpv6{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVpnBgpIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


