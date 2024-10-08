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

// InlineObject188 struct for InlineObject188
type InlineObject188 struct {
	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
	Rules []InlineResponse200190Rules `json:"rules,omitempty"`
	// Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
	AllowLanAccess *bool `json:"allowLanAccess,omitempty"`
}

// NewInlineObject188 instantiates a new InlineObject188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject188() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// NewInlineObject188WithDefaults instantiates a new InlineObject188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject188WithDefaults() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject188) GetRules() []InlineResponse200190Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200190Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetRulesOk() ([]InlineResponse200190Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject188) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200190Rules and assigns it to the Rules field.
func (o *InlineObject188) SetRules(v []InlineResponse200190Rules) {
	o.Rules = v
}

// GetAllowLanAccess returns the AllowLanAccess field value if set, zero value otherwise.
func (o *InlineObject188) GetAllowLanAccess() bool {
	if o == nil || isNil(o.AllowLanAccess) {
		var ret bool
		return ret
	}
	return *o.AllowLanAccess
}

// GetAllowLanAccessOk returns a tuple with the AllowLanAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetAllowLanAccessOk() (*bool, bool) {
	if o == nil || isNil(o.AllowLanAccess) {
    return nil, false
	}
	return o.AllowLanAccess, true
}

// HasAllowLanAccess returns a boolean if a field has been set.
func (o *InlineObject188) HasAllowLanAccess() bool {
	if o != nil && !isNil(o.AllowLanAccess) {
		return true
	}

	return false
}

// SetAllowLanAccess gets a reference to the given bool and assigns it to the AllowLanAccess field.
func (o *InlineObject188) SetAllowLanAccess(v bool) {
	o.AllowLanAccess = &v
}

func (o InlineObject188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.AllowLanAccess) {
		toSerialize["allowLanAccess"] = o.AllowLanAccess
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject188 struct {
	value *InlineObject188
	isSet bool
}

func (v NullableInlineObject188) Get() *InlineObject188 {
	return v.value
}

func (v *NullableInlineObject188) Set(val *InlineObject188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject188(val *InlineObject188) *NullableInlineObject188 {
	return &NullableInlineObject188{value: val, isSet: true}
}

func (v NullableInlineObject188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


