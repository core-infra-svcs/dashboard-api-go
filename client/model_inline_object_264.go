/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject264 struct for InlineObject264
type InlineObject264 struct {
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint,omitempty"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewInlineObject264 instantiates a new InlineObject264 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject264() *InlineObject264 {
	this := InlineObject264{}
	return &this
}

// NewInlineObject264WithDefaults instantiates a new InlineObject264 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject264WithDefaults() *InlineObject264 {
	this := InlineObject264{}
	return &this
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value if set, zero value otherwise.
func (o *InlineObject264) GetX509certSha1Fingerprint() string {
	if o == nil || isNil(o.X509certSha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject264) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil || isNil(o.X509certSha1Fingerprint) {
    return nil, false
	}
	return o.X509certSha1Fingerprint, true
}

// HasX509certSha1Fingerprint returns a boolean if a field has been set.
func (o *InlineObject264) HasX509certSha1Fingerprint() bool {
	if o != nil && !isNil(o.X509certSha1Fingerprint) {
		return true
	}

	return false
}

// SetX509certSha1Fingerprint gets a reference to the given string and assigns it to the X509certSha1Fingerprint field.
func (o *InlineObject264) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = &v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *InlineObject264) GetSloLogoutUrl() string {
	if o == nil || isNil(o.SloLogoutUrl) {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject264) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || isNil(o.SloLogoutUrl) {
    return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *InlineObject264) HasSloLogoutUrl() bool {
	if o != nil && !isNil(o.SloLogoutUrl) {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *InlineObject264) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o InlineObject264) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.X509certSha1Fingerprint) {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if !isNil(o.SloLogoutUrl) {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject264 struct {
	value *InlineObject264
	isSet bool
}

func (v NullableInlineObject264) Get() *InlineObject264 {
	return v.value
}

func (v *NullableInlineObject264) Set(val *InlineObject264) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject264) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject264) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject264(val *InlineObject264) *NullableInlineObject264 {
	return &NullableInlineObject264{value: val, isSet: true}
}

func (v NullableInlineObject264) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject264) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


