/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200265 struct for InlineResponse200265
type InlineResponse200265 struct {
	// ID associated with the SAML Identity Provider (IdP)
	IdpId *string `json:"idpId,omitempty"`
	// URL that is consuming SAML Identity Provider (IdP)
	ConsumerUrl *string `json:"consumerUrl,omitempty"`
	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint,omitempty"`
	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutUrl *string `json:"sloLogoutUrl,omitempty"`
}

// NewInlineResponse200265 instantiates a new InlineResponse200265 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200265() *InlineResponse200265 {
	this := InlineResponse200265{}
	return &this
}

// NewInlineResponse200265WithDefaults instantiates a new InlineResponse200265 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200265WithDefaults() *InlineResponse200265 {
	this := InlineResponse200265{}
	return &this
}

// GetIdpId returns the IdpId field value if set, zero value otherwise.
func (o *InlineResponse200265) GetIdpId() string {
	if o == nil || isNil(o.IdpId) {
		var ret string
		return ret
	}
	return *o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200265) GetIdpIdOk() (*string, bool) {
	if o == nil || isNil(o.IdpId) {
    return nil, false
	}
	return o.IdpId, true
}

// HasIdpId returns a boolean if a field has been set.
func (o *InlineResponse200265) HasIdpId() bool {
	if o != nil && !isNil(o.IdpId) {
		return true
	}

	return false
}

// SetIdpId gets a reference to the given string and assigns it to the IdpId field.
func (o *InlineResponse200265) SetIdpId(v string) {
	o.IdpId = &v
}

// GetConsumerUrl returns the ConsumerUrl field value if set, zero value otherwise.
func (o *InlineResponse200265) GetConsumerUrl() string {
	if o == nil || isNil(o.ConsumerUrl) {
		var ret string
		return ret
	}
	return *o.ConsumerUrl
}

// GetConsumerUrlOk returns a tuple with the ConsumerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200265) GetConsumerUrlOk() (*string, bool) {
	if o == nil || isNil(o.ConsumerUrl) {
    return nil, false
	}
	return o.ConsumerUrl, true
}

// HasConsumerUrl returns a boolean if a field has been set.
func (o *InlineResponse200265) HasConsumerUrl() bool {
	if o != nil && !isNil(o.ConsumerUrl) {
		return true
	}

	return false
}

// SetConsumerUrl gets a reference to the given string and assigns it to the ConsumerUrl field.
func (o *InlineResponse200265) SetConsumerUrl(v string) {
	o.ConsumerUrl = &v
}

// GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field value if set, zero value otherwise.
func (o *InlineResponse200265) GetX509certSha1Fingerprint() string {
	if o == nil || isNil(o.X509certSha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.X509certSha1Fingerprint
}

// GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200265) GetX509certSha1FingerprintOk() (*string, bool) {
	if o == nil || isNil(o.X509certSha1Fingerprint) {
    return nil, false
	}
	return o.X509certSha1Fingerprint, true
}

// HasX509certSha1Fingerprint returns a boolean if a field has been set.
func (o *InlineResponse200265) HasX509certSha1Fingerprint() bool {
	if o != nil && !isNil(o.X509certSha1Fingerprint) {
		return true
	}

	return false
}

// SetX509certSha1Fingerprint gets a reference to the given string and assigns it to the X509certSha1Fingerprint field.
func (o *InlineResponse200265) SetX509certSha1Fingerprint(v string) {
	o.X509certSha1Fingerprint = &v
}

// GetSloLogoutUrl returns the SloLogoutUrl field value if set, zero value otherwise.
func (o *InlineResponse200265) GetSloLogoutUrl() string {
	if o == nil || isNil(o.SloLogoutUrl) {
		var ret string
		return ret
	}
	return *o.SloLogoutUrl
}

// GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200265) GetSloLogoutUrlOk() (*string, bool) {
	if o == nil || isNil(o.SloLogoutUrl) {
    return nil, false
	}
	return o.SloLogoutUrl, true
}

// HasSloLogoutUrl returns a boolean if a field has been set.
func (o *InlineResponse200265) HasSloLogoutUrl() bool {
	if o != nil && !isNil(o.SloLogoutUrl) {
		return true
	}

	return false
}

// SetSloLogoutUrl gets a reference to the given string and assigns it to the SloLogoutUrl field.
func (o *InlineResponse200265) SetSloLogoutUrl(v string) {
	o.SloLogoutUrl = &v
}

func (o InlineResponse200265) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdpId) {
		toSerialize["idpId"] = o.IdpId
	}
	if !isNil(o.ConsumerUrl) {
		toSerialize["consumerUrl"] = o.ConsumerUrl
	}
	if !isNil(o.X509certSha1Fingerprint) {
		toSerialize["x509certSha1Fingerprint"] = o.X509certSha1Fingerprint
	}
	if !isNil(o.SloLogoutUrl) {
		toSerialize["sloLogoutUrl"] = o.SloLogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200265 struct {
	value *InlineResponse200265
	isSet bool
}

func (v NullableInlineResponse200265) Get() *InlineResponse200265 {
	return v.value
}

func (v *NullableInlineResponse200265) Set(val *InlineResponse200265) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200265) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200265) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200265(val *InlineResponse200265) *NullableInlineResponse200265 {
	return &NullableInlineResponse200265{value: val, isSet: true}
}

func (v NullableInlineResponse200265) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200265) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


