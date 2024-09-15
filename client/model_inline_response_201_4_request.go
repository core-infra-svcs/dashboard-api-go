/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2014Request The parameters of the throughput test request
type InlineResponse2014Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
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

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2014Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2014Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2014Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2014Request) SetSerial(v string) {
	o.Serial = &v
}

func (o InlineResponse2014Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
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


