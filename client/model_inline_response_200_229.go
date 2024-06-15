/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200229 struct for InlineResponse200229
type InlineResponse200229 struct {
	Usage *InlineResponse200229Usage `json:"usage,omitempty"`
	Counts *InlineResponse200229Counts `json:"counts,omitempty"`
}

// NewInlineResponse200229 instantiates a new InlineResponse200229 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200229() *InlineResponse200229 {
	this := InlineResponse200229{}
	return &this
}

// NewInlineResponse200229WithDefaults instantiates a new InlineResponse200229 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200229WithDefaults() *InlineResponse200229 {
	this := InlineResponse200229{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200229) GetUsage() InlineResponse200229Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200229Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200229) GetUsageOk() (*InlineResponse200229Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200229) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200229Usage and assigns it to the Usage field.
func (o *InlineResponse200229) SetUsage(v InlineResponse200229Usage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200229) GetCounts() InlineResponse200229Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200229Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200229) GetCountsOk() (*InlineResponse200229Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200229) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200229Counts and assigns it to the Counts field.
func (o *InlineResponse200229) SetCounts(v InlineResponse200229Counts) {
	o.Counts = &v
}

func (o InlineResponse200229) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200229 struct {
	value *InlineResponse200229
	isSet bool
}

func (v NullableInlineResponse200229) Get() *InlineResponse200229 {
	return v.value
}

func (v *NullableInlineResponse200229) Set(val *InlineResponse200229) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200229) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200229) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200229(val *InlineResponse200229) *NullableInlineResponse200229 {
	return &NullableInlineResponse200229{value: val, isSet: true}
}

func (v NullableInlineResponse200229) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200229) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


