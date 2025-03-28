/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200235Counts Counts of alerts on the organization
type InlineResponse200235Counts struct {
	// Total number of alerts on the organization
	Total int32 `json:"total"`
	// Counts of alerts on organization by severity
	BySeverity []InlineResponse200235CountsBySeverity `json:"bySeverity"`
}

// NewInlineResponse200235Counts instantiates a new InlineResponse200235Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200235Counts(total int32, bySeverity []InlineResponse200235CountsBySeverity) *InlineResponse200235Counts {
	this := InlineResponse200235Counts{}
	this.Total = total
	this.BySeverity = bySeverity
	return &this
}

// NewInlineResponse200235CountsWithDefaults instantiates a new InlineResponse200235Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200235CountsWithDefaults() *InlineResponse200235Counts {
	this := InlineResponse200235Counts{}
	return &this
}

// GetTotal returns the Total field value
func (o *InlineResponse200235Counts) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200235Counts) GetTotalOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InlineResponse200235Counts) SetTotal(v int32) {
	o.Total = v
}

// GetBySeverity returns the BySeverity field value
func (o *InlineResponse200235Counts) GetBySeverity() []InlineResponse200235CountsBySeverity {
	if o == nil {
		var ret []InlineResponse200235CountsBySeverity
		return ret
	}

	return o.BySeverity
}

// GetBySeverityOk returns a tuple with the BySeverity field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200235Counts) GetBySeverityOk() ([]InlineResponse200235CountsBySeverity, bool) {
	if o == nil {
    return nil, false
	}
	return o.BySeverity, true
}

// SetBySeverity sets field value
func (o *InlineResponse200235Counts) SetBySeverity(v []InlineResponse200235CountsBySeverity) {
	o.BySeverity = v
}

func (o InlineResponse200235Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["bySeverity"] = o.BySeverity
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200235Counts struct {
	value *InlineResponse200235Counts
	isSet bool
}

func (v NullableInlineResponse200235Counts) Get() *InlineResponse200235Counts {
	return v.value
}

func (v *NullableInlineResponse200235Counts) Set(val *InlineResponse200235Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200235Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200235Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200235Counts(val *InlineResponse200235Counts) *NullableInlineResponse200235Counts {
	return &NullableInlineResponse200235Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200235Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200235Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


