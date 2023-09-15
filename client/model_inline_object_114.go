/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject114 struct for InlineObject114
type InlineObject114 struct {
	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	Rules []NetworksNetworkIdSwitchAccessControlListsRules `json:"rules"`
}

// NewInlineObject114 instantiates a new InlineObject114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject114(rules []NetworksNetworkIdSwitchAccessControlListsRules) *InlineObject114 {
	this := InlineObject114{}
	this.Rules = rules
	return &this
}

// NewInlineObject114WithDefaults instantiates a new InlineObject114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject114WithDefaults() *InlineObject114 {
	this := InlineObject114{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject114) GetRules() []NetworksNetworkIdSwitchAccessControlListsRules {
	if o == nil {
		var ret []NetworksNetworkIdSwitchAccessControlListsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject114) GetRulesOk() ([]NetworksNetworkIdSwitchAccessControlListsRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject114) SetRules(v []NetworksNetworkIdSwitchAccessControlListsRules) {
	o.Rules = v
}

func (o InlineObject114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject114 struct {
	value *InlineObject114
	isSet bool
}

func (v NullableInlineObject114) Get() *InlineObject114 {
	return v.value
}

func (v *NullableInlineObject114) Set(val *InlineObject114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject114(val *InlineObject114) *NullableInlineObject114 {
	return &NullableInlineObject114{value: val, isSet: true}
}

func (v NullableInlineObject114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


