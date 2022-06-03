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

// NetworksNetworkIdGroupPoliciesBandwidth     The bandwidth settings for clients bound to your group policy. 
type NetworksNetworkIdGroupPoliciesBandwidth struct {
	// How bandwidth limits are enforced. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	BandwidthLimits *NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesBandwidth instantiates a new NetworksNetworkIdGroupPoliciesBandwidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesBandwidth() *NetworksNetworkIdGroupPoliciesBandwidth {
	this := NetworksNetworkIdGroupPoliciesBandwidth{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesBandwidthWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesBandwidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesBandwidthWithDefaults() *NetworksNetworkIdGroupPoliciesBandwidth {
	this := NetworksNetworkIdGroupPoliciesBandwidth{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetSettings() string {
	if o == nil || o.Settings == nil {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetSettingsOk() (*string, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) SetSettings(v string) {
	o.Settings = &v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetBandwidthLimits() NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits {
	if o == nil || o.BandwidthLimits == nil {
		var ret NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) GetBandwidthLimitsOk() (*NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits, bool) {
	if o == nil || o.BandwidthLimits == nil {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) HasBandwidthLimits() bool {
	if o != nil && o.BandwidthLimits != nil {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *NetworksNetworkIdGroupPoliciesBandwidth) SetBandwidthLimits(v NetworksNetworkIdGroupPoliciesBandwidthBandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o NetworksNetworkIdGroupPoliciesBandwidth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.BandwidthLimits != nil {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesBandwidth struct {
	value *NetworksNetworkIdGroupPoliciesBandwidth
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesBandwidth) Get() *NetworksNetworkIdGroupPoliciesBandwidth {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesBandwidth) Set(val *NetworksNetworkIdGroupPoliciesBandwidth) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesBandwidth) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesBandwidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesBandwidth(val *NetworksNetworkIdGroupPoliciesBandwidth) *NullableNetworksNetworkIdGroupPoliciesBandwidth {
	return &NullableNetworksNetworkIdGroupPoliciesBandwidth{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesBandwidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesBandwidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


