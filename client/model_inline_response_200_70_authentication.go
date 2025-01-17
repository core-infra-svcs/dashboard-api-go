/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20070Authentication Authentication settings between BGP peers
type InlineResponse20070Authentication struct {
	// Password to configure MD5 authentication between BGP peers
	Password *string `json:"password,omitempty"`
}

// NewInlineResponse20070Authentication instantiates a new InlineResponse20070Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070Authentication() *InlineResponse20070Authentication {
	this := InlineResponse20070Authentication{}
	return &this
}

// NewInlineResponse20070AuthenticationWithDefaults instantiates a new InlineResponse20070Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070AuthenticationWithDefaults() *InlineResponse20070Authentication {
	this := InlineResponse20070Authentication{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InlineResponse20070Authentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070Authentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineResponse20070Authentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineResponse20070Authentication) SetPassword(v string) {
	o.Password = &v
}

func (o InlineResponse20070Authentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20070Authentication struct {
	value *InlineResponse20070Authentication
	isSet bool
}

func (v NullableInlineResponse20070Authentication) Get() *InlineResponse20070Authentication {
	return v.value
}

func (v *NullableInlineResponse20070Authentication) Set(val *InlineResponse20070Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070Authentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070Authentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070Authentication(val *InlineResponse20070Authentication) *NullableInlineResponse20070Authentication {
	return &NullableInlineResponse20070Authentication{value: val, isSet: true}
}

func (v NullableInlineResponse20070Authentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070Authentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


