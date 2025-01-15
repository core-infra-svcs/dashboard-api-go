/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2015Request Ping request parameters
type InlineResponse2015Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// Number of pings to send. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse2015Request instantiates a new InlineResponse2015Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2015Request() *InlineResponse2015Request {
	this := InlineResponse2015Request{}
	return &this
}

// NewInlineResponse2015RequestWithDefaults instantiates a new InlineResponse2015Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2015RequestWithDefaults() *InlineResponse2015Request {
	this := InlineResponse2015Request{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2015Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2015Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2015Request) SetSerial(v string) {
	o.Serial = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse2015Request) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Request) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse2015Request) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse2015Request) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse2015Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2015Request struct {
	value *InlineResponse2015Request
	isSet bool
}

func (v NullableInlineResponse2015Request) Get() *InlineResponse2015Request {
	return v.value
}

func (v *NullableInlineResponse2015Request) Set(val *InlineResponse2015Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2015Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2015Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2015Request(val *InlineResponse2015Request) *NullableInlineResponse2015Request {
	return &NullableInlineResponse2015Request{value: val, isSet: true}
}

func (v NullableInlineResponse2015Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2015Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


