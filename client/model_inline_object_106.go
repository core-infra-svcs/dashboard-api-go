/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject106 struct for InlineObject106
type InlineObject106 struct {
	// Name of the MQTT broker.
	Name *string `json:"name,omitempty"`
	// Host name/IP address where the MQTT broker runs.
	Host *string `json:"host,omitempty"`
	// Host port though which the MQTT broker can be reached.
	Port *int32 `json:"port,omitempty"`
	Security *NetworksNetworkIdMqttBrokersSecurity1 `json:"security,omitempty"`
	Authentication *NetworksNetworkIdMqttBrokersAuthentication1 `json:"authentication,omitempty"`
}

// NewInlineObject106 instantiates a new InlineObject106 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject106() *InlineObject106 {
	this := InlineObject106{}
	return &this
}

// NewInlineObject106WithDefaults instantiates a new InlineObject106 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject106WithDefaults() *InlineObject106 {
	this := InlineObject106{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject106) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject106) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject106) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InlineObject106) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InlineObject106) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InlineObject106) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineObject106) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineObject106) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InlineObject106) SetPort(v int32) {
	o.Port = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineObject106) GetSecurity() NetworksNetworkIdMqttBrokersSecurity1 {
	if o == nil || isNil(o.Security) {
		var ret NetworksNetworkIdMqttBrokersSecurity1
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity1, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineObject106) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity1 and assigns it to the Security field.
func (o *InlineObject106) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity1) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineObject106) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication1 {
	if o == nil || isNil(o.Authentication) {
		var ret NetworksNetworkIdMqttBrokersAuthentication1
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication1, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineObject106) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NetworksNetworkIdMqttBrokersAuthentication1 and assigns it to the Authentication field.
func (o *InlineObject106) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication1) {
	o.Authentication = &v
}

func (o InlineObject106) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
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

type NullableInlineObject106 struct {
	value *InlineObject106
	isSet bool
}

func (v NullableInlineObject106) Get() *InlineObject106 {
	return v.value
}

func (v *NullableInlineObject106) Set(val *InlineObject106) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject106) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject106) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject106(val *InlineObject106) *NullableInlineObject106 {
	return &NullableInlineObject106{value: val, isSet: true}
}

func (v NullableInlineObject106) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject106) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


