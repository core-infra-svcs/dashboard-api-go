/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject195 struct for InlineObject195
type InlineObject195 struct {
	// Whether traffic shaping rules are applied to clients on your SSID.
	TrafficShapingEnabled *bool `json:"trafficShapingEnabled,omitempty"`
	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled *bool `json:"defaultRulesEnabled,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	Rules []InlineResponse200195Rules `json:"rules,omitempty"`
}

// NewInlineObject195 instantiates a new InlineObject195 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject195() *InlineObject195 {
	this := InlineObject195{}
	return &this
}

// NewInlineObject195WithDefaults instantiates a new InlineObject195 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject195WithDefaults() *InlineObject195 {
	this := InlineObject195{}
	return &this
}

// GetTrafficShapingEnabled returns the TrafficShapingEnabled field value if set, zero value otherwise.
func (o *InlineObject195) GetTrafficShapingEnabled() bool {
	if o == nil || isNil(o.TrafficShapingEnabled) {
		var ret bool
		return ret
	}
	return *o.TrafficShapingEnabled
}

// GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject195) GetTrafficShapingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.TrafficShapingEnabled) {
    return nil, false
	}
	return o.TrafficShapingEnabled, true
}

// HasTrafficShapingEnabled returns a boolean if a field has been set.
func (o *InlineObject195) HasTrafficShapingEnabled() bool {
	if o != nil && !isNil(o.TrafficShapingEnabled) {
		return true
	}

	return false
}

// SetTrafficShapingEnabled gets a reference to the given bool and assigns it to the TrafficShapingEnabled field.
func (o *InlineObject195) SetTrafficShapingEnabled(v bool) {
	o.TrafficShapingEnabled = &v
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *InlineObject195) GetDefaultRulesEnabled() bool {
	if o == nil || isNil(o.DefaultRulesEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject195) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultRulesEnabled) {
    return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *InlineObject195) HasDefaultRulesEnabled() bool {
	if o != nil && !isNil(o.DefaultRulesEnabled) {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *InlineObject195) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject195) GetRules() []InlineResponse200195Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200195Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject195) GetRulesOk() ([]InlineResponse200195Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject195) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200195Rules and assigns it to the Rules field.
func (o *InlineObject195) SetRules(v []InlineResponse200195Rules) {
	o.Rules = v
}

func (o InlineObject195) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TrafficShapingEnabled) {
		toSerialize["trafficShapingEnabled"] = o.TrafficShapingEnabled
	}
	if !isNil(o.DefaultRulesEnabled) {
		toSerialize["defaultRulesEnabled"] = o.DefaultRulesEnabled
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject195 struct {
	value *InlineObject195
	isSet bool
}

func (v NullableInlineObject195) Get() *InlineObject195 {
	return v.value
}

func (v *NullableInlineObject195) Set(val *InlineObject195) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject195) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject195) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject195(val *InlineObject195) *NullableInlineObject195 {
	return &NullableInlineObject195{value: val, isSet: true}
}

func (v NullableInlineObject195) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject195) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


