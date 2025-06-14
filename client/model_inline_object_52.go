/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject52 struct for InlineObject52
type InlineObject52 struct {
	// An array of port forwarding params
	Rules []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules `json:"rules"`
}

// NewInlineObject52 instantiates a new InlineObject52 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject52(rules []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules) *InlineObject52 {
	this := InlineObject52{}
	this.Rules = rules
	return &this
}

// NewInlineObject52WithDefaults instantiates a new InlineObject52 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject52WithDefaults() *InlineObject52 {
	this := InlineObject52{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject52) GetRules() []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject52) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallPortForwardingRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject52) SetRules(v []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules) {
	o.Rules = v
}

func (o InlineObject52) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject52 struct {
	value *InlineObject52
	isSet bool
}

func (v NullableInlineObject52) Get() *InlineObject52 {
	return v.value
}

func (v *NullableInlineObject52) Set(val *InlineObject52) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject52) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject52) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject52(val *InlineObject52) *NullableInlineObject52 {
	return &NullableInlineObject52{value: val, isSet: true}
}

func (v NullableInlineObject52) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject52) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


