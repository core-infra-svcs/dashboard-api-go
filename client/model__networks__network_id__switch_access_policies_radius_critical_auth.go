/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth Critical auth settings for when authentication is rejected by the RADIUS server
type NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth struct {
	// VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	DataVlanId *int32 `json:"dataVlanId,omitempty"`
	// VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	VoiceVlanId *int32 `json:"voiceVlanId,omitempty"`
	// Enable to suspend port bounce when RADIUS servers are unreachable
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuthWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth{}
	return &this
}

// GetDataVlanId returns the DataVlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataVlanId() int32 {
	if o == nil || isNil(o.DataVlanId) {
		var ret int32
		return ret
	}
	return *o.DataVlanId
}

// GetDataVlanIdOk returns a tuple with the DataVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetDataVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.DataVlanId) {
    return nil, false
	}
	return o.DataVlanId, true
}

// HasDataVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasDataVlanId() bool {
	if o != nil && !isNil(o.DataVlanId) {
		return true
	}

	return false
}

// SetDataVlanId gets a reference to the given int32 and assigns it to the DataVlanId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetDataVlanId(v int32) {
	o.DataVlanId = &v
}

// GetVoiceVlanId returns the VoiceVlanId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceVlanId() int32 {
	if o == nil || isNil(o.VoiceVlanId) {
		var ret int32
		return ret
	}
	return *o.VoiceVlanId
}

// GetVoiceVlanIdOk returns a tuple with the VoiceVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetVoiceVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VoiceVlanId) {
    return nil, false
	}
	return o.VoiceVlanId, true
}

// HasVoiceVlanId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasVoiceVlanId() bool {
	if o != nil && !isNil(o.VoiceVlanId) {
		return true
	}

	return false
}

// SetVoiceVlanId gets a reference to the given int32 and assigns it to the VoiceVlanId field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetVoiceVlanId(v int32) {
	o.VoiceVlanId = &v
}

// GetSuspendPortBounce returns the SuspendPortBounce field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetSuspendPortBounce() bool {
	if o == nil || isNil(o.SuspendPortBounce) {
		var ret bool
		return ret
	}
	return *o.SuspendPortBounce
}

// GetSuspendPortBounceOk returns a tuple with the SuspendPortBounce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) GetSuspendPortBounceOk() (*bool, bool) {
	if o == nil || isNil(o.SuspendPortBounce) {
    return nil, false
	}
	return o.SuspendPortBounce, true
}

// HasSuspendPortBounce returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) HasSuspendPortBounce() bool {
	if o != nil && !isNil(o.SuspendPortBounce) {
		return true
	}

	return false
}

// SetSuspendPortBounce gets a reference to the given bool and assigns it to the SuspendPortBounce field.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) SetSuspendPortBounce(v bool) {
	o.SuspendPortBounce = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DataVlanId) {
		toSerialize["dataVlanId"] = o.DataVlanId
	}
	if !isNil(o.VoiceVlanId) {
		toSerialize["voiceVlanId"] = o.VoiceVlanId
	}
	if !isNil(o.SuspendPortBounce) {
		toSerialize["suspendPortBounce"] = o.SuspendPortBounce
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) Get() *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth(val *NetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusCriticalAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


