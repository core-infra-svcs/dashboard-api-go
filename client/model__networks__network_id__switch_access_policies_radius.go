/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesRadius Object for RADIUS Settings
type NetworksNetworkIdSwitchAccessPoliciesRadius struct {
	CriticalAuth *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth `json:"criticalAuth,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadius instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadius() *NetworksNetworkIdSwitchAccessPoliciesRadius {
	this := NetworksNetworkIdSwitchAccessPoliciesRadius{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadius object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadius {
	this := NetworksNetworkIdSwitchAccessPoliciesRadius{}
	return &this
}

// GetCriticalAuth returns the CriticalAuth field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuth() NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth {
	if o == nil || o.CriticalAuth == nil {
		var ret NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth
		return ret
	}
	return *o.CriticalAuth
}

// GetCriticalAuthOk returns a tuple with the CriticalAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuthOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth, bool) {
	if o == nil || o.CriticalAuth == nil {
		return nil, false
	}
	return o.CriticalAuth, true
}

// HasCriticalAuth returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasCriticalAuth() bool {
	if o != nil && o.CriticalAuth != nil {
		return true
	}

	return false
}

// SetCriticalAuth gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth and assigns it to the CriticalAuth field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetCriticalAuth(v NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) {
	o.CriticalAuth = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadius) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CriticalAuth != nil {
		toSerialize["criticalAuth"] = o.CriticalAuth
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesRadius struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadius
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadius) Get() *NetworksNetworkIdSwitchAccessPoliciesRadius {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadius) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadius(val *NetworksNetworkIdSwitchAccessPoliciesRadius) *NullableNetworksNetworkIdSwitchAccessPoliciesRadius {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadius{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

