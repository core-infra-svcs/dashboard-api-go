/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdMqttBrokersSecurityTls TLS settings of the MQTT broker.
type NetworksNetworkIdMqttBrokersSecurityTls struct {
	// Indicates whether the CA certificate is set
	HasCaCertificate *bool `json:"hasCaCertificate,omitempty"`
	// Whether the TLS hostname verification is enabled for the MQTT broker.
	VerifyHostnames *bool `json:"verifyHostnames,omitempty"`
}

// NewNetworksNetworkIdMqttBrokersSecurityTls instantiates a new NetworksNetworkIdMqttBrokersSecurityTls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMqttBrokersSecurityTls() *NetworksNetworkIdMqttBrokersSecurityTls {
	this := NetworksNetworkIdMqttBrokersSecurityTls{}
	return &this
}

// NewNetworksNetworkIdMqttBrokersSecurityTlsWithDefaults instantiates a new NetworksNetworkIdMqttBrokersSecurityTls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMqttBrokersSecurityTlsWithDefaults() *NetworksNetworkIdMqttBrokersSecurityTls {
	this := NetworksNetworkIdMqttBrokersSecurityTls{}
	return &this
}

// GetHasCaCertificate returns the HasCaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) GetHasCaCertificate() bool {
	if o == nil || isNil(o.HasCaCertificate) {
		var ret bool
		return ret
	}
	return *o.HasCaCertificate
}

// GetHasCaCertificateOk returns a tuple with the HasCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) GetHasCaCertificateOk() (*bool, bool) {
	if o == nil || isNil(o.HasCaCertificate) {
    return nil, false
	}
	return o.HasCaCertificate, true
}

// HasHasCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) HasHasCaCertificate() bool {
	if o != nil && !isNil(o.HasCaCertificate) {
		return true
	}

	return false
}

// SetHasCaCertificate gets a reference to the given bool and assigns it to the HasCaCertificate field.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) SetHasCaCertificate(v bool) {
	o.HasCaCertificate = &v
}

// GetVerifyHostnames returns the VerifyHostnames field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) GetVerifyHostnames() bool {
	if o == nil || isNil(o.VerifyHostnames) {
		var ret bool
		return ret
	}
	return *o.VerifyHostnames
}

// GetVerifyHostnamesOk returns a tuple with the VerifyHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) GetVerifyHostnamesOk() (*bool, bool) {
	if o == nil || isNil(o.VerifyHostnames) {
    return nil, false
	}
	return o.VerifyHostnames, true
}

// HasVerifyHostnames returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) HasVerifyHostnames() bool {
	if o != nil && !isNil(o.VerifyHostnames) {
		return true
	}

	return false
}

// SetVerifyHostnames gets a reference to the given bool and assigns it to the VerifyHostnames field.
func (o *NetworksNetworkIdMqttBrokersSecurityTls) SetVerifyHostnames(v bool) {
	o.VerifyHostnames = &v
}

func (o NetworksNetworkIdMqttBrokersSecurityTls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasCaCertificate) {
		toSerialize["hasCaCertificate"] = o.HasCaCertificate
	}
	if !isNil(o.VerifyHostnames) {
		toSerialize["verifyHostnames"] = o.VerifyHostnames
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMqttBrokersSecurityTls struct {
	value *NetworksNetworkIdMqttBrokersSecurityTls
	isSet bool
}

func (v NullableNetworksNetworkIdMqttBrokersSecurityTls) Get() *NetworksNetworkIdMqttBrokersSecurityTls {
	return v.value
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurityTls) Set(val *NetworksNetworkIdMqttBrokersSecurityTls) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMqttBrokersSecurityTls) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurityTls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMqttBrokersSecurityTls(val *NetworksNetworkIdMqttBrokersSecurityTls) *NullableNetworksNetworkIdMqttBrokersSecurityTls {
	return &NullableNetworksNetworkIdMqttBrokersSecurityTls{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMqttBrokersSecurityTls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurityTls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

