/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023ResultsReplies struct for InlineResponse20023ResultsReplies
type InlineResponse20023ResultsReplies struct {
	// Sequence ID of the packet
	SequenceId *int32 `json:"sequenceId,omitempty"`
	// Size of the packet in bytes
	Size *int32 `json:"size,omitempty"`
	// Latency of the packet in milliseconds
	Latency *float32 `json:"latency,omitempty"`
}

// NewInlineResponse20023ResultsReplies instantiates a new InlineResponse20023ResultsReplies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023ResultsReplies() *InlineResponse20023ResultsReplies {
	this := InlineResponse20023ResultsReplies{}
	return &this
}

// NewInlineResponse20023ResultsRepliesWithDefaults instantiates a new InlineResponse20023ResultsReplies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023ResultsRepliesWithDefaults() *InlineResponse20023ResultsReplies {
	this := InlineResponse20023ResultsReplies{}
	return &this
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise.
func (o *InlineResponse20023ResultsReplies) GetSequenceId() int32 {
	if o == nil || isNil(o.SequenceId) {
		var ret int32
		return ret
	}
	return *o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ResultsReplies) GetSequenceIdOk() (*int32, bool) {
	if o == nil || isNil(o.SequenceId) {
    return nil, false
	}
	return o.SequenceId, true
}

// HasSequenceId returns a boolean if a field has been set.
func (o *InlineResponse20023ResultsReplies) HasSequenceId() bool {
	if o != nil && !isNil(o.SequenceId) {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given int32 and assigns it to the SequenceId field.
func (o *InlineResponse20023ResultsReplies) SetSequenceId(v int32) {
	o.SequenceId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InlineResponse20023ResultsReplies) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ResultsReplies) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InlineResponse20023ResultsReplies) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *InlineResponse20023ResultsReplies) SetSize(v int32) {
	o.Size = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *InlineResponse20023ResultsReplies) GetLatency() float32 {
	if o == nil || isNil(o.Latency) {
		var ret float32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ResultsReplies) GetLatencyOk() (*float32, bool) {
	if o == nil || isNil(o.Latency) {
    return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *InlineResponse20023ResultsReplies) HasLatency() bool {
	if o != nil && !isNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given float32 and assigns it to the Latency field.
func (o *InlineResponse20023ResultsReplies) SetLatency(v float32) {
	o.Latency = &v
}

func (o InlineResponse20023ResultsReplies) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20023ResultsReplies struct {
	value *InlineResponse20023ResultsReplies
	isSet bool
}

func (v NullableInlineResponse20023ResultsReplies) Get() *InlineResponse20023ResultsReplies {
	return v.value
}

func (v *NullableInlineResponse20023ResultsReplies) Set(val *InlineResponse20023ResultsReplies) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023ResultsReplies) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023ResultsReplies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023ResultsReplies(val *InlineResponse20023ResultsReplies) *NullableInlineResponse20023ResultsReplies {
	return &NullableInlineResponse20023ResultsReplies{value: val, isSet: true}
}

func (v NullableInlineResponse20023ResultsReplies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023ResultsReplies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

