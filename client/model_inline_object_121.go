/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject121 struct for InlineObject121
type InlineObject121 struct {
	// A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
	RuleIds []string `json:"ruleIds"`
}

// NewInlineObject121 instantiates a new InlineObject121 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject121(ruleIds []string) *InlineObject121 {
	this := InlineObject121{}
	this.RuleIds = ruleIds
	return &this
}

// NewInlineObject121WithDefaults instantiates a new InlineObject121 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject121WithDefaults() *InlineObject121 {
	this := InlineObject121{}
	return &this
}

// GetRuleIds returns the RuleIds field value
func (o *InlineObject121) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject121) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *InlineObject121) SetRuleIds(v []string) {
	o.RuleIds = v
}

func (o InlineObject121) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleIds"] = o.RuleIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject121 struct {
	value *InlineObject121
	isSet bool
}

func (v NullableInlineObject121) Get() *InlineObject121 {
	return v.value
}

func (v *NullableInlineObject121) Set(val *InlineObject121) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject121) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject121) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject121(val *InlineObject121) *NullableInlineObject121 {
	return &NullableInlineObject121{value: val, isSet: true}
}

func (v NullableInlineObject121) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject121) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


