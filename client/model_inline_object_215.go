/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject215 struct for InlineObject215
type InlineObject215 struct {
	// List of network IDs with adaptive policy enabled
	EnabledNetworks []string `json:"enabledNetworks,omitempty"`
}

// NewInlineObject215 instantiates a new InlineObject215 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject215() *InlineObject215 {
	this := InlineObject215{}
	return &this
}

// NewInlineObject215WithDefaults instantiates a new InlineObject215 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject215WithDefaults() *InlineObject215 {
	this := InlineObject215{}
	return &this
}

// GetEnabledNetworks returns the EnabledNetworks field value if set, zero value otherwise.
func (o *InlineObject215) GetEnabledNetworks() []string {
	if o == nil || isNil(o.EnabledNetworks) {
		var ret []string
		return ret
	}
	return o.EnabledNetworks
}

// GetEnabledNetworksOk returns a tuple with the EnabledNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject215) GetEnabledNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.EnabledNetworks) {
    return nil, false
	}
	return o.EnabledNetworks, true
}

// HasEnabledNetworks returns a boolean if a field has been set.
func (o *InlineObject215) HasEnabledNetworks() bool {
	if o != nil && !isNil(o.EnabledNetworks) {
		return true
	}

	return false
}

// SetEnabledNetworks gets a reference to the given []string and assigns it to the EnabledNetworks field.
func (o *InlineObject215) SetEnabledNetworks(v []string) {
	o.EnabledNetworks = v
}

func (o InlineObject215) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnabledNetworks) {
		toSerialize["enabledNetworks"] = o.EnabledNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject215 struct {
	value *InlineObject215
	isSet bool
}

func (v NullableInlineObject215) Get() *InlineObject215 {
	return v.value
}

func (v *NullableInlineObject215) Set(val *InlineObject215) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject215) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject215) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject215(val *InlineObject215) *NullableInlineObject215 {
	return &NullableInlineObject215{value: val, isSet: true}
}

func (v NullableInlineObject215) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject215) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


