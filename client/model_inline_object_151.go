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

// InlineObject151 struct for InlineObject151
type InlineObject151 struct {
	// The serial of the switch to be removed
	Serial string `json:"serial"`
}

// NewInlineObject151 instantiates a new InlineObject151 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject151(serial string) *InlineObject151 {
	this := InlineObject151{}
	this.Serial = serial
	return &this
}

// NewInlineObject151WithDefaults instantiates a new InlineObject151 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject151WithDefaults() *InlineObject151 {
	this := InlineObject151{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject151) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject151) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject151) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject151) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject151 struct {
	value *InlineObject151
	isSet bool
}

func (v NullableInlineObject151) Get() *InlineObject151 {
	return v.value
}

func (v *NullableInlineObject151) Set(val *InlineObject151) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject151) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject151) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject151(val *InlineObject151) *NullableInlineObject151 {
	return &NullableInlineObject151{value: val, isSet: true}
}

func (v NullableInlineObject151) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject151) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


