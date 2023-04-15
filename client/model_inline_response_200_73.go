/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20073 struct for InlineResponse20073
type InlineResponse20073 struct {
	// Webhook payload template Id
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The type of the payload template
	Type *string `json:"type,omitempty"`
	// The name of the payload template
	Name *string `json:"name,omitempty"`
	// The payload template headers, will be rendered as a key-value pair in the webhook.
	Headers []NetworksNetworkIdWebhooksPayloadTemplatesHeaders `json:"headers,omitempty"`
	// The body of the payload template, in liquid template
	Body *string `json:"body,omitempty"`
}

// NewInlineResponse20073 instantiates a new InlineResponse20073 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20073() *InlineResponse20073 {
	this := InlineResponse20073{}
	return &this
}

// NewInlineResponse20073WithDefaults instantiates a new InlineResponse20073 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20073WithDefaults() *InlineResponse20073 {
	this := InlineResponse20073{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *InlineResponse20073) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *InlineResponse20073) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *InlineResponse20073) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20073) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20073) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20073) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20073) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20073) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20073) SetName(v string) {
	o.Name = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineResponse20073) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	if o == nil || isNil(o.Headers) {
		var ret []NetworksNetworkIdWebhooksPayloadTemplatesHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetHeadersOk() ([]NetworksNetworkIdWebhooksPayloadTemplatesHeaders, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineResponse20073) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NetworksNetworkIdWebhooksPayloadTemplatesHeaders and assigns it to the Headers field.
func (o *InlineResponse20073) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders) {
	o.Headers = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse20073) GetBody() string {
	if o == nil || isNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20073) GetBodyOk() (*string, bool) {
	if o == nil || isNil(o.Body) {
    return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse20073) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InlineResponse20073) SetBody(v string) {
	o.Body = &v
}

func (o InlineResponse20073) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !isNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20073 struct {
	value *InlineResponse20073
	isSet bool
}

func (v NullableInlineResponse20073) Get() *InlineResponse20073 {
	return v.value
}

func (v *NullableInlineResponse20073) Set(val *InlineResponse20073) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20073) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20073) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20073(val *InlineResponse20073) *NullableInlineResponse20073 {
	return &NullableInlineResponse20073{value: val, isSet: true}
}

func (v NullableInlineResponse20073) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20073) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


