/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpServerPolicyArpInspection Dynamic ARP Inspection settings
type NetworksNetworkIdSwitchDhcpServerPolicyArpInspection struct {
	// Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspection instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspection() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspection{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpServerPolicyArpInspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpServerPolicyArpInspectionWithDefaults() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	this := NetworksNetworkIdSwitchDhcpServerPolicyArpInspection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection struct {
	value *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection) Get() *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection) Set(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection(val *NetworksNetworkIdSwitchDhcpServerPolicyArpInspection) *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection {
	return &NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpServerPolicyArpInspection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


