/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20068Authentication Authentication settings between BGP peers
type InlineResponse20068Authentication struct {
	// Password to configure MD5 authentication between BGP peers
	Password *string `json:"password,omitempty"`
}

// NewInlineResponse20068Authentication instantiates a new InlineResponse20068Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20068Authentication() *InlineResponse20068Authentication {
	this := InlineResponse20068Authentication{}
	return &this
}

// NewInlineResponse20068AuthenticationWithDefaults instantiates a new InlineResponse20068Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20068AuthenticationWithDefaults() *InlineResponse20068Authentication {
	this := InlineResponse20068Authentication{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InlineResponse20068Authentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068Authentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineResponse20068Authentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineResponse20068Authentication) SetPassword(v string) {
	o.Password = &v
}

func (o InlineResponse20068Authentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20068Authentication struct {
	value *InlineResponse20068Authentication
	isSet bool
}

func (v NullableInlineResponse20068Authentication) Get() *InlineResponse20068Authentication {
	return v.value
}

func (v *NullableInlineResponse20068Authentication) Set(val *InlineResponse20068Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20068Authentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20068Authentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20068Authentication(val *InlineResponse20068Authentication) *NullableInlineResponse20068Authentication {
	return &NullableInlineResponse20068Authentication{value: val, isSet: true}
}

func (v NullableInlineResponse20068Authentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20068Authentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


