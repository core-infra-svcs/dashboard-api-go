/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject91 struct for InlineObject91
type InlineObject91 struct {
	// Name of the MQTT broker.
	Name string `json:"name"`
	// Host name/IP address where the MQTT broker runs.
	Host string `json:"host"`
	// Host port though which the MQTT broker can be reached.
	Port int32 `json:"port"`
	Security *NetworksNetworkIdMqttBrokersSecurity `json:"security,omitempty"`
	// Authentication settings of the MQTT broker
	Authentication map[string]interface{} `json:"authentication,omitempty"`
}

// NewInlineObject91 instantiates a new InlineObject91 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject91(name string, host string, port int32) *InlineObject91 {
	this := InlineObject91{}
	this.Name = name
	this.Host = host
	this.Port = port
	return &this
}

// NewInlineObject91WithDefaults instantiates a new InlineObject91 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject91WithDefaults() *InlineObject91 {
	this := InlineObject91{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject91) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject91) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
func (o *InlineObject91) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *InlineObject91) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *InlineObject91) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *InlineObject91) SetPort(v int32) {
	o.Port = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineObject91) GetSecurity() NetworksNetworkIdMqttBrokersSecurity {
	if o == nil || isNil(o.Security) {
		var ret NetworksNetworkIdMqttBrokersSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineObject91) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity and assigns it to the Security field.
func (o *InlineObject91) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineObject91) GetAuthentication() map[string]interface{} {
	if o == nil || isNil(o.Authentication) {
		var ret map[string]interface{}
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetAuthenticationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Authentication) {
    return map[string]interface{}{}, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineObject91) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given map[string]interface{} and assigns it to the Authentication field.
func (o *InlineObject91) SetAuthentication(v map[string]interface{}) {
	o.Authentication = v
}

func (o InlineObject91) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject91 struct {
	value *InlineObject91
	isSet bool
}

func (v NullableInlineObject91) Get() *InlineObject91 {
	return v.value
}

func (v *NullableInlineObject91) Set(val *InlineObject91) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject91) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject91) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject91(val *InlineObject91) *NullableInlineObject91 {
	return &NullableInlineObject91{value: val, isSet: true}
}

func (v NullableInlineObject91) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject91) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


