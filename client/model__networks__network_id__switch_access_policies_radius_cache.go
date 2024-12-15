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

// NetworksNetworkIdSwitchAccessPoliciesRadiusCache Object for RADIUS Cache Settings
type NetworksNetworkIdSwitchAccessPoliciesRadiusCache struct {
	// Enable to cache authorization and authentication responses on the RADIUS server
	Enabled *bool `json:"enabled,omitempty"`
	// If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusCache instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCache() *NetworksNetworkIdSwitchAccessPoliciesRadiusCache {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusCache{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusCacheWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCacheWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusCache {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusCache{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadiusCache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadiusCache
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache) Get() *NetworksNetworkIdSwitchAccessPoliciesRadiusCache {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache(val *NetworksNetworkIdSwitchAccessPoliciesRadiusCache) *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


