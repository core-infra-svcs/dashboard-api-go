/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback WAN failover and failback behavior
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback struct {
	Immediate *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) GetImmediate() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) GetImmediateOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) SetImmediate(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionFailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


