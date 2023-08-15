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

// InlineResponse20087 struct for InlineResponse20087
type InlineResponse20087 struct {
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
	Sharing *NetworksNetworkIdWebhooksPayloadTemplatesSharing `json:"sharing,omitempty"`
}

// NewInlineResponse20087 instantiates a new InlineResponse20087 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20087() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// NewInlineResponse20087WithDefaults instantiates a new InlineResponse20087 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20087WithDefaults() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *InlineResponse20087) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *InlineResponse20087) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *InlineResponse20087) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20087) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20087) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20087) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20087) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20087) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20087) SetName(v string) {
	o.Name = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineResponse20087) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	if o == nil || isNil(o.Headers) {
		var ret []NetworksNetworkIdWebhooksPayloadTemplatesHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetHeadersOk() ([]NetworksNetworkIdWebhooksPayloadTemplatesHeaders, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineResponse20087) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NetworksNetworkIdWebhooksPayloadTemplatesHeaders and assigns it to the Headers field.
func (o *InlineResponse20087) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders) {
	o.Headers = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse20087) GetBody() string {
	if o == nil || isNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetBodyOk() (*string, bool) {
	if o == nil || isNil(o.Body) {
    return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse20087) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InlineResponse20087) SetBody(v string) {
	o.Body = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *InlineResponse20087) GetSharing() NetworksNetworkIdWebhooksPayloadTemplatesSharing {
	if o == nil || isNil(o.Sharing) {
		var ret NetworksNetworkIdWebhooksPayloadTemplatesSharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetSharingOk() (*NetworksNetworkIdWebhooksPayloadTemplatesSharing, bool) {
	if o == nil || isNil(o.Sharing) {
    return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *InlineResponse20087) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given NetworksNetworkIdWebhooksPayloadTemplatesSharing and assigns it to the Sharing field.
func (o *InlineResponse20087) SetSharing(v NetworksNetworkIdWebhooksPayloadTemplatesSharing) {
	o.Sharing = &v
}

func (o InlineResponse20087) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20087 struct {
	value *InlineResponse20087
	isSet bool
}

func (v NullableInlineResponse20087) Get() *InlineResponse20087 {
	return v.value
}

func (v *NullableInlineResponse20087) Set(val *InlineResponse20087) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20087) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20087) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20087(val *InlineResponse20087) *NullableInlineResponse20087 {
	return &NullableInlineResponse20087{value: val, isSet: true}
}

func (v NullableInlineResponse20087) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20087) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


