/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject152 struct for InlineObject152
type InlineObject152 struct {
	// Optional boolean to retain all the current configs given by the template.
	RetainConfigs *bool `json:"retainConfigs,omitempty"`
}

// NewInlineObject152 instantiates a new InlineObject152 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject152() *InlineObject152 {
	this := InlineObject152{}
	return &this
}

// NewInlineObject152WithDefaults instantiates a new InlineObject152 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject152WithDefaults() *InlineObject152 {
	this := InlineObject152{}
	return &this
}

// GetRetainConfigs returns the RetainConfigs field value if set, zero value otherwise.
func (o *InlineObject152) GetRetainConfigs() bool {
	if o == nil || isNil(o.RetainConfigs) {
		var ret bool
		return ret
	}
	return *o.RetainConfigs
}

// GetRetainConfigsOk returns a tuple with the RetainConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject152) GetRetainConfigsOk() (*bool, bool) {
	if o == nil || isNil(o.RetainConfigs) {
    return nil, false
	}
	return o.RetainConfigs, true
}

// HasRetainConfigs returns a boolean if a field has been set.
func (o *InlineObject152) HasRetainConfigs() bool {
	if o != nil && !isNil(o.RetainConfigs) {
		return true
	}

	return false
}

// SetRetainConfigs gets a reference to the given bool and assigns it to the RetainConfigs field.
func (o *InlineObject152) SetRetainConfigs(v bool) {
	o.RetainConfigs = &v
}

func (o InlineObject152) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RetainConfigs) {
		toSerialize["retainConfigs"] = o.RetainConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject152 struct {
	value *InlineObject152
	isSet bool
}

func (v NullableInlineObject152) Get() *InlineObject152 {
	return v.value
}

func (v *NullableInlineObject152) Set(val *InlineObject152) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject152) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject152) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject152(val *InlineObject152) *NullableInlineObject152 {
	return &NullableInlineObject152{value: val, isSet: true}
}

func (v NullableInlineObject152) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject152) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


