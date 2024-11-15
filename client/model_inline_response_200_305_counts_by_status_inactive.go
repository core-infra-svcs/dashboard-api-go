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

// InlineResponse200305CountsByStatusInactive The count data for inactive ports
type InlineResponse200305CountsByStatusInactive struct {
	// The total number of inactive ports
	Total *int32 `json:"total,omitempty"`
	ByMedia *InlineResponse200305CountsByStatusInactiveByMedia `json:"byMedia,omitempty"`
}

// NewInlineResponse200305CountsByStatusInactive instantiates a new InlineResponse200305CountsByStatusInactive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200305CountsByStatusInactive() *InlineResponse200305CountsByStatusInactive {
	this := InlineResponse200305CountsByStatusInactive{}
	return &this
}

// NewInlineResponse200305CountsByStatusInactiveWithDefaults instantiates a new InlineResponse200305CountsByStatusInactive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200305CountsByStatusInactiveWithDefaults() *InlineResponse200305CountsByStatusInactive {
	this := InlineResponse200305CountsByStatusInactive{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatusInactive) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatusInactive) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatusInactive) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200305CountsByStatusInactive) SetTotal(v int32) {
	o.Total = &v
}

// GetByMedia returns the ByMedia field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatusInactive) GetByMedia() InlineResponse200305CountsByStatusInactiveByMedia {
	if o == nil || isNil(o.ByMedia) {
		var ret InlineResponse200305CountsByStatusInactiveByMedia
		return ret
	}
	return *o.ByMedia
}

// GetByMediaOk returns a tuple with the ByMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatusInactive) GetByMediaOk() (*InlineResponse200305CountsByStatusInactiveByMedia, bool) {
	if o == nil || isNil(o.ByMedia) {
    return nil, false
	}
	return o.ByMedia, true
}

// HasByMedia returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatusInactive) HasByMedia() bool {
	if o != nil && !isNil(o.ByMedia) {
		return true
	}

	return false
}

// SetByMedia gets a reference to the given InlineResponse200305CountsByStatusInactiveByMedia and assigns it to the ByMedia field.
func (o *InlineResponse200305CountsByStatusInactive) SetByMedia(v InlineResponse200305CountsByStatusInactiveByMedia) {
	o.ByMedia = &v
}

func (o InlineResponse200305CountsByStatusInactive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByMedia) {
		toSerialize["byMedia"] = o.ByMedia
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200305CountsByStatusInactive struct {
	value *InlineResponse200305CountsByStatusInactive
	isSet bool
}

func (v NullableInlineResponse200305CountsByStatusInactive) Get() *InlineResponse200305CountsByStatusInactive {
	return v.value
}

func (v *NullableInlineResponse200305CountsByStatusInactive) Set(val *InlineResponse200305CountsByStatusInactive) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200305CountsByStatusInactive) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200305CountsByStatusInactive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200305CountsByStatusInactive(val *InlineResponse200305CountsByStatusInactive) *NullableInlineResponse200305CountsByStatusInactive {
	return &NullableInlineResponse200305CountsByStatusInactive{value: val, isSet: true}
}

func (v NullableInlineResponse200305CountsByStatusInactive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200305CountsByStatusInactive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


