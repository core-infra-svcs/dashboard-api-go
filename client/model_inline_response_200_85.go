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

// InlineResponse20085 struct for InlineResponse20085
type InlineResponse20085 struct {
	// The timestamp
	Ts *time.Time `json:"ts,omitempty"`
	// The total traffic over a time range for clients on a network
	Total *float32 `json:"total,omitempty"`
	// The upstream traffic over a time range for clients on a network
	Upstream *float32 `json:"upstream,omitempty"`
	// The downstream traffic over a time range for clients on a network
	Downstream *float32 `json:"downstream,omitempty"`
}

// NewInlineResponse20085 instantiates a new InlineResponse20085 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20085() *InlineResponse20085 {
	this := InlineResponse20085{}
	return &this
}

// NewInlineResponse20085WithDefaults instantiates a new InlineResponse20085 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20085WithDefaults() *InlineResponse20085 {
	this := InlineResponse20085{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20085) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20085) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse20085) SetTs(v time.Time) {
	o.Ts = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse20085) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse20085) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *InlineResponse20085) SetTotal(v float32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *InlineResponse20085) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *InlineResponse20085) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *InlineResponse20085) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *InlineResponse20085) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *InlineResponse20085) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *InlineResponse20085) SetDownstream(v float32) {
	o.Downstream = &v
}

func (o InlineResponse20085) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20085 struct {
	value *InlineResponse20085
	isSet bool
}

func (v NullableInlineResponse20085) Get() *InlineResponse20085 {
	return v.value
}

func (v *NullableInlineResponse20085) Set(val *InlineResponse20085) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20085) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20085) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20085(val *InlineResponse20085) *NullableInlineResponse20085 {
	return &NullableInlineResponse20085{value: val, isSet: true}
}

func (v NullableInlineResponse20085) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20085) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


