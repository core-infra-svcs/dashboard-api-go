/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject154 struct for InlineObject154
type InlineObject154 struct {
	// The name of the new stack
	Name string `json:"name"`
	// An array of switch serials to be added into the new stack
	Serials []string `json:"serials"`
}

// NewInlineObject154 instantiates a new InlineObject154 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject154(name string, serials []string) *InlineObject154 {
	this := InlineObject154{}
	this.Name = name
	this.Serials = serials
	return &this
}

// NewInlineObject154WithDefaults instantiates a new InlineObject154 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject154WithDefaults() *InlineObject154 {
	this := InlineObject154{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject154) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject154) SetName(v string) {
	o.Name = v
}

// GetSerials returns the Serials field value
func (o *InlineObject154) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineObject154) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineObject154) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject154) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject154 struct {
	value *InlineObject154
	isSet bool
}

func (v NullableInlineObject154) Get() *InlineObject154 {
	return v.value
}

func (v *NullableInlineObject154) Set(val *InlineObject154) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject154) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject154) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject154(val *InlineObject154) *NullableInlineObject154 {
	return &NullableInlineObject154{value: val, isSet: true}
}

func (v NullableInlineObject154) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject154) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


