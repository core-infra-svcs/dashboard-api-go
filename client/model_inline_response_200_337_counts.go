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

// InlineResponse200337Counts The count data of all ports
type InlineResponse200337Counts struct {
	// The total number of ports
	Total *int32 `json:"total,omitempty"`
	ByStatus *InlineResponse200337CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200337Counts instantiates a new InlineResponse200337Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200337Counts() *InlineResponse200337Counts {
	this := InlineResponse200337Counts{}
	return &this
}

// NewInlineResponse200337CountsWithDefaults instantiates a new InlineResponse200337Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200337CountsWithDefaults() *InlineResponse200337Counts {
	this := InlineResponse200337Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200337Counts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Counts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200337Counts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200337Counts) SetTotal(v int32) {
	o.Total = &v
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200337Counts) GetByStatus() InlineResponse200337CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200337CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337Counts) GetByStatusOk() (*InlineResponse200337CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200337Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200337CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200337Counts) SetByStatus(v InlineResponse200337CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200337Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200337Counts struct {
	value *InlineResponse200337Counts
	isSet bool
}

func (v NullableInlineResponse200337Counts) Get() *InlineResponse200337Counts {
	return v.value
}

func (v *NullableInlineResponse200337Counts) Set(val *InlineResponse200337Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200337Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200337Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200337Counts(val *InlineResponse200337Counts) *NullableInlineResponse200337Counts {
	return &NullableInlineResponse200337Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200337Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200337Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


