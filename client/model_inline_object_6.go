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

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	// The duration in seconds. Must be between 5 and 120. Default is 20 seconds
	Duration *int32 `json:"duration,omitempty"`
	// The period in milliseconds. Must be between 100 and 1000. Default is 160 milliseconds
	Period *int32 `json:"period,omitempty"`
	// The duty cycle as the percent active. Must be between 10 and 90. Default is 50.
	Duty *int32 `json:"duty,omitempty"`
}

// NewInlineObject6 instantiates a new InlineObject6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject6() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// NewInlineObject6WithDefaults instantiates a new InlineObject6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject6WithDefaults() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineObject6) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineObject6) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineObject6) SetDuration(v int32) {
	o.Duration = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *InlineObject6) GetPeriod() int32 {
	if o == nil || isNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.Period) {
    return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *InlineObject6) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *InlineObject6) SetPeriod(v int32) {
	o.Period = &v
}

// GetDuty returns the Duty field value if set, zero value otherwise.
func (o *InlineObject6) GetDuty() int32 {
	if o == nil || isNil(o.Duty) {
		var ret int32
		return ret
	}
	return *o.Duty
}

// GetDutyOk returns a tuple with the Duty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetDutyOk() (*int32, bool) {
	if o == nil || isNil(o.Duty) {
    return nil, false
	}
	return o.Duty, true
}

// HasDuty returns a boolean if a field has been set.
func (o *InlineObject6) HasDuty() bool {
	if o != nil && !isNil(o.Duty) {
		return true
	}

	return false
}

// SetDuty gets a reference to the given int32 and assigns it to the Duty field.
func (o *InlineObject6) SetDuty(v int32) {
	o.Duty = &v
}

func (o InlineObject6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !isNil(o.Duty) {
		toSerialize["duty"] = o.Duty
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject6 struct {
	value *InlineObject6
	isSet bool
}

func (v NullableInlineObject6) Get() *InlineObject6 {
	return v.value
}

func (v *NullableInlineObject6) Set(val *InlineObject6) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject6) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject6(val *InlineObject6) *NullableInlineObject6 {
	return &NullableInlineObject6{value: val, isSet: true}
}

func (v NullableInlineObject6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


