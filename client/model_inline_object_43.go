/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject43 struct for InlineObject43
type InlineObject43 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"`
}

// NewInlineObject43 instantiates a new InlineObject43 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject43() *InlineObject43 {
	this := InlineObject43{}
	return &this
}

// NewInlineObject43WithDefaults instantiates a new InlineObject43 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject43WithDefaults() *InlineObject43 {
	this := InlineObject43{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject43) GetRules() []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject43) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject43) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject43) SetRules(v []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) {
	o.Rules = v
}

func (o InlineObject43) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject43 struct {
	value *InlineObject43
	isSet bool
}

func (v NullableInlineObject43) Get() *InlineObject43 {
	return v.value
}

func (v *NullableInlineObject43) Set(val *InlineObject43) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject43) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject43) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject43(val *InlineObject43) *NullableInlineObject43 {
	return &NullableInlineObject43{value: val, isSet: true}
}

func (v NullableInlineObject43) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject43) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


