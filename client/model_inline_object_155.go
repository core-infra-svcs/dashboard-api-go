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

// InlineObject155 struct for InlineObject155
type InlineObject155 struct {
	// The serial of the switch to be added
	Serial string `json:"serial"`
}

// NewInlineObject155 instantiates a new InlineObject155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject155(serial string) *InlineObject155 {
	this := InlineObject155{}
	this.Serial = serial
	return &this
}

// NewInlineObject155WithDefaults instantiates a new InlineObject155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject155WithDefaults() *InlineObject155 {
	this := InlineObject155{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject155) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject155) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject155 struct {
	value *InlineObject155
	isSet bool
}

func (v NullableInlineObject155) Get() *InlineObject155 {
	return v.value
}

func (v *NullableInlineObject155) Set(val *InlineObject155) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject155) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject155(val *InlineObject155) *NullableInlineObject155 {
	return &NullableInlineObject155{value: val, isSet: true}
}

func (v NullableInlineObject155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


