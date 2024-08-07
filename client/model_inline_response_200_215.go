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

// InlineResponse200215 struct for InlineResponse200215
type InlineResponse200215 struct {
	Counts InlineResponse200215Counts `json:"counts"`
}

// NewInlineResponse200215 instantiates a new InlineResponse200215 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200215(counts InlineResponse200215Counts) *InlineResponse200215 {
	this := InlineResponse200215{}
	this.Counts = counts
	return &this
}

// NewInlineResponse200215WithDefaults instantiates a new InlineResponse200215 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200215WithDefaults() *InlineResponse200215 {
	this := InlineResponse200215{}
	return &this
}

// GetCounts returns the Counts field value
func (o *InlineResponse200215) GetCounts() InlineResponse200215Counts {
	if o == nil {
		var ret InlineResponse200215Counts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200215) GetCountsOk() (*InlineResponse200215Counts, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *InlineResponse200215) SetCounts(v InlineResponse200215Counts) {
	o.Counts = v
}

func (o InlineResponse200215) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200215 struct {
	value *InlineResponse200215
	isSet bool
}

func (v NullableInlineResponse200215) Get() *InlineResponse200215 {
	return v.value
}

func (v *NullableInlineResponse200215) Set(val *InlineResponse200215) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200215) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200215) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200215(val *InlineResponse200215) *NullableInlineResponse200215 {
	return &NullableInlineResponse200215{value: val, isSet: true}
}

func (v NullableInlineResponse200215) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200215) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


