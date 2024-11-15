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

// InlineResponse200241 struct for InlineResponse200241
type InlineResponse200241 struct {
	Usage *InlineResponse200241Usage `json:"usage,omitempty"`
	Counts *InlineResponse200241Counts `json:"counts,omitempty"`
}

// NewInlineResponse200241 instantiates a new InlineResponse200241 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200241() *InlineResponse200241 {
	this := InlineResponse200241{}
	return &this
}

// NewInlineResponse200241WithDefaults instantiates a new InlineResponse200241 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200241WithDefaults() *InlineResponse200241 {
	this := InlineResponse200241{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200241) GetUsage() InlineResponse200241Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200241Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200241) GetUsageOk() (*InlineResponse200241Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200241) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200241Usage and assigns it to the Usage field.
func (o *InlineResponse200241) SetUsage(v InlineResponse200241Usage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200241) GetCounts() InlineResponse200241Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200241Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200241) GetCountsOk() (*InlineResponse200241Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200241) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200241Counts and assigns it to the Counts field.
func (o *InlineResponse200241) SetCounts(v InlineResponse200241Counts) {
	o.Counts = &v
}

func (o InlineResponse200241) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200241 struct {
	value *InlineResponse200241
	isSet bool
}

func (v NullableInlineResponse200241) Get() *InlineResponse200241 {
	return v.value
}

func (v *NullableInlineResponse200241) Set(val *InlineResponse200241) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200241) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200241) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200241(val *InlineResponse200241) *NullableInlineResponse200241 {
	return &NullableInlineResponse200241{value: val, isSet: true}
}

func (v NullableInlineResponse200241) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200241) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


