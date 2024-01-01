/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject180 struct for InlineObject180
type InlineObject180 struct {
	// If true, the SSID outage schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	Ranges []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges `json:"ranges,omitempty"`
	// List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
	RangesInSeconds []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds `json:"rangesInSeconds,omitempty"`
}

// NewInlineObject180 instantiates a new InlineObject180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject180() *InlineObject180 {
	this := InlineObject180{}
	return &this
}

// NewInlineObject180WithDefaults instantiates a new InlineObject180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject180WithDefaults() *InlineObject180 {
	this := InlineObject180{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject180) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject180) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject180) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *InlineObject180) GetRanges() []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges {
	if o == nil || isNil(o.Ranges) {
		var ret []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetRangesOk() ([]NetworksNetworkIdWirelessSsidsNumberSchedulesRanges, bool) {
	if o == nil || isNil(o.Ranges) {
    return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *InlineObject180) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges and assigns it to the Ranges field.
func (o *InlineObject180) SetRanges(v []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) {
	o.Ranges = v
}

// GetRangesInSeconds returns the RangesInSeconds field value if set, zero value otherwise.
func (o *InlineObject180) GetRangesInSeconds() []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds {
	if o == nil || isNil(o.RangesInSeconds) {
		var ret []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds
		return ret
	}
	return o.RangesInSeconds
}

// GetRangesInSecondsOk returns a tuple with the RangesInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject180) GetRangesInSecondsOk() ([]NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds, bool) {
	if o == nil || isNil(o.RangesInSeconds) {
    return nil, false
	}
	return o.RangesInSeconds, true
}

// HasRangesInSeconds returns a boolean if a field has been set.
func (o *InlineObject180) HasRangesInSeconds() bool {
	if o != nil && !isNil(o.RangesInSeconds) {
		return true
	}

	return false
}

// SetRangesInSeconds gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds and assigns it to the RangesInSeconds field.
func (o *InlineObject180) SetRangesInSeconds(v []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) {
	o.RangesInSeconds = v
}

func (o InlineObject180) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject180 struct {
	value *InlineObject180
	isSet bool
}

func (v NullableInlineObject180) Get() *InlineObject180 {
	return v.value
}

func (v *NullableInlineObject180) Set(val *InlineObject180) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject180) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject180) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject180(val *InlineObject180) *NullableInlineObject180 {
	return &NullableInlineObject180{value: val, isSet: true}
}

func (v NullableInlineObject180) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject180) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


