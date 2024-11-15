/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200305CountsByStatusInactiveByMediaSfp The count data for inactive SFP ports
type InlineResponse200305CountsByStatusInactiveByMediaSfp struct {
	// The total number of inactive SFP ports
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse200305CountsByStatusInactiveByMediaSfp instantiates a new InlineResponse200305CountsByStatusInactiveByMediaSfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200305CountsByStatusInactiveByMediaSfp() *InlineResponse200305CountsByStatusInactiveByMediaSfp {
	this := InlineResponse200305CountsByStatusInactiveByMediaSfp{}
	return &this
}

// NewInlineResponse200305CountsByStatusInactiveByMediaSfpWithDefaults instantiates a new InlineResponse200305CountsByStatusInactiveByMediaSfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200305CountsByStatusInactiveByMediaSfpWithDefaults() *InlineResponse200305CountsByStatusInactiveByMediaSfp {
	this := InlineResponse200305CountsByStatusInactiveByMediaSfp{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatusInactiveByMediaSfp) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatusInactiveByMediaSfp) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatusInactiveByMediaSfp) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200305CountsByStatusInactiveByMediaSfp) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse200305CountsByStatusInactiveByMediaSfp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200305CountsByStatusInactiveByMediaSfp struct {
	value *InlineResponse200305CountsByStatusInactiveByMediaSfp
	isSet bool
}

func (v NullableInlineResponse200305CountsByStatusInactiveByMediaSfp) Get() *InlineResponse200305CountsByStatusInactiveByMediaSfp {
	return v.value
}

func (v *NullableInlineResponse200305CountsByStatusInactiveByMediaSfp) Set(val *InlineResponse200305CountsByStatusInactiveByMediaSfp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200305CountsByStatusInactiveByMediaSfp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200305CountsByStatusInactiveByMediaSfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200305CountsByStatusInactiveByMediaSfp(val *InlineResponse200305CountsByStatusInactiveByMediaSfp) *NullableInlineResponse200305CountsByStatusInactiveByMediaSfp {
	return &NullableInlineResponse200305CountsByStatusInactiveByMediaSfp{value: val, isSet: true}
}

func (v NullableInlineResponse200305CountsByStatusInactiveByMediaSfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200305CountsByStatusInactiveByMediaSfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


