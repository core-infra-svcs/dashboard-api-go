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

// InlineResponse200203RangesInSeconds struct for InlineResponse200203RangesInSeconds
type InlineResponse200203RangesInSeconds struct {
	// Seconds since Sunday at midnight when the outage range starts.
	Start int32 `json:"start"`
	// Seconds since Sunday at midnight when that outage range ends.
	End int32 `json:"end"`
}

// NewInlineResponse200203RangesInSeconds instantiates a new InlineResponse200203RangesInSeconds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200203RangesInSeconds(start int32, end int32) *InlineResponse200203RangesInSeconds {
	this := InlineResponse200203RangesInSeconds{}
	this.Start = start
	this.End = end
	return &this
}

// NewInlineResponse200203RangesInSecondsWithDefaults instantiates a new InlineResponse200203RangesInSeconds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200203RangesInSecondsWithDefaults() *InlineResponse200203RangesInSeconds {
	this := InlineResponse200203RangesInSeconds{}
	return &this
}

// GetStart returns the Start field value
func (o *InlineResponse200203RangesInSeconds) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200203RangesInSeconds) GetStartOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *InlineResponse200203RangesInSeconds) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *InlineResponse200203RangesInSeconds) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200203RangesInSeconds) GetEndOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *InlineResponse200203RangesInSeconds) SetEnd(v int32) {
	o.End = v
}

func (o InlineResponse200203RangesInSeconds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200203RangesInSeconds struct {
	value *InlineResponse200203RangesInSeconds
	isSet bool
}

func (v NullableInlineResponse200203RangesInSeconds) Get() *InlineResponse200203RangesInSeconds {
	return v.value
}

func (v *NullableInlineResponse200203RangesInSeconds) Set(val *InlineResponse200203RangesInSeconds) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200203RangesInSeconds) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200203RangesInSeconds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200203RangesInSeconds(val *InlineResponse200203RangesInSeconds) *NullableInlineResponse200203RangesInSeconds {
	return &NullableInlineResponse200203RangesInSeconds{value: val, isSet: true}
}

func (v NullableInlineResponse200203RangesInSeconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200203RangesInSeconds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


