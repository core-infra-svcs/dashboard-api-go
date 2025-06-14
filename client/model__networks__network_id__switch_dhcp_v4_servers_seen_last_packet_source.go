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

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource Source of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource struct {
	// Source mac address of the packet.
	Mac *string `json:"mac,omitempty"`
	Ipv4 *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 `json:"ipv4,omitempty"`
	// Source port of the packet.
	Port *int32 `json:"port,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) SetMac(v string) {
	o.Mac = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) GetIpv4() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4 and assigns it to the Ipv4 field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) SetIpv4(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSourceIpv4) {
	o.Ipv4 = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) SetPort(v int32) {
	o.Port = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


