/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject15 struct for InlineObject15
type InlineObject15 struct {
	Wan1 *DevicesSerialManagementInterfaceWan1 `json:"wan1,omitempty"`
	Wan2 *DevicesSerialManagementInterfaceWan2 `json:"wan2,omitempty"`
}

// NewInlineObject15 instantiates a new InlineObject15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject15() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// NewInlineObject15WithDefaults instantiates a new InlineObject15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject15WithDefaults() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineObject15) GetWan1() DevicesSerialManagementInterfaceWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret DevicesSerialManagementInterfaceWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetWan1Ok() (*DevicesSerialManagementInterfaceWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineObject15) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given DevicesSerialManagementInterfaceWan1 and assigns it to the Wan1 field.
func (o *InlineObject15) SetWan1(v DevicesSerialManagementInterfaceWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineObject15) GetWan2() DevicesSerialManagementInterfaceWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret DevicesSerialManagementInterfaceWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetWan2Ok() (*DevicesSerialManagementInterfaceWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineObject15) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given DevicesSerialManagementInterfaceWan2 and assigns it to the Wan2 field.
func (o *InlineObject15) SetWan2(v DevicesSerialManagementInterfaceWan2) {
	o.Wan2 = &v
}

func (o InlineObject15) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject15 struct {
	value *InlineObject15
	isSet bool
}

func (v NullableInlineObject15) Get() *InlineObject15 {
	return v.value
}

func (v *NullableInlineObject15) Set(val *InlineObject15) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject15) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject15(val *InlineObject15) *NullableInlineObject15 {
	return &NullableInlineObject15{value: val, isSet: true}
}

func (v NullableInlineObject15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


