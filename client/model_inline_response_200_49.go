/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20049 struct for InlineResponse20049
type InlineResponse20049 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []InlineResponse20049Rules `json:"rules,omitempty"`
}

// NewInlineResponse20049 instantiates a new InlineResponse20049 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049WithDefaults() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse20049) GetRules() []InlineResponse20049Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse20049Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetRulesOk() ([]InlineResponse20049Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse20049) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse20049Rules and assigns it to the Rules field.
func (o *InlineResponse20049) SetRules(v []InlineResponse20049Rules) {
	o.Rules = v
}

func (o InlineResponse20049) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20049 struct {
	value *InlineResponse20049
	isSet bool
}

func (v NullableInlineResponse20049) Get() *InlineResponse20049 {
	return v.value
}

func (v *NullableInlineResponse20049) Set(val *InlineResponse20049) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049(val *InlineResponse20049) *NullableInlineResponse20049 {
	return &NullableInlineResponse20049{value: val, isSet: true}
}

func (v NullableInlineResponse20049) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


