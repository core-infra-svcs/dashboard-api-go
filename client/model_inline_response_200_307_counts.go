/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200307Counts Number of clients on the port in a given time.
type InlineResponse200307Counts struct {
	ByStatus *InlineResponse200307CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200307Counts instantiates a new InlineResponse200307Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200307Counts() *InlineResponse200307Counts {
	this := InlineResponse200307Counts{}
	return &this
}

// NewInlineResponse200307CountsWithDefaults instantiates a new InlineResponse200307Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200307CountsWithDefaults() *InlineResponse200307Counts {
	this := InlineResponse200307Counts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200307Counts) GetByStatus() InlineResponse200307CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200307CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200307Counts) GetByStatusOk() (*InlineResponse200307CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200307Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200307CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200307Counts) SetByStatus(v InlineResponse200307CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200307Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200307Counts struct {
	value *InlineResponse200307Counts
	isSet bool
}

func (v NullableInlineResponse200307Counts) Get() *InlineResponse200307Counts {
	return v.value
}

func (v *NullableInlineResponse200307Counts) Set(val *InlineResponse200307Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200307Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200307Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200307Counts(val *InlineResponse200307Counts) *NullableInlineResponse200307Counts {
	return &NullableInlineResponse200307Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200307Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200307Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


