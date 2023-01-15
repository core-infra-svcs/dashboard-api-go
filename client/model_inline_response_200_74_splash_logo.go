/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20074SplashLogo The logo used in the splash page.
type InlineResponse20074SplashLogo struct {
	// The MD5 value of the logo file.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the logo file.
	Extension *string `json:"extension,omitempty"`
}

// NewInlineResponse20074SplashLogo instantiates a new InlineResponse20074SplashLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20074SplashLogo() *InlineResponse20074SplashLogo {
	this := InlineResponse20074SplashLogo{}
	return &this
}

// NewInlineResponse20074SplashLogoWithDefaults instantiates a new InlineResponse20074SplashLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20074SplashLogoWithDefaults() *InlineResponse20074SplashLogo {
	this := InlineResponse20074SplashLogo{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *InlineResponse20074SplashLogo) GetMd5() string {
	if o == nil || isNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074SplashLogo) GetMd5Ok() (*string, bool) {
	if o == nil || isNil(o.Md5) {
    return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *InlineResponse20074SplashLogo) HasMd5() bool {
	if o != nil && !isNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *InlineResponse20074SplashLogo) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *InlineResponse20074SplashLogo) GetExtension() string {
	if o == nil || isNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20074SplashLogo) GetExtensionOk() (*string, bool) {
	if o == nil || isNil(o.Extension) {
    return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *InlineResponse20074SplashLogo) HasExtension() bool {
	if o != nil && !isNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *InlineResponse20074SplashLogo) SetExtension(v string) {
	o.Extension = &v
}

func (o InlineResponse20074SplashLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !isNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20074SplashLogo struct {
	value *InlineResponse20074SplashLogo
	isSet bool
}

func (v NullableInlineResponse20074SplashLogo) Get() *InlineResponse20074SplashLogo {
	return v.value
}

func (v *NullableInlineResponse20074SplashLogo) Set(val *InlineResponse20074SplashLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20074SplashLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20074SplashLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20074SplashLogo(val *InlineResponse20074SplashLogo) *NullableInlineResponse20074SplashLogo {
	return &NullableInlineResponse20074SplashLogo{value: val, isSet: true}
}

func (v NullableInlineResponse20074SplashLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20074SplashLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

