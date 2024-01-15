/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 The payload template to use when posting data to the HTTP server.
type NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 struct {
	// The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The name of the payload template.
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdWebhooksHttpServersPayloadTemplate1 instantiates a new NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWebhooksHttpServersPayloadTemplate1() *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 {
	this := NetworksNetworkIdWebhooksHttpServersPayloadTemplate1{}
	return &this
}

// NewNetworksNetworkIdWebhooksHttpServersPayloadTemplate1WithDefaults instantiates a new NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWebhooksHttpServersPayloadTemplate1WithDefaults() *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 {
	this := NetworksNetworkIdWebhooksHttpServersPayloadTemplate1{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1 struct {
	value *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1
	isSet bool
}

func (v NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1) Get() *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1 {
	return v.value
}

func (v *NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1) Set(val *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1(val *NetworksNetworkIdWebhooksHttpServersPayloadTemplate1) *NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1 {
	return &NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWebhooksHttpServersPayloadTemplate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


