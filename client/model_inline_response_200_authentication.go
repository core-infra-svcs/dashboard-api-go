/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200Authentication Authentication info
type InlineResponse200Authentication struct {
	// Authentication mode
	Mode *string `json:"mode,omitempty"`
	Api *InlineResponse200AuthenticationApi `json:"api,omitempty"`
	TwoFactor *InlineResponse200AuthenticationTwoFactor `json:"twoFactor,omitempty"`
	Saml *InlineResponse200AuthenticationSaml `json:"saml,omitempty"`
}

// NewInlineResponse200Authentication instantiates a new InlineResponse200Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200Authentication() *InlineResponse200Authentication {
	this := InlineResponse200Authentication{}
	return &this
}

// NewInlineResponse200AuthenticationWithDefaults instantiates a new InlineResponse200Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200AuthenticationWithDefaults() *InlineResponse200Authentication {
	this := InlineResponse200Authentication{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse200Authentication) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200Authentication) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse200Authentication) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse200Authentication) SetMode(v string) {
	o.Mode = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *InlineResponse200Authentication) GetApi() InlineResponse200AuthenticationApi {
	if o == nil || isNil(o.Api) {
		var ret InlineResponse200AuthenticationApi
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200Authentication) GetApiOk() (*InlineResponse200AuthenticationApi, bool) {
	if o == nil || isNil(o.Api) {
    return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *InlineResponse200Authentication) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given InlineResponse200AuthenticationApi and assigns it to the Api field.
func (o *InlineResponse200Authentication) SetApi(v InlineResponse200AuthenticationApi) {
	o.Api = &v
}

// GetTwoFactor returns the TwoFactor field value if set, zero value otherwise.
func (o *InlineResponse200Authentication) GetTwoFactor() InlineResponse200AuthenticationTwoFactor {
	if o == nil || isNil(o.TwoFactor) {
		var ret InlineResponse200AuthenticationTwoFactor
		return ret
	}
	return *o.TwoFactor
}

// GetTwoFactorOk returns a tuple with the TwoFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200Authentication) GetTwoFactorOk() (*InlineResponse200AuthenticationTwoFactor, bool) {
	if o == nil || isNil(o.TwoFactor) {
    return nil, false
	}
	return o.TwoFactor, true
}

// HasTwoFactor returns a boolean if a field has been set.
func (o *InlineResponse200Authentication) HasTwoFactor() bool {
	if o != nil && !isNil(o.TwoFactor) {
		return true
	}

	return false
}

// SetTwoFactor gets a reference to the given InlineResponse200AuthenticationTwoFactor and assigns it to the TwoFactor field.
func (o *InlineResponse200Authentication) SetTwoFactor(v InlineResponse200AuthenticationTwoFactor) {
	o.TwoFactor = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *InlineResponse200Authentication) GetSaml() InlineResponse200AuthenticationSaml {
	if o == nil || isNil(o.Saml) {
		var ret InlineResponse200AuthenticationSaml
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200Authentication) GetSamlOk() (*InlineResponse200AuthenticationSaml, bool) {
	if o == nil || isNil(o.Saml) {
    return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *InlineResponse200Authentication) HasSaml() bool {
	if o != nil && !isNil(o.Saml) {
		return true
	}

	return false
}

// SetSaml gets a reference to the given InlineResponse200AuthenticationSaml and assigns it to the Saml field.
func (o *InlineResponse200Authentication) SetSaml(v InlineResponse200AuthenticationSaml) {
	o.Saml = &v
}

func (o InlineResponse200Authentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !isNil(o.TwoFactor) {
		toSerialize["twoFactor"] = o.TwoFactor
	}
	if !isNil(o.Saml) {
		toSerialize["saml"] = o.Saml
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200Authentication struct {
	value *InlineResponse200Authentication
	isSet bool
}

func (v NullableInlineResponse200Authentication) Get() *InlineResponse200Authentication {
	return v.value
}

func (v *NullableInlineResponse200Authentication) Set(val *InlineResponse200Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200Authentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200Authentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200Authentication(val *InlineResponse200Authentication) *NullableInlineResponse200Authentication {
	return &NullableInlineResponse200Authentication{value: val, isSet: true}
}

func (v NullableInlineResponse200Authentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200Authentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


