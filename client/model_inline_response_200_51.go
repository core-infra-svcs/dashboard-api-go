/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20051 struct for InlineResponse20051
type InlineResponse20051 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []InlineResponse20050Rules `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewInlineResponse20051 instantiates a new InlineResponse20051 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20051() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// NewInlineResponse20051WithDefaults instantiates a new InlineResponse20051 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20051WithDefaults() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse20051) GetRules() []InlineResponse20050Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse20050Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetRulesOk() ([]InlineResponse20050Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse20051) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse20050Rules and assigns it to the Rules field.
func (o *InlineResponse20051) SetRules(v []InlineResponse20050Rules) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *InlineResponse20051) GetSyslogDefaultRule() bool {
	if o == nil || isNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogDefaultRule) {
    return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *InlineResponse20051) HasSyslogDefaultRule() bool {
	if o != nil && !isNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *InlineResponse20051) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o InlineResponse20051) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20051 struct {
	value *InlineResponse20051
	isSet bool
}

func (v NullableInlineResponse20051) Get() *InlineResponse20051 {
	return v.value
}

func (v *NullableInlineResponse20051) Set(val *InlineResponse20051) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20051) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20051) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20051(val *InlineResponse20051) *NullableInlineResponse20051 {
	return &NullableInlineResponse20051{value: val, isSet: true}
}

func (v NullableInlineResponse20051) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20051) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


