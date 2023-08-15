/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200AuthenticationApiKey API key
type InlineResponse200AuthenticationApiKey struct {
	// If API key is created for this user
	Created *bool `json:"created,omitempty"`
}

// NewInlineResponse200AuthenticationApiKey instantiates a new InlineResponse200AuthenticationApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200AuthenticationApiKey() *InlineResponse200AuthenticationApiKey {
	this := InlineResponse200AuthenticationApiKey{}
	return &this
}

// NewInlineResponse200AuthenticationApiKeyWithDefaults instantiates a new InlineResponse200AuthenticationApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200AuthenticationApiKeyWithDefaults() *InlineResponse200AuthenticationApiKey {
	this := InlineResponse200AuthenticationApiKey{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *InlineResponse200AuthenticationApiKey) GetCreated() bool {
	if o == nil || isNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200AuthenticationApiKey) GetCreatedOk() (*bool, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *InlineResponse200AuthenticationApiKey) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *InlineResponse200AuthenticationApiKey) SetCreated(v bool) {
	o.Created = &v
}

func (o InlineResponse200AuthenticationApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200AuthenticationApiKey struct {
	value *InlineResponse200AuthenticationApiKey
	isSet bool
}

func (v NullableInlineResponse200AuthenticationApiKey) Get() *InlineResponse200AuthenticationApiKey {
	return v.value
}

func (v *NullableInlineResponse200AuthenticationApiKey) Set(val *InlineResponse200AuthenticationApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200AuthenticationApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200AuthenticationApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200AuthenticationApiKey(val *InlineResponse200AuthenticationApiKey) *NullableInlineResponse200AuthenticationApiKey {
	return &NullableInlineResponse200AuthenticationApiKey{value: val, isSet: true}
}

func (v NullableInlineResponse200AuthenticationApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200AuthenticationApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


