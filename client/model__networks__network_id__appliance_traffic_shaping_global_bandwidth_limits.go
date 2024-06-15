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

// NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits Global per-client bandwidth limit
type NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits struct {
	// The upload bandwidth limit in Kbps. (0 represents no limit.)
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The download bandwidth limit in Kbps. (0 represents no limit.)
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits instantiates a new NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits() *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits {
	this := NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimitsWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimitsWithDefaults() *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits {
	this := NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits struct {
	value *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) Get() *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) Set(val *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits(val *NetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) *NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits {
	return &NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingGlobalBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


