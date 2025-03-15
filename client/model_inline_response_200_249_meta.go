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

// InlineResponse200249Meta Meta details about the result
type InlineResponse200249Meta struct {
	Counts *InlineResponse200249MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200249Meta instantiates a new InlineResponse200249Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200249Meta() *InlineResponse200249Meta {
	this := InlineResponse200249Meta{}
	return &this
}

// NewInlineResponse200249MetaWithDefaults instantiates a new InlineResponse200249Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200249MetaWithDefaults() *InlineResponse200249Meta {
	this := InlineResponse200249Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200249Meta) GetCounts() InlineResponse200249MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200249MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200249Meta) GetCountsOk() (*InlineResponse200249MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200249Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200249MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200249Meta) SetCounts(v InlineResponse200249MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200249Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200249Meta struct {
	value *InlineResponse200249Meta
	isSet bool
}

func (v NullableInlineResponse200249Meta) Get() *InlineResponse200249Meta {
	return v.value
}

func (v *NullableInlineResponse200249Meta) Set(val *InlineResponse200249Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200249Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200249Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200249Meta(val *InlineResponse200249Meta) *NullableInlineResponse200249Meta {
	return &NullableInlineResponse200249Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200249Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200249Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


