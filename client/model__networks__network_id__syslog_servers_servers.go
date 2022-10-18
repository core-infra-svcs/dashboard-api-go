/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSyslogServersServers struct for NetworksNetworkIdSyslogServersServers
type NetworksNetworkIdSyslogServersServers struct {
	// The IP address of the syslog server
	Host string `json:"host"`
	// The port of the syslog server
	Port int32 `json:"port"`
	// A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'
	Roles []string `json:"roles"`
}

// NewNetworksNetworkIdSyslogServersServers instantiates a new NetworksNetworkIdSyslogServersServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSyslogServersServers(host string, port int32, roles []string) *NetworksNetworkIdSyslogServersServers {
	this := NetworksNetworkIdSyslogServersServers{}
	this.Host = host
	this.Port = port
	this.Roles = roles
	return &this
}

// NewNetworksNetworkIdSyslogServersServersWithDefaults instantiates a new NetworksNetworkIdSyslogServersServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSyslogServersServersWithDefaults() *NetworksNetworkIdSyslogServersServers {
	this := NetworksNetworkIdSyslogServersServers{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworksNetworkIdSyslogServersServers) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSyslogServersServers) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworksNetworkIdSyslogServersServers) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *NetworksNetworkIdSyslogServersServers) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSyslogServersServers) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NetworksNetworkIdSyslogServersServers) SetPort(v int32) {
	o.Port = v
}

// GetRoles returns the Roles field value
func (o *NetworksNetworkIdSyslogServersServers) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSyslogServersServers) GetRolesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *NetworksNetworkIdSyslogServersServers) SetRoles(v []string) {
	o.Roles = v
}

func (o NetworksNetworkIdSyslogServersServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSyslogServersServers struct {
	value *NetworksNetworkIdSyslogServersServers
	isSet bool
}

func (v NullableNetworksNetworkIdSyslogServersServers) Get() *NetworksNetworkIdSyslogServersServers {
	return v.value
}

func (v *NullableNetworksNetworkIdSyslogServersServers) Set(val *NetworksNetworkIdSyslogServersServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSyslogServersServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSyslogServersServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSyslogServersServers(val *NetworksNetworkIdSyslogServersServers) *NullableNetworksNetworkIdSyslogServersServers {
	return &NullableNetworksNetworkIdSyslogServersServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSyslogServersServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSyslogServersServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


