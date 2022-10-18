/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchRoutingMulticastDefaultSettings Default multicast setting for entire network. IGMP snooping and Flood unknown multicast traffic settings are enabled by default.
type NetworksNetworkIdSwitchRoutingMulticastDefaultSettings struct {
	// IGMP snooping setting for entire network
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled,omitempty"`
	// Flood unknown multicast traffic setting for entire network
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`
}

// NewNetworksNetworkIdSwitchRoutingMulticastDefaultSettings instantiates a new NetworksNetworkIdSwitchRoutingMulticastDefaultSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchRoutingMulticastDefaultSettings() *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	this := NetworksNetworkIdSwitchRoutingMulticastDefaultSettings{}
	return &this
}

// NewNetworksNetworkIdSwitchRoutingMulticastDefaultSettingsWithDefaults instantiates a new NetworksNetworkIdSwitchRoutingMulticastDefaultSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchRoutingMulticastDefaultSettingsWithDefaults() *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	this := NetworksNetworkIdSwitchRoutingMulticastDefaultSettings{}
	return &this
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) GetIgmpSnoopingEnabled() bool {
	if o == nil || o.IgmpSnoopingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || o.IgmpSnoopingEnabled == nil {
		return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) HasIgmpSnoopingEnabled() bool {
	if o != nil && o.IgmpSnoopingEnabled != nil {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || o.FloodUnknownMulticastTrafficEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || o.FloodUnknownMulticastTrafficEnabled == nil {
		return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && o.FloodUnknownMulticastTrafficEnabled != nil {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IgmpSnoopingEnabled != nil {
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if o.FloodUnknownMulticastTrafficEnabled != nil {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings struct {
	value *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings) Get() *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings) Set(val *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings(val *NetworksNetworkIdSwitchRoutingMulticastDefaultSettings) *NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings {
	return &NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchRoutingMulticastDefaultSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


