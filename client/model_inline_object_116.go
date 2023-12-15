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

// InlineObject116 struct for InlineObject116
type InlineObject116 struct {
	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	Rules []NetworksNetworkIdSwitchAccessControlListsRules `json:"rules"`
}

// NewInlineObject116 instantiates a new InlineObject116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject116(rules []NetworksNetworkIdSwitchAccessControlListsRules) *InlineObject116 {
	this := InlineObject116{}
	this.Rules = rules
	return &this
}

// NewInlineObject116WithDefaults instantiates a new InlineObject116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject116WithDefaults() *InlineObject116 {
	this := InlineObject116{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject116) GetRules() []NetworksNetworkIdSwitchAccessControlListsRules {
	if o == nil {
		var ret []NetworksNetworkIdSwitchAccessControlListsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetRulesOk() ([]NetworksNetworkIdSwitchAccessControlListsRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject116) SetRules(v []NetworksNetworkIdSwitchAccessControlListsRules) {
	o.Rules = v
}

func (o InlineObject116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject116 struct {
	value *InlineObject116
	isSet bool
}

func (v NullableInlineObject116) Get() *InlineObject116 {
	return v.value
}

func (v *NullableInlineObject116) Set(val *InlineObject116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject116(val *InlineObject116) *NullableInlineObject116 {
	return &NullableInlineObject116{value: val, isSet: true}
}

func (v NullableInlineObject116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


