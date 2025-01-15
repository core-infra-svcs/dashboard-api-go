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

// InlineObject234 struct for InlineObject234
type InlineObject234 struct {
	// The name of the new organization
	Name string `json:"name"`
}

// NewInlineObject234 instantiates a new InlineObject234 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject234(name string) *InlineObject234 {
	this := InlineObject234{}
	this.Name = name
	return &this
}

// NewInlineObject234WithDefaults instantiates a new InlineObject234 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject234WithDefaults() *InlineObject234 {
	this := InlineObject234{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject234) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject234) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject234) SetName(v string) {
	o.Name = v
}

func (o InlineObject234) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject234 struct {
	value *InlineObject234
	isSet bool
}

func (v NullableInlineObject234) Get() *InlineObject234 {
	return v.value
}

func (v *NullableInlineObject234) Set(val *InlineObject234) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject234) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject234) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject234(val *InlineObject234) *NullableInlineObject234 {
	return &NullableInlineObject234{value: val, isSet: true}
}

func (v NullableInlineObject234) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject234) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


