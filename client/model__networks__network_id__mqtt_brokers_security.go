/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
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
	Security *NetworksNetworkIdMqttBrokersSecuritySecurity `json:"security,omitempty"`
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

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecurity) GetSecurity() NetworksNetworkIdMqttBrokersSecuritySecurity {
	if o == nil || isNil(o.Security) {
		var ret NetworksNetworkIdMqttBrokersSecuritySecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecuritySecurity, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecurity) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecuritySecurity and assigns it to the Security field.
func (o *NetworksNetworkIdMqttBrokersSecurity) SetSecurity(v NetworksNetworkIdMqttBrokersSecuritySecurity) {
	o.Security = &v
}

func (o NetworksNetworkIdMqttBrokersSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Security) {
		toSerialize["security"] = o.Security
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


