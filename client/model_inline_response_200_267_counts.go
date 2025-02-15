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

// InlineResponse200267Counts counts
type InlineResponse200267Counts struct {
	ByStatus *InlineResponse200267CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200267Counts instantiates a new InlineResponse200267Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200267Counts() *InlineResponse200267Counts {
	this := InlineResponse200267Counts{}
	return &this
}

// NewInlineResponse200267CountsWithDefaults instantiates a new InlineResponse200267Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200267CountsWithDefaults() *InlineResponse200267Counts {
	this := InlineResponse200267Counts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200267Counts) GetByStatus() InlineResponse200267CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200267CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200267Counts) GetByStatusOk() (*InlineResponse200267CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200267Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200267CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200267Counts) SetByStatus(v InlineResponse200267CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200267Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200267Counts struct {
	value *InlineResponse200267Counts
	isSet bool
}

func (v NullableInlineResponse200267Counts) Get() *InlineResponse200267Counts {
	return v.value
}

func (v *NullableInlineResponse200267Counts) Set(val *InlineResponse200267Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200267Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200267Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200267Counts(val *InlineResponse200267Counts) *NullableInlineResponse200267Counts {
	return &NullableInlineResponse200267Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200267Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200267Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


