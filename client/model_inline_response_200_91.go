/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20091 struct for InlineResponse20091
type InlineResponse20091 struct {
	// List of the syslog servers for this network
	Servers []InlineResponse20091Servers `json:"servers,omitempty"`
}

// NewInlineResponse20091 instantiates a new InlineResponse20091 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// NewInlineResponse20091WithDefaults instantiates a new InlineResponse20091 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091WithDefaults() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *InlineResponse20091) GetServers() []InlineResponse20091Servers {
	if o == nil || isNil(o.Servers) {
		var ret []InlineResponse20091Servers
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetServersOk() ([]InlineResponse20091Servers, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *InlineResponse20091) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []InlineResponse20091Servers and assigns it to the Servers field.
func (o *InlineResponse20091) SetServers(v []InlineResponse20091Servers) {
	o.Servers = v
}

func (o InlineResponse20091) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20091 struct {
	value *InlineResponse20091
	isSet bool
}

func (v NullableInlineResponse20091) Get() *InlineResponse20091 {
	return v.value
}

func (v *NullableInlineResponse20091) Set(val *InlineResponse20091) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091(val *InlineResponse20091) *NullableInlineResponse20091 {
	return &NullableInlineResponse20091{value: val, isSet: true}
}

func (v NullableInlineResponse20091) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


