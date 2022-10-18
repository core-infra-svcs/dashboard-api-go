/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20052 struct for InlineResponse20052
type InlineResponse20052 struct {
	// Webhook payload template Id
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The type of the payload template
	Type *string `json:"type,omitempty"`
	// The name of the payload template
	Name *string `json:"name,omitempty"`
	// The payload template headers, will be rendered as a key-value pair in the webhook.
	Headers *[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders `json:"headers,omitempty"`
	// The body of the payload template, in liquid template
	Body *string `json:"body,omitempty"`
}

// NewInlineResponse20052 instantiates a new InlineResponse20052 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20052() *InlineResponse20052 {
	this := InlineResponse20052{}
	return &this
}

// NewInlineResponse20052WithDefaults instantiates a new InlineResponse20052 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20052WithDefaults() *InlineResponse20052 {
	this := InlineResponse20052{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *InlineResponse20052) GetPayloadTemplateId() string {
	if o == nil || o.PayloadTemplateId == nil {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || o.PayloadTemplateId == nil {
		return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *InlineResponse20052) HasPayloadTemplateId() bool {
	if o != nil && o.PayloadTemplateId != nil {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *InlineResponse20052) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20052) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20052) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20052) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20052) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20052) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20052) SetName(v string) {
	o.Name = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineResponse20052) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	if o == nil || o.Headers == nil {
		var ret []NetworksNetworkIdWebhooksPayloadTemplatesHeaders
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052) GetHeadersOk() (*[]NetworksNetworkIdWebhooksPayloadTemplatesHeaders, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineResponse20052) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NetworksNetworkIdWebhooksPayloadTemplatesHeaders and assigns it to the Headers field.
func (o *InlineResponse20052) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders) {
	o.Headers = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse20052) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20052) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse20052) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InlineResponse20052) SetBody(v string) {
	o.Body = &v
}

func (o InlineResponse20052) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayloadTemplateId != nil {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20052 struct {
	value *InlineResponse20052
	isSet bool
}

func (v NullableInlineResponse20052) Get() *InlineResponse20052 {
	return v.value
}

func (v *NullableInlineResponse20052) Set(val *InlineResponse20052) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20052) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20052) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20052(val *InlineResponse20052) *NullableInlineResponse20052 {
	return &NullableInlineResponse20052{value: val, isSet: true}
}

func (v NullableInlineResponse20052) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20052) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


