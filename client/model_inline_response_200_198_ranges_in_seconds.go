/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200198RangesInSeconds struct for InlineResponse200198RangesInSeconds
type InlineResponse200198RangesInSeconds struct {
	// Seconds since Sunday at midnight when the outage range starts.
	Start int32 `json:"start"`
	// Seconds since Sunday at midnight when that outage range ends.
	End int32 `json:"end"`
}

// NewInlineResponse200198RangesInSeconds instantiates a new InlineResponse200198RangesInSeconds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200198RangesInSeconds(start int32, end int32) *InlineResponse200198RangesInSeconds {
	this := InlineResponse200198RangesInSeconds{}
	this.Start = start
	this.End = end
	return &this
}

// NewInlineResponse200198RangesInSecondsWithDefaults instantiates a new InlineResponse200198RangesInSeconds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200198RangesInSecondsWithDefaults() *InlineResponse200198RangesInSeconds {
	this := InlineResponse200198RangesInSeconds{}
	return &this
}

// GetStart returns the Start field value
func (o *InlineResponse200198RangesInSeconds) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200198RangesInSeconds) GetStartOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *InlineResponse200198RangesInSeconds) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *InlineResponse200198RangesInSeconds) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200198RangesInSeconds) GetEndOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *InlineResponse200198RangesInSeconds) SetEnd(v int32) {
	o.End = v
}

func (o InlineResponse200198RangesInSeconds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200198RangesInSeconds struct {
	value *InlineResponse200198RangesInSeconds
	isSet bool
}

func (v NullableInlineResponse200198RangesInSeconds) Get() *InlineResponse200198RangesInSeconds {
	return v.value
}

func (v *NullableInlineResponse200198RangesInSeconds) Set(val *InlineResponse200198RangesInSeconds) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200198RangesInSeconds) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200198RangesInSeconds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200198RangesInSeconds(val *InlineResponse200198RangesInSeconds) *NullableInlineResponse200198RangesInSeconds {
	return &NullableInlineResponse200198RangesInSeconds{value: val, isSet: true}
}

func (v NullableInlineResponse200198RangesInSeconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200198RangesInSeconds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


