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

// InlineResponse200250 struct for InlineResponse200250
type InlineResponse200250 struct {
	// A list of serials of devices updated
	Serials []string `json:"serials"`
}

// NewInlineResponse200250 instantiates a new InlineResponse200250 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200250(serials []string) *InlineResponse200250 {
	this := InlineResponse200250{}
	this.Serials = serials
	return &this
}

// NewInlineResponse200250WithDefaults instantiates a new InlineResponse200250 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200250WithDefaults() *InlineResponse200250 {
	this := InlineResponse200250{}
	return &this
}

// GetSerials returns the Serials field value
func (o *InlineResponse200250) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200250) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineResponse200250) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse200250) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200250 struct {
	value *InlineResponse200250
	isSet bool
}

func (v NullableInlineResponse200250) Get() *InlineResponse200250 {
	return v.value
}

func (v *NullableInlineResponse200250) Set(val *InlineResponse200250) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200250) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200250) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200250(val *InlineResponse200250) *NullableInlineResponse200250 {
	return &NullableInlineResponse200250{value: val, isSet: true}
}

func (v NullableInlineResponse200250) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200250) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


