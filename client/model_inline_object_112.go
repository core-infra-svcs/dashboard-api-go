/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject112 struct for InlineObject112
type InlineObject112 struct {
	// The serial of the switch to be removed
	Serial string `json:"serial"`
}

// NewInlineObject112 instantiates a new InlineObject112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject112(serial string) *InlineObject112 {
	this := InlineObject112{}
	this.Serial = serial
	return &this
}

// NewInlineObject112WithDefaults instantiates a new InlineObject112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject112WithDefaults() *InlineObject112 {
	this := InlineObject112{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject112) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject112) GetSerialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject112) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject112 struct {
	value *InlineObject112
	isSet bool
}

func (v NullableInlineObject112) Get() *InlineObject112 {
	return v.value
}

func (v *NullableInlineObject112) Set(val *InlineObject112) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject112) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject112(val *InlineObject112) *NullableInlineObject112 {
	return &NullableInlineObject112{value: val, isSet: true}
}

func (v NullableInlineObject112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


