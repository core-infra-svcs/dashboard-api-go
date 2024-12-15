/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2021 struct for InlineResponse2021
type InlineResponse2021 struct {
	// The duration in seconds. Will be between 5 and 120. Default is 20 seconds
	Duration *int32 `json:"duration,omitempty"`
	// The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds
	Period *int32 `json:"period,omitempty"`
	// The duty cycle as the percent active. Will be between 10 and 90. Default is 50
	Duty *int32 `json:"duty,omitempty"`
}

// NewInlineResponse2021 instantiates a new InlineResponse2021 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2021() *InlineResponse2021 {
	this := InlineResponse2021{}
	return &this
}

// NewInlineResponse2021WithDefaults instantiates a new InlineResponse2021 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2021WithDefaults() *InlineResponse2021 {
	this := InlineResponse2021{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse2021) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2021) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse2021) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse2021) SetDuration(v int32) {
	o.Duration = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *InlineResponse2021) GetPeriod() int32 {
	if o == nil || isNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2021) GetPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.Period) {
    return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *InlineResponse2021) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *InlineResponse2021) SetPeriod(v int32) {
	o.Period = &v
}

// GetDuty returns the Duty field value if set, zero value otherwise.
func (o *InlineResponse2021) GetDuty() int32 {
	if o == nil || isNil(o.Duty) {
		var ret int32
		return ret
	}
	return *o.Duty
}

// GetDutyOk returns a tuple with the Duty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2021) GetDutyOk() (*int32, bool) {
	if o == nil || isNil(o.Duty) {
    return nil, false
	}
	return o.Duty, true
}

// HasDuty returns a boolean if a field has been set.
func (o *InlineResponse2021) HasDuty() bool {
	if o != nil && !isNil(o.Duty) {
		return true
	}

	return false
}

// SetDuty gets a reference to the given int32 and assigns it to the Duty field.
func (o *InlineResponse2021) SetDuty(v int32) {
	o.Duty = &v
}

func (o InlineResponse2021) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2021 struct {
	value *InlineResponse2021
	isSet bool
}

func (v NullableInlineResponse2021) Get() *InlineResponse2021 {
	return v.value
}

func (v *NullableInlineResponse2021) Set(val *InlineResponse2021) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2021) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2021) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2021(val *InlineResponse2021) *NullableInlineResponse2021 {
	return &NullableInlineResponse2021{value: val, isSet: true}
}

func (v NullableInlineResponse2021) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2021) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


