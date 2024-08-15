/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200259StatesRecentlyQueued Data for recently queued licenses
type InlineResponse200259StatesRecentlyQueued struct {
	// The number of recently queued licenses
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse200259StatesRecentlyQueued instantiates a new InlineResponse200259StatesRecentlyQueued object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200259StatesRecentlyQueued() *InlineResponse200259StatesRecentlyQueued {
	this := InlineResponse200259StatesRecentlyQueued{}
	return &this
}

// NewInlineResponse200259StatesRecentlyQueuedWithDefaults instantiates a new InlineResponse200259StatesRecentlyQueued object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200259StatesRecentlyQueuedWithDefaults() *InlineResponse200259StatesRecentlyQueued {
	this := InlineResponse200259StatesRecentlyQueued{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200259StatesRecentlyQueued) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200259StatesRecentlyQueued) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200259StatesRecentlyQueued) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200259StatesRecentlyQueued) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse200259StatesRecentlyQueued) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200259StatesRecentlyQueued struct {
	value *InlineResponse200259StatesRecentlyQueued
	isSet bool
}

func (v NullableInlineResponse200259StatesRecentlyQueued) Get() *InlineResponse200259StatesRecentlyQueued {
	return v.value
}

func (v *NullableInlineResponse200259StatesRecentlyQueued) Set(val *InlineResponse200259StatesRecentlyQueued) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200259StatesRecentlyQueued) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200259StatesRecentlyQueued) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200259StatesRecentlyQueued(val *InlineResponse200259StatesRecentlyQueued) *NullableInlineResponse200259StatesRecentlyQueued {
	return &NullableInlineResponse200259StatesRecentlyQueued{value: val, isSet: true}
}

func (v NullableInlineResponse200259StatesRecentlyQueued) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200259StatesRecentlyQueued) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


