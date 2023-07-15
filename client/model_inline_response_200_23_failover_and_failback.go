/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023FailoverAndFailback WAN failover and failback
type InlineResponse20023FailoverAndFailback struct {
	Immediate *InlineResponse20023FailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewInlineResponse20023FailoverAndFailback instantiates a new InlineResponse20023FailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023FailoverAndFailback() *InlineResponse20023FailoverAndFailback {
	this := InlineResponse20023FailoverAndFailback{}
	return &this
}

// NewInlineResponse20023FailoverAndFailbackWithDefaults instantiates a new InlineResponse20023FailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023FailoverAndFailbackWithDefaults() *InlineResponse20023FailoverAndFailback {
	this := InlineResponse20023FailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *InlineResponse20023FailoverAndFailback) GetImmediate() InlineResponse20023FailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret InlineResponse20023FailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023FailoverAndFailback) GetImmediateOk() (*InlineResponse20023FailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *InlineResponse20023FailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given InlineResponse20023FailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *InlineResponse20023FailoverAndFailback) SetImmediate(v InlineResponse20023FailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o InlineResponse20023FailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023FailoverAndFailback struct {
	value *InlineResponse20023FailoverAndFailback
	isSet bool
}

func (v NullableInlineResponse20023FailoverAndFailback) Get() *InlineResponse20023FailoverAndFailback {
	return v.value
}

func (v *NullableInlineResponse20023FailoverAndFailback) Set(val *InlineResponse20023FailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023FailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023FailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023FailoverAndFailback(val *InlineResponse20023FailoverAndFailback) *NullableInlineResponse20023FailoverAndFailback {
	return &NullableInlineResponse20023FailoverAndFailback{value: val, isSet: true}
}

func (v NullableInlineResponse20023FailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023FailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


