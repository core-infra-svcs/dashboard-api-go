/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200159 struct for InlineResponse200159
type InlineResponse200159 struct {
	// The id of the VPP Account
	Id *string `json:"id,omitempty"`
	// The VPP Account's Service Token
	VppServiceToken *string `json:"vppServiceToken,omitempty"`
}

// NewInlineResponse200159 instantiates a new InlineResponse200159 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200159() *InlineResponse200159 {
	this := InlineResponse200159{}
	return &this
}

// NewInlineResponse200159WithDefaults instantiates a new InlineResponse200159 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200159WithDefaults() *InlineResponse200159 {
	this := InlineResponse200159{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200159) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200159) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200159) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200159) SetId(v string) {
	o.Id = &v
}

// GetVppServiceToken returns the VppServiceToken field value if set, zero value otherwise.
func (o *InlineResponse200159) GetVppServiceToken() string {
	if o == nil || isNil(o.VppServiceToken) {
		var ret string
		return ret
	}
	return *o.VppServiceToken
}

// GetVppServiceTokenOk returns a tuple with the VppServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200159) GetVppServiceTokenOk() (*string, bool) {
	if o == nil || isNil(o.VppServiceToken) {
    return nil, false
	}
	return o.VppServiceToken, true
}

// HasVppServiceToken returns a boolean if a field has been set.
func (o *InlineResponse200159) HasVppServiceToken() bool {
	if o != nil && !isNil(o.VppServiceToken) {
		return true
	}

	return false
}

// SetVppServiceToken gets a reference to the given string and assigns it to the VppServiceToken field.
func (o *InlineResponse200159) SetVppServiceToken(v string) {
	o.VppServiceToken = &v
}

func (o InlineResponse200159) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.VppServiceToken) {
		toSerialize["vppServiceToken"] = o.VppServiceToken
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200159 struct {
	value *InlineResponse200159
	isSet bool
}

func (v NullableInlineResponse200159) Get() *InlineResponse200159 {
	return v.value
}

func (v *NullableInlineResponse200159) Set(val *InlineResponse200159) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200159) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200159) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200159(val *InlineResponse200159) *NullableInlineResponse200159 {
	return &NullableInlineResponse200159{value: val, isSet: true}
}

func (v NullableInlineResponse200159) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200159) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


