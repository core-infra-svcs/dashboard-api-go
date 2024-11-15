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

// InlineObject165 struct for InlineObject165
type InlineObject165 struct {
	// Optional boolean to retain all the current configs given by the template.
	RetainConfigs *bool `json:"retainConfigs,omitempty"`
}

// NewInlineObject165 instantiates a new InlineObject165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject165() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// NewInlineObject165WithDefaults instantiates a new InlineObject165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject165WithDefaults() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// GetRetainConfigs returns the RetainConfigs field value if set, zero value otherwise.
func (o *InlineObject165) GetRetainConfigs() bool {
	if o == nil || isNil(o.RetainConfigs) {
		var ret bool
		return ret
	}
	return *o.RetainConfigs
}

// GetRetainConfigsOk returns a tuple with the RetainConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetRetainConfigsOk() (*bool, bool) {
	if o == nil || isNil(o.RetainConfigs) {
    return nil, false
	}
	return o.RetainConfigs, true
}

// HasRetainConfigs returns a boolean if a field has been set.
func (o *InlineObject165) HasRetainConfigs() bool {
	if o != nil && !isNil(o.RetainConfigs) {
		return true
	}

	return false
}

// SetRetainConfigs gets a reference to the given bool and assigns it to the RetainConfigs field.
func (o *InlineObject165) SetRetainConfigs(v bool) {
	o.RetainConfigs = &v
}

func (o InlineObject165) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RetainConfigs) {
		toSerialize["retainConfigs"] = o.RetainConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject165 struct {
	value *InlineObject165
	isSet bool
}

func (v NullableInlineObject165) Get() *InlineObject165 {
	return v.value
}

func (v *NullableInlineObject165) Set(val *InlineObject165) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject165) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject165(val *InlineObject165) *NullableInlineObject165 {
	return &NullableInlineObject165{value: val, isSet: true}
}

func (v NullableInlineObject165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


