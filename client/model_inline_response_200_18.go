/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20018 struct for InlineResponse20018
type InlineResponse20018 struct {
	// An array of port forwarding params
	Rules []InlineResponse20018Rules `json:"rules,omitempty"`
}

// NewInlineResponse20018 instantiates a new InlineResponse20018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20018() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// NewInlineResponse20018WithDefaults instantiates a new InlineResponse20018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20018WithDefaults() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse20018) GetRules() []InlineResponse20018Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse20018Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetRulesOk() ([]InlineResponse20018Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse20018) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse20018Rules and assigns it to the Rules field.
func (o *InlineResponse20018) SetRules(v []InlineResponse20018Rules) {
	o.Rules = v
}

func (o InlineResponse20018) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20018 struct {
	value *InlineResponse20018
	isSet bool
}

func (v NullableInlineResponse20018) Get() *InlineResponse20018 {
	return v.value
}

func (v *NullableInlineResponse20018) Set(val *InlineResponse20018) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20018) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20018(val *InlineResponse20018) *NullableInlineResponse20018 {
	return &NullableInlineResponse20018{value: val, isSet: true}
}

func (v NullableInlineResponse20018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


