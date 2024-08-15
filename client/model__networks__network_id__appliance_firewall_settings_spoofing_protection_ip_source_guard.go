/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard IP source address spoofing settings
type NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard struct {
	// Mode of protection
	Mode *string `json:"mode,omitempty"`
}

// NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard instantiates a new NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard() *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard {
	this := NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard{}
	return &this
}

// NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuardWithDefaults instantiates a new NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuardWithDefaults() *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard {
	this := NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) SetMode(v string) {
	o.Mode = &v
}

func (o NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard struct {
	value *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) Get() *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) Set(val *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard(val *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard {
	return &NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceFirewallSettingsSpoofingProtectionIpSourceGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


