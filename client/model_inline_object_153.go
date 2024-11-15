/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject153 struct for InlineObject153
type InlineObject153 struct {
	// The name of the new stack
	Name string `json:"name"`
	// An array of switch serials to be added into the new stack
	Serials []string `json:"serials"`
}

// NewInlineObject153 instantiates a new InlineObject153 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject153(name string, serials []string) *InlineObject153 {
	this := InlineObject153{}
	this.Name = name
	this.Serials = serials
	return &this
}

// NewInlineObject153WithDefaults instantiates a new InlineObject153 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject153WithDefaults() *InlineObject153 {
	this := InlineObject153{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject153) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject153) SetName(v string) {
	o.Name = v
}

// GetSerials returns the Serials field value
func (o *InlineObject153) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineObject153) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject153) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject153 struct {
	value *InlineObject153
	isSet bool
}

func (v NullableInlineObject153) Get() *InlineObject153 {
	return v.value
}

func (v *NullableInlineObject153) Set(val *InlineObject153) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject153) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject153) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject153(val *InlineObject153) *NullableInlineObject153 {
	return &NullableInlineObject153{value: val, isSet: true}
}

func (v NullableInlineObject153) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject153) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


