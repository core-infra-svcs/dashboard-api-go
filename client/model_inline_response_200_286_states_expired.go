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

// InlineResponse200286StatesExpired Data for expired licenses
type InlineResponse200286StatesExpired struct {
	// The number of expired licenses
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse200286StatesExpired instantiates a new InlineResponse200286StatesExpired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200286StatesExpired() *InlineResponse200286StatesExpired {
	this := InlineResponse200286StatesExpired{}
	return &this
}

// NewInlineResponse200286StatesExpiredWithDefaults instantiates a new InlineResponse200286StatesExpired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200286StatesExpiredWithDefaults() *InlineResponse200286StatesExpired {
	this := InlineResponse200286StatesExpired{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200286StatesExpired) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286StatesExpired) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200286StatesExpired) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200286StatesExpired) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse200286StatesExpired) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200286StatesExpired struct {
	value *InlineResponse200286StatesExpired
	isSet bool
}

func (v NullableInlineResponse200286StatesExpired) Get() *InlineResponse200286StatesExpired {
	return v.value
}

func (v *NullableInlineResponse200286StatesExpired) Set(val *InlineResponse200286StatesExpired) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200286StatesExpired) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200286StatesExpired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200286StatesExpired(val *InlineResponse200286StatesExpired) *NullableInlineResponse200286StatesExpired {
	return &NullableInlineResponse200286StatesExpired{value: val, isSet: true}
}

func (v NullableInlineResponse200286StatesExpired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200286StatesExpired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


