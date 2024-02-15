/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius RADIUS settings. This param is only valid when authMode is 'open-with-radius' and ipAssignmentMode is not 'NAT mode'.
type NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius struct {
	GuestVlan *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan `json:"guestVlan,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius() *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusWithDefaults() *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius{}
	return &this
}

// GetGuestVlan returns the GuestVlan field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) GetGuestVlan() NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan {
	if o == nil || isNil(o.GuestVlan) {
		var ret NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan
		return ret
	}
	return *o.GuestVlan
}

// GetGuestVlanOk returns a tuple with the GuestVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) GetGuestVlanOk() (*NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan, bool) {
	if o == nil || isNil(o.GuestVlan) {
    return nil, false
	}
	return o.GuestVlan, true
}

// HasGuestVlan returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) HasGuestVlan() bool {
	if o != nil && !isNil(o.GuestVlan) {
		return true
	}

	return false
}

// SetGuestVlan gets a reference to the given NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan and assigns it to the GuestVlan field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) SetGuestVlan(v NetworksNetworkIdWirelessSsidsNumberNamedVlansRadiusGuestVlan) {
	o.GuestVlan = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GuestVlan) {
		toSerialize["guestVlan"] = o.GuestVlan
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius struct {
	value *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) Get() *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) Set(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius(val *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius {
	return &NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


