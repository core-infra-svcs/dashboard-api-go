/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject47 struct for InlineObject47
type InlineObject47 struct {
	// An array of 1:Many nat rules
	Rules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules `json:"rules"`
}

// NewInlineObject47 instantiates a new InlineObject47 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject47(rules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) *InlineObject47 {
	this := InlineObject47{}
	this.Rules = rules
	return &this
}

// NewInlineObject47WithDefaults instantiates a new InlineObject47 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject47WithDefaults() *InlineObject47 {
	this := InlineObject47{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject47) GetRules() []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject47) SetRules(v []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) {
	o.Rules = v
}

func (o InlineObject47) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject47 struct {
	value *InlineObject47
	isSet bool
}

func (v NullableInlineObject47) Get() *InlineObject47 {
	return v.value
}

func (v *NullableInlineObject47) Set(val *InlineObject47) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject47) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject47) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject47(val *InlineObject47) *NullableInlineObject47 {
	return &NullableInlineObject47{value: val, isSet: true}
}

func (v NullableInlineObject47) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject47) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


