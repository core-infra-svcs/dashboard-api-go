/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200310Meta Meta details about the result
type InlineResponse200310Meta struct {
	Counts *InlineResponse200310MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200310Meta instantiates a new InlineResponse200310Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200310Meta() *InlineResponse200310Meta {
	this := InlineResponse200310Meta{}
	return &this
}

// NewInlineResponse200310MetaWithDefaults instantiates a new InlineResponse200310Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200310MetaWithDefaults() *InlineResponse200310Meta {
	this := InlineResponse200310Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200310Meta) GetCounts() InlineResponse200310MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200310MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200310Meta) GetCountsOk() (*InlineResponse200310MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200310Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200310MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200310Meta) SetCounts(v InlineResponse200310MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200310Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200310Meta struct {
	value *InlineResponse200310Meta
	isSet bool
}

func (v NullableInlineResponse200310Meta) Get() *InlineResponse200310Meta {
	return v.value
}

func (v *NullableInlineResponse200310Meta) Set(val *InlineResponse200310Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200310Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200310Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200310Meta(val *InlineResponse200310Meta) *NullableInlineResponse200310Meta {
	return &NullableInlineResponse200310Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200310Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200310Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


