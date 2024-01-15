/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20029FailoverAndFailback WAN failover and failback
type InlineResponse20029FailoverAndFailback struct {
	Immediate *InlineResponse20029FailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewInlineResponse20029FailoverAndFailback instantiates a new InlineResponse20029FailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029FailoverAndFailback() *InlineResponse20029FailoverAndFailback {
	this := InlineResponse20029FailoverAndFailback{}
	return &this
}

// NewInlineResponse20029FailoverAndFailbackWithDefaults instantiates a new InlineResponse20029FailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029FailoverAndFailbackWithDefaults() *InlineResponse20029FailoverAndFailback {
	this := InlineResponse20029FailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *InlineResponse20029FailoverAndFailback) GetImmediate() InlineResponse20029FailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret InlineResponse20029FailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029FailoverAndFailback) GetImmediateOk() (*InlineResponse20029FailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *InlineResponse20029FailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given InlineResponse20029FailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *InlineResponse20029FailoverAndFailback) SetImmediate(v InlineResponse20029FailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o InlineResponse20029FailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029FailoverAndFailback struct {
	value *InlineResponse20029FailoverAndFailback
	isSet bool
}

func (v NullableInlineResponse20029FailoverAndFailback) Get() *InlineResponse20029FailoverAndFailback {
	return v.value
}

func (v *NullableInlineResponse20029FailoverAndFailback) Set(val *InlineResponse20029FailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029FailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029FailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029FailoverAndFailback(val *InlineResponse20029FailoverAndFailback) *NullableInlineResponse20029FailoverAndFailback {
	return &NullableInlineResponse20029FailoverAndFailback{value: val, isSet: true}
}

func (v NullableInlineResponse20029FailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029FailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


