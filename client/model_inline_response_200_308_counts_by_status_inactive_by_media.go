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

// InlineResponse200308CountsByStatusInactiveByMedia The inactive count data, indexed by media type (RJ45 or SFP)
type InlineResponse200308CountsByStatusInactiveByMedia struct {
	Rj45 *InlineResponse200308CountsByStatusInactiveByMediaRj45 `json:"rj45,omitempty"`
	Sfp *InlineResponse200308CountsByStatusInactiveByMediaSfp `json:"sfp,omitempty"`
}

// NewInlineResponse200308CountsByStatusInactiveByMedia instantiates a new InlineResponse200308CountsByStatusInactiveByMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308CountsByStatusInactiveByMedia() *InlineResponse200308CountsByStatusInactiveByMedia {
	this := InlineResponse200308CountsByStatusInactiveByMedia{}
	return &this
}

// NewInlineResponse200308CountsByStatusInactiveByMediaWithDefaults instantiates a new InlineResponse200308CountsByStatusInactiveByMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308CountsByStatusInactiveByMediaWithDefaults() *InlineResponse200308CountsByStatusInactiveByMedia {
	this := InlineResponse200308CountsByStatusInactiveByMedia{}
	return &this
}

// GetRj45 returns the Rj45 field value if set, zero value otherwise.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) GetRj45() InlineResponse200308CountsByStatusInactiveByMediaRj45 {
	if o == nil || isNil(o.Rj45) {
		var ret InlineResponse200308CountsByStatusInactiveByMediaRj45
		return ret
	}
	return *o.Rj45
}

// GetRj45Ok returns a tuple with the Rj45 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) GetRj45Ok() (*InlineResponse200308CountsByStatusInactiveByMediaRj45, bool) {
	if o == nil || isNil(o.Rj45) {
    return nil, false
	}
	return o.Rj45, true
}

// HasRj45 returns a boolean if a field has been set.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) HasRj45() bool {
	if o != nil && !isNil(o.Rj45) {
		return true
	}

	return false
}

// SetRj45 gets a reference to the given InlineResponse200308CountsByStatusInactiveByMediaRj45 and assigns it to the Rj45 field.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) SetRj45(v InlineResponse200308CountsByStatusInactiveByMediaRj45) {
	o.Rj45 = &v
}

// GetSfp returns the Sfp field value if set, zero value otherwise.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) GetSfp() InlineResponse200308CountsByStatusInactiveByMediaSfp {
	if o == nil || isNil(o.Sfp) {
		var ret InlineResponse200308CountsByStatusInactiveByMediaSfp
		return ret
	}
	return *o.Sfp
}

// GetSfpOk returns a tuple with the Sfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) GetSfpOk() (*InlineResponse200308CountsByStatusInactiveByMediaSfp, bool) {
	if o == nil || isNil(o.Sfp) {
    return nil, false
	}
	return o.Sfp, true
}

// HasSfp returns a boolean if a field has been set.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) HasSfp() bool {
	if o != nil && !isNil(o.Sfp) {
		return true
	}

	return false
}

// SetSfp gets a reference to the given InlineResponse200308CountsByStatusInactiveByMediaSfp and assigns it to the Sfp field.
func (o *InlineResponse200308CountsByStatusInactiveByMedia) SetSfp(v InlineResponse200308CountsByStatusInactiveByMediaSfp) {
	o.Sfp = &v
}

func (o InlineResponse200308CountsByStatusInactiveByMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rj45) {
		toSerialize["rj45"] = o.Rj45
	}
	if !isNil(o.Sfp) {
		toSerialize["sfp"] = o.Sfp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308CountsByStatusInactiveByMedia struct {
	value *InlineResponse200308CountsByStatusInactiveByMedia
	isSet bool
}

func (v NullableInlineResponse200308CountsByStatusInactiveByMedia) Get() *InlineResponse200308CountsByStatusInactiveByMedia {
	return v.value
}

func (v *NullableInlineResponse200308CountsByStatusInactiveByMedia) Set(val *InlineResponse200308CountsByStatusInactiveByMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308CountsByStatusInactiveByMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308CountsByStatusInactiveByMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308CountsByStatusInactiveByMedia(val *InlineResponse200308CountsByStatusInactiveByMedia) *NullableInlineResponse200308CountsByStatusInactiveByMedia {
	return &NullableInlineResponse200308CountsByStatusInactiveByMedia{value: val, isSet: true}
}

func (v NullableInlineResponse200308CountsByStatusInactiveByMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308CountsByStatusInactiveByMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

