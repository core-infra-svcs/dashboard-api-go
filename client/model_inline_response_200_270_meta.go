/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200270Meta Other metadata related to this result set.
type InlineResponse200270Meta struct {
	Counts *InlineResponse200270MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200270Meta instantiates a new InlineResponse200270Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200270Meta() *InlineResponse200270Meta {
	this := InlineResponse200270Meta{}
	return &this
}

// NewInlineResponse200270MetaWithDefaults instantiates a new InlineResponse200270Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200270MetaWithDefaults() *InlineResponse200270Meta {
	this := InlineResponse200270Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200270Meta) GetCounts() InlineResponse200270MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200270MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Meta) GetCountsOk() (*InlineResponse200270MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200270Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200270MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200270Meta) SetCounts(v InlineResponse200270MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200270Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200270Meta struct {
	value *InlineResponse200270Meta
	isSet bool
}

func (v NullableInlineResponse200270Meta) Get() *InlineResponse200270Meta {
	return v.value
}

func (v *NullableInlineResponse200270Meta) Set(val *InlineResponse200270Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200270Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200270Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200270Meta(val *InlineResponse200270Meta) *NullableInlineResponse200270Meta {
	return &NullableInlineResponse200270Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200270Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200270Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


