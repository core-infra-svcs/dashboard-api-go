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

// InlineResponse200316Meta Meta details about the result
type InlineResponse200316Meta struct {
	Counts *InlineResponse200316MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200316Meta instantiates a new InlineResponse200316Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200316Meta() *InlineResponse200316Meta {
	this := InlineResponse200316Meta{}
	return &this
}

// NewInlineResponse200316MetaWithDefaults instantiates a new InlineResponse200316Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200316MetaWithDefaults() *InlineResponse200316Meta {
	this := InlineResponse200316Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200316Meta) GetCounts() InlineResponse200316MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200316MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200316Meta) GetCountsOk() (*InlineResponse200316MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200316Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200316MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200316Meta) SetCounts(v InlineResponse200316MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200316Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200316Meta struct {
	value *InlineResponse200316Meta
	isSet bool
}

func (v NullableInlineResponse200316Meta) Get() *InlineResponse200316Meta {
	return v.value
}

func (v *NullableInlineResponse200316Meta) Set(val *InlineResponse200316Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200316Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200316Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200316Meta(val *InlineResponse200316Meta) *NullableInlineResponse200316Meta {
	return &NullableInlineResponse200316Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200316Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200316Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


