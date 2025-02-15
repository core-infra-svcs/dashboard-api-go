/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesRadius Object for RADIUS Settings
type NetworksNetworkIdSwitchAccessPoliciesRadius struct {
	CriticalAuth *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth `json:"criticalAuth,omitempty"`
	// VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	FailedAuthVlanId *int32 `json:"failedAuthVlanId,omitempty"`
	// Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval *int32 `json:"reAuthenticationInterval,omitempty"`
	Cache *NetworksNetworkIdSwitchAccessPoliciesRadiusCache `json:"cache,omitempty"`
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
	if o == nil || isNil(o.CriticalAuth) {
		var ret NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth
		return ret
	}
	return *o.CriticalAuth
}

// GetCriticalAuthOk returns a tuple with the CriticalAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCriticalAuthOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth, bool) {
	if o == nil || isNil(o.CriticalAuth) {
    return nil, false
	}
	return o.CriticalAuth, true
}

// HasCriticalAuth returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasCriticalAuth() bool {
	if o != nil && !isNil(o.CriticalAuth) {
		return true
	}

	return false
}

// SetCriticalAuth gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth and assigns it to the CriticalAuth field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetCriticalAuth(v NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) {
	o.CriticalAuth = &v
}

// GetFailedAuthVlanId returns the FailedAuthVlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthVlanId() int32 {
	if o == nil || isNil(o.FailedAuthVlanId) {
		var ret int32
		return ret
	}
	return *o.FailedAuthVlanId
}

// GetFailedAuthVlanIdOk returns a tuple with the FailedAuthVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetFailedAuthVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.FailedAuthVlanId) {
    return nil, false
	}
	return o.FailedAuthVlanId, true
}

// HasFailedAuthVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasFailedAuthVlanId() bool {
	if o != nil && !isNil(o.FailedAuthVlanId) {
		return true
	}

	return false
}

// SetFailedAuthVlanId gets a reference to the given int32 and assigns it to the FailedAuthVlanId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetFailedAuthVlanId(v int32) {
	o.FailedAuthVlanId = &v
}

// GetReAuthenticationInterval returns the ReAuthenticationInterval field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetReAuthenticationInterval() int32 {
	if o == nil || isNil(o.ReAuthenticationInterval) {
		var ret int32
		return ret
	}
	return *o.ReAuthenticationInterval
}

// GetReAuthenticationIntervalOk returns a tuple with the ReAuthenticationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetReAuthenticationIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.ReAuthenticationInterval) {
    return nil, false
	}
	return o.ReAuthenticationInterval, true
}

// HasReAuthenticationInterval returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasReAuthenticationInterval() bool {
	if o != nil && !isNil(o.ReAuthenticationInterval) {
		return true
	}

	return false
}

// SetReAuthenticationInterval gets a reference to the given int32 and assigns it to the ReAuthenticationInterval field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetReAuthenticationInterval(v int32) {
	o.ReAuthenticationInterval = &v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCache() NetworksNetworkIdSwitchAccessPoliciesRadiusCache {
	if o == nil || isNil(o.Cache) {
		var ret NetworksNetworkIdSwitchAccessPoliciesRadiusCache
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) GetCacheOk() (*NetworksNetworkIdSwitchAccessPoliciesRadiusCache, bool) {
	if o == nil || isNil(o.Cache) {
    return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) HasCache() bool {
	if o != nil && !isNil(o.Cache) {
		return true
	}

	return false
}

// SetCache gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesRadiusCache and assigns it to the Cache field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadius) SetCache(v NetworksNetworkIdSwitchAccessPoliciesRadiusCache) {
	o.Cache = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadius) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CriticalAuth) {
		toSerialize["criticalAuth"] = o.CriticalAuth
	}
	if !isNil(o.FailedAuthVlanId) {
		toSerialize["failedAuthVlanId"] = o.FailedAuthVlanId
	}
	if !isNil(o.ReAuthenticationInterval) {
		toSerialize["reAuthenticationInterval"] = o.ReAuthenticationInterval
	}
	if !isNil(o.Cache) {
		toSerialize["cache"] = o.Cache
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


