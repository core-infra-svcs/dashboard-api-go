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

// InlineResponse200307Meta Other metadata related to this result set.
type InlineResponse200307Meta struct {
	Counts *InlineResponse200307MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200307Meta instantiates a new InlineResponse200307Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200307Meta() *InlineResponse200307Meta {
	this := InlineResponse200307Meta{}
	return &this
}

// NewInlineResponse200307MetaWithDefaults instantiates a new InlineResponse200307Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200307MetaWithDefaults() *InlineResponse200307Meta {
	this := InlineResponse200307Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200307Meta) GetCounts() InlineResponse200307MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200307MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307Meta) GetCountsOk() (*InlineResponse200307MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200307Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200307MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200307Meta) SetCounts(v InlineResponse200307MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200307Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200307Meta struct {
	value *InlineResponse200307Meta
	isSet bool
}

func (v NullableInlineResponse200307Meta) Get() *InlineResponse200307Meta {
	return v.value
}

func (v *NullableInlineResponse200307Meta) Set(val *InlineResponse200307Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200307Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200307Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200307Meta(val *InlineResponse200307Meta) *NullableInlineResponse200307Meta {
	return &NullableInlineResponse200307Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200307Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200307Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


