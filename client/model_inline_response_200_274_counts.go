/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200274Counts Aggregated count data for the license type
type InlineResponse200274Counts struct {
	// The number of unassigned licenses
	Unassigned *int32 `json:"unassigned,omitempty"`
}

// NewInlineResponse200274Counts instantiates a new InlineResponse200274Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200274Counts() *InlineResponse200274Counts {
	this := InlineResponse200274Counts{}
	return &this
}

// NewInlineResponse200274CountsWithDefaults instantiates a new InlineResponse200274Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200274CountsWithDefaults() *InlineResponse200274Counts {
	this := InlineResponse200274Counts{}
	return &this
}

// GetUnassigned returns the Unassigned field value if set, zero value otherwise.
func (o *InlineResponse200274Counts) GetUnassigned() int32 {
	if o == nil || isNil(o.Unassigned) {
		var ret int32
		return ret
	}
	return *o.Unassigned
}

// GetUnassignedOk returns a tuple with the Unassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274Counts) GetUnassignedOk() (*int32, bool) {
	if o == nil || isNil(o.Unassigned) {
    return nil, false
	}
	return o.Unassigned, true
}

// HasUnassigned returns a boolean if a field has been set.
func (o *InlineResponse200274Counts) HasUnassigned() bool {
	if o != nil && !isNil(o.Unassigned) {
		return true
	}

	return false
}

// SetUnassigned gets a reference to the given int32 and assigns it to the Unassigned field.
func (o *InlineResponse200274Counts) SetUnassigned(v int32) {
	o.Unassigned = &v
}

func (o InlineResponse200274Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Unassigned) {
		toSerialize["unassigned"] = o.Unassigned
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200274Counts struct {
	value *InlineResponse200274Counts
	isSet bool
}

func (v NullableInlineResponse200274Counts) Get() *InlineResponse200274Counts {
	return v.value
}

func (v *NullableInlineResponse200274Counts) Set(val *InlineResponse200274Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200274Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200274Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200274Counts(val *InlineResponse200274Counts) *NullableInlineResponse200274Counts {
	return &NullableInlineResponse200274Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200274Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200274Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

