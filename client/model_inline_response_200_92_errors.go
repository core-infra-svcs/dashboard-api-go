/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20092Errors struct for InlineResponse20092Errors
type InlineResponse20092Errors struct {
	// The serial of the device
	Serial string `json:"serial"`
	// The errors for the device
	Errors []string `json:"errors"`
}

// NewInlineResponse20092Errors instantiates a new InlineResponse20092Errors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092Errors(serial string, errors []string) *InlineResponse20092Errors {
	this := InlineResponse20092Errors{}
	this.Serial = serial
	this.Errors = errors
	return &this
}

// NewInlineResponse20092ErrorsWithDefaults instantiates a new InlineResponse20092Errors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092ErrorsWithDefaults() *InlineResponse20092Errors {
	this := InlineResponse20092Errors{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineResponse20092Errors) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20092Errors) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineResponse20092Errors) SetSerial(v string) {
	o.Serial = v
}

// GetErrors returns the Errors field value
func (o *InlineResponse20092Errors) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20092Errors) GetErrorsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *InlineResponse20092Errors) SetErrors(v []string) {
	o.Errors = v
}

func (o InlineResponse20092Errors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092Errors struct {
	value *InlineResponse20092Errors
	isSet bool
}

func (v NullableInlineResponse20092Errors) Get() *InlineResponse20092Errors {
	return v.value
}

func (v *NullableInlineResponse20092Errors) Set(val *InlineResponse20092Errors) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092Errors) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092Errors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092Errors(val *InlineResponse20092Errors) *NullableInlineResponse20092Errors {
	return &NullableInlineResponse20092Errors{value: val, isSet: true}
}

func (v NullableInlineResponse20092Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092Errors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


