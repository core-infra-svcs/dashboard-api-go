/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	Ids DevicesSerialCameraWirelessProfilesIds `json:"ids"`
}

// NewInlineObject12 instantiates a new InlineObject12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject12(ids DevicesSerialCameraWirelessProfilesIds) *InlineObject12 {
	this := InlineObject12{}
	this.Ids = ids
	return &this
}

// NewInlineObject12WithDefaults instantiates a new InlineObject12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject12WithDefaults() *InlineObject12 {
	this := InlineObject12{}
	return &this
}

// GetIds returns the Ids field value
func (o *InlineObject12) GetIds() DevicesSerialCameraWirelessProfilesIds {
	if o == nil {
		var ret DevicesSerialCameraWirelessProfilesIds
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetIdsOk() (*DevicesSerialCameraWirelessProfilesIds, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *InlineObject12) SetIds(v DevicesSerialCameraWirelessProfilesIds) {
	o.Ids = v
}

func (o InlineObject12) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject12 struct {
	value *InlineObject12
	isSet bool
}

func (v NullableInlineObject12) Get() *InlineObject12 {
	return v.value
}

func (v *NullableInlineObject12) Set(val *InlineObject12) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject12) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject12(val *InlineObject12) *NullableInlineObject12 {
	return &NullableInlineObject12{value: val, isSet: true}
}

func (v NullableInlineObject12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


