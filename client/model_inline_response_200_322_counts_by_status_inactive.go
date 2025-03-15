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

// InlineResponse200322CountsByStatusInactive The count data for inactive ports
type InlineResponse200322CountsByStatusInactive struct {
	// The total number of inactive ports
	Total *int32 `json:"total,omitempty"`
	ByMedia *InlineResponse200322CountsByStatusInactiveByMedia `json:"byMedia,omitempty"`
}

// NewInlineResponse200322CountsByStatusInactive instantiates a new InlineResponse200322CountsByStatusInactive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200322CountsByStatusInactive() *InlineResponse200322CountsByStatusInactive {
	this := InlineResponse200322CountsByStatusInactive{}
	return &this
}

// NewInlineResponse200322CountsByStatusInactiveWithDefaults instantiates a new InlineResponse200322CountsByStatusInactive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200322CountsByStatusInactiveWithDefaults() *InlineResponse200322CountsByStatusInactive {
	this := InlineResponse200322CountsByStatusInactive{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200322CountsByStatusInactive) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322CountsByStatusInactive) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200322CountsByStatusInactive) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200322CountsByStatusInactive) SetTotal(v int32) {
	o.Total = &v
}

// GetByMedia returns the ByMedia field value if set, zero value otherwise.
func (o *InlineResponse200322CountsByStatusInactive) GetByMedia() InlineResponse200322CountsByStatusInactiveByMedia {
	if o == nil || isNil(o.ByMedia) {
		var ret InlineResponse200322CountsByStatusInactiveByMedia
		return ret
	}
	return *o.ByMedia
}

// GetByMediaOk returns a tuple with the ByMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200322CountsByStatusInactive) GetByMediaOk() (*InlineResponse200322CountsByStatusInactiveByMedia, bool) {
	if o == nil || isNil(o.ByMedia) {
    return nil, false
	}
	return o.ByMedia, true
}

// HasByMedia returns a boolean if a field has been set.
func (o *InlineResponse200322CountsByStatusInactive) HasByMedia() bool {
	if o != nil && !isNil(o.ByMedia) {
		return true
	}

	return false
}

// SetByMedia gets a reference to the given InlineResponse200322CountsByStatusInactiveByMedia and assigns it to the ByMedia field.
func (o *InlineResponse200322CountsByStatusInactive) SetByMedia(v InlineResponse200322CountsByStatusInactiveByMedia) {
	o.ByMedia = &v
}

func (o InlineResponse200322CountsByStatusInactive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByMedia) {
		toSerialize["byMedia"] = o.ByMedia
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200322CountsByStatusInactive struct {
	value *InlineResponse200322CountsByStatusInactive
	isSet bool
}

func (v NullableInlineResponse200322CountsByStatusInactive) Get() *InlineResponse200322CountsByStatusInactive {
	return v.value
}

func (v *NullableInlineResponse200322CountsByStatusInactive) Set(val *InlineResponse200322CountsByStatusInactive) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200322CountsByStatusInactive) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200322CountsByStatusInactive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200322CountsByStatusInactive(val *InlineResponse200322CountsByStatusInactive) *NullableInlineResponse200322CountsByStatusInactive {
	return &NullableInlineResponse200322CountsByStatusInactive{value: val, isSet: true}
}

func (v NullableInlineResponse200322CountsByStatusInactive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200322CountsByStatusInactive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


