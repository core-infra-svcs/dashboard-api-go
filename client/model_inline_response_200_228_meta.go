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

// InlineResponse200228Meta Meta details about the result
type InlineResponse200228Meta struct {
	Counts *InlineResponse200228MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200228Meta instantiates a new InlineResponse200228Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200228Meta() *InlineResponse200228Meta {
	this := InlineResponse200228Meta{}
	return &this
}

// NewInlineResponse200228MetaWithDefaults instantiates a new InlineResponse200228Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200228MetaWithDefaults() *InlineResponse200228Meta {
	this := InlineResponse200228Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200228Meta) GetCounts() InlineResponse200228MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200228MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200228Meta) GetCountsOk() (*InlineResponse200228MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200228Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200228MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200228Meta) SetCounts(v InlineResponse200228MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200228Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200228Meta struct {
	value *InlineResponse200228Meta
	isSet bool
}

func (v NullableInlineResponse200228Meta) Get() *InlineResponse200228Meta {
	return v.value
}

func (v *NullableInlineResponse200228Meta) Set(val *InlineResponse200228Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200228Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200228Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200228Meta(val *InlineResponse200228Meta) *NullableInlineResponse200228Meta {
	return &NullableInlineResponse200228Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200228Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200228Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


