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

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination Destination of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination struct {
	// Destination mac address of the packet.
	Mac *string `json:"mac,omitempty"`
	Ipv4 *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 `json:"ipv4,omitempty"`
	// Destination port of the packet.
	Port *int32 `json:"port,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) SetMac(v string) {
	o.Mac = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) GetIpv4() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 and assigns it to the Ipv4 field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) SetIpv4(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestinationIpv4) {
	o.Ipv4 = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) SetPort(v int32) {
	o.Port = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


