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

// InlineObject126 struct for InlineObject126
type InlineObject126 struct {
	// The URL where the test webhook will be sent
	Url string `json:"url"`
	// The shared secret the test webhook will send. Optional. Defaults to an empty string.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// The ID of the payload template of the test webhook. Defaults to the HTTP server's template ID if one exists for the given URL, or Generic template ID otherwise
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The name of the payload template.
	PayloadTemplateName *string `json:"payloadTemplateName,omitempty"`
	// The type of alert which the test webhook will send. Optional. Defaults to power_supply_down.
	AlertTypeId *string `json:"alertTypeId,omitempty"`
}

// NewInlineObject126 instantiates a new InlineObject126 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject126(url string) *InlineObject126 {
	this := InlineObject126{}
	this.Url = url
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// NewInlineObject126WithDefaults instantiates a new InlineObject126 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject126WithDefaults() *InlineObject126 {
	this := InlineObject126{}
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// GetUrl returns the Url field value
func (o *InlineObject126) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InlineObject126) SetUrl(v string) {
	o.Url = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *InlineObject126) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *InlineObject126) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *InlineObject126) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *InlineObject126) GetPayloadTemplateId() string {
	if o == nil || o.PayloadTemplateId == nil {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || o.PayloadTemplateId == nil {
		return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *InlineObject126) HasPayloadTemplateId() bool {
	if o != nil && o.PayloadTemplateId != nil {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *InlineObject126) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetPayloadTemplateName returns the PayloadTemplateName field value if set, zero value otherwise.
func (o *InlineObject126) GetPayloadTemplateName() string {
	if o == nil || o.PayloadTemplateName == nil {
		var ret string
		return ret
	}
	return *o.PayloadTemplateName
}

// GetPayloadTemplateNameOk returns a tuple with the PayloadTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetPayloadTemplateNameOk() (*string, bool) {
	if o == nil || o.PayloadTemplateName == nil {
		return nil, false
	}
	return o.PayloadTemplateName, true
}

// HasPayloadTemplateName returns a boolean if a field has been set.
func (o *InlineObject126) HasPayloadTemplateName() bool {
	if o != nil && o.PayloadTemplateName != nil {
		return true
	}

	return false
}

// SetPayloadTemplateName gets a reference to the given string and assigns it to the PayloadTemplateName field.
func (o *InlineObject126) SetPayloadTemplateName(v string) {
	o.PayloadTemplateName = &v
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *InlineObject126) GetAlertTypeId() string {
	if o == nil || o.AlertTypeId == nil {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || o.AlertTypeId == nil {
		return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *InlineObject126) HasAlertTypeId() bool {
	if o != nil && o.AlertTypeId != nil {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *InlineObject126) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

func (o InlineObject126) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.SharedSecret != nil {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if o.PayloadTemplateId != nil {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if o.PayloadTemplateName != nil {
		toSerialize["payloadTemplateName"] = o.PayloadTemplateName
	}
	if o.AlertTypeId != nil {
		toSerialize["alertTypeId"] = o.AlertTypeId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject126 struct {
	value *InlineObject126
	isSet bool
}

func (v NullableInlineObject126) Get() *InlineObject126 {
	return v.value
}

func (v *NullableInlineObject126) Set(val *InlineObject126) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject126) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject126) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject126(val *InlineObject126) *NullableInlineObject126 {
	return &NullableInlineObject126{value: val, isSet: true}
}

func (v NullableInlineObject126) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject126) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

