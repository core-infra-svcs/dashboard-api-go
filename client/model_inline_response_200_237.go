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

// InlineResponse200237 struct for InlineResponse200237
type InlineResponse200237 struct {
	// A list of serials of devices updated
	Serials []string `json:"serials"`
}

// NewInlineResponse200237 instantiates a new InlineResponse200237 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200237(serials []string) *InlineResponse200237 {
	this := InlineResponse200237{}
	this.Serials = serials
	return &this
}

// NewInlineResponse200237WithDefaults instantiates a new InlineResponse200237 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200237WithDefaults() *InlineResponse200237 {
	this := InlineResponse200237{}
	return &this
}

// GetSerials returns the Serials field value
func (o *InlineResponse200237) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200237) GetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *InlineResponse200237) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineResponse200237) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200237 struct {
	value *InlineResponse200237
	isSet bool
}

func (v NullableInlineResponse200237) Get() *InlineResponse200237 {
	return v.value
}

func (v *NullableInlineResponse200237) Set(val *InlineResponse200237) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200237) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200237) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200237(val *InlineResponse200237) *NullableInlineResponse200237 {
	return &NullableInlineResponse200237{value: val, isSet: true}
}

func (v NullableInlineResponse200237) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200237) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


