/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200254 struct for InlineResponse200254
type InlineResponse200254 struct {
	// Timestamp for the bandwidth usage snapshot.
	Ts *time.Time `json:"ts,omitempty"`
	// Total bandwidth usage, in mbps.
	Total *int32 `json:"total,omitempty"`
	// Uploaded data, in mbps.
	Upstream *int32 `json:"upstream,omitempty"`
	// Downloaded data, in mbps.
	Downstream *int32 `json:"downstream,omitempty"`
}

// NewInlineResponse200254 instantiates a new InlineResponse200254 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200254() *InlineResponse200254 {
	this := InlineResponse200254{}
	return &this
}

// NewInlineResponse200254WithDefaults instantiates a new InlineResponse200254 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200254WithDefaults() *InlineResponse200254 {
	this := InlineResponse200254{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200254) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200254) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200254) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200254) SetTs(v time.Time) {
	o.Ts = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200254) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200254) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200254) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200254) SetTotal(v int32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse200254) GetUpstream() int32 {
	if o == nil || isNil(o.Upstream) {
		var ret int32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200254) GetUpstreamOk() (*int32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse200254) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given int32 and assigns it to the Upstream field.
func (o *InlineResponse200254) SetUpstream(v int32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse200254) GetDownstream() int32 {
	if o == nil || isNil(o.Downstream) {
		var ret int32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200254) GetDownstreamOk() (*int32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse200254) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given int32 and assigns it to the Downstream field.
func (o *InlineResponse200254) SetDownstream(v int32) {
	o.Downstream = &v
}

func (o InlineResponse200254) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
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

type NullableInlineResponse200254 struct {
	value *InlineResponse200254
	isSet bool
}

func (v NullableInlineResponse200254) Get() *InlineResponse200254 {
	return v.value
}

func (v *NullableInlineResponse200254) Set(val *InlineResponse200254) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200254) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200254) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200254(val *InlineResponse200254) *NullableInlineResponse200254 {
	return &NullableInlineResponse200254{value: val, isSet: true}
}

func (v NullableInlineResponse200254) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200254) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


