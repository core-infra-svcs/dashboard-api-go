/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject104 struct for InlineObject104
type InlineObject104 struct {
	// The ids of the devices to attempt activation lock bypass.
	Ids []string `json:"ids"`
}

// NewInlineObject104 instantiates a new InlineObject104 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject104(ids []string) *InlineObject104 {
	this := InlineObject104{}
	this.Ids = ids
	return &this
}

// NewInlineObject104WithDefaults instantiates a new InlineObject104 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject104WithDefaults() *InlineObject104 {
	this := InlineObject104{}
	return &this
}

// GetIds returns the Ids field value
func (o *InlineObject104) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InlineObject104) GetIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *InlineObject104) SetIds(v []string) {
	o.Ids = v
}

func (o InlineObject104) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject104 struct {
	value *InlineObject104
	isSet bool
}

func (v NullableInlineObject104) Get() *InlineObject104 {
	return v.value
}

func (v *NullableInlineObject104) Set(val *InlineObject104) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject104) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject104) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject104(val *InlineObject104) *NullableInlineObject104 {
	return &NullableInlineObject104{value: val, isSet: true}
}

func (v NullableInlineObject104) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject104) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


