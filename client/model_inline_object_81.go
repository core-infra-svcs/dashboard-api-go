/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject81 struct for InlineObject81
type InlineObject81 struct {
	// The ids of the devices to attempt activation lock bypass.
	Ids []string `json:"ids"`
}

// NewInlineObject81 instantiates a new InlineObject81 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject81(ids []string) *InlineObject81 {
	this := InlineObject81{}
	this.Ids = ids
	return &this
}

// NewInlineObject81WithDefaults instantiates a new InlineObject81 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject81WithDefaults() *InlineObject81 {
	this := InlineObject81{}
	return &this
}

// GetIds returns the Ids field value
func (o *InlineObject81) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InlineObject81) GetIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *InlineObject81) SetIds(v []string) {
	o.Ids = v
}

func (o InlineObject81) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
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


