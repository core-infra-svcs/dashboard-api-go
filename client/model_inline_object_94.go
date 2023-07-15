/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject94 struct for InlineObject94
type InlineObject94 struct {
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

// NewInlineObject94 instantiates a new InlineObject94 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject94(name string, host string, port int32) *InlineObject94 {
	this := InlineObject94{}
	this.Name = name
	this.Host = host
	this.Port = port
	return &this
}

// NewInlineObject94WithDefaults instantiates a new InlineObject94 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject94WithDefaults() *InlineObject94 {
	this := InlineObject94{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject94) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject94) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
func (o *InlineObject94) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *InlineObject94) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *InlineObject94) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *InlineObject94) SetPort(v int32) {
	o.Port = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineObject94) GetSecurity() NetworksNetworkIdMqttBrokersSecurity {
	if o == nil || isNil(o.Security) {
		var ret NetworksNetworkIdMqttBrokersSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineObject94) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity and assigns it to the Security field.
func (o *InlineObject94) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineObject94) GetAuthentication() map[string]interface{} {
	if o == nil || isNil(o.Authentication) {
		var ret map[string]interface{}
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetAuthenticationOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Authentication) {
    return map[string]interface{}{}, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineObject94) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given map[string]interface{} and assigns it to the Authentication field.
func (o *InlineObject94) SetAuthentication(v map[string]interface{}) {
	o.Authentication = v
}

func (o InlineObject94) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject94 struct {
	value *InlineObject94
	isSet bool
}

func (v NullableInlineObject94) Get() *InlineObject94 {
	return v.value
}

func (v *NullableInlineObject94) Set(val *InlineObject94) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject94) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject94) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject94(val *InlineObject94) *NullableInlineObject94 {
	return &NullableInlineObject94{value: val, isSet: true}
}

func (v NullableInlineObject94) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject94) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


