/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject114 struct for InlineObject114
type InlineObject114 struct {
	// Name of the MQTT broker.
	Name *string `json:"name,omitempty"`
	// Host name/IP address where the MQTT broker runs.
	Host *string `json:"host,omitempty"`
	// Host port though which the MQTT broker can be reached.
	Port *int32 `json:"port,omitempty"`
	Security *NetworksNetworkIdMqttBrokersSecurity1 `json:"security,omitempty"`
	Authentication *NetworksNetworkIdMqttBrokersAuthentication1 `json:"authentication,omitempty"`
}

// NewInlineObject114 instantiates a new InlineObject114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject114() *InlineObject114 {
	this := InlineObject114{}
	return &this
}

// NewInlineObject114WithDefaults instantiates a new InlineObject114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject114WithDefaults() *InlineObject114 {
	this := InlineObject114{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject114) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject114) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject114) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InlineObject114) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InlineObject114) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InlineObject114) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineObject114) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineObject114) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InlineObject114) SetPort(v int32) {
	o.Port = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineObject114) GetSecurity() NetworksNetworkIdMqttBrokersSecurity1 {
	if o == nil || isNil(o.Security) {
		var ret NetworksNetworkIdMqttBrokersSecurity1
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity1, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineObject114) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity1 and assigns it to the Security field.
func (o *InlineObject114) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity1) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineObject114) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication1 {
	if o == nil || isNil(o.Authentication) {
		var ret NetworksNetworkIdMqttBrokersAuthentication1
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication1, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineObject114) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NetworksNetworkIdMqttBrokersAuthentication1 and assigns it to the Authentication field.
func (o *InlineObject114) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication1) {
	o.Authentication = &v
}

func (o InlineObject114) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject114 struct {
	value *InlineObject114
	isSet bool
}

func (v NullableInlineObject114) Get() *InlineObject114 {
	return v.value
}

func (v *NullableInlineObject114) Set(val *InlineObject114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject114(val *InlineObject114) *NullableInlineObject114 {
	return &NullableInlineObject114{value: val, isSet: true}
}

func (v NullableInlineObject114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


