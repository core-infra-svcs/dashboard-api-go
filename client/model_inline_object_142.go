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

// InlineObject142 struct for InlineObject142
type InlineObject142 struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []NetworksNetworkIdSwitchDscpToCosMappingsMappings `json:"mappings"`
}

// NewInlineObject142 instantiates a new InlineObject142 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject142(mappings []NetworksNetworkIdSwitchDscpToCosMappingsMappings) *InlineObject142 {
	this := InlineObject142{}
	this.Mappings = mappings
	return &this
}

// NewInlineObject142WithDefaults instantiates a new InlineObject142 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject142WithDefaults() *InlineObject142 {
	this := InlineObject142{}
	return &this
}

// GetMappings returns the Mappings field value
func (o *InlineObject142) GetMappings() []NetworksNetworkIdSwitchDscpToCosMappingsMappings {
	if o == nil {
		var ret []NetworksNetworkIdSwitchDscpToCosMappingsMappings
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *InlineObject142) GetMappingsOk() ([]NetworksNetworkIdSwitchDscpToCosMappingsMappings, bool) {
	if o == nil {
    return nil, false
	}
	return o.Mappings, true
}

// SetMappings sets field value
func (o *InlineObject142) SetMappings(v []NetworksNetworkIdSwitchDscpToCosMappingsMappings) {
	o.Mappings = v
}

func (o InlineObject142) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject142 struct {
	value *InlineObject142
	isSet bool
}

func (v NullableInlineObject142) Get() *InlineObject142 {
	return v.value
}

func (v *NullableInlineObject142) Set(val *InlineObject142) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject142) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject142) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject142(val *InlineObject142) *NullableInlineObject142 {
	return &NullableInlineObject142{value: val, isSet: true}
}

func (v NullableInlineObject142) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject142) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


