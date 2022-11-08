/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject34 struct for InlineObject34
type InlineObject34 struct {
	// An ordered array of the MX L7 firewall rules
	Rules []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules `json:"rules,omitempty"`
}

// NewInlineObject34 instantiates a new InlineObject34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject34() *InlineObject34 {
	this := InlineObject34{}
	return &this
}

// NewInlineObject34WithDefaults instantiates a new InlineObject34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject34WithDefaults() *InlineObject34 {
	this := InlineObject34{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject34) GetRules() []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject34) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallL7FirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject34) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject34) SetRules(v []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) {
	o.Rules = v
}

func (o InlineObject34) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject34 struct {
	value *InlineObject34
	isSet bool
}

func (v NullableInlineObject34) Get() *InlineObject34 {
	return v.value
}

func (v *NullableInlineObject34) Set(val *InlineObject34) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject34) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject34(val *InlineObject34) *NullableInlineObject34 {
	return &NullableInlineObject34{value: val, isSet: true}
}

func (v NullableInlineObject34) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


