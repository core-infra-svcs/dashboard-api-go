/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdVlanProfilesVlanNames struct for NetworksNetworkIdVlanProfilesVlanNames
type NetworksNetworkIdVlanProfilesVlanNames struct {
	// Name of the VLAN, string length must be from 1 to 32 characters
	Name *string `json:"name,omitempty"`
	// VLAN ID
	VlanId *string `json:"vlanId,omitempty"`
	AdaptivePolicyGroup *NetworksNetworkIdVlanProfilesAdaptivePolicyGroup `json:"adaptivePolicyGroup,omitempty"`
}

// NewNetworksNetworkIdVlanProfilesVlanNames instantiates a new NetworksNetworkIdVlanProfilesVlanNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdVlanProfilesVlanNames() *NetworksNetworkIdVlanProfilesVlanNames {
	this := NetworksNetworkIdVlanProfilesVlanNames{}
	return &this
}

// NewNetworksNetworkIdVlanProfilesVlanNamesWithDefaults instantiates a new NetworksNetworkIdVlanProfilesVlanNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdVlanProfilesVlanNamesWithDefaults() *NetworksNetworkIdVlanProfilesVlanNames {
	this := NetworksNetworkIdVlanProfilesVlanNames{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesVlanNames) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdVlanProfilesVlanNames) SetName(v string) {
	o.Name = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesVlanNames) GetVlanId() string {
	if o == nil || isNil(o.VlanId) {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames) GetVlanIdOk() (*string, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *NetworksNetworkIdVlanProfilesVlanNames) SetVlanId(v string) {
	o.VlanId = &v
}

// GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesVlanNames) GetAdaptivePolicyGroup() NetworksNetworkIdVlanProfilesAdaptivePolicyGroup {
	if o == nil || isNil(o.AdaptivePolicyGroup) {
		var ret NetworksNetworkIdVlanProfilesAdaptivePolicyGroup
		return ret
	}
	return *o.AdaptivePolicyGroup
}

// GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames) GetAdaptivePolicyGroupOk() (*NetworksNetworkIdVlanProfilesAdaptivePolicyGroup, bool) {
	if o == nil || isNil(o.AdaptivePolicyGroup) {
    return nil, false
	}
	return o.AdaptivePolicyGroup, true
}

// HasAdaptivePolicyGroup returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesVlanNames) HasAdaptivePolicyGroup() bool {
	if o != nil && !isNil(o.AdaptivePolicyGroup) {
		return true
	}

	return false
}

// SetAdaptivePolicyGroup gets a reference to the given NetworksNetworkIdVlanProfilesAdaptivePolicyGroup and assigns it to the AdaptivePolicyGroup field.
func (o *NetworksNetworkIdVlanProfilesVlanNames) SetAdaptivePolicyGroup(v NetworksNetworkIdVlanProfilesAdaptivePolicyGroup) {
	o.AdaptivePolicyGroup = &v
}

func (o NetworksNetworkIdVlanProfilesVlanNames) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.AdaptivePolicyGroup) {
		toSerialize["adaptivePolicyGroup"] = o.AdaptivePolicyGroup
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdVlanProfilesVlanNames struct {
	value *NetworksNetworkIdVlanProfilesVlanNames
	isSet bool
}

func (v NullableNetworksNetworkIdVlanProfilesVlanNames) Get() *NetworksNetworkIdVlanProfilesVlanNames {
	return v.value
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanNames) Set(val *NetworksNetworkIdVlanProfilesVlanNames) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdVlanProfilesVlanNames) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdVlanProfilesVlanNames(val *NetworksNetworkIdVlanProfilesVlanNames) *NullableNetworksNetworkIdVlanProfilesVlanNames {
	return &NullableNetworksNetworkIdVlanProfilesVlanNames{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdVlanProfilesVlanNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdVlanProfilesVlanNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


