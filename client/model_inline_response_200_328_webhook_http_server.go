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

// InlineResponse200328WebhookHttpServer The webhook receiver used for the callback webhook
type InlineResponse200328WebhookHttpServer struct {
	// The webhook receiver ID that will receive information
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse200328WebhookHttpServer instantiates a new InlineResponse200328WebhookHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200328WebhookHttpServer() *InlineResponse200328WebhookHttpServer {
	this := InlineResponse200328WebhookHttpServer{}
	return &this
}

// NewInlineResponse200328WebhookHttpServerWithDefaults instantiates a new InlineResponse200328WebhookHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200328WebhookHttpServerWithDefaults() *InlineResponse200328WebhookHttpServer {
	this := InlineResponse200328WebhookHttpServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200328WebhookHttpServer) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200328WebhookHttpServer) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200328WebhookHttpServer) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200328WebhookHttpServer) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse200328WebhookHttpServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200328WebhookHttpServer struct {
	value *InlineResponse200328WebhookHttpServer
	isSet bool
}

func (v NullableInlineResponse200328WebhookHttpServer) Get() *InlineResponse200328WebhookHttpServer {
	return v.value
}

func (v *NullableInlineResponse200328WebhookHttpServer) Set(val *InlineResponse200328WebhookHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200328WebhookHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200328WebhookHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200328WebhookHttpServer(val *InlineResponse200328WebhookHttpServer) *NullableInlineResponse200328WebhookHttpServer {
	return &NullableInlineResponse200328WebhookHttpServer{value: val, isSet: true}
}

func (v NullableInlineResponse200328WebhookHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200328WebhookHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


