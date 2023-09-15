/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20089 struct for InlineResponse20089
type InlineResponse20089 struct {
	// List of the syslog servers for this network
	Servers []InlineResponse20089Servers `json:"servers,omitempty"`
}

// NewInlineResponse20089 instantiates a new InlineResponse20089 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// NewInlineResponse20089WithDefaults instantiates a new InlineResponse20089 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089WithDefaults() *InlineResponse20089 {
	this := InlineResponse20089{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *InlineResponse20089) GetServers() []InlineResponse20089Servers {
	if o == nil || isNil(o.Servers) {
		var ret []InlineResponse20089Servers
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089) GetServersOk() ([]InlineResponse20089Servers, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *InlineResponse20089) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []InlineResponse20089Servers and assigns it to the Servers field.
func (o *InlineResponse20089) SetServers(v []InlineResponse20089Servers) {
	o.Servers = v
}

func (o InlineResponse20089) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089 struct {
	value *InlineResponse20089
	isSet bool
}

func (v NullableInlineResponse20089) Get() *InlineResponse20089 {
	return v.value
}

func (v *NullableInlineResponse20089) Set(val *InlineResponse20089) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089(val *InlineResponse20089) *NullableInlineResponse20089 {
	return &NullableInlineResponse20089{value: val, isSet: true}
}

func (v NullableInlineResponse20089) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


