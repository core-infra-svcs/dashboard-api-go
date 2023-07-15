/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWebhooksPayloadTemplatesHeaders struct for NetworksNetworkIdWebhooksPayloadTemplatesHeaders
type NetworksNetworkIdWebhooksPayloadTemplatesHeaders struct {
	// The name of the header attribute
	Name *string `json:"name,omitempty"`
	// The value returned in the header attribute, in liquid template
	Template *string `json:"template,omitempty"`
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesHeaders instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWebhooksPayloadTemplatesHeaders() *NetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	this := NetworksNetworkIdWebhooksPayloadTemplatesHeaders{}
	return &this
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesHeadersWithDefaults instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWebhooksPayloadTemplatesHeadersWithDefaults() *NetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	this := NetworksNetworkIdWebhooksPayloadTemplatesHeaders{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) SetName(v string) {
	o.Name = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) GetTemplate() string {
	if o == nil || isNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) GetTemplateOk() (*string, bool) {
	if o == nil || isNil(o.Template) {
    return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) SetTemplate(v string) {
	o.Template = &v
}

func (o NetworksNetworkIdWebhooksPayloadTemplatesHeaders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders struct {
	value *NetworksNetworkIdWebhooksPayloadTemplatesHeaders
	isSet bool
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders) Get() *NetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	return v.value
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders) Set(val *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders(val *NetworksNetworkIdWebhooksPayloadTemplatesHeaders) *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders {
	return &NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


