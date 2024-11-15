/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject45 struct for InlineObject45
type InlineObject45 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewInlineObject45 instantiates a new InlineObject45 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject45() *InlineObject45 {
	this := InlineObject45{}
	return &this
}

// NewInlineObject45WithDefaults instantiates a new InlineObject45 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject45WithDefaults() *InlineObject45 {
	this := InlineObject45{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject45) GetRules() []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject45) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject45) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject45) SetRules(v []NetworksNetworkIdApplianceFirewallCellularFirewallRulesRules) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *InlineObject45) GetSyslogDefaultRule() bool {
	if o == nil || isNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject45) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogDefaultRule) {
    return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *InlineObject45) HasSyslogDefaultRule() bool {
	if o != nil && !isNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *InlineObject45) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o InlineObject45) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject45 struct {
	value *InlineObject45
	isSet bool
}

func (v NullableInlineObject45) Get() *InlineObject45 {
	return v.value
}

func (v *NullableInlineObject45) Set(val *InlineObject45) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject45) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject45) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject45(val *InlineObject45) *NullableInlineObject45 {
	return &NullableInlineObject45{value: val, isSet: true}
}

func (v NullableInlineObject45) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject45) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


