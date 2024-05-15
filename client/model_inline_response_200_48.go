/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20048 struct for InlineResponse20048
type InlineResponse20048 struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []InlineResponse20047Rules `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewInlineResponse20048 instantiates a new InlineResponse20048 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20048() *InlineResponse20048 {
	this := InlineResponse20048{}
	return &this
}

// NewInlineResponse20048WithDefaults instantiates a new InlineResponse20048 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20048WithDefaults() *InlineResponse20048 {
	this := InlineResponse20048{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse20048) GetRules() []InlineResponse20047Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse20047Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20048) GetRulesOk() ([]InlineResponse20047Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse20048) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse20047Rules and assigns it to the Rules field.
func (o *InlineResponse20048) SetRules(v []InlineResponse20047Rules) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *InlineResponse20048) GetSyslogDefaultRule() bool {
	if o == nil || isNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20048) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || isNil(o.SyslogDefaultRule) {
    return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *InlineResponse20048) HasSyslogDefaultRule() bool {
	if o != nil && !isNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *InlineResponse20048) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o InlineResponse20048) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20048 struct {
	value *InlineResponse20048
	isSet bool
}

func (v NullableInlineResponse20048) Get() *InlineResponse20048 {
	return v.value
}

func (v *NullableInlineResponse20048) Set(val *InlineResponse20048) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20048) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20048) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20048(val *InlineResponse20048) *NullableInlineResponse20048 {
	return &NullableInlineResponse20048{value: val, isSet: true}
}

func (v NullableInlineResponse20048) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20048) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


