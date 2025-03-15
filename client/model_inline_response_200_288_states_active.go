/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200288StatesActive Data for active licenses
type InlineResponse200288StatesActive struct {
	// The number of active licenses
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse200288StatesActive instantiates a new InlineResponse200288StatesActive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200288StatesActive() *InlineResponse200288StatesActive {
	this := InlineResponse200288StatesActive{}
	return &this
}

// NewInlineResponse200288StatesActiveWithDefaults instantiates a new InlineResponse200288StatesActive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200288StatesActiveWithDefaults() *InlineResponse200288StatesActive {
	this := InlineResponse200288StatesActive{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200288StatesActive) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200288StatesActive) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200288StatesActive) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200288StatesActive) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse200288StatesActive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200288StatesActive struct {
	value *InlineResponse200288StatesActive
	isSet bool
}

func (v NullableInlineResponse200288StatesActive) Get() *InlineResponse200288StatesActive {
	return v.value
}

func (v *NullableInlineResponse200288StatesActive) Set(val *InlineResponse200288StatesActive) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200288StatesActive) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200288StatesActive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200288StatesActive(val *InlineResponse200288StatesActive) *NullableInlineResponse200288StatesActive {
	return &NullableInlineResponse200288StatesActive{value: val, isSet: true}
}

func (v NullableInlineResponse200288StatesActive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200288StatesActive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


