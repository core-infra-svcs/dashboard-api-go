/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection Spoofing protection settings
type NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection struct {
	IpSourceGuard *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard `json:"ipSourceGuard,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection instantiates a new NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection() *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection {
	this := NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection{}
	return &this
}

// NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionWithDefaults() *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection {
	this := NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection{}
	return &this
}

// GetIpSourceGuard returns the IpSourceGuard field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) GetIpSourceGuard() NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard {
	if o == nil || isNil(o.IpSourceGuard) {
		var ret NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard
		return ret
	}
	return *o.IpSourceGuard
}

// GetIpSourceGuardOk returns a tuple with the IpSourceGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) GetIpSourceGuardOk() (*NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard, bool) {
	if o == nil || isNil(o.IpSourceGuard) {
    return nil, false
	}
	return o.IpSourceGuard, true
}

// HasIpSourceGuard returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) HasIpSourceGuard() bool {
	if o != nil && !isNil(o.IpSourceGuard) {
		return true
	}

	return false
}

// SetIpSourceGuard gets a reference to the given NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard and assigns it to the IpSourceGuard field.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) SetIpSourceGuard(v NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) {
	o.IpSourceGuard = &v
}

func (o NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpSourceGuard) {
		toSerialize["ipSourceGuard"] = o.IpSourceGuard
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection struct {
	value *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) Get() *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) Set(val *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection(val *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection {
	return &NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


