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

// InlineObject273 struct for InlineObject273
type InlineObject273 struct {
	// File name. Will overwrite files with same name.
	Name *string `json:"name,omitempty"`
	// a file containing the asset content
	Content *string `json:"content,omitempty"`
}

// NewInlineObject273 instantiates a new InlineObject273 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject273() *InlineObject273 {
	this := InlineObject273{}
	return &this
}

// NewInlineObject273WithDefaults instantiates a new InlineObject273 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject273WithDefaults() *InlineObject273 {
	this := InlineObject273{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject273) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject273) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject273) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject273) SetName(v string) {
	o.Name = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *InlineObject273) GetContent() string {
	if o == nil || isNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject273) GetContentOk() (*string, bool) {
	if o == nil || isNil(o.Content) {
    return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *InlineObject273) HasContent() bool {
	if o != nil && !isNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *InlineObject273) SetContent(v string) {
	o.Content = &v
}

func (o InlineObject273) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject273 struct {
	value *InlineObject273
	isSet bool
}

func (v NullableInlineObject273) Get() *InlineObject273 {
	return v.value
}

func (v *NullableInlineObject273) Set(val *InlineObject273) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject273) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject273) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject273(val *InlineObject273) *NullableInlineObject273 {
	return &NullableInlineObject273{value: val, isSet: true}
}

func (v NullableInlineObject273) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject273) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


