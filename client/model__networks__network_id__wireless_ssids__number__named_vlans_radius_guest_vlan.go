/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan Guest VLAN settings. Used to direct traffic to a guest VLAN when none of the RADIUS servers are reachable or a client receives access-reject from the RADIUS server.
type NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan struct {
	// Whether or not RADIUS guest named VLAN is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// RADIUS guest VLAN name.
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan() *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlanWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlanWithDefaults() *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan struct {
	value *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) Get() *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) Set(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan {
	return &NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


