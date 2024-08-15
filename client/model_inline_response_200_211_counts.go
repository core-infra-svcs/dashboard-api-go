/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200211Counts counts
type InlineResponse200211Counts struct {
	ByStatus *InlineResponse200211CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200211Counts instantiates a new InlineResponse200211Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200211Counts() *InlineResponse200211Counts {
	this := InlineResponse200211Counts{}
	return &this
}

// NewInlineResponse200211CountsWithDefaults instantiates a new InlineResponse200211Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200211CountsWithDefaults() *InlineResponse200211Counts {
	this := InlineResponse200211Counts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200211Counts) GetByStatus() InlineResponse200211CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200211CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200211Counts) GetByStatusOk() (*InlineResponse200211CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200211Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200211CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200211Counts) SetByStatus(v InlineResponse200211CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200211Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200211Counts struct {
	value *InlineResponse200211Counts
	isSet bool
}

func (v NullableInlineResponse200211Counts) Get() *InlineResponse200211Counts {
	return v.value
}

func (v *NullableInlineResponse200211Counts) Set(val *InlineResponse200211Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200211Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200211Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200211Counts(val *InlineResponse200211Counts) *NullableInlineResponse200211Counts {
	return &NullableInlineResponse200211Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200211Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200211Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


