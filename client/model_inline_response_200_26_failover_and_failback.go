/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026FailoverAndFailback WAN failover and failback
type InlineResponse20026FailoverAndFailback struct {
	Immediate *InlineResponse20026FailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewInlineResponse20026FailoverAndFailback instantiates a new InlineResponse20026FailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026FailoverAndFailback() *InlineResponse20026FailoverAndFailback {
	this := InlineResponse20026FailoverAndFailback{}
	return &this
}

// NewInlineResponse20026FailoverAndFailbackWithDefaults instantiates a new InlineResponse20026FailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026FailoverAndFailbackWithDefaults() *InlineResponse20026FailoverAndFailback {
	this := InlineResponse20026FailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *InlineResponse20026FailoverAndFailback) GetImmediate() InlineResponse20026FailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret InlineResponse20026FailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026FailoverAndFailback) GetImmediateOk() (*InlineResponse20026FailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *InlineResponse20026FailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given InlineResponse20026FailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *InlineResponse20026FailoverAndFailback) SetImmediate(v InlineResponse20026FailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o InlineResponse20026FailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026FailoverAndFailback struct {
	value *InlineResponse20026FailoverAndFailback
	isSet bool
}

func (v NullableInlineResponse20026FailoverAndFailback) Get() *InlineResponse20026FailoverAndFailback {
	return v.value
}

func (v *NullableInlineResponse20026FailoverAndFailback) Set(val *InlineResponse20026FailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026FailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026FailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026FailoverAndFailback(val *InlineResponse20026FailoverAndFailback) *NullableInlineResponse20026FailoverAndFailback {
	return &NullableInlineResponse20026FailoverAndFailback{value: val, isSet: true}
}

func (v NullableInlineResponse20026FailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026FailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

