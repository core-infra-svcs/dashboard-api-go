/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject78 struct for InlineObject78
type InlineObject78 struct {
	// The size of the vMX you claim. It can be one of: small, medium, large, 100
	Size string `json:"size"`
}

// NewInlineObject78 instantiates a new InlineObject78 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject78(size string) *InlineObject78 {
	this := InlineObject78{}
	this.Size = size
	return &this
}

// NewInlineObject78WithDefaults instantiates a new InlineObject78 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject78WithDefaults() *InlineObject78 {
	this := InlineObject78{}
	return &this
}

// GetSize returns the Size field value
func (o *InlineObject78) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *InlineObject78) GetSizeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *InlineObject78) SetSize(v string) {
	o.Size = v
}

func (o InlineObject78) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject78 struct {
	value *InlineObject78
	isSet bool
}

func (v NullableInlineObject78) Get() *InlineObject78 {
	return v.value
}

func (v *NullableInlineObject78) Set(val *InlineObject78) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject78) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject78) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject78(val *InlineObject78) *NullableInlineObject78 {
	return &NullableInlineObject78{value: val, isSet: true}
}

func (v NullableInlineObject78) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject78) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


