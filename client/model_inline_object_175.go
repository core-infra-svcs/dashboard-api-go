/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject175 struct for InlineObject175
type InlineObject175 struct {
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

// NewInlineObject175 instantiates a new InlineObject175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject175(url string) *InlineObject175 {
	this := InlineObject175{}
	this.Url = url
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// NewInlineObject175WithDefaults instantiates a new InlineObject175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject175WithDefaults() *InlineObject175 {
	this := InlineObject175{}
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// GetUrl returns the Url field value
func (o *InlineObject175) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InlineObject175) SetUrl(v string) {
	o.Url = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *InlineObject175) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *InlineObject175) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *InlineObject175) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *InlineObject175) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *InlineObject175) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *InlineObject175) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetPayloadTemplateName returns the PayloadTemplateName field value if set, zero value otherwise.
func (o *InlineObject175) GetPayloadTemplateName() string {
	if o == nil || isNil(o.PayloadTemplateName) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateName
}

// GetPayloadTemplateNameOk returns a tuple with the PayloadTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetPayloadTemplateNameOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateName) {
    return nil, false
	}
	return o.PayloadTemplateName, true
}

// HasPayloadTemplateName returns a boolean if a field has been set.
func (o *InlineObject175) HasPayloadTemplateName() bool {
	if o != nil && !isNil(o.PayloadTemplateName) {
		return true
	}

	return false
}

// SetPayloadTemplateName gets a reference to the given string and assigns it to the PayloadTemplateName field.
func (o *InlineObject175) SetPayloadTemplateName(v string) {
	o.PayloadTemplateName = &v
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *InlineObject175) GetAlertTypeId() string {
	if o == nil || isNil(o.AlertTypeId) {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject175) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || isNil(o.AlertTypeId) {
    return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *InlineObject175) HasAlertTypeId() bool {
	if o != nil && !isNil(o.AlertTypeId) {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *InlineObject175) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

func (o InlineObject175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !isNil(o.PayloadTemplateName) {
		toSerialize["payloadTemplateName"] = o.PayloadTemplateName
	}
	if !isNil(o.AlertTypeId) {
		toSerialize["alertTypeId"] = o.AlertTypeId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject175 struct {
	value *InlineObject175
	isSet bool
}

func (v NullableInlineObject175) Get() *InlineObject175 {
	return v.value
}

func (v *NullableInlineObject175) Set(val *InlineObject175) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject175) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject175(val *InlineObject175) *NullableInlineObject175 {
	return &NullableInlineObject175{value: val, isSet: true}
}

func (v NullableInlineObject175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


