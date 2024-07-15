/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200187 struct for InlineResponse200187
type InlineResponse200187 struct {
	// If true, Bonjour forwarding is enabled on the SSID.
	Enabled *bool `json:"enabled,omitempty"`
	Exception *InlineResponse200187Exception `json:"exception,omitempty"`
	// Bonjour forwarding rules
	Rules []InlineResponse200187Rules `json:"rules,omitempty"`
}

// NewInlineResponse200187 instantiates a new InlineResponse200187 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200187() *InlineResponse200187 {
	this := InlineResponse200187{}
	return &this
}

// NewInlineResponse200187WithDefaults instantiates a new InlineResponse200187 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200187WithDefaults() *InlineResponse200187 {
	this := InlineResponse200187{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200187) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200187) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200187) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200187) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *InlineResponse200187) GetException() InlineResponse200187Exception {
	if o == nil || isNil(o.Exception) {
		var ret InlineResponse200187Exception
		return ret
	}
	return *o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200187) GetExceptionOk() (*InlineResponse200187Exception, bool) {
	if o == nil || isNil(o.Exception) {
    return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *InlineResponse200187) HasException() bool {
	if o != nil && !isNil(o.Exception) {
		return true
	}

	return false
}

// SetException gets a reference to the given InlineResponse200187Exception and assigns it to the Exception field.
func (o *InlineResponse200187) SetException(v InlineResponse200187Exception) {
	o.Exception = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse200187) GetRules() []InlineResponse200187Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200187Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200187) GetRulesOk() ([]InlineResponse200187Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse200187) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200187Rules and assigns it to the Rules field.
func (o *InlineResponse200187) SetRules(v []InlineResponse200187Rules) {
	o.Rules = v
}

func (o InlineResponse200187) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Exception) {
		toSerialize["exception"] = o.Exception
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200187 struct {
	value *InlineResponse200187
	isSet bool
}

func (v NullableInlineResponse200187) Get() *InlineResponse200187 {
	return v.value
}

func (v *NullableInlineResponse200187) Set(val *InlineResponse200187) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200187) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200187) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200187(val *InlineResponse200187) *NullableInlineResponse200187 {
	return &NullableInlineResponse200187{value: val, isSet: true}
}

func (v NullableInlineResponse200187) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200187) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


