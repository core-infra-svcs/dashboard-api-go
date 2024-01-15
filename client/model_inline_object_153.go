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

// InlineObject153 struct for InlineObject153
type InlineObject153 struct {
	// Name of the profile, string length must be from 1 to 255 characters
	Name string `json:"name"`
	// An array of named VLANs
	VlanNames []NetworksNetworkIdVlanProfilesVlanNames1 `json:"vlanNames"`
	// An array of VLAN groups
	VlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1 `json:"vlanGroups"`
	// IName of the profile
	Iname string `json:"iname"`
}

// NewInlineObject153 instantiates a new InlineObject153 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject153(name string, vlanNames []NetworksNetworkIdVlanProfilesVlanNames1, vlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1, iname string) *InlineObject153 {
	this := InlineObject153{}
	this.Name = name
	this.VlanNames = vlanNames
	this.VlanGroups = vlanGroups
	this.Iname = iname
	return &this
}

// NewInlineObject153WithDefaults instantiates a new InlineObject153 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject153WithDefaults() *InlineObject153 {
	this := InlineObject153{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject153) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject153) SetName(v string) {
	o.Name = v
}

// GetVlanNames returns the VlanNames field value
func (o *InlineObject153) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames1 {
	if o == nil {
		var ret []NetworksNetworkIdVlanProfilesVlanNames1
		return ret
	}

	return o.VlanNames
}

// GetVlanNamesOk returns a tuple with the VlanNames field value
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetVlanNamesOk() ([]NetworksNetworkIdVlanProfilesVlanNames1, bool) {
	if o == nil {
    return nil, false
	}
	return o.VlanNames, true
}

// SetVlanNames sets field value
func (o *InlineObject153) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames1) {
	o.VlanNames = v
}

// GetVlanGroups returns the VlanGroups field value
func (o *InlineObject153) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups1 {
	if o == nil {
		var ret []NetworksNetworkIdVlanProfilesVlanGroups1
		return ret
	}

	return o.VlanGroups
}

// GetVlanGroupsOk returns a tuple with the VlanGroups field value
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetVlanGroupsOk() ([]NetworksNetworkIdVlanProfilesVlanGroups1, bool) {
	if o == nil {
    return nil, false
	}
	return o.VlanGroups, true
}

// SetVlanGroups sets field value
func (o *InlineObject153) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups1) {
	o.VlanGroups = v
}

// GetIname returns the Iname field value
func (o *InlineObject153) GetIname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iname
}

// GetInameOk returns a tuple with the Iname field value
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetInameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Iname, true
}

// SetIname sets field value
func (o *InlineObject153) SetIname(v string) {
	o.Iname = v
}

func (o InlineObject153) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["iname"] = o.Iname
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject153 struct {
	value *InlineObject153
	isSet bool
}

func (v NullableInlineObject153) Get() *InlineObject153 {
	return v.value
}

func (v *NullableInlineObject153) Set(val *InlineObject153) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject153) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject153) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject153(val *InlineObject153) *NullableInlineObject153 {
	return &NullableInlineObject153{value: val, isSet: true}
}

func (v NullableInlineObject153) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject153) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


