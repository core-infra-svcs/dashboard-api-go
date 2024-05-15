/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject149 struct for InlineObject149
type InlineObject149 struct {
	// The serial of the switch to be added
	Serial string `json:"serial"`
}

// NewInlineObject149 instantiates a new InlineObject149 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject149(serial string) *InlineObject149 {
	this := InlineObject149{}
	this.Serial = serial
	return &this
}

// NewInlineObject149WithDefaults instantiates a new InlineObject149 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject149WithDefaults() *InlineObject149 {
	this := InlineObject149{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject149) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject149) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject149) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject149) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject149 struct {
	value *InlineObject149
	isSet bool
}

func (v NullableInlineObject149) Get() *InlineObject149 {
	return v.value
}

func (v *NullableInlineObject149) Set(val *InlineObject149) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject149) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject149) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject149(val *InlineObject149) *NullableInlineObject149 {
	return &NullableInlineObject149{value: val, isSet: true}
}

func (v NullableInlineObject149) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject149) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


