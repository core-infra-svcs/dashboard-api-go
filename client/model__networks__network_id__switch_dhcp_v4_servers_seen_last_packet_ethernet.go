/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet Additional ethernet attributes of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet struct {
	// Ethernet type of the packet.
	Type *string `json:"type,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernetWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernetWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) SetType(v string) {
	o.Type = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketEthernet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


