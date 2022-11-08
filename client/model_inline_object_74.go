/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject74 struct for InlineObject74
type InlineObject74 struct {
	// The size of the vMX you claim. It can be one of: small, medium, large, 100
	Size string `json:"size"`
}

// NewInlineObject74 instantiates a new InlineObject74 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject74(size string) *InlineObject74 {
	this := InlineObject74{}
	this.Size = size
	return &this
}

// NewInlineObject74WithDefaults instantiates a new InlineObject74 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject74WithDefaults() *InlineObject74 {
	this := InlineObject74{}
	return &this
}

// GetSize returns the Size field value
func (o *InlineObject74) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *InlineObject74) GetSizeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *InlineObject74) SetSize(v string) {
	o.Size = v
}

func (o InlineObject74) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject74 struct {
	value *InlineObject74
	isSet bool
}

func (v NullableInlineObject74) Get() *InlineObject74 {
	return v.value
}

func (v *NullableInlineObject74) Set(val *InlineObject74) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject74) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject74) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject74(val *InlineObject74) *NullableInlineObject74 {
	return &NullableInlineObject74{value: val, isSet: true}
}

func (v NullableInlineObject74) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject74) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


