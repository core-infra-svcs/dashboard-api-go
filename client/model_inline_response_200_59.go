/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20059 struct for InlineResponse20059
type InlineResponse20059 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Average latency in milliseconds
	AvgLatencyMs *int32 `json:"avgLatencyMs,omitempty"`
}

// NewInlineResponse20059 instantiates a new InlineResponse20059 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// NewInlineResponse20059WithDefaults instantiates a new InlineResponse20059 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059WithDefaults() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse20059) GetStartTs() time.Time {
	if o == nil || o.StartTs == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetStartTsOk() (*time.Time, bool) {
	if o == nil || o.StartTs == nil {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse20059) HasStartTs() bool {
	if o != nil && o.StartTs != nil {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse20059) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse20059) GetEndTs() time.Time {
	if o == nil || o.EndTs == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetEndTsOk() (*time.Time, bool) {
	if o == nil || o.EndTs == nil {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse20059) HasEndTs() bool {
	if o != nil && o.EndTs != nil {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse20059) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetAvgLatencyMs returns the AvgLatencyMs field value if set, zero value otherwise.
func (o *InlineResponse20059) GetAvgLatencyMs() int32 {
	if o == nil || o.AvgLatencyMs == nil {
		var ret int32
		return ret
	}
	return *o.AvgLatencyMs
}

// GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetAvgLatencyMsOk() (*int32, bool) {
	if o == nil || o.AvgLatencyMs == nil {
		return nil, false
	}
	return o.AvgLatencyMs, true
}

// HasAvgLatencyMs returns a boolean if a field has been set.
func (o *InlineResponse20059) HasAvgLatencyMs() bool {
	if o != nil && o.AvgLatencyMs != nil {
		return true
	}

	return false
}

// SetAvgLatencyMs gets a reference to the given int32 and assigns it to the AvgLatencyMs field.
func (o *InlineResponse20059) SetAvgLatencyMs(v int32) {
	o.AvgLatencyMs = &v
}

func (o InlineResponse20059) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartTs != nil {
		toSerialize["startTs"] = o.StartTs
	}
	if o.EndTs != nil {
		toSerialize["endTs"] = o.EndTs
	}
	if o.AvgLatencyMs != nil {
		toSerialize["avgLatencyMs"] = o.AvgLatencyMs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059 struct {
	value *InlineResponse20059
	isSet bool
}

func (v NullableInlineResponse20059) Get() *InlineResponse20059 {
	return v.value
}

func (v *NullableInlineResponse20059) Set(val *InlineResponse20059) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059(val *InlineResponse20059) *NullableInlineResponse20059 {
	return &NullableInlineResponse20059{value: val, isSet: true}
}

func (v NullableInlineResponse20059) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


