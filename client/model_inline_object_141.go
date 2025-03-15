/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject141 struct for InlineObject141
type InlineObject141 struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []NetworksNetworkIdSwitchDscpToCosMappingsMappings `json:"mappings"`
}

// NewInlineObject141 instantiates a new InlineObject141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject141(mappings []NetworksNetworkIdSwitchDscpToCosMappingsMappings) *InlineObject141 {
	this := InlineObject141{}
	this.Mappings = mappings
	return &this
}

// NewInlineObject141WithDefaults instantiates a new InlineObject141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject141WithDefaults() *InlineObject141 {
	this := InlineObject141{}
	return &this
}

// GetMappings returns the Mappings field value
func (o *InlineObject141) GetMappings() []NetworksNetworkIdSwitchDscpToCosMappingsMappings {
	if o == nil {
		var ret []NetworksNetworkIdSwitchDscpToCosMappingsMappings
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *InlineObject141) GetMappingsOk() ([]NetworksNetworkIdSwitchDscpToCosMappingsMappings, bool) {
	if o == nil {
    return nil, false
	}
	return o.Mappings, true
}

// SetMappings sets field value
func (o *InlineObject141) SetMappings(v []NetworksNetworkIdSwitchDscpToCosMappingsMappings) {
	o.Mappings = v
}

func (o InlineObject141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject141 struct {
	value *InlineObject141
	isSet bool
}

func (v NullableInlineObject141) Get() *InlineObject141 {
	return v.value
}

func (v *NullableInlineObject141) Set(val *InlineObject141) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject141) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject141(val *InlineObject141) *NullableInlineObject141 {
	return &NullableInlineObject141{value: val, isSet: true}
}

func (v NullableInlineObject141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


