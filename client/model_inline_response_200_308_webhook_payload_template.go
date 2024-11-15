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

// InlineResponse200308WebhookPayloadTemplate The payload template of the webhook used for the callback
type InlineResponse200308WebhookPayloadTemplate struct {
	// The ID of the payload template
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse200308WebhookPayloadTemplate instantiates a new InlineResponse200308WebhookPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200308WebhookPayloadTemplate() *InlineResponse200308WebhookPayloadTemplate {
	this := InlineResponse200308WebhookPayloadTemplate{}
	return &this
}

// NewInlineResponse200308WebhookPayloadTemplateWithDefaults instantiates a new InlineResponse200308WebhookPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200308WebhookPayloadTemplateWithDefaults() *InlineResponse200308WebhookPayloadTemplate {
	this := InlineResponse200308WebhookPayloadTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200308WebhookPayloadTemplate) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200308WebhookPayloadTemplate) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200308WebhookPayloadTemplate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200308WebhookPayloadTemplate) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse200308WebhookPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200308WebhookPayloadTemplate struct {
	value *InlineResponse200308WebhookPayloadTemplate
	isSet bool
}

func (v NullableInlineResponse200308WebhookPayloadTemplate) Get() *InlineResponse200308WebhookPayloadTemplate {
	return v.value
}

func (v *NullableInlineResponse200308WebhookPayloadTemplate) Set(val *InlineResponse200308WebhookPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200308WebhookPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200308WebhookPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200308WebhookPayloadTemplate(val *InlineResponse200308WebhookPayloadTemplate) *NullableInlineResponse200308WebhookPayloadTemplate {
	return &NullableInlineResponse200308WebhookPayloadTemplate{value: val, isSet: true}
}

func (v NullableInlineResponse200308WebhookPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200308WebhookPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

