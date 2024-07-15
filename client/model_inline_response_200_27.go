/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20027 struct for InlineResponse20027
type InlineResponse20027 struct {
	DdnsHostnames *InlineResponse20027DdnsHostnames `json:"ddnsHostnames,omitempty"`
	Wan1 *InlineResponse20027Wan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse20027Wan2 `json:"wan2,omitempty"`
}

// NewInlineResponse20027 instantiates a new InlineResponse20027 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20027() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20027WithDefaults() *InlineResponse20027 {
	this := InlineResponse20027{}
	return &this
}

// GetDdnsHostnames returns the DdnsHostnames field value if set, zero value otherwise.
func (o *InlineResponse20027) GetDdnsHostnames() InlineResponse20027DdnsHostnames {
	if o == nil || isNil(o.DdnsHostnames) {
		var ret InlineResponse20027DdnsHostnames
		return ret
	}
	return *o.DdnsHostnames
}

// GetDdnsHostnamesOk returns a tuple with the DdnsHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetDdnsHostnamesOk() (*InlineResponse20027DdnsHostnames, bool) {
	if o == nil || isNil(o.DdnsHostnames) {
    return nil, false
	}
	return o.DdnsHostnames, true
}

// HasDdnsHostnames returns a boolean if a field has been set.
func (o *InlineResponse20027) HasDdnsHostnames() bool {
	if o != nil && !isNil(o.DdnsHostnames) {
		return true
	}

	return false
}

// SetDdnsHostnames gets a reference to the given InlineResponse20027DdnsHostnames and assigns it to the DdnsHostnames field.
func (o *InlineResponse20027) SetDdnsHostnames(v InlineResponse20027DdnsHostnames) {
	o.DdnsHostnames = &v
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse20027) GetWan1() InlineResponse20027Wan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse20027Wan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetWan1Ok() (*InlineResponse20027Wan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse20027) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse20027Wan1 and assigns it to the Wan1 field.
func (o *InlineResponse20027) SetWan1(v InlineResponse20027Wan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse20027) GetWan2() InlineResponse20027Wan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse20027Wan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20027) GetWan2Ok() (*InlineResponse20027Wan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse20027) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse20027Wan2 and assigns it to the Wan2 field.
func (o *InlineResponse20027) SetWan2(v InlineResponse20027Wan2) {
	o.Wan2 = &v
}

func (o InlineResponse20027) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DdnsHostnames) {
		toSerialize["ddnsHostnames"] = o.DdnsHostnames
	}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20027 struct {
	value *InlineResponse20027
	isSet bool
}

func (v NullableInlineResponse20027) Get() *InlineResponse20027 {
	return v.value
}

func (v *NullableInlineResponse20027) Set(val *InlineResponse20027) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20027) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20027) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20027(val *InlineResponse20027) *NullableInlineResponse20027 {
	return &NullableInlineResponse20027{value: val, isSet: true}
}

func (v NullableInlineResponse20027) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20027) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


