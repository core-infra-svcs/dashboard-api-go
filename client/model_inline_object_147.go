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

// InlineObject147 struct for InlineObject147
type InlineObject147 struct {
	// A name for easy reference to the HTTP server
	Name *string `json:"name,omitempty"`
	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	PayloadTemplate *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewInlineObject147 instantiates a new InlineObject147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject147() *InlineObject147 {
	this := InlineObject147{}
	return &this
}

// NewInlineObject147WithDefaults instantiates a new InlineObject147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject147WithDefaults() *InlineObject147 {
	this := InlineObject147{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject147) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject147) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject147) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject147) SetName(v string) {
	o.Name = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *InlineObject147) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject147) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *InlineObject147) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *InlineObject147) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *InlineObject147) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject147) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *InlineObject147) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *InlineObject147) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o InlineObject147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject147 struct {
	value *InlineObject147
	isSet bool
}

func (v NullableInlineObject147) Get() *InlineObject147 {
	return v.value
}

func (v *NullableInlineObject147) Set(val *InlineObject147) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject147) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject147(val *InlineObject147) *NullableInlineObject147 {
	return &NullableInlineObject147{value: val, isSet: true}
}

func (v NullableInlineObject147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


