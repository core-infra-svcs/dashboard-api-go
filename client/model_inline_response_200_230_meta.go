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

// InlineResponse200230Meta Meta details about the result
type InlineResponse200230Meta struct {
	Counts *InlineResponse200230MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200230Meta instantiates a new InlineResponse200230Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200230Meta() *InlineResponse200230Meta {
	this := InlineResponse200230Meta{}
	return &this
}

// NewInlineResponse200230MetaWithDefaults instantiates a new InlineResponse200230Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200230MetaWithDefaults() *InlineResponse200230Meta {
	this := InlineResponse200230Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200230Meta) GetCounts() InlineResponse200230MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200230MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200230Meta) GetCountsOk() (*InlineResponse200230MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200230Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200230MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200230Meta) SetCounts(v InlineResponse200230MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200230Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200230Meta struct {
	value *InlineResponse200230Meta
	isSet bool
}

func (v NullableInlineResponse200230Meta) Get() *InlineResponse200230Meta {
	return v.value
}

func (v *NullableInlineResponse200230Meta) Set(val *InlineResponse200230Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200230Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200230Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200230Meta(val *InlineResponse200230Meta) *NullableInlineResponse200230Meta {
	return &NullableInlineResponse200230Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200230Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200230Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


