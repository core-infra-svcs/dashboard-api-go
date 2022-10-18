/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20067Counts Client count information
type InlineResponse20067Counts struct {
	// Total number of clients with data usage in organization
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse20067Counts instantiates a new InlineResponse20067Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20067Counts() *InlineResponse20067Counts {
	this := InlineResponse20067Counts{}
	return &this
}

// NewInlineResponse20067CountsWithDefaults instantiates a new InlineResponse20067Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20067CountsWithDefaults() *InlineResponse20067Counts {
	this := InlineResponse20067Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20067Counts) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20067Counts) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20067Counts) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse20067Counts) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse20067Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20067Counts struct {
	value *InlineResponse20067Counts
	isSet bool
}

func (v NullableInlineResponse20067Counts) Get() *InlineResponse20067Counts {
	return v.value
}

func (v *NullableInlineResponse20067Counts) Set(val *InlineResponse20067Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20067Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20067Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20067Counts(val *InlineResponse20067Counts) *NullableInlineResponse20067Counts {
	return &NullableInlineResponse20067Counts{value: val, isSet: true}
}

func (v NullableInlineResponse20067Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20067Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


