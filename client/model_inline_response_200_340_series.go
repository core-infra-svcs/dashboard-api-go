/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200340Series struct for InlineResponse200340Series
type InlineResponse200340Series struct {
	// Timestamp of the cpu load measurement
	Ts *time.Time `json:"ts,omitempty"`
	// The 5 minutes cpu load average of the device
	CpuLoad5 *int32 `json:"cpuLoad5,omitempty"`
}

// NewInlineResponse200340Series instantiates a new InlineResponse200340Series object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200340Series() *InlineResponse200340Series {
	this := InlineResponse200340Series{}
	return &this
}

// NewInlineResponse200340SeriesWithDefaults instantiates a new InlineResponse200340Series object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200340SeriesWithDefaults() *InlineResponse200340Series {
	this := InlineResponse200340Series{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200340Series) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200340Series) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200340Series) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200340Series) SetTs(v time.Time) {
	o.Ts = &v
}

// GetCpuLoad5 returns the CpuLoad5 field value if set, zero value otherwise.
func (o *InlineResponse200340Series) GetCpuLoad5() int32 {
	if o == nil || isNil(o.CpuLoad5) {
		var ret int32
		return ret
	}
	return *o.CpuLoad5
}

// GetCpuLoad5Ok returns a tuple with the CpuLoad5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200340Series) GetCpuLoad5Ok() (*int32, bool) {
	if o == nil || isNil(o.CpuLoad5) {
    return nil, false
	}
	return o.CpuLoad5, true
}

// HasCpuLoad5 returns a boolean if a field has been set.
func (o *InlineResponse200340Series) HasCpuLoad5() bool {
	if o != nil && !isNil(o.CpuLoad5) {
		return true
	}

	return false
}

// SetCpuLoad5 gets a reference to the given int32 and assigns it to the CpuLoad5 field.
func (o *InlineResponse200340Series) SetCpuLoad5(v int32) {
	o.CpuLoad5 = &v
}

func (o InlineResponse200340Series) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.CpuLoad5) {
		toSerialize["cpuLoad5"] = o.CpuLoad5
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200340Series struct {
	value *InlineResponse200340Series
	isSet bool
}

func (v NullableInlineResponse200340Series) Get() *InlineResponse200340Series {
	return v.value
}

func (v *NullableInlineResponse200340Series) Set(val *InlineResponse200340Series) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200340Series) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200340Series) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200340Series(val *InlineResponse200340Series) *NullableInlineResponse200340Series {
	return &NullableInlineResponse200340Series{value: val, isSet: true}
}

func (v NullableInlineResponse200340Series) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200340Series) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


