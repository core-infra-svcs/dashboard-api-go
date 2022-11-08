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

// InlineResponse20074UsageOverall Overall data usage of all clients across organization
type InlineResponse20074UsageOverall struct {
	// Total data usage (in kb) of all clients across organization
	Total *float32 `json:"total,omitempty"`
	// Downstream data usage (in kb) of all clients across organization
	Downstream *float32 `json:"downstream,omitempty"`
	// Upstream data usage (in kb) of all clients across organization
	Upstream *float32 `json:"upstream,omitempty"`
}

// NewInlineResponse20074UsageOverall instantiates a new InlineResponse20074UsageOverall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074UsageOverall() *InlineResponse20074UsageOverall {
	this := InlineResponse20074UsageOverall{}
	return &this
}

// NewInlineResponse20074UsageOverallWithDefaults instantiates a new InlineResponse20074UsageOverall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074UsageOverallWithDefaults() *InlineResponse20074UsageOverall {
	this := InlineResponse20074UsageOverall{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20074UsageOverall) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074UsageOverall) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20074UsageOverall) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *InlineResponse20074UsageOverall) SetTotal(v float32) {
	o.Total = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse20074UsageOverall) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074UsageOverall) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse20074UsageOverall) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *InlineResponse20074UsageOverall) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse20074UsageOverall) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074UsageOverall) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse20074UsageOverall) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *InlineResponse20074UsageOverall) SetUpstream(v float32) {
	o.Upstream = &v
}

func (o InlineResponse20074UsageOverall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074UsageOverall struct {
	value *InlineResponse20074UsageOverall
	isSet bool
}

func (v NullableInlineResponse20074UsageOverall) Get() *InlineResponse20074UsageOverall {
	return v.value
}

func (v *NullableInlineResponse20074UsageOverall) Set(val *InlineResponse20074UsageOverall) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074UsageOverall) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074UsageOverall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074UsageOverall(val *InlineResponse20074UsageOverall) *NullableInlineResponse20074UsageOverall {
	return &NullableInlineResponse20074UsageOverall{value: val, isSet: true}
}

func (v NullableInlineResponse20074UsageOverall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074UsageOverall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

