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

// NetworksNetworkIdVlanProfilesVlanNames1 struct for NetworksNetworkIdVlanProfilesVlanNames1
type NetworksNetworkIdVlanProfilesVlanNames1 struct {
	// Name of the VLAN, string length must be from 1 to 32 characters
	Name string `json:"name"`
	// VLAN ID
	VlanId string `json:"vlanId"`
	AdaptivePolicyGroup *NetworksNetworkIdVlanProfilesAdaptivePolicyGroup1 `json:"adaptivePolicyGroup,omitempty"`
}

// NewNetworksNetworkIdVlanProfilesVlanNames1 instantiates a new NetworksNetworkIdVlanProfilesVlanNames1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdVlanProfilesVlanNames1(name string, vlanId string) *NetworksNetworkIdVlanProfilesVlanNames1 {
	this := NetworksNetworkIdVlanProfilesVlanNames1{}
	this.Name = name
	this.VlanId = vlanId
	return &this
}

// NewNetworksNetworkIdVlanProfilesVlanNames1WithDefaults instantiates a new NetworksNetworkIdVlanProfilesVlanNames1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdVlanProfilesVlanNames1WithDefaults() *NetworksNetworkIdVlanProfilesVlanNames1 {
	this := NetworksNetworkIdVlanProfilesVlanNames1{}
	return &this
}

// GetName returns the Name field value
func (o *NetworksNetworkIdVlanProfilesVlanNames1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames1) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworksNetworkIdVlanProfilesVlanNames1) SetName(v string) {
	o.Name = v
}

// GetVlanId returns the VlanId field value
func (o *NetworksNetworkIdVlanProfilesVlanNames1) GetVlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames1) GetVlanIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *NetworksNetworkIdVlanProfilesVlanNames1) SetVlanId(v string) {
	o.VlanId = v
}

// GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesVlanNames1) GetAdaptivePolicyGroup() NetworksNetworkIdVlanProfilesAdaptivePolicyGroup1 {
	if o == nil || isNil(o.AdaptivePolicyGroup) {
		var ret NetworksNetworkIdVlanProfilesAdaptivePolicyGroup1
		return ret
	}
	return *o.AdaptivePolicyGroup
}

// GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames1) GetAdaptivePolicyGroupOk() (*NetworksNetworkIdVlanProfilesAdaptivePolicyGroup1, bool) {
	if o == nil || isNil(o.AdaptivePolicyGroup) {
    return nil, false
	}
	return o.AdaptivePolicyGroup, true
}

// HasAdaptivePolicyGroup returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames1) HasAdaptivePolicyGroup() bool {
	if o != nil && !isNil(o.AdaptivePolicyGroup) {
		return true
	}

	return false
}

// SetAdaptivePolicyGroup gets a reference to the given NetworksNetworkIdVlanProfilesAdaptivePolicyGroup1 and assigns it to the AdaptivePolicyGroup field.
func (o *NetworksNetworkIdVlanProfilesVlanNames1) SetAdaptivePolicyGroup(v NetworksNetworkIdVlanProfilesAdaptivePolicyGroup1) {
	o.AdaptivePolicyGroup = &v
}

func (o NetworksNetworkIdVlanProfilesVlanNames1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.AdaptivePolicyGroup) {
		toSerialize["adaptivePolicyGroup"] = o.AdaptivePolicyGroup
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdVlanProfilesVlanNames1 struct {
	value *NetworksNetworkIdVlanProfilesVlanNames1
	isSet bool
}

func (v NullableNetworksNetworkIdVlanProfilesVlanNames1) Get() *NetworksNetworkIdVlanProfilesVlanNames1 {
	return v.value
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanNames1) Set(val *NetworksNetworkIdVlanProfilesVlanNames1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdVlanProfilesVlanNames1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanNames1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdVlanProfilesVlanNames1(val *NetworksNetworkIdVlanProfilesVlanNames1) *NullableNetworksNetworkIdVlanProfilesVlanNames1 {
	return &NullableNetworksNetworkIdVlanProfilesVlanNames1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdVlanProfilesVlanNames1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanNames1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


