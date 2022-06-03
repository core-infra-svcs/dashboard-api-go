/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject77 struct for InlineObject77
type InlineObject77 struct {
	// Name of the MQTT broker.
	Name *string `json:"name,omitempty"`
	// Host name/IP address where the MQTT broker runs.
	Host *string `json:"host,omitempty"`
	// Host port though which the MQTT broker can be reached.
	Port *int32 `json:"port,omitempty"`
	Security *NetworksNetworkIdMqttBrokersSecurity `json:"security,omitempty"`
	// Authentication settings of the MQTT broker
	Authentication *map[string]interface{} `json:"authentication,omitempty"`
}

// NewInlineObject77 instantiates a new InlineObject77 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject77() *InlineObject77 {
	this := InlineObject77{}
	return &this
}

// NewInlineObject77WithDefaults instantiates a new InlineObject77 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject77WithDefaults() *InlineObject77 {
	this := InlineObject77{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject77) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject77) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject77) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InlineObject77) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InlineObject77) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InlineObject77) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineObject77) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineObject77) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InlineObject77) SetPort(v int32) {
	o.Port = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineObject77) GetSecurity() NetworksNetworkIdMqttBrokersSecurity {
	if o == nil || o.Security == nil {
		var ret NetworksNetworkIdMqttBrokersSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineObject77) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity and assigns it to the Security field.
func (o *InlineObject77) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineObject77) GetAuthentication() map[string]interface{} {
	if o == nil || o.Authentication == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetAuthenticationOk() (*map[string]interface{}, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineObject77) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given map[string]interface{} and assigns it to the Authentication field.
func (o *InlineObject77) SetAuthentication(v map[string]interface{}) {
	o.Authentication = &v
}

func (o InlineObject77) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
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

type NullableInlineObject77 struct {
	value *InlineObject77
	isSet bool
}

func (v NullableInlineObject77) Get() *InlineObject77 {
	return v.value
}

func (v *NullableInlineObject77) Set(val *InlineObject77) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject77) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject77) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject77(val *InlineObject77) *NullableInlineObject77 {
	return &NullableInlineObject77{value: val, isSet: true}
}

func (v NullableInlineObject77) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject77) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


