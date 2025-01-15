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

// InlineResponse200308CountsByStatusInactiveByMediaSfp The count data for inactive SFP ports
type InlineResponse200308CountsByStatusInactiveByMediaSfp struct {
	// The total number of inactive SFP ports
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse200308CountsByStatusInactiveByMediaSfp instantiates a new InlineResponse200308CountsByStatusInactiveByMediaSfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308CountsByStatusInactiveByMediaSfp() *InlineResponse200308CountsByStatusInactiveByMediaSfp {
	this := InlineResponse200308CountsByStatusInactiveByMediaSfp{}
	return &this
}

// NewInlineResponse200308CountsByStatusInactiveByMediaSfpWithDefaults instantiates a new InlineResponse200308CountsByStatusInactiveByMediaSfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308CountsByStatusInactiveByMediaSfpWithDefaults() *InlineResponse200308CountsByStatusInactiveByMediaSfp {
	this := InlineResponse200308CountsByStatusInactiveByMediaSfp{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200308CountsByStatusInactiveByMediaSfp) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308CountsByStatusInactiveByMediaSfp) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200308CountsByStatusInactiveByMediaSfp) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200308CountsByStatusInactiveByMediaSfp) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse200308CountsByStatusInactiveByMediaSfp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308CountsByStatusInactiveByMediaSfp struct {
	value *InlineResponse200308CountsByStatusInactiveByMediaSfp
	isSet bool
}

func (v NullableInlineResponse200308CountsByStatusInactiveByMediaSfp) Get() *InlineResponse200308CountsByStatusInactiveByMediaSfp {
	return v.value
}

func (v *NullableInlineResponse200308CountsByStatusInactiveByMediaSfp) Set(val *InlineResponse200308CountsByStatusInactiveByMediaSfp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308CountsByStatusInactiveByMediaSfp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308CountsByStatusInactiveByMediaSfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308CountsByStatusInactiveByMediaSfp(val *InlineResponse200308CountsByStatusInactiveByMediaSfp) *NullableInlineResponse200308CountsByStatusInactiveByMediaSfp {
	return &NullableInlineResponse200308CountsByStatusInactiveByMediaSfp{value: val, isSet: true}
}

func (v NullableInlineResponse200308CountsByStatusInactiveByMediaSfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308CountsByStatusInactiveByMediaSfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


