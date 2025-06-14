/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200AuthenticationTwoFactor TwoFactor authentication
type InlineResponse200AuthenticationTwoFactor struct {
	// If twoFactor authentication is enabled for this user
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200AuthenticationTwoFactor instantiates a new InlineResponse200AuthenticationTwoFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200AuthenticationTwoFactor() *InlineResponse200AuthenticationTwoFactor {
	this := InlineResponse200AuthenticationTwoFactor{}
	return &this
}

// NewInlineResponse200AuthenticationTwoFactorWithDefaults instantiates a new InlineResponse200AuthenticationTwoFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200AuthenticationTwoFactorWithDefaults() *InlineResponse200AuthenticationTwoFactor {
	this := InlineResponse200AuthenticationTwoFactor{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200AuthenticationTwoFactor) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200AuthenticationTwoFactor) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200AuthenticationTwoFactor) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200AuthenticationTwoFactor) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200AuthenticationTwoFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200AuthenticationTwoFactor struct {
	value *InlineResponse200AuthenticationTwoFactor
	isSet bool
}

func (v NullableInlineResponse200AuthenticationTwoFactor) Get() *InlineResponse200AuthenticationTwoFactor {
	return v.value
}

func (v *NullableInlineResponse200AuthenticationTwoFactor) Set(val *InlineResponse200AuthenticationTwoFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200AuthenticationTwoFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200AuthenticationTwoFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200AuthenticationTwoFactor(val *InlineResponse200AuthenticationTwoFactor) *NullableInlineResponse200AuthenticationTwoFactor {
	return &NullableInlineResponse200AuthenticationTwoFactor{value: val, isSet: true}
}

func (v NullableInlineResponse200AuthenticationTwoFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200AuthenticationTwoFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


