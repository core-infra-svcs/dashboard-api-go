/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200146Counts counts
type InlineResponse200146Counts struct {
	ByStatus *InlineResponse200146CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200146Counts instantiates a new InlineResponse200146Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200146Counts() *InlineResponse200146Counts {
	this := InlineResponse200146Counts{}
	return &this
}

// NewInlineResponse200146CountsWithDefaults instantiates a new InlineResponse200146Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200146CountsWithDefaults() *InlineResponse200146Counts {
	this := InlineResponse200146Counts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200146Counts) GetByStatus() InlineResponse200146CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200146CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146Counts) GetByStatusOk() (*InlineResponse200146CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200146Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200146CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200146Counts) SetByStatus(v InlineResponse200146CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200146Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200146Counts struct {
	value *InlineResponse200146Counts
	isSet bool
}

func (v NullableInlineResponse200146Counts) Get() *InlineResponse200146Counts {
	return v.value
}

func (v *NullableInlineResponse200146Counts) Set(val *InlineResponse200146Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200146Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200146Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200146Counts(val *InlineResponse200146Counts) *NullableInlineResponse200146Counts {
	return &NullableInlineResponse200146Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200146Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200146Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


