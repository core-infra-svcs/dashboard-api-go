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

// InlineObject49 struct for InlineObject49
type InlineObject49 struct {
	// An array of 1:Many nat rules
	Rules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules `json:"rules"`
}

// NewInlineObject49 instantiates a new InlineObject49 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject49(rules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) *InlineObject49 {
	this := InlineObject49{}
	this.Rules = rules
	return &this
}

// NewInlineObject49WithDefaults instantiates a new InlineObject49 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject49WithDefaults() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject49) GetRules() []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject49) SetRules(v []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) {
	o.Rules = v
}

func (o InlineObject49) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject49 struct {
	value *InlineObject49
	isSet bool
}

func (v NullableInlineObject49) Get() *InlineObject49 {
	return v.value
}

func (v *NullableInlineObject49) Set(val *InlineObject49) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject49) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject49) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject49(val *InlineObject49) *NullableInlineObject49 {
	return &NullableInlineObject49{value: val, isSet: true}
}

func (v NullableInlineObject49) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject49) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


