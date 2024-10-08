/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200216Counts Counts of alerts on the organization
type InlineResponse200216Counts struct {
	// Total number of alerts on the organization
	Total int32 `json:"total"`
	// Counts of alerts on organization by severity
	BySeverity []InlineResponse200216CountsBySeverity `json:"bySeverity"`
}

// NewInlineResponse200216Counts instantiates a new InlineResponse200216Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200216Counts(total int32, bySeverity []InlineResponse200216CountsBySeverity) *InlineResponse200216Counts {
	this := InlineResponse200216Counts{}
	this.Total = total
	this.BySeverity = bySeverity
	return &this
}

// NewInlineResponse200216CountsWithDefaults instantiates a new InlineResponse200216Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200216CountsWithDefaults() *InlineResponse200216Counts {
	this := InlineResponse200216Counts{}
	return &this
}

// GetTotal returns the Total field value
func (o *InlineResponse200216Counts) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Counts) GetTotalOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InlineResponse200216Counts) SetTotal(v int32) {
	o.Total = v
}

// GetBySeverity returns the BySeverity field value
func (o *InlineResponse200216Counts) GetBySeverity() []InlineResponse200216CountsBySeverity {
	if o == nil {
		var ret []InlineResponse200216CountsBySeverity
		return ret
	}

	return o.BySeverity
}

// GetBySeverityOk returns a tuple with the BySeverity field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Counts) GetBySeverityOk() ([]InlineResponse200216CountsBySeverity, bool) {
	if o == nil {
    return nil, false
	}
	return o.BySeverity, true
}

// SetBySeverity sets field value
func (o *InlineResponse200216Counts) SetBySeverity(v []InlineResponse200216CountsBySeverity) {
	o.BySeverity = v
}

func (o InlineResponse200216Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["bySeverity"] = o.BySeverity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200216Counts struct {
	value *InlineResponse200216Counts
	isSet bool
}

func (v NullableInlineResponse200216Counts) Get() *InlineResponse200216Counts {
	return v.value
}

func (v *NullableInlineResponse200216Counts) Set(val *InlineResponse200216Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200216Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200216Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200216Counts(val *InlineResponse200216Counts) *NullableInlineResponse200216Counts {
	return &NullableInlineResponse200216Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200216Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200216Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


