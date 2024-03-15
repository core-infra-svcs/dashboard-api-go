/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20031FiveGhzSettings Settings related to 5Ghz band.
type InlineResponse20031FiveGhzSettings struct {
	// Min bitrate (Mbps) of 2.4Ghz band.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Whether ax radio on 5Ghz band is on or off.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewInlineResponse20031FiveGhzSettings instantiates a new InlineResponse20031FiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20031FiveGhzSettings() *InlineResponse20031FiveGhzSettings {
	this := InlineResponse20031FiveGhzSettings{}
	return &this
}

// NewInlineResponse20031FiveGhzSettingsWithDefaults instantiates a new InlineResponse20031FiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20031FiveGhzSettingsWithDefaults() *InlineResponse20031FiveGhzSettings {
	this := InlineResponse20031FiveGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *InlineResponse20031FiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031FiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *InlineResponse20031FiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *InlineResponse20031FiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *InlineResponse20031FiveGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20031FiveGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *InlineResponse20031FiveGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *InlineResponse20031FiveGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o InlineResponse20031FiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20031FiveGhzSettings struct {
	value *InlineResponse20031FiveGhzSettings
	isSet bool
}

func (v NullableInlineResponse20031FiveGhzSettings) Get() *InlineResponse20031FiveGhzSettings {
	return v.value
}

func (v *NullableInlineResponse20031FiveGhzSettings) Set(val *InlineResponse20031FiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20031FiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20031FiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20031FiveGhzSettings(val *InlineResponse20031FiveGhzSettings) *NullableInlineResponse20031FiveGhzSettings {
	return &NullableInlineResponse20031FiveGhzSettings{value: val, isSet: true}
}

func (v NullableInlineResponse20031FiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20031FiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

