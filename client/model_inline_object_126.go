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

// InlineObject126 struct for InlineObject126
type InlineObject126 struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []NetworksNetworkIdSwitchDscpToCosMappingsMappings `json:"mappings"`
}

// NewInlineObject126 instantiates a new InlineObject126 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject126(mappings []NetworksNetworkIdSwitchDscpToCosMappingsMappings) *InlineObject126 {
	this := InlineObject126{}
	this.Mappings = mappings
	return &this
}

// NewInlineObject126WithDefaults instantiates a new InlineObject126 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject126WithDefaults() *InlineObject126 {
	this := InlineObject126{}
	return &this
}

// GetMappings returns the Mappings field value
func (o *InlineObject126) GetMappings() []NetworksNetworkIdSwitchDscpToCosMappingsMappings {
	if o == nil {
		var ret []NetworksNetworkIdSwitchDscpToCosMappingsMappings
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetMappingsOk() ([]NetworksNetworkIdSwitchDscpToCosMappingsMappings, bool) {
	if o == nil {
    return nil, false
	}
	return o.Mappings, true
}

// SetMappings sets field value
func (o *InlineObject126) SetMappings(v []NetworksNetworkIdSwitchDscpToCosMappingsMappings) {
	o.Mappings = v
}

func (o InlineObject126) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject126 struct {
	value *InlineObject126
	isSet bool
}

func (v NullableInlineObject126) Get() *InlineObject126 {
	return v.value
}

func (v *NullableInlineObject126) Set(val *InlineObject126) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject126) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject126) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject126(val *InlineObject126) *NullableInlineObject126 {
	return &NullableInlineObject126{value: val, isSet: true}
}

func (v NullableInlineObject126) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject126) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


