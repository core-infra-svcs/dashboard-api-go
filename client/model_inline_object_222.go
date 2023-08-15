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

// InlineObject222 struct for InlineObject222
type InlineObject222 struct {
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint,omitempty"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewInlineObject222 instantiates a new InlineObject222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject222() *InlineObject222 {
	this := InlineObject222{}
	return &this
}

// NewInlineObject222WithDefaults instantiates a new InlineObject222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject222WithDefaults() *InlineObject222 {
	this := InlineObject222{}
	return &this
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value if set, zero value otherwise.
func (o *InlineObject222) GetX509certSha1Fingerprint() string {
	if o == nil || isNil(o.X509certSha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject222) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil || isNil(o.X509certSha1Fingerprint) {
    return nil, false
	}
	return o.X509certSha1Fingerprint, true
}

// HasX509certSha1Fingerprint returns a boolean if a field has been set.
func (o *InlineObject222) HasX509certSha1Fingerprint() bool {
	if o != nil && !isNil(o.X509certSha1Fingerprint) {
		return true
	}

	return false
}

// SetX509certSha1Fingerprint gets a reference to the given string and assigns it to the X509certSha1Fingerprint field.
func (o *InlineObject222) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = &v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *InlineObject222) GetSloLogoutUrl() string {
	if o == nil || isNil(o.SloLogoutUrl) {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject222) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || isNil(o.SloLogoutUrl) {
    return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *InlineObject222) HasSloLogoutUrl() bool {
	if o != nil && !isNil(o.SloLogoutUrl) {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *InlineObject222) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o InlineObject222) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.X509certSha1Fingerprint) {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if !isNil(o.SloLogoutUrl) {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject222 struct {
	value *InlineObject222
	isSet bool
}

func (v NullableInlineObject222) Get() *InlineObject222 {
	return v.value
}

func (v *NullableInlineObject222) Set(val *InlineObject222) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject222) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject222) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject222(val *InlineObject222) *NullableInlineObject222 {
	return &NullableInlineObject222{value: val, isSet: true}
}

func (v NullableInlineObject222) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject222) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


