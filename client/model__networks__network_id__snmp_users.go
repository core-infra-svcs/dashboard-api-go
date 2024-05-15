/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSnmpUsers struct for NetworksNetworkIdSnmpUsers
type NetworksNetworkIdSnmpUsers struct {
	// The username for the SNMP user. Required.
	Username string `json:"username"`
	// The passphrase for the SNMP user. Required.
	Passphrase string `json:"passphrase"`
}

// NewNetworksNetworkIdSnmpUsers instantiates a new NetworksNetworkIdSnmpUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSnmpUsers(username string, passphrase string) *NetworksNetworkIdSnmpUsers {
	this := NetworksNetworkIdSnmpUsers{}
	this.Username = username
	this.Passphrase = passphrase
	return &this
}

// NewNetworksNetworkIdSnmpUsersWithDefaults instantiates a new NetworksNetworkIdSnmpUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSnmpUsersWithDefaults() *NetworksNetworkIdSnmpUsers {
	this := NetworksNetworkIdSnmpUsers{}
	return &this
}

// GetUsername returns the Username field value
func (o *NetworksNetworkIdSnmpUsers) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSnmpUsers) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NetworksNetworkIdSnmpUsers) SetUsername(v string) {
	o.Username = v
}

// GetPassphrase returns the Passphrase field value
func (o *NetworksNetworkIdSnmpUsers) GetPassphrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSnmpUsers) GetPassphraseOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *NetworksNetworkIdSnmpUsers) SetPassphrase(v string) {
	o.Passphrase = v
}

func (o NetworksNetworkIdSnmpUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSnmpUsers struct {
	value *NetworksNetworkIdSnmpUsers
	isSet bool
}

func (v NullableNetworksNetworkIdSnmpUsers) Get() *NetworksNetworkIdSnmpUsers {
	return v.value
}

func (v *NullableNetworksNetworkIdSnmpUsers) Set(val *NetworksNetworkIdSnmpUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSnmpUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSnmpUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSnmpUsers(val *NetworksNetworkIdSnmpUsers) *NullableNetworksNetworkIdSnmpUsers {
	return &NullableNetworksNetworkIdSnmpUsers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSnmpUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSnmpUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


