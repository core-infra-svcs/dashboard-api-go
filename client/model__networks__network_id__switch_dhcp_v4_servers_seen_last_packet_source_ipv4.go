/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 Source ipv4 attributes of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 struct {
	// Source ipv4 address of the packet.
	Address *string `json:"address,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4WithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4WithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


