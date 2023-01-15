/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20014FailoverAndFailback WAN failover and failback
type InlineResponse20014FailoverAndFailback struct {
	Immediate *InlineResponse20014FailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewInlineResponse20014FailoverAndFailback instantiates a new InlineResponse20014FailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014FailoverAndFailback() *InlineResponse20014FailoverAndFailback {
	this := InlineResponse20014FailoverAndFailback{}
	return &this
}

// NewInlineResponse20014FailoverAndFailbackWithDefaults instantiates a new InlineResponse20014FailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014FailoverAndFailbackWithDefaults() *InlineResponse20014FailoverAndFailback {
	this := InlineResponse20014FailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *InlineResponse20014FailoverAndFailback) GetImmediate() InlineResponse20014FailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret InlineResponse20014FailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014FailoverAndFailback) GetImmediateOk() (*InlineResponse20014FailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *InlineResponse20014FailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given InlineResponse20014FailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *InlineResponse20014FailoverAndFailback) SetImmediate(v InlineResponse20014FailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o InlineResponse20014FailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014FailoverAndFailback struct {
	value *InlineResponse20014FailoverAndFailback
	isSet bool
}

func (v NullableInlineResponse20014FailoverAndFailback) Get() *InlineResponse20014FailoverAndFailback {
	return v.value
}

func (v *NullableInlineResponse20014FailoverAndFailback) Set(val *InlineResponse20014FailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014FailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014FailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014FailoverAndFailback(val *InlineResponse20014FailoverAndFailback) *NullableInlineResponse20014FailoverAndFailback {
	return &NullableInlineResponse20014FailoverAndFailback{value: val, isSet: true}
}

func (v NullableInlineResponse20014FailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014FailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


