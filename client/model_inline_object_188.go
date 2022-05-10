/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject188 struct for InlineObject188
type InlineObject188 struct {
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint,omitempty"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewInlineObject188 instantiates a new InlineObject188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject188() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// NewInlineObject188WithDefaults instantiates a new InlineObject188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject188WithDefaults() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value if set, zero value otherwise.
func (o *InlineObject188) GetX509certSha1Fingerprint() string {
	if o == nil || o.X509certSha1Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil || o.X509certSha1Fingerprint == nil {
		return nil, false
	}
	return o.X509certSha1Fingerprint, true
}

// HasX509certSha1Fingerprint returns a boolean if a field has been set.
func (o *InlineObject188) HasX509certSha1Fingerprint() bool {
	if o != nil && o.X509certSha1Fingerprint != nil {
		return true
	}

	return false
}

// SetX509certSha1Fingerprint gets a reference to the given string and assigns it to the X509certSha1Fingerprint field.
func (o *InlineObject188) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = &v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *InlineObject188) GetSloLogoutUrl() string {
	if o == nil || o.SloLogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || o.SloLogoutUrl == nil {
		return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *InlineObject188) HasSloLogoutUrl() bool {
	if o != nil && o.SloLogoutUrl != nil {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *InlineObject188) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o InlineObject188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X509certSha1Fingerprint != nil {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if o.SloLogoutUrl != nil {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject188 struct {
	value *InlineObject188
	isSet bool
}

func (v NullableInlineObject188) Get() *InlineObject188 {
	return v.value
}

func (v *NullableInlineObject188) Set(val *InlineObject188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject188(val *InlineObject188) *NullableInlineObject188 {
	return &NullableInlineObject188{value: val, isSet: true}
}

func (v NullableInlineObject188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


