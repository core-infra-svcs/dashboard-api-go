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

// NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 Splash authorization for SSID 14
type NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14{}
	return &this
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14WithDefaults instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14WithDefaults() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) GetIsAuthorized() bool {
	if o == nil || isNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || isNil(o.IsAuthorized) {
    return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) HasIsAuthorized() bool {
	if o != nil && !isNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 struct {
	value *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14
	isSet bool
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) Get() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) Set(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 {
	return &NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


