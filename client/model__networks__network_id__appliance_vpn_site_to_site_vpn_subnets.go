/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets struct for NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets
type NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets struct {
	// The CIDR notation subnet used within the VPN
	LocalSubnet string `json:"localSubnet"`
	// Indicates the presence of the subnet in the VPN
	UseVpn *bool `json:"useVpn,omitempty"`
}

// NewNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets instantiates a new NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets(localSubnet string) *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets {
	this := NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets{}
	this.LocalSubnet = localSubnet
	return &this
}

// NewNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnetsWithDefaults instantiates a new NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnetsWithDefaults() *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets {
	this := NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets{}
	return &this
}

// GetLocalSubnet returns the LocalSubnet field value
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) GetLocalSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalSubnet
}

// GetLocalSubnetOk returns a tuple with the LocalSubnet field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) GetLocalSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalSubnet, true
}

// SetLocalSubnet sets field value
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) SetLocalSubnet(v string) {
	o.LocalSubnet = v
}

// GetUseVpn returns the UseVpn field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) GetUseVpn() bool {
	if o == nil || isNil(o.UseVpn) {
		var ret bool
		return ret
	}
	return *o.UseVpn
}

// GetUseVpnOk returns a tuple with the UseVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) GetUseVpnOk() (*bool, bool) {
	if o == nil || isNil(o.UseVpn) {
    return nil, false
	}
	return o.UseVpn, true
}

// HasUseVpn returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) HasUseVpn() bool {
	if o != nil && !isNil(o.UseVpn) {
		return true
	}

	return false
}

// SetUseVpn gets a reference to the given bool and assigns it to the UseVpn field.
func (o *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) SetUseVpn(v bool) {
	o.UseVpn = &v
}

func (o NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["localSubnet"] = o.LocalSubnet
	}
	if !isNil(o.UseVpn) {
		toSerialize["useVpn"] = o.UseVpn
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets struct {
	value *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) Get() *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) Set(val *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets(val *NetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets {
	return &NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVpnSiteToSiteVpnSubnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


