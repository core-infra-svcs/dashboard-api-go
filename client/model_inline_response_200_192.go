/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200192 struct for InlineResponse200192
type InlineResponse200192 struct {
	// If true, the SSID outage schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	Ranges []InlineResponse200192Ranges `json:"ranges,omitempty"`
	// List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
	RangesInSeconds []InlineResponse200192RangesInSeconds `json:"rangesInSeconds,omitempty"`
}

// NewInlineResponse200192 instantiates a new InlineResponse200192 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200192() *InlineResponse200192 {
	this := InlineResponse200192{}
	return &this
}

// NewInlineResponse200192WithDefaults instantiates a new InlineResponse200192 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200192WithDefaults() *InlineResponse200192 {
	this := InlineResponse200192{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200192) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200192) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200192) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200192) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *InlineResponse200192) GetRanges() []InlineResponse200192Ranges {
	if o == nil || isNil(o.Ranges) {
		var ret []InlineResponse200192Ranges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200192) GetRangesOk() ([]InlineResponse200192Ranges, bool) {
	if o == nil || isNil(o.Ranges) {
    return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *InlineResponse200192) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []InlineResponse200192Ranges and assigns it to the Ranges field.
func (o *InlineResponse200192) SetRanges(v []InlineResponse200192Ranges) {
	o.Ranges = v
}

// GetRangesInSeconds returns the RangesInSeconds field value if set, zero value otherwise.
func (o *InlineResponse200192) GetRangesInSeconds() []InlineResponse200192RangesInSeconds {
	if o == nil || isNil(o.RangesInSeconds) {
		var ret []InlineResponse200192RangesInSeconds
		return ret
	}
	return o.RangesInSeconds
}

// GetRangesInSecondsOk returns a tuple with the RangesInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200192) GetRangesInSecondsOk() ([]InlineResponse200192RangesInSeconds, bool) {
	if o == nil || isNil(o.RangesInSeconds) {
    return nil, false
	}
	return o.RangesInSeconds, true
}

// HasRangesInSeconds returns a boolean if a field has been set.
func (o *InlineResponse200192) HasRangesInSeconds() bool {
	if o != nil && !isNil(o.RangesInSeconds) {
		return true
	}

	return false
}

// SetRangesInSeconds gets a reference to the given []InlineResponse200192RangesInSeconds and assigns it to the RangesInSeconds field.
func (o *InlineResponse200192) SetRangesInSeconds(v []InlineResponse200192RangesInSeconds) {
	o.RangesInSeconds = v
}

func (o InlineResponse200192) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200192 struct {
	value *InlineResponse200192
	isSet bool
}

func (v NullableInlineResponse200192) Get() *InlineResponse200192 {
	return v.value
}

func (v *NullableInlineResponse200192) Set(val *InlineResponse200192) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200192) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200192) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200192(val *InlineResponse200192) *NullableInlineResponse200192 {
	return &NullableInlineResponse200192{value: val, isSet: true}
}

func (v NullableInlineResponse200192) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200192) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


