/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	Ids DevicesSerialCameraWirelessProfilesIds `json:"ids"`
}

// NewInlineObject9 instantiates a new InlineObject9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9(ids DevicesSerialCameraWirelessProfilesIds) *InlineObject9 {
	this := InlineObject9{}
	this.Ids = ids
	return &this
}

// NewInlineObject9WithDefaults instantiates a new InlineObject9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9WithDefaults() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// GetIds returns the Ids field value
func (o *InlineObject9) GetIds() DevicesSerialCameraWirelessProfilesIds {
	if o == nil {
		var ret DevicesSerialCameraWirelessProfilesIds
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetIdsOk() (*DevicesSerialCameraWirelessProfilesIds, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *InlineObject9) SetIds(v DevicesSerialCameraWirelessProfilesIds) {
	o.Ids = v
}

func (o InlineObject9) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject9 struct {
	value *InlineObject9
	isSet bool
}

func (v NullableInlineObject9) Get() *InlineObject9 {
	return v.value
}

func (v *NullableInlineObject9) Set(val *InlineObject9) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9(val *InlineObject9) *NullableInlineObject9 {
	return &NullableInlineObject9{value: val, isSet: true}
}

func (v NullableInlineObject9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


