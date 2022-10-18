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

// InlineObject31 struct for InlineObject31
type InlineObject31 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules *[]NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
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
	if o == nil || o.Rules == nil {
		var ret []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetRulesOk() (*[]NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject31) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject31) SetRules(v []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) {
	o.Rules = &v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *InlineObject31) GetSyslogDefaultRule() bool {
	if o == nil || o.SyslogDefaultRule == nil {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject31) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || o.SyslogDefaultRule == nil {
		return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *InlineObject31) HasSyslogDefaultRule() bool {
	if o != nil && o.SyslogDefaultRule != nil {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *InlineObject31) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o InlineObject31) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.SyslogDefaultRule != nil {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
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


