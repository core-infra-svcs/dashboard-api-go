/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject128 struct for InlineObject128
type InlineObject128 struct {
	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	Rules []NetworksNetworkIdSwitchAccessControlListsRules `json:"rules"`
}

// NewInlineObject128 instantiates a new InlineObject128 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject128(rules []NetworksNetworkIdSwitchAccessControlListsRules) *InlineObject128 {
	this := InlineObject128{}
	this.Rules = rules
	return &this
}

// NewInlineObject128WithDefaults instantiates a new InlineObject128 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject128WithDefaults() *InlineObject128 {
	this := InlineObject128{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject128) GetRules() []NetworksNetworkIdSwitchAccessControlListsRules {
	if o == nil {
		var ret []NetworksNetworkIdSwitchAccessControlListsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject128) GetRulesOk() ([]NetworksNetworkIdSwitchAccessControlListsRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject128) SetRules(v []NetworksNetworkIdSwitchAccessControlListsRules) {
	o.Rules = v
}

func (o InlineObject128) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject128 struct {
	value *InlineObject128
	isSet bool
}

func (v NullableInlineObject128) Get() *InlineObject128 {
	return v.value
}

func (v *NullableInlineObject128) Set(val *InlineObject128) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject128) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject128) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject128(val *InlineObject128) *NullableInlineObject128 {
	return &NullableInlineObject128{value: val, isSet: true}
}

func (v NullableInlineObject128) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject128) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


