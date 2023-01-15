/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 struct for NetworksNetworkIdWebhooksPayloadTemplatesHeaders1
type NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 struct {
	// The name of the header template
	Name *string `json:"name,omitempty"`
	// The liquid template for the headers
	Template *string `json:"template,omitempty"`
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesHeaders1 instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWebhooksPayloadTemplatesHeaders1() *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 {
	this := NetworksNetworkIdWebhooksPayloadTemplatesHeaders1{}
	return &this
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesHeaders1WithDefaults instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWebhooksPayloadTemplatesHeaders1WithDefaults() *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 {
	this := NetworksNetworkIdWebhooksPayloadTemplatesHeaders1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) SetName(v string) {
	o.Name = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) GetTemplate() string {
	if o == nil || isNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) GetTemplateOk() (*string, bool) {
	if o == nil || isNil(o.Template) {
    return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) SetTemplate(v string) {
	o.Template = &v
}

func (o NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1 struct {
	value *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1
	isSet bool
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1) Get() *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 {
	return v.value
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1) Set(val *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1(val *NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1 {
	return &NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesHeaders1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


