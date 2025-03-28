/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20059FiveGhzSettings Settings related to 5Ghz band.
type InlineResponse20059FiveGhzSettings struct {
	// Min bitrate (Mbps) of 2.4Ghz band.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Whether ax radio on 5Ghz band is on or off.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewInlineResponse20059FiveGhzSettings instantiates a new InlineResponse20059FiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059FiveGhzSettings() *InlineResponse20059FiveGhzSettings {
	this := InlineResponse20059FiveGhzSettings{}
	return &this
}

// NewInlineResponse20059FiveGhzSettingsWithDefaults instantiates a new InlineResponse20059FiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059FiveGhzSettingsWithDefaults() *InlineResponse20059FiveGhzSettings {
	this := InlineResponse20059FiveGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *InlineResponse20059FiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059FiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *InlineResponse20059FiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *InlineResponse20059FiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *InlineResponse20059FiveGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059FiveGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *InlineResponse20059FiveGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *InlineResponse20059FiveGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o InlineResponse20059FiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059FiveGhzSettings struct {
	value *InlineResponse20059FiveGhzSettings
	isSet bool
}

func (v NullableInlineResponse20059FiveGhzSettings) Get() *InlineResponse20059FiveGhzSettings {
	return v.value
}

func (v *NullableInlineResponse20059FiveGhzSettings) Set(val *InlineResponse20059FiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059FiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059FiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059FiveGhzSettings(val *InlineResponse20059FiveGhzSettings) *NullableInlineResponse20059FiveGhzSettings {
	return &NullableInlineResponse20059FiveGhzSettings{value: val, isSet: true}
}

func (v NullableInlineResponse20059FiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059FiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


