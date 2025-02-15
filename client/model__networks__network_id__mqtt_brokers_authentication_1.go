/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdMqttBrokersAuthentication1 Authentication settings of the MQTT broker
type NetworksNetworkIdMqttBrokersAuthentication1 struct {
	// Username for the MQTT broker.
	Username *string `json:"username,omitempty"`
	// Password for the MQTT broker.
	Password *string `json:"password,omitempty"`
}

// NewNetworksNetworkIdMqttBrokersAuthentication1 instantiates a new NetworksNetworkIdMqttBrokersAuthentication1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMqttBrokersAuthentication1() *NetworksNetworkIdMqttBrokersAuthentication1 {
	this := NetworksNetworkIdMqttBrokersAuthentication1{}
	return &this
}

// NewNetworksNetworkIdMqttBrokersAuthentication1WithDefaults instantiates a new NetworksNetworkIdMqttBrokersAuthentication1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMqttBrokersAuthentication1WithDefaults() *NetworksNetworkIdMqttBrokersAuthentication1 {
	this := NetworksNetworkIdMqttBrokersAuthentication1{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NetworksNetworkIdMqttBrokersAuthentication1) SetPassword(v string) {
	o.Password = &v
}

func (o NetworksNetworkIdMqttBrokersAuthentication1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMqttBrokersAuthentication1 struct {
	value *NetworksNetworkIdMqttBrokersAuthentication1
	isSet bool
}

func (v NullableNetworksNetworkIdMqttBrokersAuthentication1) Get() *NetworksNetworkIdMqttBrokersAuthentication1 {
	return v.value
}

func (v *NullableNetworksNetworkIdMqttBrokersAuthentication1) Set(val *NetworksNetworkIdMqttBrokersAuthentication1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMqttBrokersAuthentication1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMqttBrokersAuthentication1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMqttBrokersAuthentication1(val *NetworksNetworkIdMqttBrokersAuthentication1) *NullableNetworksNetworkIdMqttBrokersAuthentication1 {
	return &NullableNetworksNetworkIdMqttBrokersAuthentication1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMqttBrokersAuthentication1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMqttBrokersAuthentication1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


