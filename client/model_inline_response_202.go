/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse202 struct for InlineResponse202
type InlineResponse202 struct {
	// The duration in seconds. Will be between 5 and 120. Default is 20 seconds
	Duration *int32 `json:"duration,omitempty"`
	// The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds
	Period *int32 `json:"period,omitempty"`
	// The duty cycle as the percent active. Will be between 10 and 90. Default is 50
	Duty *int32 `json:"duty,omitempty"`
}

// NewInlineResponse202 instantiates a new InlineResponse202 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse202() *InlineResponse202 {
	this := InlineResponse202{}
	return &this
}

// NewInlineResponse202WithDefaults instantiates a new InlineResponse202 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse202WithDefaults() *InlineResponse202 {
	this := InlineResponse202{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse202) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse202) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse202) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse202) SetDuration(v int32) {
	o.Duration = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *InlineResponse202) GetPeriod() int32 {
	if o == nil || isNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse202) GetPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.Period) {
    return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *InlineResponse202) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *InlineResponse202) SetPeriod(v int32) {
	o.Period = &v
}

// GetDuty returns the Duty field value if set, zero value otherwise.
func (o *InlineResponse202) GetDuty() int32 {
	if o == nil || isNil(o.Duty) {
		var ret int32
		return ret
	}
	return *o.Duty
}

// GetDutyOk returns a tuple with the Duty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse202) GetDutyOk() (*int32, bool) {
	if o == nil || isNil(o.Duty) {
    return nil, false
	}
	return o.Duty, true
}

// HasDuty returns a boolean if a field has been set.
func (o *InlineResponse202) HasDuty() bool {
	if o != nil && !isNil(o.Duty) {
		return true
	}

	return false
}

// SetDuty gets a reference to the given int32 and assigns it to the Duty field.
func (o *InlineResponse202) SetDuty(v int32) {
	o.Duty = &v
}

func (o InlineResponse202) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse202 struct {
	value *InlineResponse202
	isSet bool
}

func (v NullableInlineResponse202) Get() *InlineResponse202 {
	return v.value
}

func (v *NullableInlineResponse202) Set(val *InlineResponse202) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse202) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse202) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse202(val *InlineResponse202) *NullableInlineResponse202 {
	return &NullableInlineResponse202{value: val, isSet: true}
}

func (v NullableInlineResponse202) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse202) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


