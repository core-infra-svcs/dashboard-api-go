/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention Smart Retention records footage in two qualities and intelligently retains higher quality when motion, people or vehicles are detected.
type NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention struct {
	// Boolean indicating if Smart Retention is enabled(true) or disabled(false).
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention() *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention {
	this := NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention{}
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesSmartRetentionWithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesSmartRetentionWithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention {
	this := NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) Get() *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention(val *NetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) *NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesSmartRetention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


