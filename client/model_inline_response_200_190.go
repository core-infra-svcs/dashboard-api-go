/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200190 struct for InlineResponse200190
type InlineResponse200190 struct {
	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
	Rules []InlineResponse200190Rules `json:"rules,omitempty"`
}

// NewInlineResponse200190 instantiates a new InlineResponse200190 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200190() *InlineResponse200190 {
	this := InlineResponse200190{}
	return &this
}

// NewInlineResponse200190WithDefaults instantiates a new InlineResponse200190 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200190WithDefaults() *InlineResponse200190 {
	this := InlineResponse200190{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineResponse200190) GetRules() []InlineResponse200190Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200190Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200190) GetRulesOk() ([]InlineResponse200190Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineResponse200190) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200190Rules and assigns it to the Rules field.
func (o *InlineResponse200190) SetRules(v []InlineResponse200190Rules) {
	o.Rules = v
}

func (o InlineResponse200190) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200190 struct {
	value *InlineResponse200190
	isSet bool
}

func (v NullableInlineResponse200190) Get() *InlineResponse200190 {
	return v.value
}

func (v *NullableInlineResponse200190) Set(val *InlineResponse200190) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200190) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200190) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200190(val *InlineResponse200190) *NullableInlineResponse200190 {
	return &NullableInlineResponse200190{value: val, isSet: true}
}

func (v NullableInlineResponse200190) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200190) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


