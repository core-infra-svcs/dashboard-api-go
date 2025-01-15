/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200332Readings struct for InlineResponse200332Readings
type InlineResponse200332Readings struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	Counts *InlineResponse200332Counts `json:"counts,omitempty"`
}

// NewInlineResponse200332Readings instantiates a new InlineResponse200332Readings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200332Readings() *InlineResponse200332Readings {
	this := InlineResponse200332Readings{}
	return &this
}

// NewInlineResponse200332ReadingsWithDefaults instantiates a new InlineResponse200332Readings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200332ReadingsWithDefaults() *InlineResponse200332Readings {
	this := InlineResponse200332Readings{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200332Readings) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200332Readings) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200332Readings) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200332Readings) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200332Readings) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200332Readings) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200332Readings) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200332Readings) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200332Readings) GetCounts() InlineResponse200332Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200332Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200332Readings) GetCountsOk() (*InlineResponse200332Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200332Readings) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200332Counts and assigns it to the Counts field.
func (o *InlineResponse200332Readings) SetCounts(v InlineResponse200332Counts) {
	o.Counts = &v
}

func (o InlineResponse200332Readings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200332Readings struct {
	value *InlineResponse200332Readings
	isSet bool
}

func (v NullableInlineResponse200332Readings) Get() *InlineResponse200332Readings {
	return v.value
}

func (v *NullableInlineResponse200332Readings) Set(val *InlineResponse200332Readings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200332Readings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200332Readings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200332Readings(val *InlineResponse200332Readings) *NullableInlineResponse200332Readings {
	return &NullableInlineResponse200332Readings{value: val, isSet: true}
}

func (v NullableInlineResponse200332Readings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200332Readings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

