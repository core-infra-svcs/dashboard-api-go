/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject168 struct for InlineObject168
type InlineObject168 struct {
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

// NewInlineObject168 instantiates a new InlineObject168 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject168(url string) *InlineObject168 {
	this := InlineObject168{}
	this.Url = url
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// NewInlineObject168WithDefaults instantiates a new InlineObject168 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject168WithDefaults() *InlineObject168 {
	this := InlineObject168{}
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// GetUrl returns the Url field value
func (o *InlineObject168) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InlineObject168) SetUrl(v string) {
	o.Url = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *InlineObject168) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *InlineObject168) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *InlineObject168) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *InlineObject168) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *InlineObject168) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *InlineObject168) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetPayloadTemplateName returns the PayloadTemplateName field value if set, zero value otherwise.
func (o *InlineObject168) GetPayloadTemplateName() string {
	if o == nil || isNil(o.PayloadTemplateName) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateName
}

// GetPayloadTemplateNameOk returns a tuple with the PayloadTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetPayloadTemplateNameOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateName) {
    return nil, false
	}
	return o.PayloadTemplateName, true
}

// HasPayloadTemplateName returns a boolean if a field has been set.
func (o *InlineObject168) HasPayloadTemplateName() bool {
	if o != nil && !isNil(o.PayloadTemplateName) {
		return true
	}

	return false
}

// SetPayloadTemplateName gets a reference to the given string and assigns it to the PayloadTemplateName field.
func (o *InlineObject168) SetPayloadTemplateName(v string) {
	o.PayloadTemplateName = &v
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *InlineObject168) GetAlertTypeId() string {
	if o == nil || isNil(o.AlertTypeId) {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || isNil(o.AlertTypeId) {
    return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *InlineObject168) HasAlertTypeId() bool {
	if o != nil && !isNil(o.AlertTypeId) {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *InlineObject168) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

func (o InlineObject168) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject168 struct {
	value *InlineObject168
	isSet bool
}

func (v NullableInlineObject168) Get() *InlineObject168 {
	return v.value
}

func (v *NullableInlineObject168) Set(val *InlineObject168) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject168) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject168) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject168(val *InlineObject168) *NullableInlineObject168 {
	return &NullableInlineObject168{value: val, isSet: true}
}

func (v NullableInlineObject168) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject168) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


