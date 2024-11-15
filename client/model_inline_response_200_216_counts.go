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

// InlineResponse200216Counts counts
type InlineResponse200216Counts struct {
	ByStatus *InlineResponse200216CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200216Counts instantiates a new InlineResponse200216Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200216Counts() *InlineResponse200216Counts {
	this := InlineResponse200216Counts{}
	return &this
}

// NewInlineResponse200216CountsWithDefaults instantiates a new InlineResponse200216Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200216CountsWithDefaults() *InlineResponse200216Counts {
	this := InlineResponse200216Counts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200216Counts) GetByStatus() InlineResponse200216CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200216CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200216Counts) GetByStatusOk() (*InlineResponse200216CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200216Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200216CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200216Counts) SetByStatus(v InlineResponse200216CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200216Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
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


