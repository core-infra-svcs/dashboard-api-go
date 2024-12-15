/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject193 struct for InlineObject193
type InlineObject193 struct {
	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).
	Rules []InlineResponse200196Rules `json:"rules,omitempty"`
	// Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
	AllowLanAccess *bool `json:"allowLanAccess,omitempty"`
}

// NewInlineObject193 instantiates a new InlineObject193 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject193() *InlineObject193 {
	this := InlineObject193{}
	return &this
}

// NewInlineObject193WithDefaults instantiates a new InlineObject193 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject193WithDefaults() *InlineObject193 {
	this := InlineObject193{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject193) GetRules() []InlineResponse200196Rules {
	if o == nil || isNil(o.Rules) {
		var ret []InlineResponse200196Rules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject193) GetRulesOk() ([]InlineResponse200196Rules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject193) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []InlineResponse200196Rules and assigns it to the Rules field.
func (o *InlineObject193) SetRules(v []InlineResponse200196Rules) {
	o.Rules = v
}

// GetAllowLanAccess returns the AllowLanAccess field value if set, zero value otherwise.
func (o *InlineObject193) GetAllowLanAccess() bool {
	if o == nil || isNil(o.AllowLanAccess) {
		var ret bool
		return ret
	}
	return *o.AllowLanAccess
}

// GetAllowLanAccessOk returns a tuple with the AllowLanAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject193) GetAllowLanAccessOk() (*bool, bool) {
	if o == nil || isNil(o.AllowLanAccess) {
    return nil, false
	}
	return o.AllowLanAccess, true
}

// HasAllowLanAccess returns a boolean if a field has been set.
func (o *InlineObject193) HasAllowLanAccess() bool {
	if o != nil && !isNil(o.AllowLanAccess) {
		return true
	}

	return false
}

// SetAllowLanAccess gets a reference to the given bool and assigns it to the AllowLanAccess field.
func (o *InlineObject193) SetAllowLanAccess(v bool) {
	o.AllowLanAccess = &v
}

func (o InlineObject193) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.AllowLanAccess) {
		toSerialize["allowLanAccess"] = o.AllowLanAccess
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject193 struct {
	value *InlineObject193
	isSet bool
}

func (v NullableInlineObject193) Get() *InlineObject193 {
	return v.value
}

func (v *NullableInlineObject193) Set(val *InlineObject193) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject193) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject193) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject193(val *InlineObject193) *NullableInlineObject193 {
	return &NullableInlineObject193{value: val, isSet: true}
}

func (v NullableInlineObject193) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject193) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


