/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 Destination ipv4 attributes of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 struct {
	// Destination ipv4 address of the packet.
	Address *string `json:"address,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4WithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4WithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) SetAddress(v string) {
	o.Address = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


