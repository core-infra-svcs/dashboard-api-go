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

// InlineResponse200308 struct for InlineResponse200308
type InlineResponse200308 struct {
	Counts *InlineResponse200308Counts `json:"counts,omitempty"`
}

// NewInlineResponse200308 instantiates a new InlineResponse200308 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308() *InlineResponse200308 {
	this := InlineResponse200308{}
	return &this
}

// NewInlineResponse200308WithDefaults instantiates a new InlineResponse200308 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308WithDefaults() *InlineResponse200308 {
	this := InlineResponse200308{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200308) GetCounts() InlineResponse200308Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200308Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308) GetCountsOk() (*InlineResponse200308Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200308) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200308Counts and assigns it to the Counts field.
func (o *InlineResponse200308) SetCounts(v InlineResponse200308Counts) {
	o.Counts = &v
}

func (o InlineResponse200308) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308 struct {
	value *InlineResponse200308
	isSet bool
}

func (v NullableInlineResponse200308) Get() *InlineResponse200308 {
	return v.value
}

func (v *NullableInlineResponse200308) Set(val *InlineResponse200308) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308(val *InlineResponse200308) *NullableInlineResponse200308 {
	return &NullableInlineResponse200308{value: val, isSet: true}
}

func (v NullableInlineResponse200308) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


