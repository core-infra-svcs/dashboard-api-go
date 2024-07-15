/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200253 struct for InlineResponse200253
type InlineResponse200253 struct {
	// Serials of the devices that were released
	Serials []string `json:"serials,omitempty"`
}

// NewInlineResponse200253 instantiates a new InlineResponse200253 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200253() *InlineResponse200253 {
	this := InlineResponse200253{}
	return &this
}

// NewInlineResponse200253WithDefaults instantiates a new InlineResponse200253 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200253WithDefaults() *InlineResponse200253 {
	this := InlineResponse200253{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200253) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200253) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200253) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200253) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse200253) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200253 struct {
	value *InlineResponse200253
	isSet bool
}

func (v NullableInlineResponse200253) Get() *InlineResponse200253 {
	return v.value
}

func (v *NullableInlineResponse200253) Set(val *InlineResponse200253) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200253) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200253) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200253(val *InlineResponse200253) *NullableInlineResponse200253 {
	return &NullableInlineResponse200253{value: val, isSet: true}
}

func (v NullableInlineResponse200253) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200253) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

