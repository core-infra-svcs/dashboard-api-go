/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject202 struct for InlineObject202
type InlineObject202 struct {
	// Unique name of the artifact
	Name *string `json:"name,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject202) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject202) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject202) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject202) SetName(v string) {
	o.Name = &v
}

func (o InlineObject202) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
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


