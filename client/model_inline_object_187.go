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

// InlineObject187 struct for InlineObject187
type InlineObject187 struct {
	// Unique name of the artifact
	Name *string `json:"name,omitempty"`
}

// NewInlineObject187 instantiates a new InlineObject187 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject187() *InlineObject187 {
	this := InlineObject187{}
	return &this
}

// NewInlineObject187WithDefaults instantiates a new InlineObject187 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject187WithDefaults() *InlineObject187 {
	this := InlineObject187{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject187) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject187) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject187) SetName(v string) {
	o.Name = &v
}

func (o InlineObject187) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject187 struct {
	value *InlineObject187
	isSet bool
}

func (v NullableInlineObject187) Get() *InlineObject187 {
	return v.value
}

func (v *NullableInlineObject187) Set(val *InlineObject187) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject187) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject187) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject187(val *InlineObject187) *NullableInlineObject187 {
	return &NullableInlineObject187{value: val, isSet: true}
}

func (v NullableInlineObject187) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject187) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


