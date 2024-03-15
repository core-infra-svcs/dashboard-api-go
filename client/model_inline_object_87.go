/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject87 struct for InlineObject87
type InlineObject87 struct {
	// The size of the vMX you claim. It can be one of: small, medium, large, xlarge, 100
	Size string `json:"size"`
}

// NewInlineObject87 instantiates a new InlineObject87 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject87(size string) *InlineObject87 {
	this := InlineObject87{}
	this.Size = size
	return &this
}

// NewInlineObject87WithDefaults instantiates a new InlineObject87 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject87WithDefaults() *InlineObject87 {
	this := InlineObject87{}
	return &this
}

// GetSize returns the Size field value
func (o *InlineObject87) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *InlineObject87) GetSizeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *InlineObject87) SetSize(v string) {
	o.Size = v
}

func (o InlineObject87) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject87 struct {
	value *InlineObject87
	isSet bool
}

func (v NullableInlineObject87) Get() *InlineObject87 {
	return v.value
}

func (v *NullableInlineObject87) Set(val *InlineObject87) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject87) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject87) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject87(val *InlineObject87) *NullableInlineObject87 {
	return &NullableInlineObject87{value: val, isSet: true}
}

func (v NullableInlineObject87) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject87) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


