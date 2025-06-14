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

// InlineResponse200273File Object containing information about the file
type InlineResponse200273File struct {
	// File size of packet capture file
	Size *int32 `json:"size,omitempty"`
}

// NewInlineResponse200273File instantiates a new InlineResponse200273File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200273File() *InlineResponse200273File {
	this := InlineResponse200273File{}
	return &this
}

// NewInlineResponse200273FileWithDefaults instantiates a new InlineResponse200273File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200273FileWithDefaults() *InlineResponse200273File {
	this := InlineResponse200273File{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InlineResponse200273File) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200273File) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InlineResponse200273File) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *InlineResponse200273File) SetSize(v int32) {
	o.Size = &v
}

func (o InlineResponse200273File) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200273File struct {
	value *InlineResponse200273File
	isSet bool
}

func (v NullableInlineResponse200273File) Get() *InlineResponse200273File {
	return v.value
}

func (v *NullableInlineResponse200273File) Set(val *InlineResponse200273File) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200273File) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200273File) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200273File(val *InlineResponse200273File) *NullableInlineResponse200273File {
	return &NullableInlineResponse200273File{value: val, isSet: true}
}

func (v NullableInlineResponse200273File) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200273File) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


