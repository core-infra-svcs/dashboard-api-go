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

// NetworksNetworkIdApplianceSettingsDynamicDns Dynamic DNS settings for a network
type NetworksNetworkIdApplianceSettingsDynamicDns struct {
	// Dynamic DNS url prefix. DDNS must be enabled to update
	Prefix *string `json:"prefix,omitempty"`
	// Dynamic DNS enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdApplianceSettingsDynamicDns instantiates a new NetworksNetworkIdApplianceSettingsDynamicDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSettingsDynamicDns() *NetworksNetworkIdApplianceSettingsDynamicDns {
	this := NetworksNetworkIdApplianceSettingsDynamicDns{}
	return &this
}

// NewNetworksNetworkIdApplianceSettingsDynamicDnsWithDefaults instantiates a new NetworksNetworkIdApplianceSettingsDynamicDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSettingsDynamicDnsWithDefaults() *NetworksNetworkIdApplianceSettingsDynamicDns {
	this := NetworksNetworkIdApplianceSettingsDynamicDns{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) SetPrefix(v string) {
	o.Prefix = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdApplianceSettingsDynamicDns) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdApplianceSettingsDynamicDns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSettingsDynamicDns struct {
	value *NetworksNetworkIdApplianceSettingsDynamicDns
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSettingsDynamicDns) Get() *NetworksNetworkIdApplianceSettingsDynamicDns {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSettingsDynamicDns) Set(val *NetworksNetworkIdApplianceSettingsDynamicDns) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSettingsDynamicDns) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSettingsDynamicDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSettingsDynamicDns(val *NetworksNetworkIdApplianceSettingsDynamicDns) *NullableNetworksNetworkIdApplianceSettingsDynamicDns {
	return &NullableNetworksNetworkIdApplianceSettingsDynamicDns{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSettingsDynamicDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSettingsDynamicDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


