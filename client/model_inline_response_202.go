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

// InlineResponse202 struct for InlineResponse202
type InlineResponse202 struct {
	// API key in plaintext. This value will not be accessible outside of key generation
	Key *string `json:"key,omitempty"`
}

// NewInlineResponse202 instantiates a new InlineResponse202 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse202() *InlineResponse202 {
	this := InlineResponse202{}
	return &this
}

// NewInlineResponse202WithDefaults instantiates a new InlineResponse202 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse202WithDefaults() *InlineResponse202 {
	this := InlineResponse202{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InlineResponse202) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse202) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineResponse202) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InlineResponse202) SetKey(v string) {
	o.Key = &v
}

func (o InlineResponse202) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse202 struct {
	value *InlineResponse202
	isSet bool
}

func (v NullableInlineResponse202) Get() *InlineResponse202 {
	return v.value
}

func (v *NullableInlineResponse202) Set(val *InlineResponse202) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse202) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse202) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse202(val *InlineResponse202) *NullableInlineResponse202 {
	return &NullableInlineResponse202{value: val, isSet: true}
}

func (v NullableInlineResponse202) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse202) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


