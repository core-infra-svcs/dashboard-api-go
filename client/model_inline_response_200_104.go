/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200104 struct for InlineResponse200104
type InlineResponse200104 struct {
	// Organization APNS Certificate used by devices to communication with Apple
	Certificate *string `json:"certificate,omitempty"`
}

// NewInlineResponse200104 instantiates a new InlineResponse200104 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200104() *InlineResponse200104 {
	this := InlineResponse200104{}
	return &this
}

// NewInlineResponse200104WithDefaults instantiates a new InlineResponse200104 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200104WithDefaults() *InlineResponse200104 {
	this := InlineResponse200104{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *InlineResponse200104) GetCertificate() string {
	if o == nil || isNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200104) GetCertificateOk() (*string, bool) {
	if o == nil || isNil(o.Certificate) {
    return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *InlineResponse200104) HasCertificate() bool {
	if o != nil && !isNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *InlineResponse200104) SetCertificate(v string) {
	o.Certificate = &v
}

func (o InlineResponse200104) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200104 struct {
	value *InlineResponse200104
	isSet bool
}

func (v NullableInlineResponse200104) Get() *InlineResponse200104 {
	return v.value
}

func (v *NullableInlineResponse200104) Set(val *InlineResponse200104) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200104) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200104) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200104(val *InlineResponse200104) *NullableInlineResponse200104 {
	return &NullableInlineResponse200104{value: val, isSet: true}
}

func (v NullableInlineResponse200104) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200104) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


