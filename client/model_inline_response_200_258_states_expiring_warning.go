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

// InlineResponse200258StatesExpiringWarning Data for the warning threshold
type InlineResponse200258StatesExpiringWarning struct {
	// The number of days from now denoting the warning threshold for an expiring license
	ThresholdInDays *int32 `json:"thresholdInDays,omitempty"`
	// The number of licenses that will expire in this window
	ExpiringCount *int32 `json:"expiringCount,omitempty"`
}

// NewInlineResponse200258StatesExpiringWarning instantiates a new InlineResponse200258StatesExpiringWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200258StatesExpiringWarning() *InlineResponse200258StatesExpiringWarning {
	this := InlineResponse200258StatesExpiringWarning{}
	return &this
}

// NewInlineResponse200258StatesExpiringWarningWithDefaults instantiates a new InlineResponse200258StatesExpiringWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200258StatesExpiringWarningWithDefaults() *InlineResponse200258StatesExpiringWarning {
	this := InlineResponse200258StatesExpiringWarning{}
	return &this
}

// GetThresholdInDays returns the ThresholdInDays field value if set, zero value otherwise.
func (o *InlineResponse200258StatesExpiringWarning) GetThresholdInDays() int32 {
	if o == nil || isNil(o.ThresholdInDays) {
		var ret int32
		return ret
	}
	return *o.ThresholdInDays
}

// GetThresholdInDaysOk returns a tuple with the ThresholdInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258StatesExpiringWarning) GetThresholdInDaysOk() (*int32, bool) {
	if o == nil || isNil(o.ThresholdInDays) {
    return nil, false
	}
	return o.ThresholdInDays, true
}

// HasThresholdInDays returns a boolean if a field has been set.
func (o *InlineResponse200258StatesExpiringWarning) HasThresholdInDays() bool {
	if o != nil && !isNil(o.ThresholdInDays) {
		return true
	}

	return false
}

// SetThresholdInDays gets a reference to the given int32 and assigns it to the ThresholdInDays field.
func (o *InlineResponse200258StatesExpiringWarning) SetThresholdInDays(v int32) {
	o.ThresholdInDays = &v
}

// GetExpiringCount returns the ExpiringCount field value if set, zero value otherwise.
func (o *InlineResponse200258StatesExpiringWarning) GetExpiringCount() int32 {
	if o == nil || isNil(o.ExpiringCount) {
		var ret int32
		return ret
	}
	return *o.ExpiringCount
}

// GetExpiringCountOk returns a tuple with the ExpiringCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258StatesExpiringWarning) GetExpiringCountOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiringCount) {
    return nil, false
	}
	return o.ExpiringCount, true
}

// HasExpiringCount returns a boolean if a field has been set.
func (o *InlineResponse200258StatesExpiringWarning) HasExpiringCount() bool {
	if o != nil && !isNil(o.ExpiringCount) {
		return true
	}

	return false
}

// SetExpiringCount gets a reference to the given int32 and assigns it to the ExpiringCount field.
func (o *InlineResponse200258StatesExpiringWarning) SetExpiringCount(v int32) {
	o.ExpiringCount = &v
}

func (o InlineResponse200258StatesExpiringWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ThresholdInDays) {
		toSerialize["thresholdInDays"] = o.ThresholdInDays
	}
	if !isNil(o.ExpiringCount) {
		toSerialize["expiringCount"] = o.ExpiringCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200258StatesExpiringWarning struct {
	value *InlineResponse200258StatesExpiringWarning
	isSet bool
}

func (v NullableInlineResponse200258StatesExpiringWarning) Get() *InlineResponse200258StatesExpiringWarning {
	return v.value
}

func (v *NullableInlineResponse200258StatesExpiringWarning) Set(val *InlineResponse200258StatesExpiringWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200258StatesExpiringWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200258StatesExpiringWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200258StatesExpiringWarning(val *InlineResponse200258StatesExpiringWarning) *NullableInlineResponse200258StatesExpiringWarning {
	return &NullableInlineResponse200258StatesExpiringWarning{value: val, isSet: true}
}

func (v NullableInlineResponse200258StatesExpiringWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200258StatesExpiringWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

