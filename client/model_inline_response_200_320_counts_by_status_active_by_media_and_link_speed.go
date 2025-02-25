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

// InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed The active count data, indexed by media type (RJ45 or SFP)
type InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed struct {
	Rj45 *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedRj45 `json:"rj45,omitempty"`
	Sfp *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedSfp `json:"sfp,omitempty"`
}

// NewInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed instantiates a new InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed() *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed {
	this := InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed{}
	return &this
}

// NewInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedWithDefaults instantiates a new InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedWithDefaults() *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed {
	this := InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed{}
	return &this
}

// GetRj45 returns the Rj45 field value if set, zero value otherwise.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) GetRj45() InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedRj45 {
	if o == nil || isNil(o.Rj45) {
		var ret InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedRj45
		return ret
	}
	return *o.Rj45
}

// GetRj45Ok returns a tuple with the Rj45 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) GetRj45Ok() (*InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedRj45, bool) {
	if o == nil || isNil(o.Rj45) {
    return nil, false
	}
	return o.Rj45, true
}

// HasRj45 returns a boolean if a field has been set.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) HasRj45() bool {
	if o != nil && !isNil(o.Rj45) {
		return true
	}

	return false
}

// SetRj45 gets a reference to the given InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedRj45 and assigns it to the Rj45 field.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) SetRj45(v InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedRj45) {
	o.Rj45 = &v
}

// GetSfp returns the Sfp field value if set, zero value otherwise.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) GetSfp() InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedSfp {
	if o == nil || isNil(o.Sfp) {
		var ret InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedSfp
		return ret
	}
	return *o.Sfp
}

// GetSfpOk returns a tuple with the Sfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) GetSfpOk() (*InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedSfp, bool) {
	if o == nil || isNil(o.Sfp) {
    return nil, false
	}
	return o.Sfp, true
}

// HasSfp returns a boolean if a field has been set.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) HasSfp() bool {
	if o != nil && !isNil(o.Sfp) {
		return true
	}

	return false
}

// SetSfp gets a reference to the given InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedSfp and assigns it to the Sfp field.
func (o *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) SetSfp(v InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeedSfp) {
	o.Sfp = &v
}

func (o InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rj45) {
		toSerialize["rj45"] = o.Rj45
	}
	if !isNil(o.Sfp) {
		toSerialize["sfp"] = o.Sfp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed struct {
	value *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed
	isSet bool
}

func (v NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) Get() *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed {
	return v.value
}

func (v *NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) Set(val *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed(val *InlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) *NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed {
	return &NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed{value: val, isSet: true}
}

func (v NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200320CountsByStatusActiveByMediaAndLinkSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


