/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject32 struct for InlineObject32
type InlineObject32 struct {
	// Enable or disable warm spare for a switch
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare switch
	SpareSerial *string `json:"spareSerial,omitempty"`
}

// NewInlineObject32 instantiates a new InlineObject32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject32(enabled bool) *InlineObject32 {
	this := InlineObject32{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject32WithDefaults instantiates a new InlineObject32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject32WithDefaults() *InlineObject32 {
	this := InlineObject32{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject32) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject32) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineObject32) GetSpareSerial() string {
	if o == nil || isNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetSpareSerialOk() (*string, bool) {
	if o == nil || isNil(o.SpareSerial) {
    return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineObject32) HasSpareSerial() bool {
	if o != nil && !isNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineObject32) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

func (o InlineObject32) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject32 struct {
	value *InlineObject32
	isSet bool
}

func (v NullableInlineObject32) Get() *InlineObject32 {
	return v.value
}

func (v *NullableInlineObject32) Set(val *InlineObject32) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject32) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject32(val *InlineObject32) *NullableInlineObject32 {
	return &NullableInlineObject32{value: val, isSet: true}
}

func (v NullableInlineObject32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


