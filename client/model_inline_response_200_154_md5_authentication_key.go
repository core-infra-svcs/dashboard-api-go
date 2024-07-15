/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200154Md5AuthenticationKey MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true
type InlineResponse200154Md5AuthenticationKey struct {
	// MD5 authentication key index. Key index must be between 1 to 255
	Id *int32 `json:"id,omitempty"`
	// MD5 authentication passphrase
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewInlineResponse200154Md5AuthenticationKey instantiates a new InlineResponse200154Md5AuthenticationKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200154Md5AuthenticationKey() *InlineResponse200154Md5AuthenticationKey {
	this := InlineResponse200154Md5AuthenticationKey{}
	return &this
}

// NewInlineResponse200154Md5AuthenticationKeyWithDefaults instantiates a new InlineResponse200154Md5AuthenticationKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200154Md5AuthenticationKeyWithDefaults() *InlineResponse200154Md5AuthenticationKey {
	this := InlineResponse200154Md5AuthenticationKey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200154Md5AuthenticationKey) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154Md5AuthenticationKey) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200154Md5AuthenticationKey) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InlineResponse200154Md5AuthenticationKey) SetId(v int32) {
	o.Id = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineResponse200154Md5AuthenticationKey) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200154Md5AuthenticationKey) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineResponse200154Md5AuthenticationKey) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineResponse200154Md5AuthenticationKey) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o InlineResponse200154Md5AuthenticationKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200154Md5AuthenticationKey struct {
	value *InlineResponse200154Md5AuthenticationKey
	isSet bool
}

func (v NullableInlineResponse200154Md5AuthenticationKey) Get() *InlineResponse200154Md5AuthenticationKey {
	return v.value
}

func (v *NullableInlineResponse200154Md5AuthenticationKey) Set(val *InlineResponse200154Md5AuthenticationKey) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200154Md5AuthenticationKey) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200154Md5AuthenticationKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200154Md5AuthenticationKey(val *InlineResponse200154Md5AuthenticationKey) *NullableInlineResponse200154Md5AuthenticationKey {
	return &NullableInlineResponse200154Md5AuthenticationKey{value: val, isSet: true}
}

func (v NullableInlineResponse200154Md5AuthenticationKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200154Md5AuthenticationKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


