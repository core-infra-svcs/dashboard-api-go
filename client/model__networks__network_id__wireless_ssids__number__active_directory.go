/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberActiveDirectory The current setting for Active Directory. Only valid if splashPage is 'Password-protected with Active Directory'
type NetworksNetworkIdWirelessSsidsNumberActiveDirectory struct {
	// The Active Directory servers to be used for authentication.
	Servers *[]NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers `json:"servers,omitempty"`
	Credentials *NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials `json:"credentials,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberActiveDirectory instantiates a new NetworksNetworkIdWirelessSsidsNumberActiveDirectory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberActiveDirectory() *NetworksNetworkIdWirelessSsidsNumberActiveDirectory {
	this := NetworksNetworkIdWirelessSsidsNumberActiveDirectory{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberActiveDirectory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberActiveDirectoryWithDefaults() *NetworksNetworkIdWirelessSsidsNumberActiveDirectory {
	this := NetworksNetworkIdWirelessSsidsNumberActiveDirectory{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) GetServers() []NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers {
	if o == nil || o.Servers == nil {
		var ret []NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) GetServersOk() (*[]NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers and assigns it to the Servers field.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) SetServers(v []NetworksNetworkIdWirelessSsidsNumberActiveDirectoryServers) {
	o.Servers = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) GetCredentials() NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials {
	if o == nil || o.Credentials == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) GetCredentialsOk() (*NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials and assigns it to the Credentials field.
func (o *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) SetCredentials(v NetworksNetworkIdWirelessSsidsNumberActiveDirectoryCredentials) {
	o.Credentials = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberActiveDirectory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory struct {
	value *NetworksNetworkIdWirelessSsidsNumberActiveDirectory
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory) Get() *NetworksNetworkIdWirelessSsidsNumberActiveDirectory {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory) Set(val *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory(val *NetworksNetworkIdWirelessSsidsNumberActiveDirectory) *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory {
	return &NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberActiveDirectory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


