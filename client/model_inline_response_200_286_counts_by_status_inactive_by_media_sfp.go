/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200286CountsByStatusInactiveByMediaSfp The count data for inactive SFP ports
type InlineResponse200286CountsByStatusInactiveByMediaSfp struct {
	// The total number of inactive SFP ports
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse200286CountsByStatusInactiveByMediaSfp instantiates a new InlineResponse200286CountsByStatusInactiveByMediaSfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200286CountsByStatusInactiveByMediaSfp() *InlineResponse200286CountsByStatusInactiveByMediaSfp {
	this := InlineResponse200286CountsByStatusInactiveByMediaSfp{}
	return &this
}

// NewInlineResponse200286CountsByStatusInactiveByMediaSfpWithDefaults instantiates a new InlineResponse200286CountsByStatusInactiveByMediaSfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200286CountsByStatusInactiveByMediaSfpWithDefaults() *InlineResponse200286CountsByStatusInactiveByMediaSfp {
	this := InlineResponse200286CountsByStatusInactiveByMediaSfp{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200286CountsByStatusInactiveByMediaSfp) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286CountsByStatusInactiveByMediaSfp) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200286CountsByStatusInactiveByMediaSfp) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200286CountsByStatusInactiveByMediaSfp) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse200286CountsByStatusInactiveByMediaSfp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200286CountsByStatusInactiveByMediaSfp struct {
	value *InlineResponse200286CountsByStatusInactiveByMediaSfp
	isSet bool
}

func (v NullableInlineResponse200286CountsByStatusInactiveByMediaSfp) Get() *InlineResponse200286CountsByStatusInactiveByMediaSfp {
	return v.value
}

func (v *NullableInlineResponse200286CountsByStatusInactiveByMediaSfp) Set(val *InlineResponse200286CountsByStatusInactiveByMediaSfp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200286CountsByStatusInactiveByMediaSfp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200286CountsByStatusInactiveByMediaSfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200286CountsByStatusInactiveByMediaSfp(val *InlineResponse200286CountsByStatusInactiveByMediaSfp) *NullableInlineResponse200286CountsByStatusInactiveByMediaSfp {
	return &NullableInlineResponse200286CountsByStatusInactiveByMediaSfp{value: val, isSet: true}
}

func (v NullableInlineResponse200286CountsByStatusInactiveByMediaSfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200286CountsByStatusInactiveByMediaSfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


