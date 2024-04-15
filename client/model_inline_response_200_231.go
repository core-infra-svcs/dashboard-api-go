/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200231 struct for InlineResponse200231
type InlineResponse200231 struct {
	Counts *InlineResponse200231Counts `json:"counts,omitempty"`
}

// NewInlineResponse200231 instantiates a new InlineResponse200231 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200231() *InlineResponse200231 {
	this := InlineResponse200231{}
	return &this
}

// NewInlineResponse200231WithDefaults instantiates a new InlineResponse200231 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200231WithDefaults() *InlineResponse200231 {
	this := InlineResponse200231{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200231) GetCounts() InlineResponse200231Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200231Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200231) GetCountsOk() (*InlineResponse200231Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200231) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200231Counts and assigns it to the Counts field.
func (o *InlineResponse200231) SetCounts(v InlineResponse200231Counts) {
	o.Counts = &v
}

func (o InlineResponse200231) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200231 struct {
	value *InlineResponse200231
	isSet bool
}

func (v NullableInlineResponse200231) Get() *InlineResponse200231 {
	return v.value
}

func (v *NullableInlineResponse200231) Set(val *InlineResponse200231) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200231) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200231) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200231(val *InlineResponse200231) *NullableInlineResponse200231 {
	return &NullableInlineResponse200231{value: val, isSet: true}
}

func (v NullableInlineResponse200231) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200231) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


