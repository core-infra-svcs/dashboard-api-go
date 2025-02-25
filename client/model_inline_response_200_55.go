/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20055 struct for InlineResponse20055
type InlineResponse20055 struct {
	// An array of port forwarding rules
	Rules []InlineResponse20055Rules `json:"rules,omitempty"`
}

// NewInlineResponse20055 instantiates a new InlineResponse20055 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20055() *InlineResponse20055 {
	this := InlineResponse20055{}
	return &this
}

// NewInlineResponse20055WithDefaults instantiates a new InlineResponse20055 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20055WithDefaults() *InlineResponse20055 {
	this := InlineResponse20055{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse20055) GetRules() []InlineResponse20055Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse20055Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055) GetRulesOk() ([]InlineResponse20055Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse20055) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse20055Rules and assigns it to the Rules field.
func (o *InlineResponse20055) SetRules(v []InlineResponse20055Rules) {
	o.Rules = v
}

func (o InlineResponse20055) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20055 struct {
	value *InlineResponse20055
	isSet bool
}

func (v NullableInlineResponse20055) Get() *InlineResponse20055 {
	return v.value
}

func (v *NullableInlineResponse20055) Set(val *InlineResponse20055) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20055) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20055) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20055(val *InlineResponse20055) *NullableInlineResponse20055 {
	return &NullableInlineResponse20055{value: val, isSet: true}
}

func (v NullableInlineResponse20055) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20055) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


