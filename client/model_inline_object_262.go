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

// InlineObject262 struct for InlineObject262
type InlineObject262 struct {
	// theme name
	Name *string `json:"name,omitempty"`
	// base theme id 
	BaseTheme *string `json:"baseTheme,omitempty"`
}

// NewInlineObject262 instantiates a new InlineObject262 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject262() *InlineObject262 {
	this := InlineObject262{}
	return &this
}

// NewInlineObject262WithDefaults instantiates a new InlineObject262 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject262WithDefaults() *InlineObject262 {
	this := InlineObject262{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject262) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject262) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject262) SetName(v string) {
	o.Name = &v
}

// GetBaseTheme returns the BaseTheme field value if set, zero value otherwise.
func (o *InlineObject262) GetBaseTheme() string {
	if o == nil || isNil(o.BaseTheme) {
		var ret string
		return ret
	}
	return *o.BaseTheme
}

// GetBaseThemeOk returns a tuple with the BaseTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetBaseThemeOk() (*string, bool) {
	if o == nil || isNil(o.BaseTheme) {
    return nil, false
	}
	return o.BaseTheme, true
}

// HasBaseTheme returns a boolean if a field has been set.
func (o *InlineObject262) HasBaseTheme() bool {
	if o != nil && !isNil(o.BaseTheme) {
		return true
	}

	return false
}

// SetBaseTheme gets a reference to the given string and assigns it to the BaseTheme field.
func (o *InlineObject262) SetBaseTheme(v string) {
	o.BaseTheme = &v
}

func (o InlineObject262) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.BaseTheme) {
		toSerialize["baseTheme"] = o.BaseTheme
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject262 struct {
	value *InlineObject262
	isSet bool
}

func (v NullableInlineObject262) Get() *InlineObject262 {
	return v.value
}

func (v *NullableInlineObject262) Set(val *InlineObject262) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject262) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject262) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject262(val *InlineObject262) *NullableInlineObject262 {
	return &NullableInlineObject262{value: val, isSet: true}
}

func (v NullableInlineObject262) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject262) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


