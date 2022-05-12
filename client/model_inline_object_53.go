/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject53 struct for InlineObject53
type InlineObject53 struct {
	// The ID of the template to which the network should be bound.
	ConfigTemplateId string `json:"configTemplateId"`
	// Optional boolean indicating whether the network's switches should automatically bind to profiles of the same model. Defaults to false if left unspecified. This option only affects switch networks and switch templates. Auto-bind is not valid unless the switch template has at least one profile and has at most one profile per switch model.
	AutoBind *bool `json:"autoBind,omitempty"`
}

// NewInlineObject53 instantiates a new InlineObject53 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject53(configTemplateId string) *InlineObject53 {
	this := InlineObject53{}
	this.ConfigTemplateId = configTemplateId
	return &this
}

// NewInlineObject53WithDefaults instantiates a new InlineObject53 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject53WithDefaults() *InlineObject53 {
	this := InlineObject53{}
	return &this
}

// GetConfigTemplateId returns the ConfigTemplateId field value
func (o *InlineObject53) GetConfigTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigTemplateId
}

// GetConfigTemplateIdOk returns a tuple with the ConfigTemplateId field value
// and a boolean to check if the value has been set.
func (o *InlineObject53) GetConfigTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigTemplateId, true
}

// SetConfigTemplateId sets field value
func (o *InlineObject53) SetConfigTemplateId(v string) {
	o.ConfigTemplateId = v
}

// GetAutoBind returns the AutoBind field value if set, zero value otherwise.
func (o *InlineObject53) GetAutoBind() bool {
	if o == nil || o.AutoBind == nil {
		var ret bool
		return ret
	}
	return *o.AutoBind
}

// GetAutoBindOk returns a tuple with the AutoBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject53) GetAutoBindOk() (*bool, bool) {
	if o == nil || o.AutoBind == nil {
		return nil, false
	}
	return o.AutoBind, true
}

// HasAutoBind returns a boolean if a field has been set.
func (o *InlineObject53) HasAutoBind() bool {
	if o != nil && o.AutoBind != nil {
		return true
	}

	return false
}

// SetAutoBind gets a reference to the given bool and assigns it to the AutoBind field.
func (o *InlineObject53) SetAutoBind(v bool) {
	o.AutoBind = &v
}

func (o InlineObject53) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configTemplateId"] = o.ConfigTemplateId
	}
	if o.AutoBind != nil {
		toSerialize["autoBind"] = o.AutoBind
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject53 struct {
	value *InlineObject53
	isSet bool
}

func (v NullableInlineObject53) Get() *InlineObject53 {
	return v.value
}

func (v *NullableInlineObject53) Set(val *InlineObject53) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject53) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject53) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject53(val *InlineObject53) *NullableInlineObject53 {
	return &NullableInlineObject53{value: val, isSet: true}
}

func (v NullableInlineObject53) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject53) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

