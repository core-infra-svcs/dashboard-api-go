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

// InlineResponse200170 struct for InlineResponse200170
type InlineResponse200170 struct {
	// List of the syslog servers for this network
	Servers []InlineResponse200170Servers `json:"servers,omitempty"`
}

// NewInlineResponse200170 instantiates a new InlineResponse200170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200170() *InlineResponse200170 {
	this := InlineResponse200170{}
	return &this
}

// NewInlineResponse200170WithDefaults instantiates a new InlineResponse200170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200170WithDefaults() *InlineResponse200170 {
	this := InlineResponse200170{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *InlineResponse200170) GetServers() []InlineResponse200170Servers {
	if o == nil || isNil(o.Servers) {
		var ret []InlineResponse200170Servers
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200170) GetServersOk() ([]InlineResponse200170Servers, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *InlineResponse200170) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []InlineResponse200170Servers and assigns it to the Servers field.
func (o *InlineResponse200170) SetServers(v []InlineResponse200170Servers) {
	o.Servers = v
}

func (o InlineResponse200170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200170 struct {
	value *InlineResponse200170
	isSet bool
}

func (v NullableInlineResponse200170) Get() *InlineResponse200170 {
	return v.value
}

func (v *NullableInlineResponse200170) Set(val *InlineResponse200170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200170(val *InlineResponse200170) *NullableInlineResponse200170 {
	return &NullableInlineResponse200170{value: val, isSet: true}
}

func (v NullableInlineResponse200170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


