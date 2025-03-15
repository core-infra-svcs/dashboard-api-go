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

// InlineResponse200205 struct for InlineResponse200205
type InlineResponse200205 struct {
	// Whether traffic shaping rules are applied to clients on your SSID.
	TrafficShapingEnabled *bool `json:"trafficShapingEnabled,omitempty"`
	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled *bool `json:"defaultRulesEnabled,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	Rules []InlineResponse200205Rules `json:"rules,omitempty"`
}

// NewInlineResponse200205 instantiates a new InlineResponse200205 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200205() *InlineResponse200205 {
	this := InlineResponse200205{}
	return &this
}

// NewInlineResponse200205WithDefaults instantiates a new InlineResponse200205 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200205WithDefaults() *InlineResponse200205 {
	this := InlineResponse200205{}
	return &this
}

// GetTrafficShapingEnabled returns the TrafficShapingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200205) GetTrafficShapingEnabled() bool {
	if o == nil || isNil(o.TrafficShapingEnabled) {
		var ret bool
		return ret
	}
	return *o.TrafficShapingEnabled
}

// GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetTrafficShapingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.TrafficShapingEnabled) {
    return nil, false
	}
	return o.TrafficShapingEnabled, true
}

// HasTrafficShapingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200205) HasTrafficShapingEnabled() bool {
	if o != nil && !isNil(o.TrafficShapingEnabled) {
		return true
	}

	return false
}

// SetTrafficShapingEnabled gets a reference to the given bool and assigns it to the TrafficShapingEnabled field.
func (o *InlineResponse200205) SetTrafficShapingEnabled(v bool) {
	o.TrafficShapingEnabled = &v
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *InlineResponse200205) GetDefaultRulesEnabled() bool {
	if o == nil || isNil(o.DefaultRulesEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultRulesEnabled) {
    return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *InlineResponse200205) HasDefaultRulesEnabled() bool {
	if o != nil && !isNil(o.DefaultRulesEnabled) {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *InlineResponse200205) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse200205) GetRules() []InlineResponse200205Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200205Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200205) GetRulesOk() ([]InlineResponse200205Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse200205) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200205Rules and assigns it to the Rules field.
func (o *InlineResponse200205) SetRules(v []InlineResponse200205Rules) {
	o.Rules = v
}

func (o InlineResponse200205) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200205 struct {
	value *InlineResponse200205
	isSet bool
}

func (v NullableInlineResponse200205) Get() *InlineResponse200205 {
	return v.value
}

func (v *NullableInlineResponse200205) Set(val *InlineResponse200205) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200205) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200205) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200205(val *InlineResponse200205) *NullableInlineResponse200205 {
	return &NullableInlineResponse200205{value: val, isSet: true}
}

func (v NullableInlineResponse200205) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200205) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


