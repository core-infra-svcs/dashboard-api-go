/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200109 struct for InlineResponse200109
type InlineResponse200109 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Average latency in milliseconds
	AvgLatencyMs *int32 `json:"avgLatencyMs,omitempty"`
}

// NewInlineResponse200109 instantiates a new InlineResponse200109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200109() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// NewInlineResponse200109WithDefaults instantiates a new InlineResponse200109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200109WithDefaults() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200109) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200109) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200109) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200109) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200109) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200109) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetAvgLatencyMs returns the AvgLatencyMs field value if set, zero value otherwise.
func (o *InlineResponse200109) GetAvgLatencyMs() int32 {
	if o == nil || isNil(o.AvgLatencyMs) {
		var ret int32
		return ret
	}
	return *o.AvgLatencyMs
}

// GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetAvgLatencyMsOk() (*int32, bool) {
	if o == nil || isNil(o.AvgLatencyMs) {
    return nil, false
	}
	return o.AvgLatencyMs, true
}

// HasAvgLatencyMs returns a boolean if a field has been set.
func (o *InlineResponse200109) HasAvgLatencyMs() bool {
	if o != nil && !isNil(o.AvgLatencyMs) {
		return true
	}

	return false
}

// SetAvgLatencyMs gets a reference to the given int32 and assigns it to the AvgLatencyMs field.
func (o *InlineResponse200109) SetAvgLatencyMs(v int32) {
	o.AvgLatencyMs = &v
}

func (o InlineResponse200109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.AvgLatencyMs) {
		toSerialize["avgLatencyMs"] = o.AvgLatencyMs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200109 struct {
	value *InlineResponse200109
	isSet bool
}

func (v NullableInlineResponse200109) Get() *InlineResponse200109 {
	return v.value
}

func (v *NullableInlineResponse200109) Set(val *InlineResponse200109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200109(val *InlineResponse200109) *NullableInlineResponse200109 {
	return &NullableInlineResponse200109{value: val, isSet: true}
}

func (v NullableInlineResponse200109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


