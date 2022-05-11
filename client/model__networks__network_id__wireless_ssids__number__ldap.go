/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberLdap The current setting for LDAP. Only valid if splashPage is 'Password-protected with LDAP'.
type NetworksNetworkIdWirelessSsidsNumberLdap struct {
	// The LDAP servers to be used for authentication.
	Servers *[]NetworksNetworkIdWirelessSsidsNumberLdapServers `json:"servers,omitempty"`
	Credentials *NetworksNetworkIdWirelessSsidsNumberLdapCredentials `json:"credentials,omitempty"`
	// The base distinguished name of users on the LDAP server.
	BaseDistinguishedName *string `json:"baseDistinguishedName,omitempty"`
	ServerCaCertificate *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate `json:"serverCaCertificate,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLdap instantiates a new NetworksNetworkIdWirelessSsidsNumberLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLdap() *NetworksNetworkIdWirelessSsidsNumberLdap {
	this := NetworksNetworkIdWirelessSsidsNumberLdap{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLdapWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLdap {
	this := NetworksNetworkIdWirelessSsidsNumberLdap{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServers() []NetworksNetworkIdWirelessSsidsNumberLdapServers {
	if o == nil || o.Servers == nil {
		var ret []NetworksNetworkIdWirelessSsidsNumberLdapServers
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServersOk() (*[]NetworksNetworkIdWirelessSsidsNumberLdapServers, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberLdapServers and assigns it to the Servers field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetServers(v []NetworksNetworkIdWirelessSsidsNumberLdapServers) {
	o.Servers = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetCredentials() NetworksNetworkIdWirelessSsidsNumberLdapCredentials {
	if o == nil || o.Credentials == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberLdapCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetCredentialsOk() (*NetworksNetworkIdWirelessSsidsNumberLdapCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NetworksNetworkIdWirelessSsidsNumberLdapCredentials and assigns it to the Credentials field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetCredentials(v NetworksNetworkIdWirelessSsidsNumberLdapCredentials) {
	o.Credentials = &v
}

// GetBaseDistinguishedName returns the BaseDistinguishedName field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetBaseDistinguishedName() string {
	if o == nil || o.BaseDistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.BaseDistinguishedName
}

// GetBaseDistinguishedNameOk returns a tuple with the BaseDistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetBaseDistinguishedNameOk() (*string, bool) {
	if o == nil || o.BaseDistinguishedName == nil {
		return nil, false
	}
	return o.BaseDistinguishedName, true
}

// HasBaseDistinguishedName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasBaseDistinguishedName() bool {
	if o != nil && o.BaseDistinguishedName != nil {
		return true
	}

	return false
}

// SetBaseDistinguishedName gets a reference to the given string and assigns it to the BaseDistinguishedName field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetBaseDistinguishedName(v string) {
	o.BaseDistinguishedName = &v
}

// GetServerCaCertificate returns the ServerCaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServerCaCertificate() NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate {
	if o == nil || o.ServerCaCertificate == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate
		return ret
	}
	return *o.ServerCaCertificate
}

// GetServerCaCertificateOk returns a tuple with the ServerCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) GetServerCaCertificateOk() (*NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate, bool) {
	if o == nil || o.ServerCaCertificate == nil {
		return nil, false
	}
	return o.ServerCaCertificate, true
}

// HasServerCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) HasServerCaCertificate() bool {
	if o != nil && o.ServerCaCertificate != nil {
		return true
	}

	return false
}

// SetServerCaCertificate gets a reference to the given NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate and assigns it to the ServerCaCertificate field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdap) SetServerCaCertificate(v NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) {
	o.ServerCaCertificate = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLdap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.BaseDistinguishedName != nil {
		toSerialize["baseDistinguishedName"] = o.BaseDistinguishedName
	}
	if o.ServerCaCertificate != nil {
		toSerialize["serverCaCertificate"] = o.ServerCaCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLdap struct {
	value *NetworksNetworkIdWirelessSsidsNumberLdap
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdap) Get() *NetworksNetworkIdWirelessSsidsNumberLdap {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdap) Set(val *NetworksNetworkIdWirelessSsidsNumberLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLdap(val *NetworksNetworkIdWirelessSsidsNumberLdap) *NullableNetworksNetworkIdWirelessSsidsNumberLdap {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLdap{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


