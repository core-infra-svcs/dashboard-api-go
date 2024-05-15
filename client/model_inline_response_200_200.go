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

// InlineResponse200200 struct for InlineResponse200200
type InlineResponse200200 struct {
	// List of network IDs with adaptive policy enabled
	EnabledNetworks []string `json:"enabledNetworks,omitempty"`
}

// NewInlineResponse200200 instantiates a new InlineResponse200200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200() *InlineResponse200200 {
	this := InlineResponse200200{}
	return &this
}

// NewInlineResponse200200WithDefaults instantiates a new InlineResponse200200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200WithDefaults() *InlineResponse200200 {
	this := InlineResponse200200{}
	return &this
}

// GetEnabledNetworks returns the EnabledNetworks field value if set, zero value otherwise.
func (o *InlineResponse200200) GetEnabledNetworks() []string {
	if o == nil || isNil(o.EnabledNetworks) {
		var ret []string
		return ret
	}
	return o.EnabledNetworks
}

// GetEnabledNetworksOk returns a tuple with the EnabledNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetEnabledNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.EnabledNetworks) {
    return nil, false
	}
	return o.EnabledNetworks, true
}

// HasEnabledNetworks returns a boolean if a field has been set.
func (o *InlineResponse200200) HasEnabledNetworks() bool {
	if o != nil && !isNil(o.EnabledNetworks) {
		return true
	}

	return false
}

// SetEnabledNetworks gets a reference to the given []string and assigns it to the EnabledNetworks field.
func (o *InlineResponse200200) SetEnabledNetworks(v []string) {
	o.EnabledNetworks = v
}

func (o InlineResponse200200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnabledNetworks) {
		toSerialize["enabledNetworks"] = o.EnabledNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200200 struct {
	value *InlineResponse200200
	isSet bool
}

func (v NullableInlineResponse200200) Get() *InlineResponse200200 {
	return v.value
}

func (v *NullableInlineResponse200200) Set(val *InlineResponse200200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200(val *InlineResponse200200) *NullableInlineResponse200200 {
	return &NullableInlineResponse200200{value: val, isSet: true}
}

func (v NullableInlineResponse200200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


