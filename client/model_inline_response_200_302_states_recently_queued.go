/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200302StatesRecentlyQueued Data for recently queued licenses
type InlineResponse200302StatesRecentlyQueued struct {
	// The number of recently queued licenses
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse200302StatesRecentlyQueued instantiates a new InlineResponse200302StatesRecentlyQueued object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200302StatesRecentlyQueued() *InlineResponse200302StatesRecentlyQueued {
	this := InlineResponse200302StatesRecentlyQueued{}
	return &this
}

// NewInlineResponse200302StatesRecentlyQueuedWithDefaults instantiates a new InlineResponse200302StatesRecentlyQueued object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200302StatesRecentlyQueuedWithDefaults() *InlineResponse200302StatesRecentlyQueued {
	this := InlineResponse200302StatesRecentlyQueued{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200302StatesRecentlyQueued) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200302StatesRecentlyQueued) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200302StatesRecentlyQueued) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200302StatesRecentlyQueued) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse200302StatesRecentlyQueued) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200302StatesRecentlyQueued struct {
	value *InlineResponse200302StatesRecentlyQueued
	isSet bool
}

func (v NullableInlineResponse200302StatesRecentlyQueued) Get() *InlineResponse200302StatesRecentlyQueued {
	return v.value
}

func (v *NullableInlineResponse200302StatesRecentlyQueued) Set(val *InlineResponse200302StatesRecentlyQueued) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200302StatesRecentlyQueued) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200302StatesRecentlyQueued) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200302StatesRecentlyQueued(val *InlineResponse200302StatesRecentlyQueued) *NullableInlineResponse200302StatesRecentlyQueued {
	return &NullableInlineResponse200302StatesRecentlyQueued{value: val, isSet: true}
}

func (v NullableInlineResponse200302StatesRecentlyQueued) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200302StatesRecentlyQueued) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


