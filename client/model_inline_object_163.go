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

// InlineObject163 struct for InlineObject163
type InlineObject163 struct {
	// Name of the profile, string length must be from 1 to 255 characters
	Name string `json:"name"`
	// An array of named VLANs
	VlanNames []NetworksNetworkIdVlanProfilesVlanNames1 `json:"vlanNames"`
	// An array of VLAN groups
	VlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1 `json:"vlanGroups"`
}

// NewInlineObject163 instantiates a new InlineObject163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject163(name string, vlanNames []NetworksNetworkIdVlanProfilesVlanNames1, vlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1) *InlineObject163 {
	this := InlineObject163{}
	this.Name = name
	this.VlanNames = vlanNames
	this.VlanGroups = vlanGroups
	return &this
}

// NewInlineObject163WithDefaults instantiates a new InlineObject163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject163WithDefaults() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject163) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject163) SetName(v string) {
	o.Name = v
}

// GetVlanNames returns the VlanNames field value
func (o *InlineObject163) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames1 {
	if o == nil {
		var ret []NetworksNetworkIdVlanProfilesVlanNames1
		return ret
	}

	return o.VlanNames
}

// GetVlanNamesOk returns a tuple with the VlanNames field value
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetVlanNamesOk() ([]NetworksNetworkIdVlanProfilesVlanNames1, bool) {
	if o == nil {
    return nil, false
	}
	return o.VlanNames, true
}

// SetVlanNames sets field value
func (o *InlineObject163) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames1) {
	o.VlanNames = v
}

// GetVlanGroups returns the VlanGroups field value
func (o *InlineObject163) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups1 {
	if o == nil {
		var ret []NetworksNetworkIdVlanProfilesVlanGroups1
		return ret
	}

	return o.VlanGroups
}

// GetVlanGroupsOk returns a tuple with the VlanGroups field value
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetVlanGroupsOk() ([]NetworksNetworkIdVlanProfilesVlanGroups1, bool) {
	if o == nil {
    return nil, false
	}
	return o.VlanGroups, true
}

// SetVlanGroups sets field value
func (o *InlineObject163) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups1) {
	o.VlanGroups = v
}

func (o InlineObject163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["vlanNames"] = o.VlanNames
	}
	if true {
		toSerialize["vlanGroups"] = o.VlanGroups
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject163 struct {
	value *InlineObject163
	isSet bool
}

func (v NullableInlineObject163) Get() *InlineObject163 {
	return v.value
}

func (v *NullableInlineObject163) Set(val *InlineObject163) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject163) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject163(val *InlineObject163) *NullableInlineObject163 {
	return &NullableInlineObject163{value: val, isSet: true}
}

func (v NullableInlineObject163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


