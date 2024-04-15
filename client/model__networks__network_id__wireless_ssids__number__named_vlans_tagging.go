/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging VLAN tagging settings. This param is only valid when ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'.
type NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging struct {
	// Whether or not traffic should be directed to use specific VLAN names.
	Enabled *bool `json:"enabled,omitempty"`
	// The default VLAN name used to tag traffic in the absence of a matching AP tag.
	DefaultVlanName *string `json:"defaultVlanName,omitempty"`
	// The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag.
	ByApTags []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags `json:"byApTags,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingWithDefaults() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefaultVlanName returns the DefaultVlanName field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetDefaultVlanName() string {
	if o == nil || isNil(o.DefaultVlanName) {
		var ret string
		return ret
	}
	return *o.DefaultVlanName
}

// GetDefaultVlanNameOk returns a tuple with the DefaultVlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetDefaultVlanNameOk() (*string, bool) {
	if o == nil || isNil(o.DefaultVlanName) {
    return nil, false
	}
	return o.DefaultVlanName, true
}

// HasDefaultVlanName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) HasDefaultVlanName() bool {
	if o != nil && !isNil(o.DefaultVlanName) {
		return true
	}

	return false
}

// SetDefaultVlanName gets a reference to the given string and assigns it to the DefaultVlanName field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) SetDefaultVlanName(v string) {
	o.DefaultVlanName = &v
}

// GetByApTags returns the ByApTags field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetByApTags() []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags {
	if o == nil || isNil(o.ByApTags) {
		var ret []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags
		return ret
	}
	return o.ByApTags
}

// GetByApTagsOk returns a tuple with the ByApTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) GetByApTagsOk() ([]NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags, bool) {
	if o == nil || isNil(o.ByApTags) {
    return nil, false
	}
	return o.ByApTags, true
}

// HasByApTags returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) HasByApTags() bool {
	if o != nil && !isNil(o.ByApTags) {
		return true
	}

	return false
}

// SetByApTags gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags and assigns it to the ByApTags field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) SetByApTags(v []NetworksNetworkIdWirelessSsidsNumberNamedVlansTaggingByApTags) {
	o.ByApTags = v
}

func (o NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DefaultVlanName) {
		toSerialize["defaultVlanName"] = o.DefaultVlanName
	}
	if !isNil(o.ByApTags) {
		toSerialize["byApTags"] = o.ByApTags
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging struct {
	value *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) Get() *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) Set(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging {
	return &NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


