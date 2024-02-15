/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage An object containing disk usage details.
type NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage struct {
	C *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC `json:"c,omitempty"`
}

// NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage instantiates a new NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage() *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage {
	this := NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage{}
	return &this
}

// NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageWithDefaults instantiates a new NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageWithDefaults() *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage {
	this := NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage{}
	return &this
}

// GetC returns the C field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) GetC() NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC {
	if o == nil || isNil(o.C) {
		var ret NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) GetCOk() (*NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC, bool) {
	if o == nil || isNil(o.C) {
    return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) HasC() bool {
	if o != nil && !isNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC and assigns it to the C field.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) SetC(v NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) {
	o.C = &v
}

func (o NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.C) {
		toSerialize["c"] = o.C
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage struct {
	value *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage
	isSet bool
}

func (v NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) Get() *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage {
	return v.value
}

func (v *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) Set(val *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage(val *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage {
	return &NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


