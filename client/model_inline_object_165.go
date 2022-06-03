/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject165 struct for InlineObject165
type InlineObject165 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules *[]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewInlineObject165 instantiates a new InlineObject165 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject165() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// NewInlineObject165WithDefaults instantiates a new InlineObject165 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject165WithDefaults() *InlineObject165 {
	this := InlineObject165{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject165) GetRules() []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules {
	if o == nil || o.Rules == nil {
		var ret []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetRulesOk() (*[]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject165) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject165) SetRules(v []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) {
	o.Rules = &v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *InlineObject165) GetSyslogDefaultRule() bool {
	if o == nil || o.SyslogDefaultRule == nil {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject165) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || o.SyslogDefaultRule == nil {
		return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *InlineObject165) HasSyslogDefaultRule() bool {
	if o != nil && o.SyslogDefaultRule != nil {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *InlineObject165) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o InlineObject165) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.SyslogDefaultRule != nil {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject165 struct {
	value *InlineObject165
	isSet bool
}

func (v NullableInlineObject165) Get() *InlineObject165 {
	return v.value
}

func (v *NullableInlineObject165) Set(val *InlineObject165) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject165) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject165) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject165(val *InlineObject165) *NullableInlineObject165 {
	return &NullableInlineObject165{value: val, isSet: true}
}

func (v NullableInlineObject165) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject165) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


