/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject204 struct for InlineObject204
type InlineObject204 struct {
	// A boolean representing whether or not the batch has been confirmed. This property cannot be unset once it is true.
	Confirmed *bool `json:"confirmed,omitempty"`
	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch.
	Synchronous *bool `json:"synchronous,omitempty"`
}

// NewInlineObject204 instantiates a new InlineObject204 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject204() *InlineObject204 {
	this := InlineObject204{}
	return &this
}

// NewInlineObject204WithDefaults instantiates a new InlineObject204 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject204WithDefaults() *InlineObject204 {
	this := InlineObject204{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineObject204) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineObject204) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineObject204) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *InlineObject204) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject204) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *InlineObject204) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *InlineObject204) SetSynchronous(v bool) {
	o.Synchronous = &v
}

func (o InlineObject204) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !isNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject204 struct {
	value *InlineObject204
	isSet bool
}

func (v NullableInlineObject204) Get() *InlineObject204 {
	return v.value
}

func (v *NullableInlineObject204) Set(val *InlineObject204) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject204) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject204) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject204(val *InlineObject204) *NullableInlineObject204 {
	return &NullableInlineObject204{value: val, isSet: true}
}

func (v NullableInlineObject204) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject204) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


