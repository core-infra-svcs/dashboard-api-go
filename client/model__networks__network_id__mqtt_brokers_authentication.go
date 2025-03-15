/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdMqttBrokersAuthentication Authentication settings of the MQTT broker
type NetworksNetworkIdMqttBrokersAuthentication struct {
	// Username for the MQTT broker.
	Username *string `json:"username,omitempty"`
}

// NewNetworksNetworkIdMqttBrokersAuthentication instantiates a new NetworksNetworkIdMqttBrokersAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMqttBrokersAuthentication() *NetworksNetworkIdMqttBrokersAuthentication {
	this := NetworksNetworkIdMqttBrokersAuthentication{}
	return &this
}

// NewNetworksNetworkIdMqttBrokersAuthenticationWithDefaults instantiates a new NetworksNetworkIdMqttBrokersAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMqttBrokersAuthenticationWithDefaults() *NetworksNetworkIdMqttBrokersAuthentication {
	this := NetworksNetworkIdMqttBrokersAuthentication{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersAuthentication) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersAuthentication) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NetworksNetworkIdMqttBrokersAuthentication) SetUsername(v string) {
	o.Username = &v
}

func (o NetworksNetworkIdMqttBrokersAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMqttBrokersAuthentication struct {
	value *NetworksNetworkIdMqttBrokersAuthentication
	isSet bool
}

func (v NullableNetworksNetworkIdMqttBrokersAuthentication) Get() *NetworksNetworkIdMqttBrokersAuthentication {
	return v.value
}

func (v *NullableNetworksNetworkIdMqttBrokersAuthentication) Set(val *NetworksNetworkIdMqttBrokersAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMqttBrokersAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMqttBrokersAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMqttBrokersAuthentication(val *NetworksNetworkIdMqttBrokersAuthentication) *NullableNetworksNetworkIdMqttBrokersAuthentication {
	return &NullableNetworksNetworkIdMqttBrokersAuthentication{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMqttBrokersAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMqttBrokersAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


