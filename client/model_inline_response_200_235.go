/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200235 struct for InlineResponse200235
type InlineResponse200235 struct {
	// Counts of devices per model
	Counts []InlineResponse200235Counts `json:"counts,omitempty"`
}

// NewInlineResponse200235 instantiates a new InlineResponse200235 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200235() *InlineResponse200235 {
	this := InlineResponse200235{}
	return &this
}

// NewInlineResponse200235WithDefaults instantiates a new InlineResponse200235 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200235WithDefaults() *InlineResponse200235 {
	this := InlineResponse200235{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200235) GetCounts() []InlineResponse200235Counts {
	if o == nil || isNil(o.Counts) {
		var ret []InlineResponse200235Counts
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200235) GetCountsOk() ([]InlineResponse200235Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200235) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []InlineResponse200235Counts and assigns it to the Counts field.
func (o *InlineResponse200235) SetCounts(v []InlineResponse200235Counts) {
	o.Counts = v
}

func (o InlineResponse200235) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200235 struct {
	value *InlineResponse200235
	isSet bool
}

func (v NullableInlineResponse200235) Get() *InlineResponse200235 {
	return v.value
}

func (v *NullableInlineResponse200235) Set(val *InlineResponse200235) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200235) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200235) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200235(val *InlineResponse200235) *NullableInlineResponse200235 {
	return &NullableInlineResponse200235{value: val, isSet: true}
}

func (v NullableInlineResponse200235) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200235) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


