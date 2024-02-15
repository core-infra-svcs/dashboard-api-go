/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	// [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// [optional] If set to \"true\" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.
	Fullframe *bool `json:"fullframe,omitempty"`
}

// NewInlineObject8 instantiates a new InlineObject8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// NewInlineObject8WithDefaults instantiates a new InlineObject8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8WithDefaults() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *InlineObject8) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InlineObject8) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *InlineObject8) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetFullframe returns the Fullframe field value if set, zero value otherwise.
func (o *InlineObject8) GetFullframe() bool {
	if o == nil || isNil(o.Fullframe) {
		var ret bool
		return ret
	}
	return *o.Fullframe
}

// GetFullframeOk returns a tuple with the Fullframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetFullframeOk() (*bool, bool) {
	if o == nil || isNil(o.Fullframe) {
    return nil, false
	}
	return o.Fullframe, true
}

// HasFullframe returns a boolean if a field has been set.
func (o *InlineObject8) HasFullframe() bool {
	if o != nil && !isNil(o.Fullframe) {
		return true
	}

	return false
}

// SetFullframe gets a reference to the given bool and assigns it to the Fullframe field.
func (o *InlineObject8) SetFullframe(v bool) {
	o.Fullframe = &v
}

func (o InlineObject8) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Fullframe) {
		toSerialize["fullframe"] = o.Fullframe
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject8 struct {
	value *InlineObject8
	isSet bool
}

func (v NullableInlineObject8) Get() *InlineObject8 {
	return v.value
}

func (v *NullableInlineObject8) Set(val *InlineObject8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8(val *InlineObject8) *NullableInlineObject8 {
	return &NullableInlineObject8{value: val, isSet: true}
}

func (v NullableInlineObject8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


