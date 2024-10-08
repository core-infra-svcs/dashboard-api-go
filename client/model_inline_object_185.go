/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject185 struct for InlineObject185
type InlineObject185 struct {
	// If true, Bonjour forwarding is enabled on this SSID.
	Enabled *bool `json:"enabled,omitempty"`
	// List of bonjour forwarding rules.
	Rules []NetworksNetworkIdGroupPoliciesBonjourForwardingRules `json:"rules,omitempty"`
	Exception *InlineResponse200188Exception `json:"exception,omitempty"`
}

// NewInlineObject185 instantiates a new InlineObject185 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject185() *InlineObject185 {
	this := InlineObject185{}
	return &this
}

// NewInlineObject185WithDefaults instantiates a new InlineObject185 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject185WithDefaults() *InlineObject185 {
	this := InlineObject185{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject185) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject185) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject185) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject185) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdGroupPoliciesBonjourForwardingRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetRulesOk() ([]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject185) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdGroupPoliciesBonjourForwardingRules and assigns it to the Rules field.
func (o *InlineObject185) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules) {
	o.Rules = v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *InlineObject185) GetException() InlineResponse200188Exception {
	if o == nil || isNil(o.Exception) {
		var ret InlineResponse200188Exception
		return ret
	}
	return *o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject185) GetExceptionOk() (*InlineResponse200188Exception, bool) {
	if o == nil || isNil(o.Exception) {
    return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *InlineObject185) HasException() bool {
	if o != nil && !isNil(o.Exception) {
		return true
	}

	return false
}

// SetException gets a reference to the given InlineResponse200188Exception and assigns it to the Exception field.
func (o *InlineObject185) SetException(v InlineResponse200188Exception) {
	o.Exception = &v
}

func (o InlineObject185) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject185 struct {
	value *InlineObject185
	isSet bool
}

func (v NullableInlineObject185) Get() *InlineObject185 {
	return v.value
}

func (v *NullableInlineObject185) Set(val *InlineObject185) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject185) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject185) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject185(val *InlineObject185) *NullableInlineObject185 {
	return &NullableInlineObject185{value: val, isSet: true}
}

func (v NullableInlineObject185) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject185) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


