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

// NetworksNetworkIdMqttBrokersSecurity Security settings of the MQTT broker.
type NetworksNetworkIdMqttBrokersSecurity struct {
	// Security protocol of the MQTT broker.
	Mode *string `json:"mode,omitempty"`
	Tls *NetworksNetworkIdMqttBrokersSecurityTls `json:"tls,omitempty"`
}

// NewNetworksNetworkIdMqttBrokersSecurity instantiates a new NetworksNetworkIdMqttBrokersSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMqttBrokersSecurity() *NetworksNetworkIdMqttBrokersSecurity {
	this := NetworksNetworkIdMqttBrokersSecurity{}
	return &this
}

// NewNetworksNetworkIdMqttBrokersSecurityWithDefaults instantiates a new NetworksNetworkIdMqttBrokersSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMqttBrokersSecurityWithDefaults() *NetworksNetworkIdMqttBrokersSecurity {
	this := NetworksNetworkIdMqttBrokersSecurity{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurity) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *NetworksNetworkIdMqttBrokersSecurity) SetMode(v string) {
	o.Mode = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurity) GetTls() NetworksNetworkIdMqttBrokersSecurityTls {
	if o == nil || isNil(o.Tls) {
		var ret NetworksNetworkIdMqttBrokersSecurityTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity) GetTlsOk() (*NetworksNetworkIdMqttBrokersSecurityTls, bool) {
	if o == nil || isNil(o.Tls) {
    return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity) HasTls() bool {
	if o != nil && !isNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given NetworksNetworkIdMqttBrokersSecurityTls and assigns it to the Tls field.
func (o *NetworksNetworkIdMqttBrokersSecurity) SetTls(v NetworksNetworkIdMqttBrokersSecurityTls) {
	o.Tls = &v
}

func (o NetworksNetworkIdMqttBrokersSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMqttBrokersSecurity struct {
	value *NetworksNetworkIdMqttBrokersSecurity
	isSet bool
}

func (v NullableNetworksNetworkIdMqttBrokersSecurity) Get() *NetworksNetworkIdMqttBrokersSecurity {
	return v.value
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurity) Set(val *NetworksNetworkIdMqttBrokersSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMqttBrokersSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMqttBrokersSecurity(val *NetworksNetworkIdMqttBrokersSecurity) *NullableNetworksNetworkIdMqttBrokersSecurity {
	return &NullableNetworksNetworkIdMqttBrokersSecurity{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMqttBrokersSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMqttBrokersSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


