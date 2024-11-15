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

// InlineResponse2014Request Ping request parameters
type InlineResponse2014Request struct {
	// Number of pings to send. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse2014Request instantiates a new InlineResponse2014Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2014Request() *InlineResponse2014Request {
	this := InlineResponse2014Request{}
	return &this
}

// NewInlineResponse2014RequestWithDefaults instantiates a new InlineResponse2014Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2014RequestWithDefaults() *InlineResponse2014Request {
	this := InlineResponse2014Request{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse2014Request) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2014Request) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse2014Request) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse2014Request) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse2014Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2014Request struct {
	value *InlineResponse2014Request
	isSet bool
}

func (v NullableInlineResponse2014Request) Get() *InlineResponse2014Request {
	return v.value
}

func (v *NullableInlineResponse2014Request) Set(val *InlineResponse2014Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2014Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2014Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2014Request(val *InlineResponse2014Request) *NullableInlineResponse2014Request {
	return &NullableInlineResponse2014Request{value: val, isSet: true}
}

func (v NullableInlineResponse2014Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2014Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


