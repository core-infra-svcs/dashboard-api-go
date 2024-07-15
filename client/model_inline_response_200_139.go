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

// InlineResponse200139 struct for InlineResponse200139
type InlineResponse200139 struct {
	// An ordered array of the access control list rules
	Rules []InlineResponse200139Rules `json:"rules,omitempty"`
}

// NewInlineResponse200139 instantiates a new InlineResponse200139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200139() *InlineResponse200139 {
	this := InlineResponse200139{}
	return &this
}

// NewInlineResponse200139WithDefaults instantiates a new InlineResponse200139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200139WithDefaults() *InlineResponse200139 {
	this := InlineResponse200139{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse200139) GetRules() []InlineResponse200139Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200139Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200139) GetRulesOk() ([]InlineResponse200139Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse200139) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200139Rules and assigns it to the Rules field.
func (o *InlineResponse200139) SetRules(v []InlineResponse200139Rules) {
	o.Rules = v
}

func (o InlineResponse200139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200139 struct {
	value *InlineResponse200139
	isSet bool
}

func (v NullableInlineResponse200139) Get() *InlineResponse200139 {
	return v.value
}

func (v *NullableInlineResponse200139) Set(val *InlineResponse200139) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200139) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200139(val *InlineResponse200139) *NullableInlineResponse200139 {
	return &NullableInlineResponse200139{value: val, isSet: true}
}

func (v NullableInlineResponse200139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


