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

// InlineResponse200330Meta Meta details about the result
type InlineResponse200330Meta struct {
	Counts *InlineResponse200330MetaCounts `json:"counts,omitempty"`
}

// NewInlineResponse200330Meta instantiates a new InlineResponse200330Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200330Meta() *InlineResponse200330Meta {
	this := InlineResponse200330Meta{}
	return &this
}

// NewInlineResponse200330MetaWithDefaults instantiates a new InlineResponse200330Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200330MetaWithDefaults() *InlineResponse200330Meta {
	this := InlineResponse200330Meta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200330Meta) GetCounts() InlineResponse200330MetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200330MetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200330Meta) GetCountsOk() (*InlineResponse200330MetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200330Meta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200330MetaCounts and assigns it to the Counts field.
func (o *InlineResponse200330Meta) SetCounts(v InlineResponse200330MetaCounts) {
	o.Counts = &v
}

func (o InlineResponse200330Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200330Meta struct {
	value *InlineResponse200330Meta
	isSet bool
}

func (v NullableInlineResponse200330Meta) Get() *InlineResponse200330Meta {
	return v.value
}

func (v *NullableInlineResponse200330Meta) Set(val *InlineResponse200330Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200330Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200330Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200330Meta(val *InlineResponse200330Meta) *NullableInlineResponse200330Meta {
	return &NullableInlineResponse200330Meta{value: val, isSet: true}
}

func (v NullableInlineResponse200330Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200330Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


