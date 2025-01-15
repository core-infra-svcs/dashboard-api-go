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

// InlineResponse200222 struct for InlineResponse200222
type InlineResponse200222 struct {
	Counts InlineResponse200222Counts `json:"counts"`
}

// NewInlineResponse200222 instantiates a new InlineResponse200222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200222(counts InlineResponse200222Counts) *InlineResponse200222 {
	this := InlineResponse200222{}
	this.Counts = counts
	return &this
}

// NewInlineResponse200222WithDefaults instantiates a new InlineResponse200222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200222WithDefaults() *InlineResponse200222 {
	this := InlineResponse200222{}
	return &this
}

// GetCounts returns the Counts field value
func (o *InlineResponse200222) GetCounts() InlineResponse200222Counts {
	if o == nil {
		var ret InlineResponse200222Counts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200222) GetCountsOk() (*InlineResponse200222Counts, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *InlineResponse200222) SetCounts(v InlineResponse200222Counts) {
	o.Counts = v
}

func (o InlineResponse200222) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200222 struct {
	value *InlineResponse200222
	isSet bool
}

func (v NullableInlineResponse200222) Get() *InlineResponse200222 {
	return v.value
}

func (v *NullableInlineResponse200222) Set(val *InlineResponse200222) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200222) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200222) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200222(val *InlineResponse200222) *NullableInlineResponse200222 {
	return &NullableInlineResponse200222{value: val, isSet: true}
}

func (v NullableInlineResponse200222) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200222) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


