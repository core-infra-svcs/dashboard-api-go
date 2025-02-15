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

// InlineResponse2016Request The parameters of the throughput test request
type InlineResponse2016Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
}

// NewInlineResponse2016Request instantiates a new InlineResponse2016Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2016Request() *InlineResponse2016Request {
	this := InlineResponse2016Request{}
	return &this
}

// NewInlineResponse2016RequestWithDefaults instantiates a new InlineResponse2016Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2016RequestWithDefaults() *InlineResponse2016Request {
	this := InlineResponse2016Request{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2016Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2016Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2016Request) SetSerial(v string) {
	o.Serial = &v
}

func (o InlineResponse2016Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2016Request struct {
	value *InlineResponse2016Request
	isSet bool
}

func (v NullableInlineResponse2016Request) Get() *InlineResponse2016Request {
	return v.value
}

func (v *NullableInlineResponse2016Request) Set(val *InlineResponse2016Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2016Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2016Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2016Request(val *InlineResponse2016Request) *NullableInlineResponse2016Request {
	return &NullableInlineResponse2016Request{value: val, isSet: true}
}

func (v NullableInlineResponse2016Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2016Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


