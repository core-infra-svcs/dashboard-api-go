/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject263 struct for InlineObject263
type InlineObject263 struct {
	// File name. Will overwrite files with same name.
	Name *string `json:"name,omitempty"`
	// a file containing the asset content
	Content *string `json:"content,omitempty"`
}

// NewInlineObject263 instantiates a new InlineObject263 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject263() *InlineObject263 {
	this := InlineObject263{}
	return &this
}

// NewInlineObject263WithDefaults instantiates a new InlineObject263 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject263WithDefaults() *InlineObject263 {
	this := InlineObject263{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject263) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject263) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject263) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject263) SetName(v string) {
	o.Name = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *InlineObject263) GetContent() string {
	if o == nil || isNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject263) GetContentOk() (*string, bool) {
	if o == nil || isNil(o.Content) {
    return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *InlineObject263) HasContent() bool {
	if o != nil && !isNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *InlineObject263) SetContent(v string) {
	o.Content = &v
}

func (o InlineObject263) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject263 struct {
	value *InlineObject263
	isSet bool
}

func (v NullableInlineObject263) Get() *InlineObject263 {
	return v.value
}

func (v *NullableInlineObject263) Set(val *InlineObject263) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject263) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject263) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject263(val *InlineObject263) *NullableInlineObject263 {
	return &NullableInlineObject263{value: val, isSet: true}
}

func (v NullableInlineObject263) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject263) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


