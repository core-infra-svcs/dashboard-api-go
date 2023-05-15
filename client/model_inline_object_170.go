/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject170 struct for InlineObject170
type InlineObject170 struct {
	// A boolean representing whether or not the batch has been confirmed. This property cannot be unset once it is true.
	Confirmed *bool `json:"confirmed,omitempty"`
	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch.
	Synchronous *bool `json:"synchronous,omitempty"`
}

// NewInlineObject170 instantiates a new InlineObject170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject170() *InlineObject170 {
	this := InlineObject170{}
	return &this
}

// NewInlineObject170WithDefaults instantiates a new InlineObject170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject170WithDefaults() *InlineObject170 {
	this := InlineObject170{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineObject170) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineObject170) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineObject170) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *InlineObject170) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *InlineObject170) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *InlineObject170) SetSynchronous(v bool) {
	o.Synchronous = &v
}

func (o InlineObject170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !isNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject170 struct {
	value *InlineObject170
	isSet bool
}

func (v NullableInlineObject170) Get() *InlineObject170 {
	return v.value
}

func (v *NullableInlineObject170) Set(val *InlineObject170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject170(val *InlineObject170) *NullableInlineObject170 {
	return &NullableInlineObject170{value: val, isSet: true}
}

func (v NullableInlineObject170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


