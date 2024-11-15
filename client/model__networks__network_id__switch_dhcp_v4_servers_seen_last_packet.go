/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket Last packet the server received.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket struct {
	Source *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource `json:"source,omitempty"`
	Destination *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination `json:"destination,omitempty"`
	// Packet type.
	Type *string `json:"type,omitempty"`
	Ethernet *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet `json:"ethernet,omitempty"`
	Ip *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp `json:"ip,omitempty"`
	Udp *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp `json:"udp,omitempty"`
	Fields *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields `json:"fields,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetSource() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource {
	if o == nil || isNil(o.Source) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetSourceOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource and assigns it to the Source field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetSource(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketSource) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetDestination() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination {
	if o == nil || isNil(o.Destination) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetDestinationOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination and assigns it to the Destination field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetDestination(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketDestination) {
	o.Destination = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetType(v string) {
	o.Type = &v
}

// GetEthernet returns the Ethernet field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetEthernet() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet {
	if o == nil || isNil(o.Ethernet) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet
		return ret
	}
	return *o.Ethernet
}

// GetEthernetOk returns a tuple with the Ethernet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetEthernetOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet, bool) {
	if o == nil || isNil(o.Ethernet) {
    return nil, false
	}
	return o.Ethernet, true
}

// HasEthernet returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasEthernet() bool {
	if o != nil && !isNil(o.Ethernet) {
		return true
	}

	return false
}

// SetEthernet gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet and assigns it to the Ethernet field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetEthernet(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) {
	o.Ethernet = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetIp() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp {
	if o == nil || isNil(o.Ip) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetIpOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp and assigns it to the Ip field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetIp(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) {
	o.Ip = &v
}

// GetUdp returns the Udp field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetUdp() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp {
	if o == nil || isNil(o.Udp) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp
		return ret
	}
	return *o.Udp
}

// GetUdpOk returns a tuple with the Udp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetUdpOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp, bool) {
	if o == nil || isNil(o.Udp) {
    return nil, false
	}
	return o.Udp, true
}

// HasUdp returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasUdp() bool {
	if o != nil && !isNil(o.Udp) {
		return true
	}

	return false
}

// SetUdp gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp and assigns it to the Udp field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetUdp(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketUdp) {
	o.Udp = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetFields() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields {
	if o == nil || isNil(o.Fields) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) GetFieldsOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields and assigns it to the Fields field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) SetFields(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketFields) {
	o.Fields = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Ethernet) {
		toSerialize["ethernet"] = o.Ethernet
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Udp) {
		toSerialize["udp"] = o.Udp
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


