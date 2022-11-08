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

// InlineResponse200AuthenticationApi API authentication
type InlineResponse200AuthenticationApi struct {
	Key *InlineResponse200AuthenticationApiKey `json:"key,omitempty"`
}

// NewInlineResponse200AuthenticationApi instantiates a new InlineResponse200AuthenticationApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200AuthenticationApi() *InlineResponse200AuthenticationApi {
	this := InlineResponse200AuthenticationApi{}
	return &this
}

// NewInlineResponse200AuthenticationApiWithDefaults instantiates a new InlineResponse200AuthenticationApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200AuthenticationApiWithDefaults() *InlineResponse200AuthenticationApi {
	this := InlineResponse200AuthenticationApi{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InlineResponse200AuthenticationApi) GetKey() InlineResponse200AuthenticationApiKey {
	if o == nil || isNil(o.Key) {
		var ret InlineResponse200AuthenticationApiKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200AuthenticationApi) GetKeyOk() (*InlineResponse200AuthenticationApiKey, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineResponse200AuthenticationApi) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given InlineResponse200AuthenticationApiKey and assigns it to the Key field.
func (o *InlineResponse200AuthenticationApi) SetKey(v InlineResponse200AuthenticationApiKey) {
	o.Key = &v
}

func (o InlineResponse200AuthenticationApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200AuthenticationApi struct {
	value *InlineResponse200AuthenticationApi
	isSet bool
}

func (v NullableInlineResponse200AuthenticationApi) Get() *InlineResponse200AuthenticationApi {
	return v.value
}

func (v *NullableInlineResponse200AuthenticationApi) Set(val *InlineResponse200AuthenticationApi) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200AuthenticationApi) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200AuthenticationApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200AuthenticationApi(val *InlineResponse200AuthenticationApi) *NullableInlineResponse200AuthenticationApi {
	return &NullableInlineResponse200AuthenticationApi{value: val, isSet: true}
}

func (v NullableInlineResponse200AuthenticationApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200AuthenticationApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

