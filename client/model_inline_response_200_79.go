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

// InlineResponse20079 struct for InlineResponse20079
type InlineResponse20079 struct {
	Counts *InlineResponse20079Counts `json:"counts,omitempty"`
	Usages *InlineResponse20079Usages `json:"usages,omitempty"`
}

// NewInlineResponse20079 instantiates a new InlineResponse20079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079WithDefaults() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse20079) GetCounts() InlineResponse20079Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse20079Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetCountsOk() (*InlineResponse20079Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse20079) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse20079Counts and assigns it to the Counts field.
func (o *InlineResponse20079) SetCounts(v InlineResponse20079Counts) {
	o.Counts = &v
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *InlineResponse20079) GetUsages() InlineResponse20079Usages {
	if o == nil || isNil(o.Usages) {
		var ret InlineResponse20079Usages
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetUsagesOk() (*InlineResponse20079Usages, bool) {
	if o == nil || isNil(o.Usages) {
    return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *InlineResponse20079) HasUsages() bool {
	if o != nil && !isNil(o.Usages) {
		return true
	}

	return false
}

// SetUsages gets a reference to the given InlineResponse20079Usages and assigns it to the Usages field.
func (o *InlineResponse20079) SetUsages(v InlineResponse20079Usages) {
	o.Usages = &v
}

func (o InlineResponse20079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !isNil(o.Usages) {
		toSerialize["usages"] = o.Usages
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079 struct {
	value *InlineResponse20079
	isSet bool
}

func (v NullableInlineResponse20079) Get() *InlineResponse20079 {
	return v.value
}

func (v *NullableInlineResponse20079) Set(val *InlineResponse20079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079(val *InlineResponse20079) *NullableInlineResponse20079 {
	return &NullableInlineResponse20079{value: val, isSet: true}
}

func (v NullableInlineResponse20079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


