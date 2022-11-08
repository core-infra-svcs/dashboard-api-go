/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject129 struct for InlineObject129
type InlineObject129 struct {
	// The serial of the switch to be added
	Serial string `json:"serial"`
}

// NewInlineObject129 instantiates a new InlineObject129 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject129(serial string) *InlineObject129 {
	this := InlineObject129{}
	this.Serial = serial
	return &this
}

// NewInlineObject129WithDefaults instantiates a new InlineObject129 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject129WithDefaults() *InlineObject129 {
	this := InlineObject129{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject129) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject129) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject129) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject129) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject129 struct {
	value *InlineObject129
	isSet bool
}

func (v NullableInlineObject129) Get() *InlineObject129 {
	return v.value
}

func (v *NullableInlineObject129) Set(val *InlineObject129) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject129) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject129) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject129(val *InlineObject129) *NullableInlineObject129 {
	return &NullableInlineObject129{value: val, isSet: true}
}

func (v NullableInlineObject129) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject129) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


