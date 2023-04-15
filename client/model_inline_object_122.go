/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject122 struct for InlineObject122
type InlineObject122 struct {
	// A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
	RuleIds []string `json:"ruleIds"`
}

// NewInlineObject122 instantiates a new InlineObject122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject122(ruleIds []string) *InlineObject122 {
	this := InlineObject122{}
	this.RuleIds = ruleIds
	return &this
}

// NewInlineObject122WithDefaults instantiates a new InlineObject122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject122WithDefaults() *InlineObject122 {
	this := InlineObject122{}
	return &this
}

// GetRuleIds returns the RuleIds field value
func (o *InlineObject122) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject122) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *InlineObject122) SetRuleIds(v []string) {
	o.RuleIds = v
}

func (o InlineObject122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleIds"] = o.RuleIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject122 struct {
	value *InlineObject122
	isSet bool
}

func (v NullableInlineObject122) Get() *InlineObject122 {
	return v.value
}

func (v *NullableInlineObject122) Set(val *InlineObject122) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject122) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject122(val *InlineObject122) *NullableInlineObject122 {
	return &NullableInlineObject122{value: val, isSet: true}
}

func (v NullableInlineObject122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


