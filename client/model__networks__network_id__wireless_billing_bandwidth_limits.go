/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessBillingBandwidthLimits The uplink bandwidth settings for the pricing plan.
type NetworksNetworkIdWirelessBillingBandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewNetworksNetworkIdWirelessBillingBandwidthLimits instantiates a new NetworksNetworkIdWirelessBillingBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessBillingBandwidthLimits() *NetworksNetworkIdWirelessBillingBandwidthLimits {
	this := NetworksNetworkIdWirelessBillingBandwidthLimits{}
	return &this
}

// NewNetworksNetworkIdWirelessBillingBandwidthLimitsWithDefaults instantiates a new NetworksNetworkIdWirelessBillingBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessBillingBandwidthLimitsWithDefaults() *NetworksNetworkIdWirelessBillingBandwidthLimits {
	this := NetworksNetworkIdWirelessBillingBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *NetworksNetworkIdWirelessBillingBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o NetworksNetworkIdWirelessBillingBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessBillingBandwidthLimits struct {
	value *NetworksNetworkIdWirelessBillingBandwidthLimits
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessBillingBandwidthLimits) Get() *NetworksNetworkIdWirelessBillingBandwidthLimits {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessBillingBandwidthLimits) Set(val *NetworksNetworkIdWirelessBillingBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessBillingBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessBillingBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessBillingBandwidthLimits(val *NetworksNetworkIdWirelessBillingBandwidthLimits) *NullableNetworksNetworkIdWirelessBillingBandwidthLimits {
	return &NullableNetworksNetworkIdWirelessBillingBandwidthLimits{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessBillingBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessBillingBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


