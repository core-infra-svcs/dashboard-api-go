/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200132 struct for InlineResponse200132
type InlineResponse200132 struct {
	Usage *InlineResponse200132Usage `json:"usage,omitempty"`
	Counts *InlineResponse200132Counts `json:"counts,omitempty"`
}

// NewInlineResponse200132 instantiates a new InlineResponse200132 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200132() *InlineResponse200132 {
	this := InlineResponse200132{}
	return &this
}

// NewInlineResponse200132WithDefaults instantiates a new InlineResponse200132 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200132WithDefaults() *InlineResponse200132 {
	this := InlineResponse200132{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200132) GetUsage() InlineResponse200132Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200132Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetUsageOk() (*InlineResponse200132Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200132) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200132Usage and assigns it to the Usage field.
func (o *InlineResponse200132) SetUsage(v InlineResponse200132Usage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200132) GetCounts() InlineResponse200132Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200132Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetCountsOk() (*InlineResponse200132Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200132) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200132Counts and assigns it to the Counts field.
func (o *InlineResponse200132) SetCounts(v InlineResponse200132Counts) {
	o.Counts = &v
}

func (o InlineResponse200132) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200132 struct {
	value *InlineResponse200132
	isSet bool
}

func (v NullableInlineResponse200132) Get() *InlineResponse200132 {
	return v.value
}

func (v *NullableInlineResponse200132) Set(val *InlineResponse200132) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200132) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200132) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200132(val *InlineResponse200132) *NullableInlineResponse200132 {
	return &NullableInlineResponse200132{value: val, isSet: true}
}

func (v NullableInlineResponse200132) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200132) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


