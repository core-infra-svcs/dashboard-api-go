/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject157 struct for InlineObject157
type InlineObject157 struct {
	// The serial of the switch to be removed
	Serial string `json:"serial"`
}

// NewInlineObject157 instantiates a new InlineObject157 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject157(serial string) *InlineObject157 {
	this := InlineObject157{}
	this.Serial = serial
	return &this
}

// NewInlineObject157WithDefaults instantiates a new InlineObject157 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject157WithDefaults() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject157) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject157) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject157) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject157 struct {
	value *InlineObject157
	isSet bool
}

func (v NullableInlineObject157) Get() *InlineObject157 {
	return v.value
}

func (v *NullableInlineObject157) Set(val *InlineObject157) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject157) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject157) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject157(val *InlineObject157) *NullableInlineObject157 {
	return &NullableInlineObject157{value: val, isSet: true}
}

func (v NullableInlineObject157) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject157) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


