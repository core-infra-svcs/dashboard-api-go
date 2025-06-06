/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200276Meta Other metadata related to this result set.
type InlineResponse200276Meta struct {
	Counts *InlineResponse200276MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200276Meta instantiates a new InlineResponse200276Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200276Meta() *InlineResponse200276Meta {
	this := InlineResponse200276Meta{}
	return &this
}

// NewInlineResponse200276MetaWithDefaults instantiates a new InlineResponse200276Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200276MetaWithDefaults() *InlineResponse200276Meta {
	this := InlineResponse200276Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200276Meta) GetCounts() InlineResponse200276MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200276MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276Meta) GetCountsOk() (*InlineResponse200276MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200276Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200276MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200276Meta) SetCounts(v InlineResponse200276MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200276Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200276Meta struct {
	value *InlineResponse200276Meta
	isSet bool
}

func (v NullableInlineResponse200276Meta) Get() *InlineResponse200276Meta {
	return v.value
}

func (v *NullableInlineResponse200276Meta) Set(val *InlineResponse200276Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200276Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200276Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200276Meta(val *InlineResponse200276Meta) *NullableInlineResponse200276Meta {
	return &NullableInlineResponse200276Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200276Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200276Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


