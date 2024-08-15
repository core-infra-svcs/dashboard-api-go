/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject193 struct for InlineObject193
type InlineObject193 struct {
	// If true, the SSID outage schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	Ranges []InlineResponse200193Ranges `json:"ranges,omitempty"`
	// List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
	RangesInSeconds []InlineResponse200193RangesInSeconds `json:"rangesInSeconds,omitempty"`
}

// NewInlineObject193 instantiates a new InlineObject193 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject193() *InlineObject193 {
	this := InlineObject193{}
	return &this
}

// NewInlineObject193WithDefaults instantiates a new InlineObject193 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject193WithDefaults() *InlineObject193 {
	this := InlineObject193{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject193) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject193) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject193) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject193) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *InlineObject193) GetRanges() []InlineResponse200193Ranges {
	if o == nil || isNil(o.Ranges) {
		var ret []InlineResponse200193Ranges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject193) GetRangesOk() ([]InlineResponse200193Ranges, bool) {
	if o == nil || isNil(o.Ranges) {
    return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *InlineObject193) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []InlineResponse200193Ranges and assigns it to the Ranges field.
func (o *InlineObject193) SetRanges(v []InlineResponse200193Ranges) {
	o.Ranges = v
}

// GetRangesInSeconds returns the RangesInSeconds field value if set, zero value otherwise.
func (o *InlineObject193) GetRangesInSeconds() []InlineResponse200193RangesInSeconds {
	if o == nil || isNil(o.RangesInSeconds) {
		var ret []InlineResponse200193RangesInSeconds
		return ret
	}
	return o.RangesInSeconds
}

// GetRangesInSecondsOk returns a tuple with the RangesInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject193) GetRangesInSecondsOk() ([]InlineResponse200193RangesInSeconds, bool) {
	if o == nil || isNil(o.RangesInSeconds) {
    return nil, false
	}
	return o.RangesInSeconds, true
}

// HasRangesInSeconds returns a boolean if a field has been set.
func (o *InlineObject193) HasRangesInSeconds() bool {
	if o != nil && !isNil(o.RangesInSeconds) {
		return true
	}

	return false
}

// SetRangesInSeconds gets a reference to the given []InlineResponse200193RangesInSeconds and assigns it to the RangesInSeconds field.
func (o *InlineObject193) SetRangesInSeconds(v []InlineResponse200193RangesInSeconds) {
	o.RangesInSeconds = v
}

func (o InlineObject193) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject193 struct {
	value *InlineObject193
	isSet bool
}

func (v NullableInlineObject193) Get() *InlineObject193 {
	return v.value
}

func (v *NullableInlineObject193) Set(val *InlineObject193) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject193) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject193) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject193(val *InlineObject193) *NullableInlineObject193 {
	return &NullableInlineObject193{value: val, isSet: true}
}

func (v NullableInlineObject193) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject193) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


