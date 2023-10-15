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

// NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail Email alert settings for DHCP servers
type NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail struct {
	// When enabled, send an email if a new DHCP server is seen. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail() *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail {
	this := NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmailWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmailWithDefaults() *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail {
	this := NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail struct {
	value *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) Get() *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) Set(val *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail(val *NetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail {
	return &NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyAlertsEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


