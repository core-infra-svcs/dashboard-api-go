/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject177 struct for InlineObject177
type InlineObject177 struct {
	// If true, Bonjour forwarding is enabled on this SSID.
	Enabled *bool `json:"enabled,omitempty"`
	// List of bonjour forwarding rules.
	Rules []NetworksNetworkIdGroupPoliciesBonjourForwardingRules `json:"rules,omitempty"`
	Exception *NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException `json:"exception,omitempty"`
}

// NewInlineObject177 instantiates a new InlineObject177 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject177() *InlineObject177 {
	this := InlineObject177{}
	return &this
}

// NewInlineObject177WithDefaults instantiates a new InlineObject177 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject177WithDefaults() *InlineObject177 {
	this := InlineObject177{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject177) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject177) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject177) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject177) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdGroupPoliciesBonjourForwardingRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetRulesOk() ([]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject177) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdGroupPoliciesBonjourForwardingRules and assigns it to the Rules field.
func (o *InlineObject177) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules) {
	o.Rules = v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *InlineObject177) GetException() NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException {
	if o == nil || isNil(o.Exception) {
		var ret NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException
		return ret
	}
	return *o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetExceptionOk() (*NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException, bool) {
	if o == nil || isNil(o.Exception) {
    return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *InlineObject177) HasException() bool {
	if o != nil && !isNil(o.Exception) {
		return true
	}

	return false
}

// SetException gets a reference to the given NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException and assigns it to the Exception field.
func (o *InlineObject177) SetException(v NetworksNetworkIdWirelessSsidsNumberBonjourForwardingException) {
	o.Exception = &v
}

func (o InlineObject177) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.Exception) {
		toSerialize["exception"] = o.Exception
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject177 struct {
	value *InlineObject177
	isSet bool
}

func (v NullableInlineObject177) Get() *InlineObject177 {
	return v.value
}

func (v *NullableInlineObject177) Set(val *InlineObject177) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject177) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject177) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject177(val *InlineObject177) *NullableInlineObject177 {
	return &NullableInlineObject177{value: val, isSet: true}
}

func (v NullableInlineObject177) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject177) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


