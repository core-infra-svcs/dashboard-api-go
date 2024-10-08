/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200216 struct for InlineResponse200216
type InlineResponse200216 struct {
	Counts InlineResponse200216Counts `json:"counts"`
}

// NewInlineResponse200216 instantiates a new InlineResponse200216 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200216(counts InlineResponse200216Counts) *InlineResponse200216 {
	this := InlineResponse200216{}
	this.Counts = counts
	return &this
}

// NewInlineResponse200216WithDefaults instantiates a new InlineResponse200216 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200216WithDefaults() *InlineResponse200216 {
	this := InlineResponse200216{}
	return &this
}

// GetCounts returns the Counts field value
func (o *InlineResponse200216) GetCounts() InlineResponse200216Counts {
	if o == nil {
		var ret InlineResponse200216Counts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216) GetCountsOk() (*InlineResponse200216Counts, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *InlineResponse200216) SetCounts(v InlineResponse200216Counts) {
	o.Counts = v
}

func (o InlineResponse200216) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200216 struct {
	value *InlineResponse200216
	isSet bool
}

func (v NullableInlineResponse200216) Get() *InlineResponse200216 {
	return v.value
}

func (v *NullableInlineResponse200216) Set(val *InlineResponse200216) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200216) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200216) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200216(val *InlineResponse200216) *NullableInlineResponse200216 {
	return &NullableInlineResponse200216{value: val, isSet: true}
}

func (v NullableInlineResponse200216) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200216) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


