/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject207 struct for InlineObject207
type InlineObject207 struct {
	// The name of the new organization
	Name string `json:"name"`
}

// NewInlineObject207 instantiates a new InlineObject207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject207(name string) *InlineObject207 {
	this := InlineObject207{}
	this.Name = name
	return &this
}

// NewInlineObject207WithDefaults instantiates a new InlineObject207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject207WithDefaults() *InlineObject207 {
	this := InlineObject207{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject207) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject207) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject207) SetName(v string) {
	o.Name = v
}

func (o InlineObject207) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject207 struct {
	value *InlineObject207
	isSet bool
}

func (v NullableInlineObject207) Get() *InlineObject207 {
	return v.value
}

func (v *NullableInlineObject207) Set(val *InlineObject207) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject207) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject207) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject207(val *InlineObject207) *NullableInlineObject207 {
	return &NullableInlineObject207{value: val, isSet: true}
}

func (v NullableInlineObject207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


