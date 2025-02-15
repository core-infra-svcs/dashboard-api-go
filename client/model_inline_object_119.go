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

// InlineObject119 struct for InlineObject119
type InlineObject119 struct {
	// The ids of the devices to attempt activation lock bypass.
	Ids []string `json:"ids"`
}

// NewInlineObject119 instantiates a new InlineObject119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject119(ids []string) *InlineObject119 {
	this := InlineObject119{}
	this.Ids = ids
	return &this
}

// NewInlineObject119WithDefaults instantiates a new InlineObject119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject119WithDefaults() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// GetIds returns the Ids field value
func (o *InlineObject119) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *InlineObject119) SetIds(v []string) {
	o.Ids = v
}

func (o InlineObject119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject119 struct {
	value *InlineObject119
	isSet bool
}

func (v NullableInlineObject119) Get() *InlineObject119 {
	return v.value
}

func (v *NullableInlineObject119) Set(val *InlineObject119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject119(val *InlineObject119) *NullableInlineObject119 {
	return &NullableInlineObject119{value: val, isSet: true}
}

func (v NullableInlineObject119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


