/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits A mapping of uplinks to their bandwidth settings (be sure to check which uplinks are supported for your network)
type NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits struct {
	Wan1 *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 `json:"wan1,omitempty"`
	Wan2 *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 `json:"wan2,omitempty"`
	Cellular *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular `json:"cellular,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits() *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) GetWan1() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 {
	if o == nil || o.Wan1 == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) GetWan1Ok() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1, bool) {
	if o == nil || o.Wan1 == nil {
		return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) HasWan1() bool {
	if o != nil && o.Wan1 != nil {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 and assigns it to the Wan1 field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) SetWan1(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) GetWan2() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 {
	if o == nil || o.Wan2 == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) GetWan2Ok() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2, bool) {
	if o == nil || o.Wan2 == nil {
		return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) HasWan2() bool {
	if o != nil && o.Wan2 != nil {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 and assigns it to the Wan2 field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) SetWan2(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2) {
	o.Wan2 = &v
}

// GetCellular returns the Cellular field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) GetCellular() NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular {
	if o == nil || o.Cellular == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular
		return ret
	}
	return *o.Cellular
}

// GetCellularOk returns a tuple with the Cellular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) GetCellularOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular, bool) {
	if o == nil || o.Cellular == nil {
		return nil, false
	}
	return o.Cellular, true
}

// HasCellular returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) HasCellular() bool {
	if o != nil && o.Cellular != nil {
		return true
	}

	return false
}

// SetCellular gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular and assigns it to the Cellular field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) SetCellular(v NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular) {
	o.Cellular = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Wan1 != nil {
		toSerialize["wan1"] = o.Wan1
	}
	if o.Wan2 != nil {
		toSerialize["wan2"] = o.Wan2
	}
	if o.Cellular != nil {
		toSerialize["cellular"] = o.Cellular
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits(val *NetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkBandwidthBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


