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

// InlineObject241 struct for InlineObject241
type InlineObject241 struct {
	// Service provider account name used on the Meraki UI
	Title *string `json:"title,omitempty"`
	// Service provider account API key
	ApiKey *string `json:"apiKey,omitempty"`
}

// NewInlineObject241 instantiates a new InlineObject241 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject241() *InlineObject241 {
	this := InlineObject241{}
	return &this
}

// NewInlineObject241WithDefaults instantiates a new InlineObject241 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject241WithDefaults() *InlineObject241 {
	this := InlineObject241{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineObject241) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject241) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineObject241) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineObject241) SetTitle(v string) {
	o.Title = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *InlineObject241) GetApiKey() string {
	if o == nil || isNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject241) GetApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.ApiKey) {
    return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *InlineObject241) HasApiKey() bool {
	if o != nil && !isNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *InlineObject241) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o InlineObject241) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject241 struct {
	value *InlineObject241
	isSet bool
}

func (v NullableInlineObject241) Get() *InlineObject241 {
	return v.value
}

func (v *NullableInlineObject241) Set(val *InlineObject241) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject241) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject241) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject241(val *InlineObject241) *NullableInlineObject241 {
	return &NullableInlineObject241{value: val, isSet: true}
}

func (v NullableInlineObject241) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject241) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


