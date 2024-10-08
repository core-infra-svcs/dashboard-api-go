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

// InlineObject150 struct for InlineObject150
type InlineObject150 struct {
	// The serial of the switch to be added
	Serial string `json:"serial"`
}

// NewInlineObject150 instantiates a new InlineObject150 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject150(serial string) *InlineObject150 {
	this := InlineObject150{}
	this.Serial = serial
	return &this
}

// NewInlineObject150WithDefaults instantiates a new InlineObject150 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject150WithDefaults() *InlineObject150 {
	this := InlineObject150{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject150) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject150) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject150) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject150) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject150 struct {
	value *InlineObject150
	isSet bool
}

func (v NullableInlineObject150) Get() *InlineObject150 {
	return v.value
}

func (v *NullableInlineObject150) Set(val *InlineObject150) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject150) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject150) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject150(val *InlineObject150) *NullableInlineObject150 {
	return &NullableInlineObject150{value: val, isSet: true}
}

func (v NullableInlineObject150) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject150) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


