/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200291CountsByStatusInactive The count data for inactive ports
type InlineResponse200291CountsByStatusInactive struct {
	// The total number of inactive ports
	Total *int32 `json:"total,omitempty"`
	ByMedia *InlineResponse200291CountsByStatusInactiveByMedia `json:"byMedia,omitempty"`
}

// NewInlineResponse200291CountsByStatusInactive instantiates a new InlineResponse200291CountsByStatusInactive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200291CountsByStatusInactive() *InlineResponse200291CountsByStatusInactive {
	this := InlineResponse200291CountsByStatusInactive{}
	return &this
}

// NewInlineResponse200291CountsByStatusInactiveWithDefaults instantiates a new InlineResponse200291CountsByStatusInactive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200291CountsByStatusInactiveWithDefaults() *InlineResponse200291CountsByStatusInactive {
	this := InlineResponse200291CountsByStatusInactive{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200291CountsByStatusInactive) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291CountsByStatusInactive) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200291CountsByStatusInactive) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200291CountsByStatusInactive) SetTotal(v int32) {
	o.Total = &v
}

// GetByMedia returns the ByMedia field value if set, zero value otherwise.
func (o *InlineResponse200291CountsByStatusInactive) GetByMedia() InlineResponse200291CountsByStatusInactiveByMedia {
	if o == nil || isNil(o.ByMedia) {
		var ret InlineResponse200291CountsByStatusInactiveByMedia
		return ret
	}
	return *o.ByMedia
}

// GetByMediaOk returns a tuple with the ByMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200291CountsByStatusInactive) GetByMediaOk() (*InlineResponse200291CountsByStatusInactiveByMedia, bool) {
	if o == nil || isNil(o.ByMedia) {
    return nil, false
	}
	return o.ByMedia, true
}

// HasByMedia returns a boolean if a field has been set.
func (o *InlineResponse200291CountsByStatusInactive) HasByMedia() bool {
	if o != nil && !isNil(o.ByMedia) {
		return true
	}

	return false
}

// SetByMedia gets a reference to the given InlineResponse200291CountsByStatusInactiveByMedia and assigns it to the ByMedia field.
func (o *InlineResponse200291CountsByStatusInactive) SetByMedia(v InlineResponse200291CountsByStatusInactiveByMedia) {
	o.ByMedia = &v
}

func (o InlineResponse200291CountsByStatusInactive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByMedia) {
		toSerialize["byMedia"] = o.ByMedia
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200291CountsByStatusInactive struct {
	value *InlineResponse200291CountsByStatusInactive
	isSet bool
}

func (v NullableInlineResponse200291CountsByStatusInactive) Get() *InlineResponse200291CountsByStatusInactive {
	return v.value
}

func (v *NullableInlineResponse200291CountsByStatusInactive) Set(val *InlineResponse200291CountsByStatusInactive) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200291CountsByStatusInactive) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200291CountsByStatusInactive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200291CountsByStatusInactive(val *InlineResponse200291CountsByStatusInactive) *NullableInlineResponse200291CountsByStatusInactive {
	return &NullableInlineResponse200291CountsByStatusInactive{value: val, isSet: true}
}

func (v NullableInlineResponse200291CountsByStatusInactive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200291CountsByStatusInactive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


