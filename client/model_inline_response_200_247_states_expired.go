/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200247StatesExpired Data for expired licenses
type InlineResponse200247StatesExpired struct {
	// The number of expired licenses
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse200247StatesExpired instantiates a new InlineResponse200247StatesExpired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200247StatesExpired() *InlineResponse200247StatesExpired {
	this := InlineResponse200247StatesExpired{}
	return &this
}

// NewInlineResponse200247StatesExpiredWithDefaults instantiates a new InlineResponse200247StatesExpired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200247StatesExpiredWithDefaults() *InlineResponse200247StatesExpired {
	this := InlineResponse200247StatesExpired{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200247StatesExpired) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200247StatesExpired) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200247StatesExpired) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200247StatesExpired) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse200247StatesExpired) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200247StatesExpired struct {
	value *InlineResponse200247StatesExpired
	isSet bool
}

func (v NullableInlineResponse200247StatesExpired) Get() *InlineResponse200247StatesExpired {
	return v.value
}

func (v *NullableInlineResponse200247StatesExpired) Set(val *InlineResponse200247StatesExpired) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200247StatesExpired) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200247StatesExpired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200247StatesExpired(val *InlineResponse200247StatesExpired) *NullableInlineResponse200247StatesExpired {
	return &NullableInlineResponse200247StatesExpired{value: val, isSet: true}
}

func (v NullableInlineResponse200247StatesExpired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200247StatesExpired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


