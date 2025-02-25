/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200286StatesUnusedActiveOldestActivation Information about the oldest historical license activation
type InlineResponse200286StatesUnusedActiveOldestActivation struct {
	// The oldest license activation date
	ActivationDate *string `json:"activationDate,omitempty"`
	// The number of licenses that activated on this date
	ActiveCount *int32 `json:"activeCount,omitempty"`
}

// NewInlineResponse200286StatesUnusedActiveOldestActivation instantiates a new InlineResponse200286StatesUnusedActiveOldestActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200286StatesUnusedActiveOldestActivation() *InlineResponse200286StatesUnusedActiveOldestActivation {
	this := InlineResponse200286StatesUnusedActiveOldestActivation{}
	return &this
}

// NewInlineResponse200286StatesUnusedActiveOldestActivationWithDefaults instantiates a new InlineResponse200286StatesUnusedActiveOldestActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200286StatesUnusedActiveOldestActivationWithDefaults() *InlineResponse200286StatesUnusedActiveOldestActivation {
	this := InlineResponse200286StatesUnusedActiveOldestActivation{}
	return &this
}

// GetActivationDate returns the ActivationDate field value if set, zero value otherwise.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) GetActivationDate() string {
	if o == nil || isNil(o.ActivationDate) {
		var ret string
		return ret
	}
	return *o.ActivationDate
}

// GetActivationDateOk returns a tuple with the ActivationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) GetActivationDateOk() (*string, bool) {
	if o == nil || isNil(o.ActivationDate) {
    return nil, false
	}
	return o.ActivationDate, true
}

// HasActivationDate returns a boolean if a field has been set.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) HasActivationDate() bool {
	if o != nil && !isNil(o.ActivationDate) {
		return true
	}

	return false
}

// SetActivationDate gets a reference to the given string and assigns it to the ActivationDate field.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) SetActivationDate(v string) {
	o.ActivationDate = &v
}

// GetActiveCount returns the ActiveCount field value if set, zero value otherwise.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) GetActiveCount() int32 {
	if o == nil || isNil(o.ActiveCount) {
		var ret int32
		return ret
	}
	return *o.ActiveCount
}

// GetActiveCountOk returns a tuple with the ActiveCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) GetActiveCountOk() (*int32, bool) {
	if o == nil || isNil(o.ActiveCount) {
    return nil, false
	}
	return o.ActiveCount, true
}

// HasActiveCount returns a boolean if a field has been set.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) HasActiveCount() bool {
	if o != nil && !isNil(o.ActiveCount) {
		return true
	}

	return false
}

// SetActiveCount gets a reference to the given int32 and assigns it to the ActiveCount field.
func (o *InlineResponse200286StatesUnusedActiveOldestActivation) SetActiveCount(v int32) {
	o.ActiveCount = &v
}

func (o InlineResponse200286StatesUnusedActiveOldestActivation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActivationDate) {
		toSerialize["activationDate"] = o.ActivationDate
	}
	if !isNil(o.ActiveCount) {
		toSerialize["activeCount"] = o.ActiveCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200286StatesUnusedActiveOldestActivation struct {
	value *InlineResponse200286StatesUnusedActiveOldestActivation
	isSet bool
}

func (v NullableInlineResponse200286StatesUnusedActiveOldestActivation) Get() *InlineResponse200286StatesUnusedActiveOldestActivation {
	return v.value
}

func (v *NullableInlineResponse200286StatesUnusedActiveOldestActivation) Set(val *InlineResponse200286StatesUnusedActiveOldestActivation) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200286StatesUnusedActiveOldestActivation) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200286StatesUnusedActiveOldestActivation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200286StatesUnusedActiveOldestActivation(val *InlineResponse200286StatesUnusedActiveOldestActivation) *NullableInlineResponse200286StatesUnusedActiveOldestActivation {
	return &NullableInlineResponse200286StatesUnusedActiveOldestActivation{value: val, isSet: true}
}

func (v NullableInlineResponse200286StatesUnusedActiveOldestActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200286StatesUnusedActiveOldestActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


