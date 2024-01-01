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

// NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags struct for NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags
type NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags struct {
	// List of AP tags.
	Tags []string `json:"tags,omitempty"`
	// VLAN name that will be used to tag traffic.
	VlanName *string `json:"vlanName,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTagsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTagsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) SetTags(v []string) {
	o.Tags = v
}

// GetVlanName returns the VlanName field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) GetVlanName() string {
	if o == nil || isNil(o.VlanName) {
		var ret string
		return ret
	}
	return *o.VlanName
}

// GetVlanNameOk returns a tuple with the VlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) GetVlanNameOk() (*string, bool) {
	if o == nil || isNil(o.VlanName) {
    return nil, false
	}
	return o.VlanName, true
}

// HasVlanName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) HasVlanName() bool {
	if o != nil && !isNil(o.VlanName) {
		return true
	}

	return false
}

// SetVlanName gets a reference to the given string and assigns it to the VlanName field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) SetVlanName(v string) {
	o.VlanName = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.VlanName) {
		toSerialize["vlanName"] = o.VlanName
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags struct {
	value *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) Get() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) Set(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags {
	return &NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


