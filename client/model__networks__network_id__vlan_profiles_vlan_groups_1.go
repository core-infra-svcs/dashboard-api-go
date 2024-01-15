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

// NetworksNetworkIdVlanProfilesVlanGroups1 struct for NetworksNetworkIdVlanProfilesVlanGroups1
type NetworksNetworkIdVlanProfilesVlanGroups1 struct {
	// Name of the VLAN, string length must be from 1 to 32 characters
	Name string `json:"name"`
	// Comma-separated VLAN IDs or ID ranges
	VlanIds string `json:"vlanIds"`
}

// NewNetworksNetworkIdVlanProfilesVlanGroups1 instantiates a new NetworksNetworkIdVlanProfilesVlanGroups1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdVlanProfilesVlanGroups1(name string, vlanIds string) *NetworksNetworkIdVlanProfilesVlanGroups1 {
	this := NetworksNetworkIdVlanProfilesVlanGroups1{}
	this.Name = name
	this.VlanIds = vlanIds
	return &this
}

// NewNetworksNetworkIdVlanProfilesVlanGroups1WithDefaults instantiates a new NetworksNetworkIdVlanProfilesVlanGroups1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdVlanProfilesVlanGroups1WithDefaults() *NetworksNetworkIdVlanProfilesVlanGroups1 {
	this := NetworksNetworkIdVlanProfilesVlanGroups1{}
	return &this
}

// GetName returns the Name field value
func (o *NetworksNetworkIdVlanProfilesVlanGroups1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanGroups1) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworksNetworkIdVlanProfilesVlanGroups1) SetName(v string) {
	o.Name = v
}

// GetVlanIds returns the VlanIds field value
func (o *NetworksNetworkIdVlanProfilesVlanGroups1) GetVlanIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VlanIds
}

// GetVlanIdsOk returns a tuple with the VlanIds field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanGroups1) GetVlanIdsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlanIds, true
}

// SetVlanIds sets field value
func (o *NetworksNetworkIdVlanProfilesVlanGroups1) SetVlanIds(v string) {
	o.VlanIds = v
}

func (o NetworksNetworkIdVlanProfilesVlanGroups1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["vlanIds"] = o.VlanIds
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdVlanProfilesVlanGroups1 struct {
	value *NetworksNetworkIdVlanProfilesVlanGroups1
	isSet bool
}

func (v NullableNetworksNetworkIdVlanProfilesVlanGroups1) Get() *NetworksNetworkIdVlanProfilesVlanGroups1 {
	return v.value
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanGroups1) Set(val *NetworksNetworkIdVlanProfilesVlanGroups1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdVlanProfilesVlanGroups1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanGroups1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdVlanProfilesVlanGroups1(val *NetworksNetworkIdVlanProfilesVlanGroups1) *NullableNetworksNetworkIdVlanProfilesVlanGroups1 {
	return &NullableNetworksNetworkIdVlanProfilesVlanGroups1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdVlanProfilesVlanGroups1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanGroups1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


