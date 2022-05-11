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

// NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits     An object describing the bandwidth settings for your rule. 
type NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits struct {
	// How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	BandwidthLimits *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits instantiates a new NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits() *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	this := NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsWithDefaults() *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	this := NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) GetSettings() string {
	if o == nil || o.Settings == nil {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) GetSettingsOk() (*string, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) SetSettings(v string) {
	o.Settings = &v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) GetBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits {
	if o == nil || o.BandwidthLimits == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) GetBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits, bool) {
	if o == nil || o.BandwidthLimits == nil {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) HasBandwidthLimits() bool {
	if o != nil && o.BandwidthLimits != nil {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) SetBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimitsBandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.BandwidthLimits != nil {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits struct {
	value *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) Get() *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) Set(val *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits(val *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) *NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	return &NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


