/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject130 struct for InlineObject130
type InlineObject130 struct {
	// The name of the new stack
	Name string `json:"name"`
	// An array of switch serials to be added into the new stack
	Serials []string `json:"serials"`
}

// NewInlineObject130 instantiates a new InlineObject130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject130(name string, serials []string) *InlineObject130 {
	this := InlineObject130{}
	this.Name = name
	this.Serials = serials
	return &this
}

// NewInlineObject130WithDefaults instantiates a new InlineObject130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject130WithDefaults() *InlineObject130 {
	this := InlineObject130{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject130) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject130) SetName(v string) {
	o.Name = v
}

// GetSerials returns the Serials field value
func (o *InlineObject130) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetSerialsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Serials, true
}

// SetSerials sets field value
func (o *InlineObject130) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject130 struct {
	value *InlineObject130
	isSet bool
}

func (v NullableInlineObject130) Get() *InlineObject130 {
	return v.value
}

func (v *NullableInlineObject130) Set(val *InlineObject130) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject130) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject130(val *InlineObject130) *NullableInlineObject130 {
	return &NullableInlineObject130{value: val, isSet: true}
}

func (v NullableInlineObject130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


