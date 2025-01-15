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

// InlineResponse20082 struct for InlineResponse20082
type InlineResponse20082 struct {
	Counts *InlineResponse20082Counts `json:"counts,omitempty"`
	Usages *InlineResponse20082Usages `json:"usages,omitempty"`
}

// NewInlineResponse20082 instantiates a new InlineResponse20082 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20082() *InlineResponse20082 {
	this := InlineResponse20082{}
	return &this
}

// NewInlineResponse20082WithDefaults instantiates a new InlineResponse20082 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20082WithDefaults() *InlineResponse20082 {
	this := InlineResponse20082{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse20082) GetCounts() InlineResponse20082Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse20082Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082) GetCountsOk() (*InlineResponse20082Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse20082) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse20082Counts and assigns it to the Counts field.
func (o *InlineResponse20082) SetCounts(v InlineResponse20082Counts) {
	o.Counts = &v
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *InlineResponse20082) GetUsages() InlineResponse20082Usages {
	if o == nil || isNil(o.Usages) {
		var ret InlineResponse20082Usages
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082) GetUsagesOk() (*InlineResponse20082Usages, bool) {
	if o == nil || isNil(o.Usages) {
    return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *InlineResponse20082) HasUsages() bool {
	if o != nil && !isNil(o.Usages) {
		return true
	}

	return false
}

// SetUsages gets a reference to the given InlineResponse20082Usages and assigns it to the Usages field.
func (o *InlineResponse20082) SetUsages(v InlineResponse20082Usages) {
	o.Usages = &v
}

func (o InlineResponse20082) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !isNil(o.Usages) {
		toSerialize["usages"] = o.Usages
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20082 struct {
	value *InlineResponse20082
	isSet bool
}

func (v NullableInlineResponse20082) Get() *InlineResponse20082 {
	return v.value
}

func (v *NullableInlineResponse20082) Set(val *InlineResponse20082) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20082) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20082) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20082(val *InlineResponse20082) *NullableInlineResponse20082 {
	return &NullableInlineResponse20082{value: val, isSet: true}
}

func (v NullableInlineResponse20082) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20082) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


