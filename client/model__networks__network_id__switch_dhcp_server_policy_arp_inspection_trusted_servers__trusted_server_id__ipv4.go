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

// NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 The updated IPv4 attributes of the trusted server
type NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 struct {
	// The updated IPv4 address of the trusted server
	Address *string `json:"address,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4WithDefaults instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4WithDefaults() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 struct {
	value *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) Get() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) Set(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 {
	return &NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


