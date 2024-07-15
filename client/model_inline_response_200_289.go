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

// InlineResponse200289 struct for InlineResponse200289
type InlineResponse200289 struct {
	// Serial number of the source switch (must be on a network not bound to a template)
	SourceSerial *string `json:"sourceSerial,omitempty"`
	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials,omitempty"`
}

// NewInlineResponse200289 instantiates a new InlineResponse200289 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200289() *InlineResponse200289 {
	this := InlineResponse200289{}
	return &this
}

// NewInlineResponse200289WithDefaults instantiates a new InlineResponse200289 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200289WithDefaults() *InlineResponse200289 {
	this := InlineResponse200289{}
	return &this
}

// GetSourceSerial returns the SourceSerial field value if set, zero value otherwise.
func (o *InlineResponse200289) GetSourceSerial() string {
	if o == nil || isNil(o.SourceSerial) {
		var ret string
		return ret
	}
	return *o.SourceSerial
}

// GetSourceSerialOk returns a tuple with the SourceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetSourceSerialOk() (*string, bool) {
	if o == nil || isNil(o.SourceSerial) {
    return nil, false
	}
	return o.SourceSerial, true
}

// HasSourceSerial returns a boolean if a field has been set.
func (o *InlineResponse200289) HasSourceSerial() bool {
	if o != nil && !isNil(o.SourceSerial) {
		return true
	}

	return false
}

// SetSourceSerial gets a reference to the given string and assigns it to the SourceSerial field.
func (o *InlineResponse200289) SetSourceSerial(v string) {
	o.SourceSerial = &v
}

// GetTargetSerials returns the TargetSerials field value if set, zero value otherwise.
func (o *InlineResponse200289) GetTargetSerials() []string {
	if o == nil || isNil(o.TargetSerials) {
		var ret []string
		return ret
	}
	return o.TargetSerials
}

// GetTargetSerialsOk returns a tuple with the TargetSerials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200289) GetTargetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.TargetSerials) {
    return nil, false
	}
	return o.TargetSerials, true
}

// HasTargetSerials returns a boolean if a field has been set.
func (o *InlineResponse200289) HasTargetSerials() bool {
	if o != nil && !isNil(o.TargetSerials) {
		return true
	}

	return false
}

// SetTargetSerials gets a reference to the given []string and assigns it to the TargetSerials field.
func (o *InlineResponse200289) SetTargetSerials(v []string) {
	o.TargetSerials = v
}

func (o InlineResponse200289) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceSerial) {
		toSerialize["sourceSerial"] = o.SourceSerial
	}
	if !isNil(o.TargetSerials) {
		toSerialize["targetSerials"] = o.TargetSerials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200289 struct {
	value *InlineResponse200289
	isSet bool
}

func (v NullableInlineResponse200289) Get() *InlineResponse200289 {
	return v.value
}

func (v *NullableInlineResponse200289) Set(val *InlineResponse200289) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200289) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200289) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200289(val *InlineResponse200289) *NullableInlineResponse200289 {
	return &NullableInlineResponse200289{value: val, isSet: true}
}

func (v NullableInlineResponse200289) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200289) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

