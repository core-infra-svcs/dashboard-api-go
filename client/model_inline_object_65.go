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

// InlineObject65 struct for InlineObject65
type InlineObject65 struct {
	// Boolean indicating whether to enable (true) or disable (false) VLANs for the network
	VlansEnabled *bool `json:"vlansEnabled,omitempty"`
}

// NewInlineObject65 instantiates a new InlineObject65 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject65() *InlineObject65 {
	this := InlineObject65{}
	return &this
}

// NewInlineObject65WithDefaults instantiates a new InlineObject65 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject65WithDefaults() *InlineObject65 {
	this := InlineObject65{}
	return &this
}

// GetVlansEnabled returns the VlansEnabled field value if set, zero value otherwise.
func (o *InlineObject65) GetVlansEnabled() bool {
	if o == nil || isNil(o.VlansEnabled) {
		var ret bool
		return ret
	}
	return *o.VlansEnabled
}

// GetVlansEnabledOk returns a tuple with the VlansEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject65) GetVlansEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.VlansEnabled) {
    return nil, false
	}
	return o.VlansEnabled, true
}

// HasVlansEnabled returns a boolean if a field has been set.
func (o *InlineObject65) HasVlansEnabled() bool {
	if o != nil && !isNil(o.VlansEnabled) {
		return true
	}

	return false
}

// SetVlansEnabled gets a reference to the given bool and assigns it to the VlansEnabled field.
func (o *InlineObject65) SetVlansEnabled(v bool) {
	o.VlansEnabled = &v
}

func (o InlineObject65) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VlansEnabled) {
		toSerialize["vlansEnabled"] = o.VlansEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject65 struct {
	value *InlineObject65
	isSet bool
}

func (v NullableInlineObject65) Get() *InlineObject65 {
	return v.value
}

func (v *NullableInlineObject65) Set(val *InlineObject65) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject65) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject65) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject65(val *InlineObject65) *NullableInlineObject65 {
	return &NullableInlineObject65{value: val, isSet: true}
}

func (v NullableInlineObject65) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject65) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


