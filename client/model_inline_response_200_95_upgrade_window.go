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

// InlineResponse20095UpgradeWindow Upgrade window for devices in network
type InlineResponse20095UpgradeWindow struct {
	// Day of the week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
	// Hour of the day
	HourOfDay *string `json:"hourOfDay,omitempty"`
}

// NewInlineResponse20095UpgradeWindow instantiates a new InlineResponse20095UpgradeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095UpgradeWindow() *InlineResponse20095UpgradeWindow {
	this := InlineResponse20095UpgradeWindow{}
	return &this
}

// NewInlineResponse20095UpgradeWindowWithDefaults instantiates a new InlineResponse20095UpgradeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095UpgradeWindowWithDefaults() *InlineResponse20095UpgradeWindow {
	this := InlineResponse20095UpgradeWindow{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *InlineResponse20095UpgradeWindow) GetDayOfWeek() string {
	if o == nil || isNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095UpgradeWindow) GetDayOfWeekOk() (*string, bool) {
	if o == nil || isNil(o.DayOfWeek) {
    return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *InlineResponse20095UpgradeWindow) HasDayOfWeek() bool {
	if o != nil && !isNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *InlineResponse20095UpgradeWindow) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetHourOfDay returns the HourOfDay field value if set, zero value otherwise.
func (o *InlineResponse20095UpgradeWindow) GetHourOfDay() string {
	if o == nil || isNil(o.HourOfDay) {
		var ret string
		return ret
	}
	return *o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095UpgradeWindow) GetHourOfDayOk() (*string, bool) {
	if o == nil || isNil(o.HourOfDay) {
    return nil, false
	}
	return o.HourOfDay, true
}

// HasHourOfDay returns a boolean if a field has been set.
func (o *InlineResponse20095UpgradeWindow) HasHourOfDay() bool {
	if o != nil && !isNil(o.HourOfDay) {
		return true
	}

	return false
}

// SetHourOfDay gets a reference to the given string and assigns it to the HourOfDay field.
func (o *InlineResponse20095UpgradeWindow) SetHourOfDay(v string) {
	o.HourOfDay = &v
}

func (o InlineResponse20095UpgradeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !isNil(o.HourOfDay) {
		toSerialize["hourOfDay"] = o.HourOfDay
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095UpgradeWindow struct {
	value *InlineResponse20095UpgradeWindow
	isSet bool
}

func (v NullableInlineResponse20095UpgradeWindow) Get() *InlineResponse20095UpgradeWindow {
	return v.value
}

func (v *NullableInlineResponse20095UpgradeWindow) Set(val *InlineResponse20095UpgradeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095UpgradeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095UpgradeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095UpgradeWindow(val *InlineResponse20095UpgradeWindow) *NullableInlineResponse20095UpgradeWindow {
	return &NullableInlineResponse20095UpgradeWindow{value: val, isSet: true}
}

func (v NullableInlineResponse20095UpgradeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095UpgradeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


