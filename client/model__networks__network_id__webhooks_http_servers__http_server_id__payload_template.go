/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate The payload template to use when posting data to the HTTP server.
type NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate struct {
	// The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
}

// NewNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate instantiates a new NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate() *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate {
	this := NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate{}
	return &this
}

// NewNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplateWithDefaults instantiates a new NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplateWithDefaults() *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate {
	this := NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

func (o NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate struct {
	value *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate
	isSet bool
}

func (v NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) Get() *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate {
	return v.value
}

func (v *NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) Set(val *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate(val *NetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) *NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate {
	return &NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWebhooksHttpServersHttpServerIdPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


