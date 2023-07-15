/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject31 struct for InlineObject31
type InlineObject31 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"`
}

// NewInlineObject31 instantiates a new InlineObject31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject31() *InlineObject31 {
	this := InlineObject31{}
	return &this
}

// NewInlineObject31WithDefaults instantiates a new InlineObject31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject31WithDefaults() *InlineObject31 {
	this := InlineObject31{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject31) GetRules() []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject31) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject31) SetRules(v []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) {
	o.Rules = v
}

func (o InlineObject31) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject31 struct {
	value *InlineObject31
	isSet bool
}

func (v NullableInlineObject31) Get() *InlineObject31 {
	return v.value
}

func (v *NullableInlineObject31) Set(val *InlineObject31) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject31) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject31(val *InlineObject31) *NullableInlineObject31 {
	return &NullableInlineObject31{value: val, isSet: true}
}

func (v NullableInlineObject31) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


