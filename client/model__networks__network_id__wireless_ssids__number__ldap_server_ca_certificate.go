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

// NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate The CA certificate used to sign the LDAP server's key.
type NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate struct {
	// The contents of the CA certificate. Must be in PEM or DER format.
	Contents *string `json:"contents,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate instantiates a new NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate() *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate {
	this := NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificateWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificateWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate {
	this := NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) GetContents() string {
	if o == nil || o.Contents == nil {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) GetContentsOk() (*string, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) SetContents(v string) {
	o.Contents = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate struct {
	value *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) Get() *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) Set(val *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate(val *NetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) *NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapServerCaCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


