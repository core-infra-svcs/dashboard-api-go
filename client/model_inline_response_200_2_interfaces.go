/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2002Interfaces Interface settings.
type InlineResponse2002Interfaces struct {
	Wan1 *InlineResponse2002InterfacesWan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse2002InterfacesWan2 `json:"wan2,omitempty"`
}

// NewInlineResponse2002Interfaces instantiates a new InlineResponse2002Interfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Interfaces() *InlineResponse2002Interfaces {
	this := InlineResponse2002Interfaces{}
	return &this
}

// NewInlineResponse2002InterfacesWithDefaults instantiates a new InlineResponse2002Interfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002InterfacesWithDefaults() *InlineResponse2002Interfaces {
	this := InlineResponse2002Interfaces{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse2002Interfaces) GetWan1() InlineResponse2002InterfacesWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse2002InterfacesWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Interfaces) GetWan1Ok() (*InlineResponse2002InterfacesWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse2002Interfaces) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse2002InterfacesWan1 and assigns it to the Wan1 field.
func (o *InlineResponse2002Interfaces) SetWan1(v InlineResponse2002InterfacesWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse2002Interfaces) GetWan2() InlineResponse2002InterfacesWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse2002InterfacesWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Interfaces) GetWan2Ok() (*InlineResponse2002InterfacesWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse2002Interfaces) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse2002InterfacesWan2 and assigns it to the Wan2 field.
func (o *InlineResponse2002Interfaces) SetWan2(v InlineResponse2002InterfacesWan2) {
	o.Wan2 = &v
}

func (o InlineResponse2002Interfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002Interfaces struct {
	value *InlineResponse2002Interfaces
	isSet bool
}

func (v NullableInlineResponse2002Interfaces) Get() *InlineResponse2002Interfaces {
	return v.value
}

func (v *NullableInlineResponse2002Interfaces) Set(val *InlineResponse2002Interfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Interfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Interfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Interfaces(val *InlineResponse2002Interfaces) *NullableInlineResponse2002Interfaces {
	return &NullableInlineResponse2002Interfaces{value: val, isSet: true}
}

func (v NullableInlineResponse2002Interfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Interfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


