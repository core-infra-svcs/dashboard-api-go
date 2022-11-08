/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2002ResultsReplies struct for InlineResponse2002ResultsReplies
type InlineResponse2002ResultsReplies struct {
	// Sequence ID of the packet
	SequenceId *int32 `json:"sequenceId,omitempty"`
	// Size of the packet in bytes
	Size *int32 `json:"size,omitempty"`
	// Latency of the packet in milliseconds
	Latency *float32 `json:"latency,omitempty"`
}

// NewInlineResponse2002ResultsReplies instantiates a new InlineResponse2002ResultsReplies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002ResultsReplies() *InlineResponse2002ResultsReplies {
	this := InlineResponse2002ResultsReplies{}
	return &this
}

// NewInlineResponse2002ResultsRepliesWithDefaults instantiates a new InlineResponse2002ResultsReplies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002ResultsRepliesWithDefaults() *InlineResponse2002ResultsReplies {
	this := InlineResponse2002ResultsReplies{}
	return &this
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise.
func (o *InlineResponse2002ResultsReplies) GetSequenceId() int32 {
	if o == nil || isNil(o.SequenceId) {
		var ret int32
		return ret
	}
	return *o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ResultsReplies) GetSequenceIdOk() (*int32, bool) {
	if o == nil || isNil(o.SequenceId) {
    return nil, false
	}
	return o.SequenceId, true
}

// HasSequenceId returns a boolean if a field has been set.
func (o *InlineResponse2002ResultsReplies) HasSequenceId() bool {
	if o != nil && !isNil(o.SequenceId) {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given int32 and assigns it to the SequenceId field.
func (o *InlineResponse2002ResultsReplies) SetSequenceId(v int32) {
	o.SequenceId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InlineResponse2002ResultsReplies) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ResultsReplies) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InlineResponse2002ResultsReplies) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *InlineResponse2002ResultsReplies) SetSize(v int32) {
	o.Size = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *InlineResponse2002ResultsReplies) GetLatency() float32 {
	if o == nil || isNil(o.Latency) {
		var ret float32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ResultsReplies) GetLatencyOk() (*float32, bool) {
	if o == nil || isNil(o.Latency) {
    return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *InlineResponse2002ResultsReplies) HasLatency() bool {
	if o != nil && !isNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given float32 and assigns it to the Latency field.
func (o *InlineResponse2002ResultsReplies) SetLatency(v float32) {
	o.Latency = &v
}

func (o InlineResponse2002ResultsReplies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SequenceId) {
		toSerialize["sequenceId"] = o.SequenceId
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002ResultsReplies struct {
	value *InlineResponse2002ResultsReplies
	isSet bool
}

func (v NullableInlineResponse2002ResultsReplies) Get() *InlineResponse2002ResultsReplies {
	return v.value
}

func (v *NullableInlineResponse2002ResultsReplies) Set(val *InlineResponse2002ResultsReplies) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002ResultsReplies) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002ResultsReplies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002ResultsReplies(val *InlineResponse2002ResultsReplies) *NullableInlineResponse2002ResultsReplies {
	return &NullableInlineResponse2002ResultsReplies{value: val, isSet: true}
}

func (v NullableInlineResponse2002ResultsReplies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002ResultsReplies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


