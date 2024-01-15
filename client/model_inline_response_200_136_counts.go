/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200136Counts Client count information
type InlineResponse200136Counts struct {
	// Total number of clients with data usage in organization
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse200136Counts instantiates a new InlineResponse200136Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200136Counts() *InlineResponse200136Counts {
	this := InlineResponse200136Counts{}
	return &this
}

// NewInlineResponse200136CountsWithDefaults instantiates a new InlineResponse200136Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200136CountsWithDefaults() *InlineResponse200136Counts {
	this := InlineResponse200136Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200136Counts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136Counts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200136Counts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200136Counts) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse200136Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200136Counts struct {
	value *InlineResponse200136Counts
	isSet bool
}

func (v NullableInlineResponse200136Counts) Get() *InlineResponse200136Counts {
	return v.value
}

func (v *NullableInlineResponse200136Counts) Set(val *InlineResponse200136Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200136Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200136Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200136Counts(val *InlineResponse200136Counts) *NullableInlineResponse200136Counts {
	return &NullableInlineResponse200136Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200136Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200136Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

