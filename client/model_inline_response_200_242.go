/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200242 struct for InlineResponse200242
type InlineResponse200242 struct {
	Usage *InlineResponse200242Usage `json:"usage,omitempty"`
	Counts *InlineResponse200242Counts `json:"counts,omitempty"`
}

// NewInlineResponse200242 instantiates a new InlineResponse200242 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200242() *InlineResponse200242 {
	this := InlineResponse200242{}
	return &this
}

// NewInlineResponse200242WithDefaults instantiates a new InlineResponse200242 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200242WithDefaults() *InlineResponse200242 {
	this := InlineResponse200242{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200242) GetUsage() InlineResponse200242Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200242Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200242) GetUsageOk() (*InlineResponse200242Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200242) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200242Usage and assigns it to the Usage field.
func (o *InlineResponse200242) SetUsage(v InlineResponse200242Usage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200242) GetCounts() InlineResponse200242Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200242Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200242) GetCountsOk() (*InlineResponse200242Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200242) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200242Counts and assigns it to the Counts field.
func (o *InlineResponse200242) SetCounts(v InlineResponse200242Counts) {
	o.Counts = &v
}

func (o InlineResponse200242) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200242 struct {
	value *InlineResponse200242
	isSet bool
}

func (v NullableInlineResponse200242) Get() *InlineResponse200242 {
	return v.value
}

func (v *NullableInlineResponse200242) Set(val *InlineResponse200242) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200242) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200242) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200242(val *InlineResponse200242) *NullableInlineResponse200242 {
	return &NullableInlineResponse200242{value: val, isSet: true}
}

func (v NullableInlineResponse200242) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200242) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


