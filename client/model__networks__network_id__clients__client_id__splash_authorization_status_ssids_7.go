/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 Splash authorization for SSID 7
type NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7{}
	return &this
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7WithDefaults instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7WithDefaults() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) GetIsAuthorized() bool {
	if o == nil || o.IsAuthorized == nil {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || o.IsAuthorized == nil {
		return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) HasIsAuthorized() bool {
	if o != nil && o.IsAuthorized != nil {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAuthorized != nil {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 struct {
	value *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7
	isSet bool
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) Get() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) Set(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 {
	return &NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


