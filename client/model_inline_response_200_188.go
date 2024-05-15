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

// InlineResponse200188 struct for InlineResponse200188
type InlineResponse200188 struct {
	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
	Rules []InlineResponse200188Rules `json:"rules,omitempty"`
}

// NewInlineResponse200188 instantiates a new InlineResponse200188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200188() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// NewInlineResponse200188WithDefaults instantiates a new InlineResponse200188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200188WithDefaults() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse200188) GetRules() []InlineResponse200188Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200188Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetRulesOk() ([]InlineResponse200188Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse200188) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200188Rules and assigns it to the Rules field.
func (o *InlineResponse200188) SetRules(v []InlineResponse200188Rules) {
	o.Rules = v
}

func (o InlineResponse200188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200188 struct {
	value *InlineResponse200188
	isSet bool
}

func (v NullableInlineResponse200188) Get() *InlineResponse200188 {
	return v.value
}

func (v *NullableInlineResponse200188) Set(val *InlineResponse200188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200188(val *InlineResponse200188) *NullableInlineResponse200188 {
	return &NullableInlineResponse200188{value: val, isSet: true}
}

func (v NullableInlineResponse200188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


