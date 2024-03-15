/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject148 struct for InlineObject148
type InlineObject148 struct {
	// The serial of the switch to be removed
	Serial string `json:"serial"`
}

// NewInlineObject148 instantiates a new InlineObject148 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject148(serial string) *InlineObject148 {
	this := InlineObject148{}
	this.Serial = serial
	return &this
}

// NewInlineObject148WithDefaults instantiates a new InlineObject148 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject148WithDefaults() *InlineObject148 {
	this := InlineObject148{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject148) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject148) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject148) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject148) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject148 struct {
	value *InlineObject148
	isSet bool
}

func (v NullableInlineObject148) Get() *InlineObject148 {
	return v.value
}

func (v *NullableInlineObject148) Set(val *InlineObject148) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject148) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject148) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject148(val *InlineObject148) *NullableInlineObject148 {
	return &NullableInlineObject148{value: val, isSet: true}
}

func (v NullableInlineObject148) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject148) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


