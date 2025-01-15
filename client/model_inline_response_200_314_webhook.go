/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200314Webhook The webhook receiver used by the callback to send results
type InlineResponse200314Webhook struct {
	// The webhook receiver URL where the callback will be sent
	Url *string `json:"url,omitempty"`
	HttpServer *InlineResponse200314WebhookHttpServer `json:"httpServer,omitempty"`
	PayloadTemplate *InlineResponse200314WebhookPayloadTemplate `json:"payloadTemplate,omitempty"`
	// The timestamp the callback was sent to the webhook receiver
	SentAt *time.Time `json:"sentAt,omitempty"`
}

// NewInlineResponse200314Webhook instantiates a new InlineResponse200314Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200314Webhook() *InlineResponse200314Webhook {
	this := InlineResponse200314Webhook{}
	return &this
}

// NewInlineResponse200314WebhookWithDefaults instantiates a new InlineResponse200314Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200314WebhookWithDefaults() *InlineResponse200314Webhook {
	this := InlineResponse200314Webhook{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200314Webhook) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314Webhook) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200314Webhook) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200314Webhook) SetUrl(v string) {
	o.Url = &v
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise.
func (o *InlineResponse200314Webhook) GetHttpServer() InlineResponse200314WebhookHttpServer {
	if o == nil || isNil(o.HttpServer) {
		var ret InlineResponse200314WebhookHttpServer
		return ret
	}
	return *o.HttpServer
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314Webhook) GetHttpServerOk() (*InlineResponse200314WebhookHttpServer, bool) {
	if o == nil || isNil(o.HttpServer) {
    return nil, false
	}
	return o.HttpServer, true
}

// HasHttpServer returns a boolean if a field has been set.
func (o *InlineResponse200314Webhook) HasHttpServer() bool {
	if o != nil && !isNil(o.HttpServer) {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given InlineResponse200314WebhookHttpServer and assigns it to the HttpServer field.
func (o *InlineResponse200314Webhook) SetHttpServer(v InlineResponse200314WebhookHttpServer) {
	o.HttpServer = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *InlineResponse200314Webhook) GetPayloadTemplate() InlineResponse200314WebhookPayloadTemplate {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret InlineResponse200314WebhookPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314Webhook) GetPayloadTemplateOk() (*InlineResponse200314WebhookPayloadTemplate, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *InlineResponse200314Webhook) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given InlineResponse200314WebhookPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *InlineResponse200314Webhook) SetPayloadTemplate(v InlineResponse200314WebhookPayloadTemplate) {
	o.PayloadTemplate = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *InlineResponse200314Webhook) GetSentAt() time.Time {
	if o == nil || isNil(o.SentAt) {
		var ret time.Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314Webhook) GetSentAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *InlineResponse200314Webhook) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given time.Time and assigns it to the SentAt field.
func (o *InlineResponse200314Webhook) SetSentAt(v time.Time) {
	o.SentAt = &v
}

func (o InlineResponse200314Webhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.HttpServer) {
		toSerialize["httpServer"] = o.HttpServer
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200314Webhook struct {
	value *InlineResponse200314Webhook
	isSet bool
}

func (v NullableInlineResponse200314Webhook) Get() *InlineResponse200314Webhook {
	return v.value
}

func (v *NullableInlineResponse200314Webhook) Set(val *InlineResponse200314Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200314Webhook) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200314Webhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200314Webhook(val *InlineResponse200314Webhook) *NullableInlineResponse200314Webhook {
	return &NullableInlineResponse200314Webhook{value: val, isSet: true}
}

func (v NullableInlineResponse200314Webhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200314Webhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

