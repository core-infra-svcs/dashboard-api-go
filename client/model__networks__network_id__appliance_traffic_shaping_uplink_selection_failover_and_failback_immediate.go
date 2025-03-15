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

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate Immediate WAN transition terminates all flows (new and existing) on current WAN when it is deemed unreliable.
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate struct {
	// Toggle for enabling or disabling immediate WAN failover and failback
	Enabled bool `json:"enabled"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate(enabled bool) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate{}
	this.Enabled = enabled
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediateWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediateWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) SetEnabled(v bool) {
	o.Enabled = v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


