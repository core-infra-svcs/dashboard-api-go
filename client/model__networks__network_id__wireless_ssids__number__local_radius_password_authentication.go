/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication The current setting for password-based authentication.
type NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication struct {
	// Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthenticationWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthenticationWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication struct {
	value *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) Get() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) Set(val *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication(val *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


