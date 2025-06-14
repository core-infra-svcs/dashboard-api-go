/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200275Meta Meta data details about result
type InlineResponse200275Meta struct {
	Counts *InlineResponse200275MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200275Meta instantiates a new InlineResponse200275Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200275Meta() *InlineResponse200275Meta {
	this := InlineResponse200275Meta{}
	return &this
}

// NewInlineResponse200275MetaWithDefaults instantiates a new InlineResponse200275Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200275MetaWithDefaults() *InlineResponse200275Meta {
	this := InlineResponse200275Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200275Meta) GetCounts() InlineResponse200275MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200275MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200275Meta) GetCountsOk() (*InlineResponse200275MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200275Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200275MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200275Meta) SetCounts(v InlineResponse200275MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200275Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200275Meta struct {
	value *InlineResponse200275Meta
	isSet bool
}

func (v NullableInlineResponse200275Meta) Get() *InlineResponse200275Meta {
	return v.value
}

func (v *NullableInlineResponse200275Meta) Set(val *InlineResponse200275Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200275Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200275Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200275Meta(val *InlineResponse200275Meta) *NullableInlineResponse200275Meta {
	return &NullableInlineResponse200275Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200275Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200275Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


