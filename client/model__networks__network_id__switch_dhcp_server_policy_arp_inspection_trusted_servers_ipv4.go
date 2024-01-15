/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 IPv4 attributes of the trusted server.
type NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 struct {
	// IPv4 address of the trusted server.
	Address *string `json:"address,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4WithDefaults instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4WithDefaults() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 struct {
	value *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) Get() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) Set(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 {
	return &NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


