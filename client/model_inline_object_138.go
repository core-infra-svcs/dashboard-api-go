/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject138 struct for InlineObject138
type InlineObject138 struct {
	// A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
	RuleIds []string `json:"ruleIds"`
}

// NewInlineObject138 instantiates a new InlineObject138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject138(ruleIds []string) *InlineObject138 {
	this := InlineObject138{}
	this.RuleIds = ruleIds
	return &this
}

// NewInlineObject138WithDefaults instantiates a new InlineObject138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject138WithDefaults() *InlineObject138 {
	this := InlineObject138{}
	return &this
}

// GetRuleIds returns the RuleIds field value
func (o *InlineObject138) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *InlineObject138) SetRuleIds(v []string) {
	o.RuleIds = v
}

func (o InlineObject138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleIds"] = o.RuleIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject138 struct {
	value *InlineObject138
	isSet bool
}

func (v NullableInlineObject138) Get() *InlineObject138 {
	return v.value
}

func (v *NullableInlineObject138) Set(val *InlineObject138) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject138) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject138(val *InlineObject138) *NullableInlineObject138 {
	return &NullableInlineObject138{value: val, isSet: true}
}

func (v NullableInlineObject138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


