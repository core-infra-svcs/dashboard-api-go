/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200175 struct for InlineResponse200175
type InlineResponse200175 struct {
	// A Base64 encoded ID.
	Id *string `json:"id,omitempty"`
	// A name for easy reference to the HTTP server
	Name *string `json:"name,omitempty"`
	// The URL of the HTTP server.
	Url *string `json:"url,omitempty"`
	// A Meraki network ID.
	NetworkId *string `json:"networkId,omitempty"`
	PayloadTemplate *NetworksNetworkIdWebhooksHttpServersPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewInlineResponse200175 instantiates a new InlineResponse200175 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200175() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// NewInlineResponse200175WithDefaults instantiates a new InlineResponse200175 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200175WithDefaults() *InlineResponse200175 {
	this := InlineResponse200175{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200175) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200175) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200175) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200175) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200175) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200175) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200175) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200175) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200175) SetUrl(v string) {
	o.Url = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200175) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200175) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200175) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *InlineResponse200175) GetPayloadTemplate() NetworksNetworkIdWebhooksHttpServersPayloadTemplate {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret NetworksNetworkIdWebhooksHttpServersPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200175) GetPayloadTemplateOk() (*NetworksNetworkIdWebhooksHttpServersPayloadTemplate, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *InlineResponse200175) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given NetworksNetworkIdWebhooksHttpServersPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *InlineResponse200175) SetPayloadTemplate(v NetworksNetworkIdWebhooksHttpServersPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o InlineResponse200175) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200175 struct {
	value *InlineResponse200175
	isSet bool
}

func (v NullableInlineResponse200175) Get() *InlineResponse200175 {
	return v.value
}

func (v *NullableInlineResponse200175) Set(val *InlineResponse200175) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200175) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200175) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200175(val *InlineResponse200175) *NullableInlineResponse200175 {
	return &NullableInlineResponse200175{value: val, isSet: true}
}

func (v NullableInlineResponse200175) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200175) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


