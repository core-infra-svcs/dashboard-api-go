/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject20 struct for InlineObject20
type InlineObject20 struct {
	// Enable or disable warm spare for a switch
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare switch
	SpareSerial *string `json:"spareSerial,omitempty"`
}

// NewInlineObject20 instantiates a new InlineObject20 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject20(enabled bool) *InlineObject20 {
	this := InlineObject20{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject20WithDefaults instantiates a new InlineObject20 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject20WithDefaults() *InlineObject20 {
	this := InlineObject20{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject20) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject20) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineObject20) GetSpareSerial() string {
	if o == nil || o.SpareSerial == nil {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject20) GetSpareSerialOk() (*string, bool) {
	if o == nil || o.SpareSerial == nil {
		return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineObject20) HasSpareSerial() bool {
	if o != nil && o.SpareSerial != nil {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineObject20) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

func (o InlineObject20) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.SpareSerial != nil {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject20 struct {
	value *InlineObject20
	isSet bool
}

func (v NullableInlineObject20) Get() *InlineObject20 {
	return v.value
}

func (v *NullableInlineObject20) Set(val *InlineObject20) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject20) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject20) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject20(val *InlineObject20) *NullableInlineObject20 {
	return &NullableInlineObject20{value: val, isSet: true}
}

func (v NullableInlineObject20) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject20) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


