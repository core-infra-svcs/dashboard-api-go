/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject157 struct for InlineObject157
type InlineObject157 struct {
	// An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration.
	Rules []NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules `json:"rules,omitempty"`
}

// NewInlineObject157 instantiates a new InlineObject157 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject157() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// NewInlineObject157WithDefaults instantiates a new InlineObject157 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject157WithDefaults() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject157) GetRules() []NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetRulesOk() ([]NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject157) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject157) SetRules(v []NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) {
	o.Rules = v
}

func (o InlineObject157) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject157 struct {
	value *InlineObject157
	isSet bool
}

func (v NullableInlineObject157) Get() *InlineObject157 {
	return v.value
}

func (v *NullableInlineObject157) Set(val *InlineObject157) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject157) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject157) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject157(val *InlineObject157) *NullableInlineObject157 {
	return &NullableInlineObject157{value: val, isSet: true}
}

func (v NullableInlineObject157) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject157) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


