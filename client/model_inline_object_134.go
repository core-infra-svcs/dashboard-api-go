/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject134 struct for InlineObject134
type InlineObject134 struct {
	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	Rules []NetworksNetworkIdSwitchAccessControlListsRules `json:"rules"`
}

// NewInlineObject134 instantiates a new InlineObject134 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject134(rules []NetworksNetworkIdSwitchAccessControlListsRules) *InlineObject134 {
	this := InlineObject134{}
	this.Rules = rules
	return &this
}

// NewInlineObject134WithDefaults instantiates a new InlineObject134 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject134WithDefaults() *InlineObject134 {
	this := InlineObject134{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject134) GetRules() []NetworksNetworkIdSwitchAccessControlListsRules {
	if o == nil {
		var ret []NetworksNetworkIdSwitchAccessControlListsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject134) GetRulesOk() ([]NetworksNetworkIdSwitchAccessControlListsRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject134) SetRules(v []NetworksNetworkIdSwitchAccessControlListsRules) {
	o.Rules = v
}

func (o InlineObject134) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject134 struct {
	value *InlineObject134
	isSet bool
}

func (v NullableInlineObject134) Get() *InlineObject134 {
	return v.value
}

func (v *NullableInlineObject134) Set(val *InlineObject134) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject134) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject134) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject134(val *InlineObject134) *NullableInlineObject134 {
	return &NullableInlineObject134{value: val, isSet: true}
}

func (v NullableInlineObject134) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject134) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


