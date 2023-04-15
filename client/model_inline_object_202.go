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

// InlineObject202 struct for InlineObject202
type InlineObject202 struct {
	// Serials of the devices that should be released
	Serials []string `json:"serials,omitempty"`
}

// NewInlineObject202 instantiates a new InlineObject202 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject202() *InlineObject202 {
	this := InlineObject202{}
	return &this
}

// NewInlineObject202WithDefaults instantiates a new InlineObject202 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject202WithDefaults() *InlineObject202 {
	this := InlineObject202{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject202) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject202) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject202) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject202) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject202) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject202 struct {
	value *InlineObject202
	isSet bool
}

func (v NullableInlineObject202) Get() *InlineObject202 {
	return v.value
}

func (v *NullableInlineObject202) Set(val *InlineObject202) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject202) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject202) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject202(val *InlineObject202) *NullableInlineObject202 {
	return &NullableInlineObject202{value: val, isSet: true}
}

func (v NullableInlineObject202) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject202) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


