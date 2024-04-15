/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200276 struct for InlineResponse200276
type InlineResponse200276 struct {
	Counts *InlineResponse200276Counts `json:"counts,omitempty"`
}

// NewInlineResponse200276 instantiates a new InlineResponse200276 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200276() *InlineResponse200276 {
	this := InlineResponse200276{}
	return &this
}

// NewInlineResponse200276WithDefaults instantiates a new InlineResponse200276 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200276WithDefaults() *InlineResponse200276 {
	this := InlineResponse200276{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200276) GetCounts() InlineResponse200276Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200276Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276) GetCountsOk() (*InlineResponse200276Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200276) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200276Counts and assigns it to the Counts field.
func (o *InlineResponse200276) SetCounts(v InlineResponse200276Counts) {
	o.Counts = &v
}

func (o InlineResponse200276) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200276 struct {
	value *InlineResponse200276
	isSet bool
}

func (v NullableInlineResponse200276) Get() *InlineResponse200276 {
	return v.value
}

func (v *NullableInlineResponse200276) Set(val *InlineResponse200276) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200276) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200276) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200276(val *InlineResponse200276) *NullableInlineResponse200276 {
	return &NullableInlineResponse200276{value: val, isSet: true}
}

func (v NullableInlineResponse200276) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200276) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


