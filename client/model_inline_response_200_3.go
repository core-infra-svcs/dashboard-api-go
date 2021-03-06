/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	// Timestamp for the bandwidth usage snapshot.
	Ts *time.Time `json:"ts,omitempty"`
	// Total bandwidth usage, in mbps.
	Total *int32 `json:"total,omitempty"`
	// Uploaded data, in mbps.
	Upstream *int32 `json:"upstream,omitempty"`
	// Downloaded data, in mbps.
	Downstream *int32 `json:"downstream,omitempty"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse2003) GetTs() time.Time {
	if o == nil || o.Ts == nil {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetTsOk() (*time.Time, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse2003) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse2003) SetTs(v time.Time) {
	o.Ts = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse2003) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse2003) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse2003) SetTotal(v int32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse2003) GetUpstream() int32 {
	if o == nil || o.Upstream == nil {
		var ret int32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetUpstreamOk() (*int32, bool) {
	if o == nil || o.Upstream == nil {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse2003) HasUpstream() bool {
	if o != nil && o.Upstream != nil {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given int32 and assigns it to the Upstream field.
func (o *InlineResponse2003) SetUpstream(v int32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse2003) GetDownstream() int32 {
	if o == nil || o.Downstream == nil {
		var ret int32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetDownstreamOk() (*int32, bool) {
	if o == nil || o.Downstream == nil {
		return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse2003) HasDownstream() bool {
	if o != nil && o.Downstream != nil {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given int32 and assigns it to the Downstream field.
func (o *InlineResponse2003) SetDownstream(v int32) {
	o.Downstream = &v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Upstream != nil {
		toSerialize["upstream"] = o.Upstream
	}
	if o.Downstream != nil {
		toSerialize["downstream"] = o.Downstream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


