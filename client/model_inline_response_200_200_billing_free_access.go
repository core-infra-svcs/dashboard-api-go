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

// InlineResponse200200BillingFreeAccess Details associated with a free access plan with limits
type InlineResponse200200BillingFreeAccess struct {
	// Whether or not free access is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// How long a device can use a network for free.
	DurationInMinutes *int32 `json:"durationInMinutes,omitempty"`
}

// NewInlineResponse200200BillingFreeAccess instantiates a new InlineResponse200200BillingFreeAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200BillingFreeAccess() *InlineResponse200200BillingFreeAccess {
	this := InlineResponse200200BillingFreeAccess{}
	return &this
}

// NewInlineResponse200200BillingFreeAccessWithDefaults instantiates a new InlineResponse200200BillingFreeAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200BillingFreeAccessWithDefaults() *InlineResponse200200BillingFreeAccess {
	this := InlineResponse200200BillingFreeAccess{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200200BillingFreeAccess) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200BillingFreeAccess) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200200BillingFreeAccess) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200200BillingFreeAccess) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDurationInMinutes returns the DurationInMinutes field value if set, zero value otherwise.
func (o *InlineResponse200200BillingFreeAccess) GetDurationInMinutes() int32 {
	if o == nil || isNil(o.DurationInMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200BillingFreeAccess) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.DurationInMinutes) {
    return nil, false
	}
	return o.DurationInMinutes, true
}

// HasDurationInMinutes returns a boolean if a field has been set.
func (o *InlineResponse200200BillingFreeAccess) HasDurationInMinutes() bool {
	if o != nil && !isNil(o.DurationInMinutes) {
		return true
	}

	return false
}

// SetDurationInMinutes gets a reference to the given int32 and assigns it to the DurationInMinutes field.
func (o *InlineResponse200200BillingFreeAccess) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = &v
}

func (o InlineResponse200200BillingFreeAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DurationInMinutes) {
		toSerialize["durationInMinutes"] = o.DurationInMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200200BillingFreeAccess struct {
	value *InlineResponse200200BillingFreeAccess
	isSet bool
}

func (v NullableInlineResponse200200BillingFreeAccess) Get() *InlineResponse200200BillingFreeAccess {
	return v.value
}

func (v *NullableInlineResponse200200BillingFreeAccess) Set(val *InlineResponse200200BillingFreeAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200BillingFreeAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200BillingFreeAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200BillingFreeAccess(val *InlineResponse200200BillingFreeAccess) *NullableInlineResponse200200BillingFreeAccess {
	return &NullableInlineResponse200200BillingFreeAccess{value: val, isSet: true}
}

func (v NullableInlineResponse200200BillingFreeAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200BillingFreeAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

