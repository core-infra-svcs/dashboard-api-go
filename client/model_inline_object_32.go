/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject32 struct for InlineObject32
type InlineObject32 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"`
}

// NewInlineObject32 instantiates a new InlineObject32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject32() *InlineObject32 {
	this := InlineObject32{}
	return &this
}

// NewInlineObject32WithDefaults instantiates a new InlineObject32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject32WithDefaults() *InlineObject32 {
	this := InlineObject32{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject32) GetRules() []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject32) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject32) SetRules(v []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) {
	o.Rules = v
}

func (o InlineObject32) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject32 struct {
	value *InlineObject32
	isSet bool
}

func (v NullableInlineObject32) Get() *InlineObject32 {
	return v.value
}

func (v *NullableInlineObject32) Set(val *InlineObject32) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject32) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject32(val *InlineObject32) *NullableInlineObject32 {
	return &NullableInlineObject32{value: val, isSet: true}
}

func (v NullableInlineObject32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


