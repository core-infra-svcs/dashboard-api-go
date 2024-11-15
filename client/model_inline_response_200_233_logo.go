/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200233Logo Service Provider logo data.
type InlineResponse200233Logo struct {
	// URL of service provider's logo.
	Url *string `json:"url,omitempty"`
}

// NewInlineResponse200233Logo instantiates a new InlineResponse200233Logo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200233Logo() *InlineResponse200233Logo {
	this := InlineResponse200233Logo{}
	return &this
}

// NewInlineResponse200233LogoWithDefaults instantiates a new InlineResponse200233Logo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200233LogoWithDefaults() *InlineResponse200233Logo {
	this := InlineResponse200233Logo{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200233Logo) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200233Logo) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200233Logo) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200233Logo) SetUrl(v string) {
	o.Url = &v
}

func (o InlineResponse200233Logo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200233Logo struct {
	value *InlineResponse200233Logo
	isSet bool
}

func (v NullableInlineResponse200233Logo) Get() *InlineResponse200233Logo {
	return v.value
}

func (v *NullableInlineResponse200233Logo) Set(val *InlineResponse200233Logo) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200233Logo) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200233Logo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200233Logo(val *InlineResponse200233Logo) *NullableInlineResponse200233Logo {
	return &NullableInlineResponse200233Logo{value: val, isSet: true}
}

func (v NullableInlineResponse200233Logo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200233Logo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

