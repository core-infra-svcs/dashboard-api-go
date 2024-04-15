/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200188 struct for InlineResponse200188
type InlineResponse200188 struct {
	// If true, the SSID outage schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	Ranges []InlineResponse200188Ranges `json:"ranges,omitempty"`
	// List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
	RangesInSeconds []InlineResponse200188RangesInSeconds `json:"rangesInSeconds,omitempty"`
}

// NewInlineResponse200188 instantiates a new InlineResponse200188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200188() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// NewInlineResponse200188WithDefaults instantiates a new InlineResponse200188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200188WithDefaults() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200188) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200188) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200188) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *InlineResponse200188) GetRanges() []InlineResponse200188Ranges {
	if o == nil || isNil(o.Ranges) {
		var ret []InlineResponse200188Ranges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetRangesOk() ([]InlineResponse200188Ranges, bool) {
	if o == nil || isNil(o.Ranges) {
    return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *InlineResponse200188) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []InlineResponse200188Ranges and assigns it to the Ranges field.
func (o *InlineResponse200188) SetRanges(v []InlineResponse200188Ranges) {
	o.Ranges = v
}

// GetRangesInSeconds returns the RangesInSeconds field value if set, zero value otherwise.
func (o *InlineResponse200188) GetRangesInSeconds() []InlineResponse200188RangesInSeconds {
	if o == nil || isNil(o.RangesInSeconds) {
		var ret []InlineResponse200188RangesInSeconds
		return ret
	}
	return o.RangesInSeconds
}

// GetRangesInSecondsOk returns a tuple with the RangesInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetRangesInSecondsOk() ([]InlineResponse200188RangesInSeconds, bool) {
	if o == nil || isNil(o.RangesInSeconds) {
    return nil, false
	}
	return o.RangesInSeconds, true
}

// HasRangesInSeconds returns a boolean if a field has been set.
func (o *InlineResponse200188) HasRangesInSeconds() bool {
	if o != nil && !isNil(o.RangesInSeconds) {
		return true
	}

	return false
}

// SetRangesInSeconds gets a reference to the given []InlineResponse200188RangesInSeconds and assigns it to the RangesInSeconds field.
func (o *InlineResponse200188) SetRangesInSeconds(v []InlineResponse200188RangesInSeconds) {
	o.RangesInSeconds = v
}

func (o InlineResponse200188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	if !isNil(o.RangesInSeconds) {
		toSerialize["rangesInSeconds"] = o.RangesInSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200188 struct {
	value *InlineResponse200188
	isSet bool
}

func (v NullableInlineResponse200188) Get() *InlineResponse200188 {
	return v.value
}

func (v *NullableInlineResponse200188) Set(val *InlineResponse200188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200188(val *InlineResponse200188) *NullableInlineResponse200188 {
	return &NullableInlineResponse200188{value: val, isSet: true}
}

func (v NullableInlineResponse200188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


