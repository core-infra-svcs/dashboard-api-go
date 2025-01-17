/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject231 struct for InlineObject231
type InlineObject231 struct {
	// Service provider account name used on the Meraki UI
	Title *string `json:"title,omitempty"`
	// Service provider account API key
	ApiKey *string `json:"apiKey,omitempty"`
}

// NewInlineObject231 instantiates a new InlineObject231 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject231() *InlineObject231 {
	this := InlineObject231{}
	return &this
}

// NewInlineObject231WithDefaults instantiates a new InlineObject231 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject231WithDefaults() *InlineObject231 {
	this := InlineObject231{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineObject231) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject231) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineObject231) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineObject231) SetTitle(v string) {
	o.Title = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *InlineObject231) GetApiKey() string {
	if o == nil || isNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject231) GetApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.ApiKey) {
    return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *InlineObject231) HasApiKey() bool {
	if o != nil && !isNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *InlineObject231) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o InlineObject231) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject231 struct {
	value *InlineObject231
	isSet bool
}

func (v NullableInlineObject231) Get() *InlineObject231 {
	return v.value
}

func (v *NullableInlineObject231) Set(val *InlineObject231) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject231) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject231) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject231(val *InlineObject231) *NullableInlineObject231 {
	return &NullableInlineObject231{value: val, isSet: true}
}

func (v NullableInlineObject231) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject231) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


