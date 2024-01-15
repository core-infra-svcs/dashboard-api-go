/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers struct for NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers
type NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers struct {
	// IP address (or FQDN) of your Active Directory server.
	Host string `json:"host"`
	// (Optional) UDP port the Active Directory server listens on. By default, uses port 3268.
	Port *int32 `json:"port,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers instantiates a new NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers(host string) *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers {
	this := NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers{}
	this.Host = host
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServersWithDefaults() *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers {
	this := NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) SetPort(v int32) {
	o.Port = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers struct {
	value *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) Get() *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) Set(val *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers(val *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers {
	return &NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


