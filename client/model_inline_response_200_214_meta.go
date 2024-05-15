/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200214Meta Metadata about the response
type InlineResponse200214Meta struct {
	Counts InlineResponse200214MetaCounts `json:"counts"`
}

// NewInlineResponse200214Meta instantiates a new InlineResponse200214Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200214Meta(counts InlineResponse200214MetaCounts) *InlineResponse200214Meta {
	this := InlineResponse200214Meta{}
	this.Counts = counts
	return &this
}

// NewInlineResponse200214MetaWithDefaults instantiates a new InlineResponse200214Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200214MetaWithDefaults() *InlineResponse200214Meta {
	this := InlineResponse200214Meta{}
	return &this
}

// GetCounts returns the Counts field value
func (o *InlineResponse200214Meta) GetCounts() InlineResponse200214MetaCounts {
	if o == nil {
		var ret InlineResponse200214MetaCounts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200214Meta) GetCountsOk() (*InlineResponse200214MetaCounts, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *InlineResponse200214Meta) SetCounts(v InlineResponse200214MetaCounts) {
	o.Counts = v
}

func (o InlineResponse200214Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200214Meta struct {
	value *InlineResponse200214Meta
	isSet bool
}

func (v NullableInlineResponse200214Meta) Get() *InlineResponse200214Meta {
	return v.value
}

func (v *NullableInlineResponse200214Meta) Set(val *InlineResponse200214Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200214Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200214Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200214Meta(val *InlineResponse200214Meta) *NullableInlineResponse200214Meta {
	return &NullableInlineResponse200214Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200214Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200214Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


