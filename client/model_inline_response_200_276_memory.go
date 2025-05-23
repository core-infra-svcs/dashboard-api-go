/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200276Memory Information regarding memory usage and availability on the device
type InlineResponse200276Memory struct {
	Used *InlineResponse200276MemoryUsed `json:"used,omitempty"`
	Free *InlineResponse200276MemoryFree `json:"free,omitempty"`
}

// NewInlineResponse200276Memory instantiates a new InlineResponse200276Memory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200276Memory() *InlineResponse200276Memory {
	this := InlineResponse200276Memory{}
	return &this
}

// NewInlineResponse200276MemoryWithDefaults instantiates a new InlineResponse200276Memory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200276MemoryWithDefaults() *InlineResponse200276Memory {
	this := InlineResponse200276Memory{}
	return &this
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *InlineResponse200276Memory) GetUsed() InlineResponse200276MemoryUsed {
	if o == nil || isNil(o.Used) {
		var ret InlineResponse200276MemoryUsed
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276Memory) GetUsedOk() (*InlineResponse200276MemoryUsed, bool) {
	if o == nil || isNil(o.Used) {
    return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *InlineResponse200276Memory) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given InlineResponse200276MemoryUsed and assigns it to the Used field.
func (o *InlineResponse200276Memory) SetUsed(v InlineResponse200276MemoryUsed) {
	o.Used = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *InlineResponse200276Memory) GetFree() InlineResponse200276MemoryFree {
	if o == nil || isNil(o.Free) {
		var ret InlineResponse200276MemoryFree
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276Memory) GetFreeOk() (*InlineResponse200276MemoryFree, bool) {
	if o == nil || isNil(o.Free) {
    return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *InlineResponse200276Memory) HasFree() bool {
	if o != nil && !isNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given InlineResponse200276MemoryFree and assigns it to the Free field.
func (o *InlineResponse200276Memory) SetFree(v InlineResponse200276MemoryFree) {
	o.Free = &v
}

func (o InlineResponse200276Memory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !isNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200276Memory struct {
	value *InlineResponse200276Memory
	isSet bool
}

func (v NullableInlineResponse200276Memory) Get() *InlineResponse200276Memory {
	return v.value
}

func (v *NullableInlineResponse200276Memory) Set(val *InlineResponse200276Memory) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200276Memory) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200276Memory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200276Memory(val *InlineResponse200276Memory) *NullableInlineResponse200276Memory {
	return &NullableInlineResponse200276Memory{value: val, isSet: true}
}

func (v NullableInlineResponse200276Memory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200276Memory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


