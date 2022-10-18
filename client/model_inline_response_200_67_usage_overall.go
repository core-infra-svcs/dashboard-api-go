/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20067UsageOverall Overall data usage of all clients across organization
type InlineResponse20067UsageOverall struct {
	// Total data usage (in kb) of all clients across organization
	Total *float32 `json:"total,omitempty"`
	// Downstream data usage (in kb) of all clients across organization
	Downstream *float32 `json:"downstream,omitempty"`
	// Upstream data usage (in kb) of all clients across organization
	Upstream *float32 `json:"upstream,omitempty"`
}

// NewInlineResponse20067UsageOverall instantiates a new InlineResponse20067UsageOverall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20067UsageOverall() *InlineResponse20067UsageOverall {
	this := InlineResponse20067UsageOverall{}
	return &this
}

// NewInlineResponse20067UsageOverallWithDefaults instantiates a new InlineResponse20067UsageOverall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20067UsageOverallWithDefaults() *InlineResponse20067UsageOverall {
	this := InlineResponse20067UsageOverall{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20067UsageOverall) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20067UsageOverall) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20067UsageOverall) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *InlineResponse20067UsageOverall) SetTotal(v float32) {
	o.Total = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse20067UsageOverall) GetDownstream() float32 {
	if o == nil || o.Downstream == nil {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20067UsageOverall) GetDownstreamOk() (*float32, bool) {
	if o == nil || o.Downstream == nil {
		return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse20067UsageOverall) HasDownstream() bool {
	if o != nil && o.Downstream != nil {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *InlineResponse20067UsageOverall) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse20067UsageOverall) GetUpstream() float32 {
	if o == nil || o.Upstream == nil {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20067UsageOverall) GetUpstreamOk() (*float32, bool) {
	if o == nil || o.Upstream == nil {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse20067UsageOverall) HasUpstream() bool {
	if o != nil && o.Upstream != nil {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *InlineResponse20067UsageOverall) SetUpstream(v float32) {
	o.Upstream = &v
}

func (o InlineResponse20067UsageOverall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Downstream != nil {
		toSerialize["downstream"] = o.Downstream
	}
	if o.Upstream != nil {
		toSerialize["upstream"] = o.Upstream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20067UsageOverall struct {
	value *InlineResponse20067UsageOverall
	isSet bool
}

func (v NullableInlineResponse20067UsageOverall) Get() *InlineResponse20067UsageOverall {
	return v.value
}

func (v *NullableInlineResponse20067UsageOverall) Set(val *InlineResponse20067UsageOverall) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20067UsageOverall) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20067UsageOverall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20067UsageOverall(val *InlineResponse20067UsageOverall) *NullableInlineResponse20067UsageOverall {
	return &NullableInlineResponse20067UsageOverall{value: val, isSet: true}
}

func (v NullableInlineResponse20067UsageOverall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20067UsageOverall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


