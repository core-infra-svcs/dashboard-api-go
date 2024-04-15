/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject231 struct for InlineObject231
type InlineObject231 struct {
	// Serials of the devices that should be released
	Serials []string `json:"serials,omitempty"`
}

// NewInlineObject231 instantiates a new InlineObject231 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject231() *InlineObject231 {
	this := InlineObject231{}
	return &this
}

// NewInlineObject231WithDefaults instantiates a new InlineObject231 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject231WithDefaults() *InlineObject231 {
	this := InlineObject231{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject231) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject231) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject231) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject231) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject231) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject231 struct {
	value *InlineObject231
	isSet bool
}

func (v NullableInlineObject231) Get() *InlineObject231 {
	return v.value
}

func (v *NullableInlineObject231) Set(val *InlineObject231) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject231) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject231) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject231(val *InlineObject231) *NullableInlineObject231 {
	return &NullableInlineObject231{value: val, isSet: true}
}

func (v NullableInlineObject231) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject231) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


