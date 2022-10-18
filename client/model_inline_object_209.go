/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject209 struct for InlineObject209
type InlineObject209 struct {
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint string `json:"x509certSha1Fingerprint"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewInlineObject209 instantiates a new InlineObject209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject209(x509certSha1Fingerprint string) *InlineObject209 {
	this := InlineObject209{}
	this.X509certSha1Fingerprint = x509certSha1Fingerprint
	return &this
}

// NewInlineObject209WithDefaults instantiates a new InlineObject209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject209WithDefaults() *InlineObject209 {
	this := InlineObject209{}
	return &this
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value
func (o *InlineObject209) GetX509certSha1Fingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.X509certSha1Fingerprint, true
}

// SetX509certSha1Fingerprint sets field value
func (o *InlineObject209) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *InlineObject209) GetSloLogoutUrl() string {
	if o == nil || o.SloLogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || o.SloLogoutUrl == nil {
		return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *InlineObject209) HasSloLogoutUrl() bool {
	if o != nil && o.SloLogoutUrl != nil {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *InlineObject209) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o InlineObject209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if o.SloLogoutUrl != nil {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject209 struct {
	value *InlineObject209
	isSet bool
}

func (v NullableInlineObject209) Get() *InlineObject209 {
	return v.value
}

func (v *NullableInlineObject209) Set(val *InlineObject209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject209(val *InlineObject209) *NullableInlineObject209 {
	return &NullableInlineObject209{value: val, isSet: true}
}

func (v NullableInlineObject209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


