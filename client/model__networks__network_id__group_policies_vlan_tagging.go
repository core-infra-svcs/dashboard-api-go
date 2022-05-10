/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesVlanTagging The VLAN tagging settings for your group policy. Only available if your network has a wireless configuration.
type NetworksNetworkIdGroupPoliciesVlanTagging struct {
	// How VLAN tagging is applied. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	// The ID of the vlan you want to tag. This only applies if 'settings' is set to 'custom'.
	VlanId *string `json:"vlanId,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesVlanTagging instantiates a new NetworksNetworkIdGroupPoliciesVlanTagging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesVlanTagging() *NetworksNetworkIdGroupPoliciesVlanTagging {
	this := NetworksNetworkIdGroupPoliciesVlanTagging{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesVlanTaggingWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesVlanTagging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesVlanTaggingWithDefaults() *NetworksNetworkIdGroupPoliciesVlanTagging {
	this := NetworksNetworkIdGroupPoliciesVlanTagging{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetSettings() string {
	if o == nil || o.Settings == nil {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetSettingsOk() (*string, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) SetSettings(v string) {
	o.Settings = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetVlanId() string {
	if o == nil || o.VlanId == nil {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) GetVlanIdOk() (*string, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *NetworksNetworkIdGroupPoliciesVlanTagging) SetVlanId(v string) {
	o.VlanId = &v
}

func (o NetworksNetworkIdGroupPoliciesVlanTagging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.VlanId != nil {
		toSerialize["vlanId"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesVlanTagging struct {
	value *NetworksNetworkIdGroupPoliciesVlanTagging
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesVlanTagging) Get() *NetworksNetworkIdGroupPoliciesVlanTagging {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesVlanTagging) Set(val *NetworksNetworkIdGroupPoliciesVlanTagging) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesVlanTagging) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesVlanTagging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesVlanTagging(val *NetworksNetworkIdGroupPoliciesVlanTagging) *NullableNetworksNetworkIdGroupPoliciesVlanTagging {
	return &NullableNetworksNetworkIdGroupPoliciesVlanTagging{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesVlanTagging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesVlanTagging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


