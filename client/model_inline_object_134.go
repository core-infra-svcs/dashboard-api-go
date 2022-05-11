/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject134 struct for InlineObject134
type InlineObject134 struct {
	// If true, Bonjour forwarding is enabled on this SSID.
	Enabled *bool `json:"enabled,omitempty"`
	// List of bonjour forwarding rules.
	Rules *[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules `json:"rules,omitempty"`
}

// NewInlineObject134 instantiates a new InlineObject134 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject134() *InlineObject134 {
	this := InlineObject134{}
	return &this
}

// NewInlineObject134WithDefaults instantiates a new InlineObject134 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject134WithDefaults() *InlineObject134 {
	this := InlineObject134{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject134) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject134) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject134) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject134) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject134) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	if o == nil || o.Rules == nil {
		var ret []NetworksNetworkIdGroupPoliciesBonjourForwardingRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject134) GetRulesOk() (*[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject134) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdGroupPoliciesBonjourForwardingRules and assigns it to the Rules field.
func (o *InlineObject134) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules) {
	o.Rules = &v
}

func (o InlineObject134) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject134 struct {
	value *InlineObject134
	isSet bool
}

func (v NullableInlineObject134) Get() *InlineObject134 {
	return v.value
}

func (v *NullableInlineObject134) Set(val *InlineObject134) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject134) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject134) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject134(val *InlineObject134) *NullableInlineObject134 {
	return &NullableInlineObject134{value: val, isSet: true}
}

func (v NullableInlineObject134) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject134) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


