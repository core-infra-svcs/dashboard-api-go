/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdMqttBrokersSecurity1 Security settings of the MQTT broker.
type NetworksNetworkIdMqttBrokersSecurity1 struct {
	// Security protocol of the MQTT broker.
	Mode *string `json:"mode,omitempty"`
	Tls *NetworksNetworkIdMqttBrokersSecurity1Tls `json:"tls,omitempty"`
}

// NewNetworksNetworkIdMqttBrokersSecurity1 instantiates a new NetworksNetworkIdMqttBrokersSecurity1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMqttBrokersSecurity1() *NetworksNetworkIdMqttBrokersSecurity1 {
	this := NetworksNetworkIdMqttBrokersSecurity1{}
	return &this
}

// NewNetworksNetworkIdMqttBrokersSecurity1WithDefaults instantiates a new NetworksNetworkIdMqttBrokersSecurity1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMqttBrokersSecurity1WithDefaults() *NetworksNetworkIdMqttBrokersSecurity1 {
	this := NetworksNetworkIdMqttBrokersSecurity1{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurity1) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity1) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity1) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *NetworksNetworkIdMqttBrokersSecurity1) SetMode(v string) {
	o.Mode = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurity1) GetTls() NetworksNetworkIdMqttBrokersSecurity1Tls {
	if o == nil || isNil(o.Tls) {
		var ret NetworksNetworkIdMqttBrokersSecurity1Tls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity1) GetTlsOk() (*NetworksNetworkIdMqttBrokersSecurity1Tls, bool) {
	if o == nil || isNil(o.Tls) {
    return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity1) HasTls() bool {
	if o != nil && !isNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given NetworksNetworkIdMqttBrokersSecurity1Tls and assigns it to the Tls field.
func (o *NetworksNetworkIdMqttBrokersSecurity1) SetTls(v NetworksNetworkIdMqttBrokersSecurity1Tls) {
	o.Tls = &v
}

func (o NetworksNetworkIdMqttBrokersSecurity1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMqttBrokersSecurity1 struct {
	value *NetworksNetworkIdMqttBrokersSecurity1
	isSet bool
}

func (v NullableNetworksNetworkIdMqttBrokersSecurity1) Get() *NetworksNetworkIdMqttBrokersSecurity1 {
	return v.value
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurity1) Set(val *NetworksNetworkIdMqttBrokersSecurity1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMqttBrokersSecurity1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurity1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMqttBrokersSecurity1(val *NetworksNetworkIdMqttBrokersSecurity1) *NullableNetworksNetworkIdMqttBrokersSecurity1 {
	return &NullableNetworksNetworkIdMqttBrokersSecurity1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMqttBrokersSecurity1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurity1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


