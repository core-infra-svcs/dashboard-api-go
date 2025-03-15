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

// InlineResponse200308 struct for InlineResponse200308
type InlineResponse200308 struct {
	// Timestamp of the start of the interval.
	Ts *time.Time `json:"ts,omitempty"`
	// The PoE power draw in watts for all switch ports in the organization for the given interval.
	Draw *float32 `json:"draw,omitempty"`
}

// NewInlineResponse200308 instantiates a new InlineResponse200308 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308() *InlineResponse200308 {
	this := InlineResponse200308{}
	return &this
}

// NewInlineResponse200308WithDefaults instantiates a new InlineResponse200308 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308WithDefaults() *InlineResponse200308 {
	this := InlineResponse200308{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200308) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200308) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200308) SetTs(v time.Time) {
	o.Ts = &v
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *InlineResponse200308) GetDraw() float32 {
	if o == nil || isNil(o.Draw) {
		var ret float32
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308) GetDrawOk() (*float32, bool) {
	if o == nil || isNil(o.Draw) {
    return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *InlineResponse200308) HasDraw() bool {
	if o != nil && !isNil(o.Draw) {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float32 and assigns it to the Draw field.
func (o *InlineResponse200308) SetDraw(v float32) {
	o.Draw = &v
}

func (o InlineResponse200308) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Draw) {
		toSerialize["draw"] = o.Draw
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308 struct {
	value *InlineResponse200308
	isSet bool
}

func (v NullableInlineResponse200308) Get() *InlineResponse200308 {
	return v.value
}

func (v *NullableInlineResponse200308) Set(val *InlineResponse200308) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308(val *InlineResponse200308) *NullableInlineResponse200308 {
	return &NullableInlineResponse200308{value: val, isSet: true}
}

func (v NullableInlineResponse200308) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


