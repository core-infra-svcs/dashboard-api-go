/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200200 struct for InlineResponse200200
type InlineResponse200200 struct {
	Counts *InlineResponse200200Counts `json:"counts,omitempty"`
	Limits *InlineResponse200200Limits `json:"limits,omitempty"`
}

// NewInlineResponse200200 instantiates a new InlineResponse200200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200() *InlineResponse200200 {
	this := InlineResponse200200{}
	return &this
}

// NewInlineResponse200200WithDefaults instantiates a new InlineResponse200200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200WithDefaults() *InlineResponse200200 {
	this := InlineResponse200200{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200200) GetCounts() InlineResponse200200Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200200Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetCountsOk() (*InlineResponse200200Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200200) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200200Counts and assigns it to the Counts field.
func (o *InlineResponse200200) SetCounts(v InlineResponse200200Counts) {
	o.Counts = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *InlineResponse200200) GetLimits() InlineResponse200200Limits {
	if o == nil || isNil(o.Limits) {
		var ret InlineResponse200200Limits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetLimitsOk() (*InlineResponse200200Limits, bool) {
	if o == nil || isNil(o.Limits) {
    return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *InlineResponse200200) HasLimits() bool {
	if o != nil && !isNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given InlineResponse200200Limits and assigns it to the Limits field.
func (o *InlineResponse200200) SetLimits(v InlineResponse200200Limits) {
	o.Limits = &v
}

func (o InlineResponse200200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !isNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200200 struct {
	value *InlineResponse200200
	isSet bool
}

func (v NullableInlineResponse200200) Get() *InlineResponse200200 {
	return v.value
}

func (v *NullableInlineResponse200200) Set(val *InlineResponse200200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200(val *InlineResponse200200) *NullableInlineResponse200200 {
	return &NullableInlineResponse200200{value: val, isSet: true}
}

func (v NullableInlineResponse200200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


