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

// NetworksNetworkIdWirelessSsidsNumberLdapServers struct for NetworksNetworkIdWirelessSsidsNumberLdapServers
type NetworksNetworkIdWirelessSsidsNumberLdapServers struct {
	// IP address (or FQDN) of your LDAP server.
	Host string `json:"host"`
	// UDP port the LDAP server listens on.
	Port int32 `json:"port"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapServers instantiates a new NetworksNetworkIdWirelessSsidsNumberLdapServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLdapServers(host string, port int32) *NetworksNetworkIdWirelessSsidsNumberLdapServers {
	this := NetworksNetworkIdWirelessSsidsNumberLdapServers{}
	this.Host = host
	this.Port = port
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLdapServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLdapServersWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLdapServers {
	this := NetworksNetworkIdWirelessSsidsNumberLdapServers{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServers) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServers) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServers) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServers) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServers) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServers) SetPort(v int32) {
	o.Port = v
}

func (o NetworksNetworkIdWirelessSsidsNumberLdapServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLdapServers struct {
	value *NetworksNetworkIdWirelessSsidsNumberLdapServers
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapServers) Get() *NetworksNetworkIdWirelessSsidsNumberLdapServers {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapServers) Set(val *NetworksNetworkIdWirelessSsidsNumberLdapServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLdapServers(val *NetworksNetworkIdWirelessSsidsNumberLdapServers) *NullableNetworksNetworkIdWirelessSsidsNumberLdapServers {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLdapServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


