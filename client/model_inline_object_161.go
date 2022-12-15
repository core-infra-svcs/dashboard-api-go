/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject161 struct for InlineObject161
type InlineObject161 struct {
	// If true, the SSID outage schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	Ranges []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges `json:"ranges,omitempty"`
	// List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
	RangesInSeconds []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds `json:"rangesInSeconds,omitempty"`
}

// NewInlineObject161 instantiates a new InlineObject161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject161() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// NewInlineObject161WithDefaults instantiates a new InlineObject161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject161WithDefaults() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject161) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject161) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject161) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *InlineObject161) GetRanges() []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges {
	if o == nil || isNil(o.Ranges) {
		var ret []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetRangesOk() ([]NetworksNetworkIdWirelessSsidsNumberSchedulesRanges, bool) {
	if o == nil || isNil(o.Ranges) {
    return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *InlineObject161) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges and assigns it to the Ranges field.
func (o *InlineObject161) SetRanges(v []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) {
	o.Ranges = v
}

// GetRangesInSeconds returns the RangesInSeconds field value if set, zero value otherwise.
func (o *InlineObject161) GetRangesInSeconds() []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds {
	if o == nil || isNil(o.RangesInSeconds) {
		var ret []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds
		return ret
	}
	return o.RangesInSeconds
}

// GetRangesInSecondsOk returns a tuple with the RangesInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetRangesInSecondsOk() ([]NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds, bool) {
	if o == nil || isNil(o.RangesInSeconds) {
    return nil, false
	}
	return o.RangesInSeconds, true
}

// HasRangesInSeconds returns a boolean if a field has been set.
func (o *InlineObject161) HasRangesInSeconds() bool {
	if o != nil && !isNil(o.RangesInSeconds) {
		return true
	}

	return false
}

// SetRangesInSeconds gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds and assigns it to the RangesInSeconds field.
func (o *InlineObject161) SetRangesInSeconds(v []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) {
	o.RangesInSeconds = v
}

func (o InlineObject161) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject161 struct {
	value *InlineObject161
	isSet bool
}

func (v NullableInlineObject161) Get() *InlineObject161 {
	return v.value
}

func (v *NullableInlineObject161) Set(val *InlineObject161) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject161) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject161(val *InlineObject161) *NullableInlineObject161 {
	return &NullableInlineObject161{value: val, isSet: true}
}

func (v NullableInlineObject161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


