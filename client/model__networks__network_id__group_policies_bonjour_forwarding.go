/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesBonjourForwarding The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.
type NetworksNetworkIdGroupPoliciesBonjourForwarding struct {
	// How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	// A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.
	Rules *[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules `json:"rules,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesBonjourForwarding instantiates a new NetworksNetworkIdGroupPoliciesBonjourForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesBonjourForwarding() *NetworksNetworkIdGroupPoliciesBonjourForwarding {
	this := NetworksNetworkIdGroupPoliciesBonjourForwarding{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesBonjourForwardingWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesBonjourForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesBonjourForwardingWithDefaults() *NetworksNetworkIdGroupPoliciesBonjourForwarding {
	this := NetworksNetworkIdGroupPoliciesBonjourForwarding{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) GetSettings() string {
	if o == nil || o.Settings == nil {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) GetSettingsOk() (*string, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) SetSettings(v string) {
	o.Settings = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	if o == nil || o.Rules == nil {
		var ret []NetworksNetworkIdGroupPoliciesBonjourForwardingRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) GetRulesOk() (*[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdGroupPoliciesBonjourForwardingRules and assigns it to the Rules field.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwarding) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules) {
	o.Rules = &v
}

func (o NetworksNetworkIdGroupPoliciesBonjourForwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesBonjourForwarding struct {
	value *NetworksNetworkIdGroupPoliciesBonjourForwarding
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesBonjourForwarding) Get() *NetworksNetworkIdGroupPoliciesBonjourForwarding {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesBonjourForwarding) Set(val *NetworksNetworkIdGroupPoliciesBonjourForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesBonjourForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesBonjourForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesBonjourForwarding(val *NetworksNetworkIdGroupPoliciesBonjourForwarding) *NullableNetworksNetworkIdGroupPoliciesBonjourForwarding {
	return &NullableNetworksNetworkIdGroupPoliciesBonjourForwarding{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesBonjourForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesBonjourForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


