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

// InlineResponse200320CountsByStatusInactiveByMedia The inactive count data, indexed by media type (RJ45 or SFP)
type InlineResponse200320CountsByStatusInactiveByMedia struct {
	Rj45 *InlineResponse200320CountsByStatusInactiveByMediaRj45 `json:"rj45,omitempty"`
	Sfp *InlineResponse200320CountsByStatusInactiveByMediaSfp `json:"sfp,omitempty"`
}

// NewInlineResponse200320CountsByStatusInactiveByMedia instantiates a new InlineResponse200320CountsByStatusInactiveByMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200320CountsByStatusInactiveByMedia() *InlineResponse200320CountsByStatusInactiveByMedia {
	this := InlineResponse200320CountsByStatusInactiveByMedia{}
	return &this
}

// NewInlineResponse200320CountsByStatusInactiveByMediaWithDefaults instantiates a new InlineResponse200320CountsByStatusInactiveByMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200320CountsByStatusInactiveByMediaWithDefaults() *InlineResponse200320CountsByStatusInactiveByMedia {
	this := InlineResponse200320CountsByStatusInactiveByMedia{}
	return &this
}

// GetRj45 returns the Rj45 field value if set, zero value otherwise.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) GetRj45() InlineResponse200320CountsByStatusInactiveByMediaRj45 {
	if o == nil || isNil(o.Rj45) {
		var ret InlineResponse200320CountsByStatusInactiveByMediaRj45
		return ret
	}
	return *o.Rj45
}

// GetRj45Ok returns a tuple with the Rj45 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) GetRj45Ok() (*InlineResponse200320CountsByStatusInactiveByMediaRj45, bool) {
	if o == nil || isNil(o.Rj45) {
    return nil, false
	}
	return o.Rj45, true
}

// HasRj45 returns a boolean if a field has been set.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) HasRj45() bool {
	if o != nil && !isNil(o.Rj45) {
		return true
	}

	return false
}

// SetRj45 gets a reference to the given InlineResponse200320CountsByStatusInactiveByMediaRj45 and assigns it to the Rj45 field.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) SetRj45(v InlineResponse200320CountsByStatusInactiveByMediaRj45) {
	o.Rj45 = &v
}

// GetSfp returns the Sfp field value if set, zero value otherwise.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) GetSfp() InlineResponse200320CountsByStatusInactiveByMediaSfp {
	if o == nil || isNil(o.Sfp) {
		var ret InlineResponse200320CountsByStatusInactiveByMediaSfp
		return ret
	}
	return *o.Sfp
}

// GetSfpOk returns a tuple with the Sfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) GetSfpOk() (*InlineResponse200320CountsByStatusInactiveByMediaSfp, bool) {
	if o == nil || isNil(o.Sfp) {
    return nil, false
	}
	return o.Sfp, true
}

// HasSfp returns a boolean if a field has been set.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) HasSfp() bool {
	if o != nil && !isNil(o.Sfp) {
		return true
	}

	return false
}

// SetSfp gets a reference to the given InlineResponse200320CountsByStatusInactiveByMediaSfp and assigns it to the Sfp field.
func (o *InlineResponse200320CountsByStatusInactiveByMedia) SetSfp(v InlineResponse200320CountsByStatusInactiveByMediaSfp) {
	o.Sfp = &v
}

func (o InlineResponse200320CountsByStatusInactiveByMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rj45) {
		toSerialize["rj45"] = o.Rj45
	}
	if !isNil(o.Sfp) {
		toSerialize["sfp"] = o.Sfp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200320CountsByStatusInactiveByMedia struct {
	value *InlineResponse200320CountsByStatusInactiveByMedia
	isSet bool
}

func (v NullableInlineResponse200320CountsByStatusInactiveByMedia) Get() *InlineResponse200320CountsByStatusInactiveByMedia {
	return v.value
}

func (v *NullableInlineResponse200320CountsByStatusInactiveByMedia) Set(val *InlineResponse200320CountsByStatusInactiveByMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200320CountsByStatusInactiveByMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200320CountsByStatusInactiveByMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200320CountsByStatusInactiveByMedia(val *InlineResponse200320CountsByStatusInactiveByMedia) *NullableInlineResponse200320CountsByStatusInactiveByMedia {
	return &NullableInlineResponse200320CountsByStatusInactiveByMedia{value: val, isSet: true}
}

func (v NullableInlineResponse200320CountsByStatusInactiveByMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200320CountsByStatusInactiveByMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


