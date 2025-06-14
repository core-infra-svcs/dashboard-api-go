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

// InlineResponse200273Counts Object containing counts
type InlineResponse200273Counts struct {
	Packets *InlineResponse200273CountsPackets `json:"packets,omitempty"`
}

// NewInlineResponse200273Counts instantiates a new InlineResponse200273Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200273Counts() *InlineResponse200273Counts {
	this := InlineResponse200273Counts{}
	return &this
}

// NewInlineResponse200273CountsWithDefaults instantiates a new InlineResponse200273Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200273CountsWithDefaults() *InlineResponse200273Counts {
	this := InlineResponse200273Counts{}
	return &this
}

// GetPackets returns the Packets field value if set, zero value otherwise.
func (o *InlineResponse200273Counts) GetPackets() InlineResponse200273CountsPackets {
	if o == nil || isNil(o.Packets) {
		var ret InlineResponse200273CountsPackets
		return ret
	}
	return *o.Packets
}

// GetPacketsOk returns a tuple with the Packets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200273Counts) GetPacketsOk() (*InlineResponse200273CountsPackets, bool) {
	if o == nil || isNil(o.Packets) {
    return nil, false
	}
	return o.Packets, true
}

// HasPackets returns a boolean if a field has been set.
func (o *InlineResponse200273Counts) HasPackets() bool {
	if o != nil && !isNil(o.Packets) {
		return true
	}

	return false
}

// SetPackets gets a reference to the given InlineResponse200273CountsPackets and assigns it to the Packets field.
func (o *InlineResponse200273Counts) SetPackets(v InlineResponse200273CountsPackets) {
	o.Packets = &v
}

func (o InlineResponse200273Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Packets) {
		toSerialize["packets"] = o.Packets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200273Counts struct {
	value *InlineResponse200273Counts
	isSet bool
}

func (v NullableInlineResponse200273Counts) Get() *InlineResponse200273Counts {
	return v.value
}

func (v *NullableInlineResponse200273Counts) Set(val *InlineResponse200273Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200273Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200273Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200273Counts(val *InlineResponse200273Counts) *NullableInlineResponse200273Counts {
	return &NullableInlineResponse200273Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200273Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200273Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


