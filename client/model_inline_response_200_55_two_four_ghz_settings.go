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

// InlineResponse20055TwoFourGhzSettings Settings related to 2.4Ghz band.
type InlineResponse20055TwoFourGhzSettings struct {
	// Min bitrate (Mbps) of 2.4Ghz band.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Whether ax radio on 2.4Ghz band is on or off.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewInlineResponse20055TwoFourGhzSettings instantiates a new InlineResponse20055TwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20055TwoFourGhzSettings() *InlineResponse20055TwoFourGhzSettings {
	this := InlineResponse20055TwoFourGhzSettings{}
	return &this
}

// NewInlineResponse20055TwoFourGhzSettingsWithDefaults instantiates a new InlineResponse20055TwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20055TwoFourGhzSettingsWithDefaults() *InlineResponse20055TwoFourGhzSettings {
	this := InlineResponse20055TwoFourGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *InlineResponse20055TwoFourGhzSettings) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055TwoFourGhzSettings) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *InlineResponse20055TwoFourGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *InlineResponse20055TwoFourGhzSettings) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *InlineResponse20055TwoFourGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055TwoFourGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *InlineResponse20055TwoFourGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *InlineResponse20055TwoFourGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o InlineResponse20055TwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20055TwoFourGhzSettings struct {
	value *InlineResponse20055TwoFourGhzSettings
	isSet bool
}

func (v NullableInlineResponse20055TwoFourGhzSettings) Get() *InlineResponse20055TwoFourGhzSettings {
	return v.value
}

func (v *NullableInlineResponse20055TwoFourGhzSettings) Set(val *InlineResponse20055TwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20055TwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20055TwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20055TwoFourGhzSettings(val *InlineResponse20055TwoFourGhzSettings) *NullableInlineResponse20055TwoFourGhzSettings {
	return &NullableInlineResponse20055TwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableInlineResponse20055TwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20055TwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


