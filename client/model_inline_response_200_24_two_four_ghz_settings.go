/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20024TwoFourGhzSettings Settings related to 2.4Ghz band.
type InlineResponse20024TwoFourGhzSettings struct {
	// Min bitrate (Mbps) of 2.4Ghz band.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Whether ax radio on 2.4Ghz band is on or off.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewInlineResponse20024TwoFourGhzSettings instantiates a new InlineResponse20024TwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20024TwoFourGhzSettings() *InlineResponse20024TwoFourGhzSettings {
	this := InlineResponse20024TwoFourGhzSettings{}
	return &this
}

// NewInlineResponse20024TwoFourGhzSettingsWithDefaults instantiates a new InlineResponse20024TwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20024TwoFourGhzSettingsWithDefaults() *InlineResponse20024TwoFourGhzSettings {
	this := InlineResponse20024TwoFourGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *InlineResponse20024TwoFourGhzSettings) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024TwoFourGhzSettings) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *InlineResponse20024TwoFourGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *InlineResponse20024TwoFourGhzSettings) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *InlineResponse20024TwoFourGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024TwoFourGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *InlineResponse20024TwoFourGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *InlineResponse20024TwoFourGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o InlineResponse20024TwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20024TwoFourGhzSettings struct {
	value *InlineResponse20024TwoFourGhzSettings
	isSet bool
}

func (v NullableInlineResponse20024TwoFourGhzSettings) Get() *InlineResponse20024TwoFourGhzSettings {
	return v.value
}

func (v *NullableInlineResponse20024TwoFourGhzSettings) Set(val *InlineResponse20024TwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20024TwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20024TwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20024TwoFourGhzSettings(val *InlineResponse20024TwoFourGhzSettings) *NullableInlineResponse20024TwoFourGhzSettings {
	return &NullableInlineResponse20024TwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableInlineResponse20024TwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20024TwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


