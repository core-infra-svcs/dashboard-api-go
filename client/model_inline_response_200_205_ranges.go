/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200205Ranges struct for InlineResponse200205Ranges
type InlineResponse200205Ranges struct {
	// Day of when the outage starts. Can be either full day name, or three letter abbreviation.
	StartDay string `json:"startDay"`
	// 24 hour time when the outage starts.
	StartTime string `json:"startTime"`
	// Day of when the outage ends. Can be either full day name, or three letter abbreviation
	EndDay string `json:"endDay"`
	// 24 hour time when the outage ends.
	EndTime string `json:"endTime"`
}

// NewInlineResponse200205Ranges instantiates a new InlineResponse200205Ranges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200205Ranges(startDay string, startTime string, endDay string, endTime string) *InlineResponse200205Ranges {
	this := InlineResponse200205Ranges{}
	this.StartDay = startDay
	this.StartTime = startTime
	this.EndDay = endDay
	this.EndTime = endTime
	return &this
}

// NewInlineResponse200205RangesWithDefaults instantiates a new InlineResponse200205Ranges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200205RangesWithDefaults() *InlineResponse200205Ranges {
	this := InlineResponse200205Ranges{}
	return &this
}

// GetStartDay returns the StartDay field value
func (o *InlineResponse200205Ranges) GetStartDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDay
}

// GetStartDayOk returns a tuple with the StartDay field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200205Ranges) GetStartDayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StartDay, true
}

// SetStartDay sets field value
func (o *InlineResponse200205Ranges) SetStartDay(v string) {
	o.StartDay = v
}

// GetStartTime returns the StartTime field value
func (o *InlineResponse200205Ranges) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200205Ranges) GetStartTimeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *InlineResponse200205Ranges) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndDay returns the EndDay field value
func (o *InlineResponse200205Ranges) GetEndDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDay
}

// GetEndDayOk returns a tuple with the EndDay field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200205Ranges) GetEndDayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EndDay, true
}

// SetEndDay sets field value
func (o *InlineResponse200205Ranges) SetEndDay(v string) {
	o.EndDay = v
}

// GetEndTime returns the EndTime field value
func (o *InlineResponse200205Ranges) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200205Ranges) GetEndTimeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *InlineResponse200205Ranges) SetEndTime(v string) {
	o.EndTime = v
}

func (o InlineResponse200205Ranges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startDay"] = o.StartDay
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endDay"] = o.EndDay
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200205Ranges struct {
	value *InlineResponse200205Ranges
	isSet bool
}

func (v NullableInlineResponse200205Ranges) Get() *InlineResponse200205Ranges {
	return v.value
}

func (v *NullableInlineResponse200205Ranges) Set(val *InlineResponse200205Ranges) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200205Ranges) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200205Ranges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200205Ranges(val *InlineResponse200205Ranges) *NullableInlineResponse200205Ranges {
	return &NullableInlineResponse200205Ranges{value: val, isSet: true}
}

func (v NullableInlineResponse200205Ranges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200205Ranges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


