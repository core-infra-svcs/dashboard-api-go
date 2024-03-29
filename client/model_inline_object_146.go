/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject146 struct for InlineObject146
type InlineObject146 struct {
	// The name of the new stack
	Name string `json:"name"`
	// An array of switch serials to be added into the new stack
	Serials []string `json:"serials"`
}

// NewInlineObject146 instantiates a new InlineObject146 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject146(name string, serials []string) *InlineObject146 {
	this := InlineObject146{}
	this.Name = name
	this.Serials = serials
	return &this
}

// NewInlineObject146WithDefaults instantiates a new InlineObject146 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject146WithDefaults() *InlineObject146 {
	this := InlineObject146{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject146) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject146) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject146) SetName(v string) {
	o.Name = v
}

// GetSerials returns the Serials field value
func (o *InlineObject146) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineObject146) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineObject146) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject146) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject146 struct {
	value *InlineObject146
	isSet bool
}

func (v NullableInlineObject146) Get() *InlineObject146 {
	return v.value
}

func (v *NullableInlineObject146) Set(val *InlineObject146) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject146) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject146) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject146(val *InlineObject146) *NullableInlineObject146 {
	return &NullableInlineObject146{value: val, isSet: true}
}

func (v NullableInlineObject146) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject146) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


