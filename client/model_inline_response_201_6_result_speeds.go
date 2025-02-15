/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2016ResultSpeeds Shows the speeds (Mbps)
type InlineResponse2016ResultSpeeds struct {
	// Shows the download speed from shard (Mbps)
	Downstream *float32 `json:"downstream,omitempty"`
}

// NewInlineResponse2016ResultSpeeds instantiates a new InlineResponse2016ResultSpeeds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2016ResultSpeeds() *InlineResponse2016ResultSpeeds {
	this := InlineResponse2016ResultSpeeds{}
	return &this
}

// NewInlineResponse2016ResultSpeedsWithDefaults instantiates a new InlineResponse2016ResultSpeeds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2016ResultSpeedsWithDefaults() *InlineResponse2016ResultSpeeds {
	this := InlineResponse2016ResultSpeeds{}
	return &this
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse2016ResultSpeeds) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016ResultSpeeds) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse2016ResultSpeeds) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *InlineResponse2016ResultSpeeds) SetDownstream(v float32) {
	o.Downstream = &v
}

func (o InlineResponse2016ResultSpeeds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2016ResultSpeeds struct {
	value *InlineResponse2016ResultSpeeds
	isSet bool
}

func (v NullableInlineResponse2016ResultSpeeds) Get() *InlineResponse2016ResultSpeeds {
	return v.value
}

func (v *NullableInlineResponse2016ResultSpeeds) Set(val *InlineResponse2016ResultSpeeds) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2016ResultSpeeds) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2016ResultSpeeds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2016ResultSpeeds(val *InlineResponse2016ResultSpeeds) *NullableInlineResponse2016ResultSpeeds {
	return &NullableInlineResponse2016ResultSpeeds{value: val, isSet: true}
}

func (v NullableInlineResponse2016ResultSpeeds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2016ResultSpeeds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


