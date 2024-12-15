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

// InlineObject170 struct for InlineObject170
type InlineObject170 struct {
	// A name for easy reference to the HTTP server
	Name string `json:"name"`
	// The URL of the HTTP server. Once set, cannot be updated.
	Url string `json:"url"`
	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	PayloadTemplate *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 `json:"payloadTemplate,omitempty"`
}

// NewInlineObject170 instantiates a new InlineObject170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject170(name string, url string) *InlineObject170 {
	this := InlineObject170{}
	this.Name = name
	this.Url = url
	return &this
}

// NewInlineObject170WithDefaults instantiates a new InlineObject170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject170WithDefaults() *InlineObject170 {
	this := InlineObject170{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject170) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject170) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *InlineObject170) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InlineObject170) SetUrl(v string) {
	o.Url = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *InlineObject170) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *InlineObject170) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *InlineObject170) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *InlineObject170) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret NetworksNetworkIdWebhooksHttpServersPayloadTemplate1
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersPayloadTemplate1, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *InlineObject170) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 and assigns it to the PayloadTemplate field.
func (o *InlineObject170) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) {
	o.PayloadTemplate = &v
}

func (o InlineObject170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject170 struct {
	value *InlineObject170
	isSet bool
}

func (v NullableInlineObject170) Get() *InlineObject170 {
	return v.value
}

func (v *NullableInlineObject170) Set(val *InlineObject170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject170(val *InlineObject170) *NullableInlineObject170 {
	return &NullableInlineObject170{value: val, isSet: true}
}

func (v NullableInlineObject170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


