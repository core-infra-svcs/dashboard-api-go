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

// InlineResponse200305CountsByStatusActive The count data for active ports
type InlineResponse200305CountsByStatusActive struct {
	// The total number of active ports
	Total *int32 `json:"total,omitempty"`
	ByMediaAndLinkSpeed *InlineResponse200305CountsByStatusActiveByMediaAndLinkSpeed `json:"byMediaAndLinkSpeed,omitempty"`
}

// NewInlineResponse200305CountsByStatusActive instantiates a new InlineResponse200305CountsByStatusActive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200305CountsByStatusActive() *InlineResponse200305CountsByStatusActive {
	this := InlineResponse200305CountsByStatusActive{}
	return &this
}

// NewInlineResponse200305CountsByStatusActiveWithDefaults instantiates a new InlineResponse200305CountsByStatusActive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200305CountsByStatusActiveWithDefaults() *InlineResponse200305CountsByStatusActive {
	this := InlineResponse200305CountsByStatusActive{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatusActive) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatusActive) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatusActive) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200305CountsByStatusActive) SetTotal(v int32) {
	o.Total = &v
}

// GetByMediaAndLinkSpeed returns the ByMediaAndLinkSpeed field value if set, zero value otherwise.
func (o *InlineResponse200305CountsByStatusActive) GetByMediaAndLinkSpeed() InlineResponse200305CountsByStatusActiveByMediaAndLinkSpeed {
	if o == nil || isNil(o.ByMediaAndLinkSpeed) {
		var ret InlineResponse200305CountsByStatusActiveByMediaAndLinkSpeed
		return ret
	}
	return *o.ByMediaAndLinkSpeed
}

// GetByMediaAndLinkSpeedOk returns a tuple with the ByMediaAndLinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200305CountsByStatusActive) GetByMediaAndLinkSpeedOk() (*InlineResponse200305CountsByStatusActiveByMediaAndLinkSpeed, bool) {
	if o == nil || isNil(o.ByMediaAndLinkSpeed) {
    return nil, false
	}
	return o.ByMediaAndLinkSpeed, true
}

// HasByMediaAndLinkSpeed returns a boolean if a field has been set.
func (o *InlineResponse200305CountsByStatusActive) HasByMediaAndLinkSpeed() bool {
	if o != nil && !isNil(o.ByMediaAndLinkSpeed) {
		return true
	}

	return false
}

// SetByMediaAndLinkSpeed gets a reference to the given InlineResponse200305CountsByStatusActiveByMediaAndLinkSpeed and assigns it to the ByMediaAndLinkSpeed field.
func (o *InlineResponse200305CountsByStatusActive) SetByMediaAndLinkSpeed(v InlineResponse200305CountsByStatusActiveByMediaAndLinkSpeed) {
	o.ByMediaAndLinkSpeed = &v
}

func (o InlineResponse200305CountsByStatusActive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByMediaAndLinkSpeed) {
		toSerialize["byMediaAndLinkSpeed"] = o.ByMediaAndLinkSpeed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200305CountsByStatusActive struct {
	value *InlineResponse200305CountsByStatusActive
	isSet bool
}

func (v NullableInlineResponse200305CountsByStatusActive) Get() *InlineResponse200305CountsByStatusActive {
	return v.value
}

func (v *NullableInlineResponse200305CountsByStatusActive) Set(val *InlineResponse200305CountsByStatusActive) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200305CountsByStatusActive) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200305CountsByStatusActive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200305CountsByStatusActive(val *InlineResponse200305CountsByStatusActive) *NullableInlineResponse200305CountsByStatusActive {
	return &NullableInlineResponse200305CountsByStatusActive{value: val, isSet: true}
}

func (v NullableInlineResponse200305CountsByStatusActive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200305CountsByStatusActive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


