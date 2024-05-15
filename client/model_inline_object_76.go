/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject76 struct for InlineObject76
type InlineObject76 struct {
	// The ID of the template to which the network should be bound.
	ConfigTemplateId string `json:"configTemplateId"`
	// Optional boolean indicating whether the network's switches should automatically bind to profiles of the same model. Defaults to false if left unspecified. This option only affects switch networks and switch templates. Auto-bind is not valid unless the switch template has at least one profile and has at most one profile per switch model.
	AutoBind *bool `json:"autoBind,omitempty"`
}

// NewInlineObject76 instantiates a new InlineObject76 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject76(configTemplateId string) *InlineObject76 {
	this := InlineObject76{}
	this.ConfigTemplateId = configTemplateId
	return &this
}

// NewInlineObject76WithDefaults instantiates a new InlineObject76 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject76WithDefaults() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// GetConfigTemplateId returns the ConfigTemplateId field value
func (o *InlineObject76) GetConfigTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigTemplateId
}

// GetConfigTemplateIdOk returns a tuple with the ConfigTemplateId field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetConfigTemplateIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConfigTemplateId, true
}

// SetConfigTemplateId sets field value
func (o *InlineObject76) SetConfigTemplateId(v string) {
	o.ConfigTemplateId = v
}

// GetAutoBind returns the AutoBind field value if set, zero value otherwise.
func (o *InlineObject76) GetAutoBind() bool {
	if o == nil || isNil(o.AutoBind) {
		var ret bool
		return ret
	}
	return *o.AutoBind
}

// GetAutoBindOk returns a tuple with the AutoBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetAutoBindOk() (*bool, bool) {
	if o == nil || isNil(o.AutoBind) {
    return nil, false
	}
	return o.AutoBind, true
}

// HasAutoBind returns a boolean if a field has been set.
func (o *InlineObject76) HasAutoBind() bool {
	if o != nil && !isNil(o.AutoBind) {
		return true
	}

	return false
}

// SetAutoBind gets a reference to the given bool and assigns it to the AutoBind field.
func (o *InlineObject76) SetAutoBind(v bool) {
	o.AutoBind = &v
}

func (o InlineObject76) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configTemplateId"] = o.ConfigTemplateId
	}
	if !isNil(o.AutoBind) {
		toSerialize["autoBind"] = o.AutoBind
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject76 struct {
	value *InlineObject76
	isSet bool
}

func (v NullableInlineObject76) Get() *InlineObject76 {
	return v.value
}

func (v *NullableInlineObject76) Set(val *InlineObject76) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject76) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject76) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject76(val *InlineObject76) *NullableInlineObject76 {
	return &NullableInlineObject76{value: val, isSet: true}
}

func (v NullableInlineObject76) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject76) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


