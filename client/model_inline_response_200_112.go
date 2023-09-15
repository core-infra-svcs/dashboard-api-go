/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200112 struct for InlineResponse200112
type InlineResponse200112 struct {
	Counts *InlineResponse200112Counts `json:"counts,omitempty"`
	Limits *InlineResponse200112Limits `json:"limits,omitempty"`
}

// NewInlineResponse200112 instantiates a new InlineResponse200112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200112() *InlineResponse200112 {
	this := InlineResponse200112{}
	return &this
}

// NewInlineResponse200112WithDefaults instantiates a new InlineResponse200112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200112WithDefaults() *InlineResponse200112 {
	this := InlineResponse200112{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200112) GetCounts() InlineResponse200112Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200112Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetCountsOk() (*InlineResponse200112Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200112) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200112Counts and assigns it to the Counts field.
func (o *InlineResponse200112) SetCounts(v InlineResponse200112Counts) {
	o.Counts = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *InlineResponse200112) GetLimits() InlineResponse200112Limits {
	if o == nil || isNil(o.Limits) {
		var ret InlineResponse200112Limits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetLimitsOk() (*InlineResponse200112Limits, bool) {
	if o == nil || isNil(o.Limits) {
    return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *InlineResponse200112) HasLimits() bool {
	if o != nil && !isNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given InlineResponse200112Limits and assigns it to the Limits field.
func (o *InlineResponse200112) SetLimits(v InlineResponse200112Limits) {
	o.Limits = &v
}

func (o InlineResponse200112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !isNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200112 struct {
	value *InlineResponse200112
	isSet bool
}

func (v NullableInlineResponse200112) Get() *InlineResponse200112 {
	return v.value
}

func (v *NullableInlineResponse200112) Set(val *InlineResponse200112) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200112) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200112(val *InlineResponse200112) *NullableInlineResponse200112 {
	return &NullableInlineResponse200112{value: val, isSet: true}
}

func (v NullableInlineResponse200112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


