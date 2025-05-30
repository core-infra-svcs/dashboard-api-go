/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject158 struct for InlineObject158
type InlineObject158 struct {
	// The serial of the switch to be removed
	Serial string `json:"serial"`
}

// NewInlineObject158 instantiates a new InlineObject158 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject158(serial string) *InlineObject158 {
	this := InlineObject158{}
	this.Serial = serial
	return &this
}

// NewInlineObject158WithDefaults instantiates a new InlineObject158 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject158WithDefaults() *InlineObject158 {
	this := InlineObject158{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject158) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject158) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject158) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject158) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject158 struct {
	value *InlineObject158
	isSet bool
}

func (v NullableInlineObject158) Get() *InlineObject158 {
	return v.value
}

func (v *NullableInlineObject158) Set(val *InlineObject158) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject158) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject158) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject158(val *InlineObject158) *NullableInlineObject158 {
	return &NullableInlineObject158{value: val, isSet: true}
}

func (v NullableInlineObject158) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject158) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


