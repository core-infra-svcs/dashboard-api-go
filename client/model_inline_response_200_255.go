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

// InlineResponse200255 struct for InlineResponse200255
type InlineResponse200255 struct {
	Usage *InlineResponse200255Usage `json:"usage,omitempty"`
	Counts *InlineResponse200255Counts `json:"counts,omitempty"`
}

// NewInlineResponse200255 instantiates a new InlineResponse200255 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200255() *InlineResponse200255 {
	this := InlineResponse200255{}
	return &this
}

// NewInlineResponse200255WithDefaults instantiates a new InlineResponse200255 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200255WithDefaults() *InlineResponse200255 {
	this := InlineResponse200255{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200255) GetUsage() InlineResponse200255Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200255Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200255) GetUsageOk() (*InlineResponse200255Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200255) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200255Usage and assigns it to the Usage field.
func (o *InlineResponse200255) SetUsage(v InlineResponse200255Usage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200255) GetCounts() InlineResponse200255Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200255Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200255) GetCountsOk() (*InlineResponse200255Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200255) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200255Counts and assigns it to the Counts field.
func (o *InlineResponse200255) SetCounts(v InlineResponse200255Counts) {
	o.Counts = &v
}

func (o InlineResponse200255) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200255 struct {
	value *InlineResponse200255
	isSet bool
}

func (v NullableInlineResponse200255) Get() *InlineResponse200255 {
	return v.value
}

func (v *NullableInlineResponse200255) Set(val *InlineResponse200255) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200255) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200255) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200255(val *InlineResponse200255) *NullableInlineResponse200255 {
	return &NullableInlineResponse200255{value: val, isSet: true}
}

func (v NullableInlineResponse200255) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200255) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


