/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject33 struct for InlineObject33
type InlineObject33 struct {
	// An ordered array of the MX L7 firewall rules
	Rules *[]NetworksNetworkIdApplianceFirewallL7FirewallRulesRules `json:"rules,omitempty"`
}

// NewInlineObject33 instantiates a new InlineObject33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject33() *InlineObject33 {
	this := InlineObject33{}
	return &this
}

// NewInlineObject33WithDefaults instantiates a new InlineObject33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject33WithDefaults() *InlineObject33 {
	this := InlineObject33{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject33) GetRules() []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules {
	if o == nil || o.Rules == nil {
		var ret []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject33) GetRulesOk() (*[]NetworksNetworkIdApplianceFirewallL7FirewallRulesRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject33) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject33) SetRules(v []NetworksNetworkIdApplianceFirewallL7FirewallRulesRules) {
	o.Rules = &v
}

func (o InlineObject33) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject33 struct {
	value *InlineObject33
	isSet bool
}

func (v NullableInlineObject33) Get() *InlineObject33 {
	return v.value
}

func (v *NullableInlineObject33) Set(val *InlineObject33) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject33) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject33(val *InlineObject33) *NullableInlineObject33 {
	return &NullableInlineObject33{value: val, isSet: true}
}

func (v NullableInlineObject33) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


