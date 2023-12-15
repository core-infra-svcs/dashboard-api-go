/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 The IPv4 attributes of the trusted server being added
type NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 struct {
	// The IPv4 address of the trusted server being added
	Address *string `json:"address,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41WithDefaults instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41WithDefaults() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) SetAddress(v string) {
	o.Address = &v
}

func (o NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 struct {
	value *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) Get() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) Set(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	return &NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


