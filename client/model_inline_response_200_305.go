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

// InlineResponse200305 struct for InlineResponse200305
type InlineResponse200305 struct {
	Counts *InlineResponse200305Counts `json:"counts,omitempty"`
}

// NewInlineResponse200305 instantiates a new InlineResponse200305 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200305() *InlineResponse200305 {
	this := InlineResponse200305{}
	return &this
}

// NewInlineResponse200305WithDefaults instantiates a new InlineResponse200305 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200305WithDefaults() *InlineResponse200305 {
	this := InlineResponse200305{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200305) GetCounts() InlineResponse200305Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200305Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305) GetCountsOk() (*InlineResponse200305Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200305) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200305Counts and assigns it to the Counts field.
func (o *InlineResponse200305) SetCounts(v InlineResponse200305Counts) {
	o.Counts = &v
}

func (o InlineResponse200305) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200305 struct {
	value *InlineResponse200305
	isSet bool
}

func (v NullableInlineResponse200305) Get() *InlineResponse200305 {
	return v.value
}

func (v *NullableInlineResponse200305) Set(val *InlineResponse200305) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200305) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200305) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200305(val *InlineResponse200305) *NullableInlineResponse200305 {
	return &NullableInlineResponse200305{value: val, isSet: true}
}

func (v NullableInlineResponse200305) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200305) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


