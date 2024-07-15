/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20064FailoverAndFailback WAN failover and failback
type InlineResponse20064FailoverAndFailback struct {
	Immediate *InlineResponse20064FailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewInlineResponse20064FailoverAndFailback instantiates a new InlineResponse20064FailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20064FailoverAndFailback() *InlineResponse20064FailoverAndFailback {
	this := InlineResponse20064FailoverAndFailback{}
	return &this
}

// NewInlineResponse20064FailoverAndFailbackWithDefaults instantiates a new InlineResponse20064FailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20064FailoverAndFailbackWithDefaults() *InlineResponse20064FailoverAndFailback {
	this := InlineResponse20064FailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *InlineResponse20064FailoverAndFailback) GetImmediate() InlineResponse20064FailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret InlineResponse20064FailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20064FailoverAndFailback) GetImmediateOk() (*InlineResponse20064FailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *InlineResponse20064FailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given InlineResponse20064FailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *InlineResponse20064FailoverAndFailback) SetImmediate(v InlineResponse20064FailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o InlineResponse20064FailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20064FailoverAndFailback struct {
	value *InlineResponse20064FailoverAndFailback
	isSet bool
}

func (v NullableInlineResponse20064FailoverAndFailback) Get() *InlineResponse20064FailoverAndFailback {
	return v.value
}

func (v *NullableInlineResponse20064FailoverAndFailback) Set(val *InlineResponse20064FailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20064FailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20064FailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20064FailoverAndFailback(val *InlineResponse20064FailoverAndFailback) *NullableInlineResponse20064FailoverAndFailback {
	return &NullableInlineResponse20064FailoverAndFailback{value: val, isSet: true}
}

func (v NullableInlineResponse20064FailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20064FailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

