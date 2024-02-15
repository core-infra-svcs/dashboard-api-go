/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200131Counts counts
type InlineResponse200131Counts struct {
	ByStatus *InlineResponse200131CountsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200131Counts instantiates a new InlineResponse200131Counts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200131Counts() *InlineResponse200131Counts {
	this := InlineResponse200131Counts{}
	return &this
}

// NewInlineResponse200131CountsWithDefaults instantiates a new InlineResponse200131Counts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200131CountsWithDefaults() *InlineResponse200131Counts {
	this := InlineResponse200131Counts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200131Counts) GetByStatus() InlineResponse200131CountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200131CountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131Counts) GetByStatusOk() (*InlineResponse200131CountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200131Counts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200131CountsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200131Counts) SetByStatus(v InlineResponse200131CountsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200131Counts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200131Counts struct {
	value *InlineResponse200131Counts
	isSet bool
}

func (v NullableInlineResponse200131Counts) Get() *InlineResponse200131Counts {
	return v.value
}

func (v *NullableInlineResponse200131Counts) Set(val *InlineResponse200131Counts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200131Counts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200131Counts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200131Counts(val *InlineResponse200131Counts) *NullableInlineResponse200131Counts {
	return &NullableInlineResponse200131Counts{value: val, isSet: true}
}

func (v NullableInlineResponse200131Counts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200131Counts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


