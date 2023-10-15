/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject81 struct for InlineObject81
type InlineObject81 struct {
	// The serial of a device
	Serial string `json:"serial"`
}

// NewInlineObject81 instantiates a new InlineObject81 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject81(serial string) *InlineObject81 {
	this := InlineObject81{}
	this.Serial = serial
	return &this
}

// NewInlineObject81WithDefaults instantiates a new InlineObject81 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject81WithDefaults() *InlineObject81 {
	this := InlineObject81{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject81) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject81) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject81) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject81) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject81 struct {
	value *InlineObject81
	isSet bool
}

func (v NullableInlineObject81) Get() *InlineObject81 {
	return v.value
}

func (v *NullableInlineObject81) Set(val *InlineObject81) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject81) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject81) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject81(val *InlineObject81) *NullableInlineObject81 {
	return &NullableInlineObject81{value: val, isSet: true}
}

func (v NullableInlineObject81) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject81) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


