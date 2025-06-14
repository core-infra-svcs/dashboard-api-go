/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200243 struct for InlineResponse200243
type InlineResponse200243 struct {
	// Category type
	Type string `json:"type"`
	// Category title
	Title string `json:"title"`
}

// NewInlineResponse200243 instantiates a new InlineResponse200243 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200243(type_ string, title string) *InlineResponse200243 {
	this := InlineResponse200243{}
	this.Type = type_
	this.Title = title
	return &this
}

// NewInlineResponse200243WithDefaults instantiates a new InlineResponse200243 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200243WithDefaults() *InlineResponse200243 {
	this := InlineResponse200243{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200243) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200243) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200243) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *InlineResponse200243) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200243) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineResponse200243) SetTitle(v string) {
	o.Title = v
}

func (o InlineResponse200243) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200243 struct {
	value *InlineResponse200243
	isSet bool
}

func (v NullableInlineResponse200243) Get() *InlineResponse200243 {
	return v.value
}

func (v *NullableInlineResponse200243) Set(val *InlineResponse200243) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200243) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200243) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200243(val *InlineResponse200243) *NullableInlineResponse200243 {
	return &NullableInlineResponse200243{value: val, isSet: true}
}

func (v NullableInlineResponse200243) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200243) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


