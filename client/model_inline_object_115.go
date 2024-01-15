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

// InlineObject115 struct for InlineObject115
type InlineObject115 struct {
	// ids of applications to be uninstalled
	AppIds []string `json:"appIds"`
}

// NewInlineObject115 instantiates a new InlineObject115 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject115(appIds []string) *InlineObject115 {
	this := InlineObject115{}
	this.AppIds = appIds
	return &this
}

// NewInlineObject115WithDefaults instantiates a new InlineObject115 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject115WithDefaults() *InlineObject115 {
	this := InlineObject115{}
	return &this
}

// GetAppIds returns the AppIds field value
func (o *InlineObject115) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject115) GetAppIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppIds, true
}

// SetAppIds sets field value
func (o *InlineObject115) SetAppIds(v []string) {
	o.AppIds = v
}

func (o InlineObject115) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appIds"] = o.AppIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject115 struct {
	value *InlineObject115
	isSet bool
}

func (v NullableInlineObject115) Get() *InlineObject115 {
	return v.value
}

func (v *NullableInlineObject115) Set(val *InlineObject115) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject115) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject115) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject115(val *InlineObject115) *NullableInlineObject115 {
	return &NullableInlineObject115{value: val, isSet: true}
}

func (v NullableInlineObject115) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject115) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


