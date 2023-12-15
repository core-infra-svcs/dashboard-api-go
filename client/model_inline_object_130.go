/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject130 struct for InlineObject130
type InlineObject130 struct {
	// A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
	RuleIds []string `json:"ruleIds"`
}

// NewInlineObject130 instantiates a new InlineObject130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject130(ruleIds []string) *InlineObject130 {
	this := InlineObject130{}
	this.RuleIds = ruleIds
	return &this
}

// NewInlineObject130WithDefaults instantiates a new InlineObject130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject130WithDefaults() *InlineObject130 {
	this := InlineObject130{}
	return &this
}

// GetRuleIds returns the RuleIds field value
func (o *InlineObject130) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *InlineObject130) SetRuleIds(v []string) {
	o.RuleIds = v
}

func (o InlineObject130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleIds"] = o.RuleIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject130 struct {
	value *InlineObject130
	isSet bool
}

func (v NullableInlineObject130) Get() *InlineObject130 {
	return v.value
}

func (v *NullableInlineObject130) Set(val *InlineObject130) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject130) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject130(val *InlineObject130) *NullableInlineObject130 {
	return &NullableInlineObject130{value: val, isSet: true}
}

func (v NullableInlineObject130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


