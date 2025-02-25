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

// InlineResponse200268Memory Information regarding memory usage and availability on the device
type InlineResponse200268Memory struct {
	Used *InlineResponse200268MemoryUsed `json:"used,omitempty"`
	Free *InlineResponse200268MemoryFree `json:"free,omitempty"`
}

// NewInlineResponse200268Memory instantiates a new InlineResponse200268Memory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200268Memory() *InlineResponse200268Memory {
	this := InlineResponse200268Memory{}
	return &this
}

// NewInlineResponse200268MemoryWithDefaults instantiates a new InlineResponse200268Memory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200268MemoryWithDefaults() *InlineResponse200268Memory {
	this := InlineResponse200268Memory{}
	return &this
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *InlineResponse200268Memory) GetUsed() InlineResponse200268MemoryUsed {
	if o == nil || isNil(o.Used) {
		var ret InlineResponse200268MemoryUsed
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Memory) GetUsedOk() (*InlineResponse200268MemoryUsed, bool) {
	if o == nil || isNil(o.Used) {
    return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *InlineResponse200268Memory) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given InlineResponse200268MemoryUsed and assigns it to the Used field.
func (o *InlineResponse200268Memory) SetUsed(v InlineResponse200268MemoryUsed) {
	o.Used = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *InlineResponse200268Memory) GetFree() InlineResponse200268MemoryFree {
	if o == nil || isNil(o.Free) {
		var ret InlineResponse200268MemoryFree
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Memory) GetFreeOk() (*InlineResponse200268MemoryFree, bool) {
	if o == nil || isNil(o.Free) {
    return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *InlineResponse200268Memory) HasFree() bool {
	if o != nil && !isNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given InlineResponse200268MemoryFree and assigns it to the Free field.
func (o *InlineResponse200268Memory) SetFree(v InlineResponse200268MemoryFree) {
	o.Free = &v
}

func (o InlineResponse200268Memory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !isNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200268Memory struct {
	value *InlineResponse200268Memory
	isSet bool
}

func (v NullableInlineResponse200268Memory) Get() *InlineResponse200268Memory {
	return v.value
}

func (v *NullableInlineResponse200268Memory) Set(val *InlineResponse200268Memory) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200268Memory) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200268Memory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200268Memory(val *InlineResponse200268Memory) *NullableInlineResponse200268Memory {
	return &NullableInlineResponse200268Memory{value: val, isSet: true}
}

func (v NullableInlineResponse200268Memory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200268Memory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


