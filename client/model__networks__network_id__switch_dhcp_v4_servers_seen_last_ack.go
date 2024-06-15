/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck Attributes of the server's last ack.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck struct {
	// Last time the server was acked.
	Ts *time.Time `json:"ts,omitempty"`
	Ipv4 *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAckIpv4 `json:"ipv4,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastAckWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastAckWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) SetTs(v time.Time) {
	o.Ts = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) GetIpv4() NetworksNetworkIdSwitchDhcpV4ServersSeenLastAckIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastAckIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastAckIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastAckIpv4 and assigns it to the Ipv4 field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) SetIpv4(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastAckIpv4) {
	o.Ipv4 = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastAck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


