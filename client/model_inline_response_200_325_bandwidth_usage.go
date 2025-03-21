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

// InlineResponse200325BandwidthUsage Bandwidth usage data for the given interval.
type InlineResponse200325BandwidthUsage struct {
	// The average speed of the data sent and received (in kilobits-per-second).
	Total *float32 `json:"total,omitempty"`
	// The average speed of the data sent (in kilobits-per-second).
	Upstream *float32 `json:"upstream,omitempty"`
	// The average speed of the data received (in kilobits-per-second).
	Downstream *float32 `json:"downstream,omitempty"`
}

// NewInlineResponse200325BandwidthUsage instantiates a new InlineResponse200325BandwidthUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200325BandwidthUsage() *InlineResponse200325BandwidthUsage {
	this := InlineResponse200325BandwidthUsage{}
	return &this
}

// NewInlineResponse200325BandwidthUsageWithDefaults instantiates a new InlineResponse200325BandwidthUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200325BandwidthUsageWithDefaults() *InlineResponse200325BandwidthUsage {
	this := InlineResponse200325BandwidthUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200325BandwidthUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325BandwidthUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200325BandwidthUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *InlineResponse200325BandwidthUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse200325BandwidthUsage) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325BandwidthUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse200325BandwidthUsage) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *InlineResponse200325BandwidthUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse200325BandwidthUsage) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325BandwidthUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse200325BandwidthUsage) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *InlineResponse200325BandwidthUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

func (o InlineResponse200325BandwidthUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200325BandwidthUsage struct {
	value *InlineResponse200325BandwidthUsage
	isSet bool
}

func (v NullableInlineResponse200325BandwidthUsage) Get() *InlineResponse200325BandwidthUsage {
	return v.value
}

func (v *NullableInlineResponse200325BandwidthUsage) Set(val *InlineResponse200325BandwidthUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200325BandwidthUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200325BandwidthUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200325BandwidthUsage(val *InlineResponse200325BandwidthUsage) *NullableInlineResponse200325BandwidthUsage {
	return &NullableInlineResponse200325BandwidthUsage{value: val, isSet: true}
}

func (v NullableInlineResponse200325BandwidthUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200325BandwidthUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


