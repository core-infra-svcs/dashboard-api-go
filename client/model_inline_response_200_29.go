/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20029 struct for InlineResponse20029
type InlineResponse20029 struct {
	DdnsHostnames *InlineResponse20029DdnsHostnames `json:"ddnsHostnames,omitempty"`
	Wan1 *InlineResponse20029Wan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse20029Wan2 `json:"wan2,omitempty"`
}

// NewInlineResponse20029 instantiates a new InlineResponse20029 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029WithDefaults() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// GetDdnsHostnames returns the DdnsHostnames field value if set, zero value otherwise.
func (o *InlineResponse20029) GetDdnsHostnames() InlineResponse20029DdnsHostnames {
	if o == nil || isNil(o.DdnsHostnames) {
		var ret InlineResponse20029DdnsHostnames
		return ret
	}
	return *o.DdnsHostnames
}

// GetDdnsHostnamesOk returns a tuple with the DdnsHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetDdnsHostnamesOk() (*InlineResponse20029DdnsHostnames, bool) {
	if o == nil || isNil(o.DdnsHostnames) {
    return nil, false
	}
	return o.DdnsHostnames, true
}

// HasDdnsHostnames returns a boolean if a field has been set.
func (o *InlineResponse20029) HasDdnsHostnames() bool {
	if o != nil && !isNil(o.DdnsHostnames) {
		return true
	}

	return false
}

// SetDdnsHostnames gets a reference to the given InlineResponse20029DdnsHostnames and assigns it to the DdnsHostnames field.
func (o *InlineResponse20029) SetDdnsHostnames(v InlineResponse20029DdnsHostnames) {
	o.DdnsHostnames = &v
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse20029) GetWan1() InlineResponse20029Wan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse20029Wan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetWan1Ok() (*InlineResponse20029Wan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse20029) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse20029Wan1 and assigns it to the Wan1 field.
func (o *InlineResponse20029) SetWan1(v InlineResponse20029Wan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse20029) GetWan2() InlineResponse20029Wan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse20029Wan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetWan2Ok() (*InlineResponse20029Wan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse20029) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse20029Wan2 and assigns it to the Wan2 field.
func (o *InlineResponse20029) SetWan2(v InlineResponse20029Wan2) {
	o.Wan2 = &v
}

func (o InlineResponse20029) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20029 struct {
	value *InlineResponse20029
	isSet bool
}

func (v NullableInlineResponse20029) Get() *InlineResponse20029 {
	return v.value
}

func (v *NullableInlineResponse20029) Set(val *InlineResponse20029) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029(val *InlineResponse20029) *NullableInlineResponse20029 {
	return &NullableInlineResponse20029{value: val, isSet: true}
}

func (v NullableInlineResponse20029) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


