/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject285 struct for InlineObject285
type InlineObject285 struct {
	// Serial number of the source switch (must be on a network not bound to a template)
	SourceSerial string `json:"sourceSerial"`
	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials"`
}

// NewInlineObject285 instantiates a new InlineObject285 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject285(sourceSerial string, targetSerials []string) *InlineObject285 {
	this := InlineObject285{}
	this.SourceSerial = sourceSerial
	this.TargetSerials = targetSerials
	return &this
}

// NewInlineObject285WithDefaults instantiates a new InlineObject285 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject285WithDefaults() *InlineObject285 {
	this := InlineObject285{}
	return &this
}

// GetSourceSerial returns the SourceSerial field value
func (o *InlineObject285) GetSourceSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceSerial
}

// GetSourceSerialOk returns a tuple with the SourceSerial field value
// and a boolean to check if the value has been set.
func (o *InlineObject285) GetSourceSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceSerial, true
}

// SetSourceSerial sets field value
func (o *InlineObject285) SetSourceSerial(v string) {
	o.SourceSerial = v
}

// GetTargetSerials returns the TargetSerials field value
func (o *InlineObject285) GetTargetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetSerials
}

// GetTargetSerialsOk returns a tuple with the TargetSerials field value
// and a boolean to check if the value has been set.
func (o *InlineObject285) GetTargetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TargetSerials, true
}

// SetTargetSerials sets field value
func (o *InlineObject285) SetTargetSerials(v []string) {
	o.TargetSerials = v
}

func (o InlineObject285) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceSerial"] = o.SourceSerial
	}
	if true {
		toSerialize["targetSerials"] = o.TargetSerials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject285 struct {
	value *InlineObject285
	isSet bool
}

func (v NullableInlineObject285) Get() *InlineObject285 {
	return v.value
}

func (v *NullableInlineObject285) Set(val *InlineObject285) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject285) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject285) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject285(val *InlineObject285) *NullableInlineObject285 {
	return &NullableInlineObject285{value: val, isSet: true}
}

func (v NullableInlineObject285) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject285) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


