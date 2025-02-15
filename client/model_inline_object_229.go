/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject229 struct for InlineObject229
type InlineObject229 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewInlineObject229 instantiates a new InlineObject229 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject229() *InlineObject229 {
	this := InlineObject229{}
	return &this
}

// NewInlineObject229WithDefaults instantiates a new InlineObject229 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject229WithDefaults() *InlineObject229 {
	this := InlineObject229{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject229) GetRules() []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetRulesOk() ([]OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject229) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject229) SetRules(v []OrganizationsOrganizationIdApplianceVpnVpnFirewallRulesRules) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *InlineObject229) GetSyslogDefaultRule() bool {
	if o == nil || isNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject229) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogDefaultRule) {
    return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *InlineObject229) HasSyslogDefaultRule() bool {
	if o != nil && !isNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *InlineObject229) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o InlineObject229) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject229 struct {
	value *InlineObject229
	isSet bool
}

func (v NullableInlineObject229) Get() *InlineObject229 {
	return v.value
}

func (v *NullableInlineObject229) Set(val *InlineObject229) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject229) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject229) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject229(val *InlineObject229) *NullableInlineObject229 {
	return &NullableInlineObject229{value: val, isSet: true}
}

func (v NullableInlineObject229) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject229) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


