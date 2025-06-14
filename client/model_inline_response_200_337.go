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

// InlineResponse200337 struct for InlineResponse200337
type InlineResponse200337 struct {
	Counts *InlineResponse200337Counts `json:"counts,omitempty"`
}

// NewInlineResponse200337 instantiates a new InlineResponse200337 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200337() *InlineResponse200337 {
	this := InlineResponse200337{}
	return &this
}

// NewInlineResponse200337WithDefaults instantiates a new InlineResponse200337 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200337WithDefaults() *InlineResponse200337 {
	this := InlineResponse200337{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200337) GetCounts() InlineResponse200337Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200337Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337) GetCountsOk() (*InlineResponse200337Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200337) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200337Counts and assigns it to the Counts field.
func (o *InlineResponse200337) SetCounts(v InlineResponse200337Counts) {
	o.Counts = &v
}

func (o InlineResponse200337) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200337 struct {
	value *InlineResponse200337
	isSet bool
}

func (v NullableInlineResponse200337) Get() *InlineResponse200337 {
	return v.value
}

func (v *NullableInlineResponse200337) Set(val *InlineResponse200337) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200337) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200337) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200337(val *InlineResponse200337) *NullableInlineResponse200337 {
	return &NullableInlineResponse200337{value: val, isSet: true}
}

func (v NullableInlineResponse200337) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200337) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


