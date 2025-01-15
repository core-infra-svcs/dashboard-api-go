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

// InlineResponse200274StatesUnusedActive Data for unused, active licenses
type InlineResponse200274StatesUnusedActive struct {
	// The number of unused, active licenses
	Count *int32 `json:"count,omitempty"`
	OldestActivation *InlineResponse200274StatesUnusedActiveOldestActivation `json:"oldestActivation,omitempty"`
}

// NewInlineResponse200274StatesUnusedActive instantiates a new InlineResponse200274StatesUnusedActive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200274StatesUnusedActive() *InlineResponse200274StatesUnusedActive {
	this := InlineResponse200274StatesUnusedActive{}
	return &this
}

// NewInlineResponse200274StatesUnusedActiveWithDefaults instantiates a new InlineResponse200274StatesUnusedActive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200274StatesUnusedActiveWithDefaults() *InlineResponse200274StatesUnusedActive {
	this := InlineResponse200274StatesUnusedActive{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200274StatesUnusedActive) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274StatesUnusedActive) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200274StatesUnusedActive) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200274StatesUnusedActive) SetCount(v int32) {
	o.Count = &v
}

// GetOldestActivation returns the OldestActivation field value if set, zero value otherwise.
func (o *InlineResponse200274StatesUnusedActive) GetOldestActivation() InlineResponse200274StatesUnusedActiveOldestActivation {
	if o == nil || isNil(o.OldestActivation) {
		var ret InlineResponse200274StatesUnusedActiveOldestActivation
		return ret
	}
	return *o.OldestActivation
}

// GetOldestActivationOk returns a tuple with the OldestActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274StatesUnusedActive) GetOldestActivationOk() (*InlineResponse200274StatesUnusedActiveOldestActivation, bool) {
	if o == nil || isNil(o.OldestActivation) {
    return nil, false
	}
	return o.OldestActivation, true
}

// HasOldestActivation returns a boolean if a field has been set.
func (o *InlineResponse200274StatesUnusedActive) HasOldestActivation() bool {
	if o != nil && !isNil(o.OldestActivation) {
		return true
	}

	return false
}

// SetOldestActivation gets a reference to the given InlineResponse200274StatesUnusedActiveOldestActivation and assigns it to the OldestActivation field.
func (o *InlineResponse200274StatesUnusedActive) SetOldestActivation(v InlineResponse200274StatesUnusedActiveOldestActivation) {
	o.OldestActivation = &v
}

func (o InlineResponse200274StatesUnusedActive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.OldestActivation) {
		toSerialize["oldestActivation"] = o.OldestActivation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200274StatesUnusedActive struct {
	value *InlineResponse200274StatesUnusedActive
	isSet bool
}

func (v NullableInlineResponse200274StatesUnusedActive) Get() *InlineResponse200274StatesUnusedActive {
	return v.value
}

func (v *NullableInlineResponse200274StatesUnusedActive) Set(val *InlineResponse200274StatesUnusedActive) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200274StatesUnusedActive) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200274StatesUnusedActive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200274StatesUnusedActive(val *InlineResponse200274StatesUnusedActive) *NullableInlineResponse200274StatesUnusedActive {
	return &NullableInlineResponse200274StatesUnusedActive{value: val, isSet: true}
}

func (v NullableInlineResponse200274StatesUnusedActive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200274StatesUnusedActive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

