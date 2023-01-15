/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess Details associated with a free access plan with limits.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess struct {
	// Whether or not free access is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// How long a device can use a network for free.
	DurationInMinutes *int32 `json:"durationInMinutes,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccessWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccessWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDurationInMinutes returns the DurationInMinutes field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) GetDurationInMinutes() int32 {
	if o == nil || isNil(o.DurationInMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.DurationInMinutes) {
    return nil, false
	}
	return o.DurationInMinutes, true
}

// HasDurationInMinutes returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) HasDurationInMinutes() bool {
	if o != nil && !isNil(o.DurationInMinutes) {
		return true
	}

	return false
}

// SetDurationInMinutes gets a reference to the given int32 and assigns it to the DurationInMinutes field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DurationInMinutes) {
		toSerialize["durationInMinutes"] = o.DurationInMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


