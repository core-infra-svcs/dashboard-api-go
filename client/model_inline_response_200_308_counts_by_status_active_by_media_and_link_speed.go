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

// InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed The active count data, indexed by media type (RJ45 or SFP)
type InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed struct {
	Rj45 *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedRj45 `json:"rj45,omitempty"`
	Sfp *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedSfp `json:"sfp,omitempty"`
}

// NewInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed instantiates a new InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed() *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed {
	this := InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed{}
	return &this
}

// NewInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedWithDefaults instantiates a new InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedWithDefaults() *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed {
	this := InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed{}
	return &this
}

// GetRj45 returns the Rj45 field value if set, zero value otherwise.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) GetRj45() InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedRj45 {
	if o == nil || isNil(o.Rj45) {
		var ret InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedRj45
		return ret
	}
	return *o.Rj45
}

// GetRj45Ok returns a tuple with the Rj45 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) GetRj45Ok() (*InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedRj45, bool) {
	if o == nil || isNil(o.Rj45) {
    return nil, false
	}
	return o.Rj45, true
}

// HasRj45 returns a boolean if a field has been set.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) HasRj45() bool {
	if o != nil && !isNil(o.Rj45) {
		return true
	}

	return false
}

// SetRj45 gets a reference to the given InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedRj45 and assigns it to the Rj45 field.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) SetRj45(v InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedRj45) {
	o.Rj45 = &v
}

// GetSfp returns the Sfp field value if set, zero value otherwise.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) GetSfp() InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedSfp {
	if o == nil || isNil(o.Sfp) {
		var ret InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedSfp
		return ret
	}
	return *o.Sfp
}

// GetSfpOk returns a tuple with the Sfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) GetSfpOk() (*InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedSfp, bool) {
	if o == nil || isNil(o.Sfp) {
    return nil, false
	}
	return o.Sfp, true
}

// HasSfp returns a boolean if a field has been set.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) HasSfp() bool {
	if o != nil && !isNil(o.Sfp) {
		return true
	}

	return false
}

// SetSfp gets a reference to the given InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedSfp and assigns it to the Sfp field.
func (o *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) SetSfp(v InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeedSfp) {
	o.Sfp = &v
}

func (o InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rj45) {
		toSerialize["rj45"] = o.Rj45
	}
	if !isNil(o.Sfp) {
		toSerialize["sfp"] = o.Sfp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed struct {
	value *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed
	isSet bool
}

func (v NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) Get() *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed {
	return v.value
}

func (v *NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) Set(val *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed(val *InlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) *NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed {
	return &NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed{value: val, isSet: true}
}

func (v NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308CountsByStatusActiveByMediaAndLinkSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


