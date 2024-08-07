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

// InlineResponse200216Meta Metadata about the response
type InlineResponse200216Meta struct {
	Counts InlineResponse200216MetaCounts `json:"counts"`
}

// NewInlineResponse200216Meta instantiates a new InlineResponse200216Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200216Meta(counts InlineResponse200216MetaCounts) *InlineResponse200216Meta {
	this := InlineResponse200216Meta{}
	this.Counts = counts
	return &this
}

// NewInlineResponse200216MetaWithDefaults instantiates a new InlineResponse200216Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200216MetaWithDefaults() *InlineResponse200216Meta {
	this := InlineResponse200216Meta{}
	return &this
}

// GetCounts returns the Counts field value
func (o *InlineResponse200216Meta) GetCounts() InlineResponse200216MetaCounts {
	if o == nil {
		var ret InlineResponse200216MetaCounts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Meta) GetCountsOk() (*InlineResponse200216MetaCounts, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *InlineResponse200216Meta) SetCounts(v InlineResponse200216MetaCounts) {
	o.Counts = v
}

func (o InlineResponse200216Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200216Meta struct {
	value *InlineResponse200216Meta
	isSet bool
}

func (v NullableInlineResponse200216Meta) Get() *InlineResponse200216Meta {
	return v.value
}

func (v *NullableInlineResponse200216Meta) Set(val *InlineResponse200216Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200216Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200216Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200216Meta(val *InlineResponse200216Meta) *NullableInlineResponse200216Meta {
	return &NullableInlineResponse200216Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200216Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200216Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


