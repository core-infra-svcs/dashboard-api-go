/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject60 struct for InlineObject60
type InlineObject60 struct {
	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled *bool `json:"defaultRulesEnabled,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	Rules []NetworksNetworkIdApplianceTrafficShapingRulesRules `json:"rules,omitempty"`
}

// NewInlineObject60 instantiates a new InlineObject60 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject60() *InlineObject60 {
	this := InlineObject60{}
	return &this
}

// NewInlineObject60WithDefaults instantiates a new InlineObject60 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject60WithDefaults() *InlineObject60 {
	this := InlineObject60{}
	return &this
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *InlineObject60) GetDefaultRulesEnabled() bool {
	if o == nil || isNil(o.DefaultRulesEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject60) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultRulesEnabled) {
    return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *InlineObject60) HasDefaultRulesEnabled() bool {
	if o != nil && !isNil(o.DefaultRulesEnabled) {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *InlineObject60) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject60) GetRules() []NetworksNetworkIdApplianceTrafficShapingRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdApplianceTrafficShapingRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject60) GetRulesOk() ([]NetworksNetworkIdApplianceTrafficShapingRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject60) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingRulesRules and assigns it to the Rules field.
func (o *InlineObject60) SetRules(v []NetworksNetworkIdApplianceTrafficShapingRulesRules) {
	o.Rules = v
}

func (o InlineObject60) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultRulesEnabled) {
		toSerialize["defaultRulesEnabled"] = o.DefaultRulesEnabled
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject60 struct {
	value *InlineObject60
	isSet bool
}

func (v NullableInlineObject60) Get() *InlineObject60 {
	return v.value
}

func (v *NullableInlineObject60) Set(val *InlineObject60) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject60) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject60) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject60(val *InlineObject60) *NullableInlineObject60 {
	return &NullableInlineObject60{value: val, isSet: true}
}

func (v NullableInlineObject60) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject60) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


