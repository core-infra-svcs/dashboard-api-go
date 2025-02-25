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

// InlineResponse200320Counts The count data of all ports
type InlineResponse200320Counts struct {
	// The total number of ports
	Total *int32 `json:"total,omitempty"`
	ByStatus *InlineResponse200320CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200320Counts instantiates a new InlineResponse200320Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200320Counts() *InlineResponse200320Counts {
	this := InlineResponse200320Counts{}
	return &this
}

// NewInlineResponse200320CountsWithDefaults instantiates a new InlineResponse200320Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200320CountsWithDefaults() *InlineResponse200320Counts {
	this := InlineResponse200320Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200320Counts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200320Counts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200320Counts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200320Counts) SetTotal(v int32) {
	o.Total = &v
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200320Counts) GetByStatus() InlineResponse200320CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200320CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200320Counts) GetByStatusOk() (*InlineResponse200320CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200320Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200320CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200320Counts) SetByStatus(v InlineResponse200320CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200320Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200320Counts struct {
	value *InlineResponse200320Counts
	isSet bool
}

func (v NullableInlineResponse200320Counts) Get() *InlineResponse200320Counts {
	return v.value
}

func (v *NullableInlineResponse200320Counts) Set(val *InlineResponse200320Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200320Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200320Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200320Counts(val *InlineResponse200320Counts) *NullableInlineResponse200320Counts {
	return &NullableInlineResponse200320Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200320Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200320Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


