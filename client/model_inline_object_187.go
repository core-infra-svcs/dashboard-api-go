/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject187 struct for InlineObject187
type InlineObject187 struct {
	// An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration.
	Rules []InlineResponse200188Rules `json:"rules,omitempty"`
}

// NewInlineObject187 instantiates a new InlineObject187 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject187() *InlineObject187 {
	this := InlineObject187{}
	return &this
}

// NewInlineObject187WithDefaults instantiates a new InlineObject187 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject187WithDefaults() *InlineObject187 {
	this := InlineObject187{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject187) GetRules() []InlineResponse200188Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200188Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject187) GetRulesOk() ([]InlineResponse200188Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject187) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200188Rules and assigns it to the Rules field.
func (o *InlineObject187) SetRules(v []InlineResponse200188Rules) {
	o.Rules = v
}

func (o InlineObject187) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject187 struct {
	value *InlineObject187
	isSet bool
}

func (v NullableInlineObject187) Get() *InlineObject187 {
	return v.value
}

func (v *NullableInlineObject187) Set(val *InlineObject187) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject187) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject187) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject187(val *InlineObject187) *NullableInlineObject187 {
	return &NullableInlineObject187{value: val, isSet: true}
}

func (v NullableInlineObject187) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject187) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


