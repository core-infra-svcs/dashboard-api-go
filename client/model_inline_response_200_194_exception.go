/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200194Exception Bonjour forwarding exception
type InlineResponse200194Exception struct {
	// If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200194Exception instantiates a new InlineResponse200194Exception object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200194Exception() *InlineResponse200194Exception {
	this := InlineResponse200194Exception{}
	return &this
}

// NewInlineResponse200194ExceptionWithDefaults instantiates a new InlineResponse200194Exception object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200194ExceptionWithDefaults() *InlineResponse200194Exception {
	this := InlineResponse200194Exception{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200194Exception) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200194Exception) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200194Exception) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200194Exception) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200194Exception) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200194Exception struct {
	value *InlineResponse200194Exception
	isSet bool
}

func (v NullableInlineResponse200194Exception) Get() *InlineResponse200194Exception {
	return v.value
}

func (v *NullableInlineResponse200194Exception) Set(val *InlineResponse200194Exception) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200194Exception) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200194Exception) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200194Exception(val *InlineResponse200194Exception) *NullableInlineResponse200194Exception {
	return &NullableInlineResponse200194Exception{value: val, isSet: true}
}

func (v NullableInlineResponse200194Exception) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200194Exception) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


