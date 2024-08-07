/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200291Counts The count data of all ports
type InlineResponse200291Counts struct {
	// The total number of ports
	Total *int32 `json:"total,omitempty"`
	ByStatus *InlineResponse200291CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200291Counts instantiates a new InlineResponse200291Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200291Counts() *InlineResponse200291Counts {
	this := InlineResponse200291Counts{}
	return &this
}

// NewInlineResponse200291CountsWithDefaults instantiates a new InlineResponse200291Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200291CountsWithDefaults() *InlineResponse200291Counts {
	this := InlineResponse200291Counts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200291Counts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291Counts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200291Counts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200291Counts) SetTotal(v int32) {
	o.Total = &v
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200291Counts) GetByStatus() InlineResponse200291CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200291CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291Counts) GetByStatusOk() (*InlineResponse200291CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200291Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200291CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200291Counts) SetByStatus(v InlineResponse200291CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200291Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200291Counts struct {
	value *InlineResponse200291Counts
	isSet bool
}

func (v NullableInlineResponse200291Counts) Get() *InlineResponse200291Counts {
	return v.value
}

func (v *NullableInlineResponse200291Counts) Set(val *InlineResponse200291Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200291Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200291Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200291Counts(val *InlineResponse200291Counts) *NullableInlineResponse200291Counts {
	return &NullableInlineResponse200291Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200291Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200291Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


