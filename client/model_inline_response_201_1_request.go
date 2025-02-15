/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2011Request ARP table request parameters
type InlineResponse2011Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
}

// NewInlineResponse2011Request instantiates a new InlineResponse2011Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2011Request() *InlineResponse2011Request {
	this := InlineResponse2011Request{}
	return &this
}

// NewInlineResponse2011RequestWithDefaults instantiates a new InlineResponse2011Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2011RequestWithDefaults() *InlineResponse2011Request {
	this := InlineResponse2011Request{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2011Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2011Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2011Request) SetSerial(v string) {
	o.Serial = &v
}

func (o InlineResponse2011Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2011Request struct {
	value *InlineResponse2011Request
	isSet bool
}

func (v NullableInlineResponse2011Request) Get() *InlineResponse2011Request {
	return v.value
}

func (v *NullableInlineResponse2011Request) Set(val *InlineResponse2011Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2011Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2011Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2011Request(val *InlineResponse2011Request) *NullableInlineResponse2011Request {
	return &NullableInlineResponse2011Request{value: val, isSet: true}
}

func (v NullableInlineResponse2011Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2011Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


