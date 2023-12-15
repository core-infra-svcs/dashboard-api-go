/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds struct for NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds
type NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds struct {
	// Array of AP tags
	Tags []string `json:"tags,omitempty"`
	// Numerical identifier that is assigned to the VLAN
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds instantiates a new NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds() *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds {
	this := NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIdsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIdsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds {
	this := NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) SetTags(v []string) {
	o.Tags = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds struct {
	value *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) Get() *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) Set(val *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds(val *NetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) *NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds {
	return &NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberApTagsAndVlanIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


