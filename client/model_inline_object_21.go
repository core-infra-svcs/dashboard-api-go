/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject21 struct for InlineObject21
type InlineObject21 struct {
	// Enable or disable warm spare for a switch
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare switch
	SpareSerial *string `json:"spareSerial,omitempty"`
}

// NewInlineObject21 instantiates a new InlineObject21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject21(enabled bool) *InlineObject21 {
	this := InlineObject21{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject21WithDefaults instantiates a new InlineObject21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject21WithDefaults() *InlineObject21 {
	this := InlineObject21{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject21) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject21) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineObject21) GetSpareSerial() string {
	if o == nil || o.SpareSerial == nil {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject21) GetSpareSerialOk() (*string, bool) {
	if o == nil || o.SpareSerial == nil {
		return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineObject21) HasSpareSerial() bool {
	if o != nil && o.SpareSerial != nil {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineObject21) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

func (o InlineObject21) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.SpareSerial != nil {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject21 struct {
	value *InlineObject21
	isSet bool
}

func (v NullableInlineObject21) Get() *InlineObject21 {
	return v.value
}

func (v *NullableInlineObject21) Set(val *InlineObject21) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject21) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject21(val *InlineObject21) *NullableInlineObject21 {
	return &NullableInlineObject21{value: val, isSet: true}
}

func (v NullableInlineObject21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


