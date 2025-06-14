/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject222 struct for InlineObject222
type InlineObject222 struct {
	// Name of profile
	Name string `json:"name"`
}

// NewInlineObject222 instantiates a new InlineObject222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject222(name string) *InlineObject222 {
	this := InlineObject222{}
	this.Name = name
	return &this
}

// NewInlineObject222WithDefaults instantiates a new InlineObject222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject222WithDefaults() *InlineObject222 {
	this := InlineObject222{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject222) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject222) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject222) SetName(v string) {
	o.Name = v
}

func (o InlineObject222) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject222 struct {
	value *InlineObject222
	isSet bool
}

func (v NullableInlineObject222) Get() *InlineObject222 {
	return v.value
}

func (v *NullableInlineObject222) Set(val *InlineObject222) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject222) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject222) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject222(val *InlineObject222) *NullableInlineObject222 {
	return &NullableInlineObject222{value: val, isSet: true}
}

func (v NullableInlineObject222) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject222) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


