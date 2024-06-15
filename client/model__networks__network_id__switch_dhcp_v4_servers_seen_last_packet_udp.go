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

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp UDP attributes of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp struct {
	// UDP length of the packet.
	Length *int32 `json:"length,omitempty"`
	// UDP checksum of the packet.
	Checksum *string `json:"checksum,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdpWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdpWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) GetLength() int32 {
	if o == nil || isNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) GetLengthOk() (*int32, bool) {
	if o == nil || isNil(o.Length) {
    return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) HasLength() bool {
	if o != nil && !isNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) SetLength(v int32) {
	o.Length = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) GetChecksum() string {
	if o == nil || isNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) GetChecksumOk() (*string, bool) {
	if o == nil || isNil(o.Checksum) {
    return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) HasChecksum() bool {
	if o != nil && !isNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) SetChecksum(v string) {
	o.Checksum = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !isNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


