/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject76 struct for InlineObject76
type InlineObject76 struct {
	// Name of the MQTT broker.
	Name string `json:"name"`
	// Host name/IP address where the MQTT broker runs.
	Host string `json:"host"`
	// Host port though which the MQTT broker can be reached.
	Port int32 `json:"port"`
	Security *NetworksNetworkIdMqttBrokersSecurity `json:"security,omitempty"`
	// Authentication settings of the MQTT broker
	Authentication *map[string]interface{} `json:"authentication,omitempty"`
}

// NewInlineObject76 instantiates a new InlineObject76 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject76(name string, host string, port int32) *InlineObject76 {
	this := InlineObject76{}
	this.Name = name
	this.Host = host
	this.Port = port
	return &this
}

// NewInlineObject76WithDefaults instantiates a new InlineObject76 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject76WithDefaults() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject76) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject76) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
func (o *InlineObject76) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *InlineObject76) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *InlineObject76) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *InlineObject76) SetPort(v int32) {
	o.Port = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineObject76) GetSecurity() NetworksNetworkIdMqttBrokersSecurity {
	if o == nil || o.Security == nil {
		var ret NetworksNetworkIdMqttBrokersSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineObject76) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity and assigns it to the Security field.
func (o *InlineObject76) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineObject76) GetAuthentication() map[string]interface{} {
	if o == nil || o.Authentication == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetAuthenticationOk() (*map[string]interface{}, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineObject76) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given map[string]interface{} and assigns it to the Authentication field.
func (o *InlineObject76) SetAuthentication(v map[string]interface{}) {
	o.Authentication = &v
}

func (o InlineObject76) MarshalJSON() ([]byte, error) {
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
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject76 struct {
	value *InlineObject76
	isSet bool
}

func (v NullableInlineObject76) Get() *InlineObject76 {
	return v.value
}

func (v *NullableInlineObject76) Set(val *InlineObject76) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject76) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject76) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject76(val *InlineObject76) *NullableInlineObject76 {
	return &NullableInlineObject76{value: val, isSet: true}
}

func (v NullableInlineObject76) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject76) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


