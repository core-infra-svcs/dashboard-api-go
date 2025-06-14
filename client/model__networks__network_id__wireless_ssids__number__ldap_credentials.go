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

// NetworksNetworkIdWirelessSsidsNumberLdapCredentials (Optional) The credentials of the user account to be used by the AP to bind to your LDAP server. The LDAP account should have permissions on all your LDAP servers.
type NetworksNetworkIdWirelessSsidsNumberLdapCredentials struct {
	// The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	// The password of the LDAP user account.
	Password *string `json:"password,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapCredentials instantiates a new NetworksNetworkIdWirelessSsidsNumberLdapCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLdapCredentials() *NetworksNetworkIdWirelessSsidsNumberLdapCredentials {
	this := NetworksNetworkIdWirelessSsidsNumberLdapCredentials{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLdapCredentialsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLdapCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLdapCredentialsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLdapCredentials {
	this := NetworksNetworkIdWirelessSsidsNumberLdapCredentials{}
	return &this
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) GetDistinguishedName() string {
	if o == nil || isNil(o.DistinguishedName) {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || isNil(o.DistinguishedName) {
    return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) HasDistinguishedName() bool {
	if o != nil && !isNil(o.DistinguishedName) {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) SetPassword(v string) {
	o.Password = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLdapCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DistinguishedName) {
		toSerialize["distinguishedName"] = o.DistinguishedName
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials struct {
	value *NetworksNetworkIdWirelessSsidsNumberLdapCredentials
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials) Get() *NetworksNetworkIdWirelessSsidsNumberLdapCredentials {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials) Set(val *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials(val *NetworksNetworkIdWirelessSsidsNumberLdapCredentials) *NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLdapCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


