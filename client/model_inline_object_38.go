/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject38 struct for InlineObject38
type InlineObject38 struct {
	// An array of 1:1 nat rules
	Rules []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules `json:"rules"`
}

// NewInlineObject38 instantiates a new InlineObject38 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject38(rules []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) *InlineObject38 {
	this := InlineObject38{}
	this.Rules = rules
	return &this
}

// NewInlineObject38WithDefaults instantiates a new InlineObject38 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject38WithDefaults() *InlineObject38 {
	this := InlineObject38{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject38) GetRules() []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject38) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject38) SetRules(v []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) {
	o.Rules = v
}

func (o InlineObject38) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject38 struct {
	value *InlineObject38
	isSet bool
}

func (v NullableInlineObject38) Get() *InlineObject38 {
	return v.value
}

func (v *NullableInlineObject38) Set(val *InlineObject38) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject38) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject38) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject38(val *InlineObject38) *NullableInlineObject38 {
	return &NullableInlineObject38{value: val, isSet: true}
}

func (v NullableInlineObject38) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject38) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


