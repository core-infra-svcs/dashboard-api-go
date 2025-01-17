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

// InlineResponse200287 struct for InlineResponse200287
type InlineResponse200287 struct {
	// Organization APNS Certificate used by devices to communication with Apple
	Certificate *string `json:"certificate,omitempty"`
}

// NewInlineResponse200287 instantiates a new InlineResponse200287 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200287() *InlineResponse200287 {
	this := InlineResponse200287{}
	return &this
}

// NewInlineResponse200287WithDefaults instantiates a new InlineResponse200287 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200287WithDefaults() *InlineResponse200287 {
	this := InlineResponse200287{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *InlineResponse200287) GetCertificate() string {
	if o == nil || isNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetCertificateOk() (*string, bool) {
	if o == nil || isNil(o.Certificate) {
    return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *InlineResponse200287) HasCertificate() bool {
	if o != nil && !isNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *InlineResponse200287) SetCertificate(v string) {
	o.Certificate = &v
}

func (o InlineResponse200287) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200287 struct {
	value *InlineResponse200287
	isSet bool
}

func (v NullableInlineResponse200287) Get() *InlineResponse200287 {
	return v.value
}

func (v *NullableInlineResponse200287) Set(val *InlineResponse200287) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200287) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200287) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200287(val *InlineResponse200287) *NullableInlineResponse200287 {
	return &NullableInlineResponse200287{value: val, isSet: true}
}

func (v NullableInlineResponse200287) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200287) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


