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

// InlineObject201 struct for InlineObject201
type InlineObject201 struct {
	// Whether traffic shaping rules are applied to clients on your SSID.
	TrafficShapingEnabled *bool `json:"trafficShapingEnabled,omitempty"`
	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled *bool `json:"defaultRulesEnabled,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	Rules []InlineResponse200205Rules `json:"rules,omitempty"`
}

// NewInlineObject201 instantiates a new InlineObject201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject201() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// NewInlineObject201WithDefaults instantiates a new InlineObject201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject201WithDefaults() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// GetTrafficShapingEnabled returns the TrafficShapingEnabled field value if set, zero value otherwise.
func (o *InlineObject201) GetTrafficShapingEnabled() bool {
	if o == nil || isNil(o.TrafficShapingEnabled) {
		var ret bool
		return ret
	}
	return *o.TrafficShapingEnabled
}

// GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetTrafficShapingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.TrafficShapingEnabled) {
    return nil, false
	}
	return o.TrafficShapingEnabled, true
}

// HasTrafficShapingEnabled returns a boolean if a field has been set.
func (o *InlineObject201) HasTrafficShapingEnabled() bool {
	if o != nil && !isNil(o.TrafficShapingEnabled) {
		return true
	}

	return false
}

// SetTrafficShapingEnabled gets a reference to the given bool and assigns it to the TrafficShapingEnabled field.
func (o *InlineObject201) SetTrafficShapingEnabled(v bool) {
	o.TrafficShapingEnabled = &v
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *InlineObject201) GetDefaultRulesEnabled() bool {
	if o == nil || isNil(o.DefaultRulesEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultRulesEnabled) {
    return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *InlineObject201) HasDefaultRulesEnabled() bool {
	if o != nil && !isNil(o.DefaultRulesEnabled) {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *InlineObject201) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject201) GetRules() []InlineResponse200205Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200205Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetRulesOk() ([]InlineResponse200205Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject201) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200205Rules and assigns it to the Rules field.
func (o *InlineObject201) SetRules(v []InlineResponse200205Rules) {
	o.Rules = v
}

func (o InlineObject201) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject201 struct {
	value *InlineObject201
	isSet bool
}

func (v NullableInlineObject201) Get() *InlineObject201 {
	return v.value
}

func (v *NullableInlineObject201) Set(val *InlineObject201) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject201) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject201(val *InlineObject201) *NullableInlineObject201 {
	return &NullableInlineObject201{value: val, isSet: true}
}

func (v NullableInlineObject201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


