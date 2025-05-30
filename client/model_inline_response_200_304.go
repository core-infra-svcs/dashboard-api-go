/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200304 struct for InlineResponse200304
type InlineResponse200304 struct {
	// Toggle depicting if SAML SSO settings are enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200304 instantiates a new InlineResponse200304 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200304() *InlineResponse200304 {
	this := InlineResponse200304{}
	return &this
}

// NewInlineResponse200304WithDefaults instantiates a new InlineResponse200304 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200304WithDefaults() *InlineResponse200304 {
	this := InlineResponse200304{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200304) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200304) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200304) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200304) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200304) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200304 struct {
	value *InlineResponse200304
	isSet bool
}

func (v NullableInlineResponse200304) Get() *InlineResponse200304 {
	return v.value
}

func (v *NullableInlineResponse200304) Set(val *InlineResponse200304) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200304) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200304) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200304(val *InlineResponse200304) *NullableInlineResponse200304 {
	return &NullableInlineResponse200304{value: val, isSet: true}
}

func (v NullableInlineResponse200304) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200304) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


