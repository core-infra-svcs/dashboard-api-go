/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance Current metrics on how the mesh is performing.
type NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance struct {
	// Average Mbps.
	Mbps *int32 `json:"mbps,omitempty"`
	// Represents the quality of the entire route from the repeater access point to its gateway access point.
	Metric *int32 `json:"metric,omitempty"`
	// Mesh utilization as a percentage.
	UsagePercentage *string `json:"usagePercentage,omitempty"`
}

// NewNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance instantiates a new NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance() *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance {
	this := NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance{}
	return &this
}

// NewNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformanceWithDefaults instantiates a new NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformanceWithDefaults() *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance {
	this := NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance{}
	return &this
}

// GetMbps returns the Mbps field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) GetMbps() int32 {
	if o == nil || isNil(o.Mbps) {
		var ret int32
		return ret
	}
	return *o.Mbps
}

// GetMbpsOk returns a tuple with the Mbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) GetMbpsOk() (*int32, bool) {
	if o == nil || isNil(o.Mbps) {
    return nil, false
	}
	return o.Mbps, true
}

// HasMbps returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) HasMbps() bool {
	if o != nil && !isNil(o.Mbps) {
		return true
	}

	return false
}

// SetMbps gets a reference to the given int32 and assigns it to the Mbps field.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) SetMbps(v int32) {
	o.Mbps = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) GetMetric() int32 {
	if o == nil || isNil(o.Metric) {
		var ret int32
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) GetMetricOk() (*int32, bool) {
	if o == nil || isNil(o.Metric) {
    return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) HasMetric() bool {
	if o != nil && !isNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given int32 and assigns it to the Metric field.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) SetMetric(v int32) {
	o.Metric = &v
}

// GetUsagePercentage returns the UsagePercentage field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) GetUsagePercentage() string {
	if o == nil || isNil(o.UsagePercentage) {
		var ret string
		return ret
	}
	return *o.UsagePercentage
}

// GetUsagePercentageOk returns a tuple with the UsagePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) GetUsagePercentageOk() (*string, bool) {
	if o == nil || isNil(o.UsagePercentage) {
    return nil, false
	}
	return o.UsagePercentage, true
}

// HasUsagePercentage returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) HasUsagePercentage() bool {
	if o != nil && !isNil(o.UsagePercentage) {
		return true
	}

	return false
}

// SetUsagePercentage gets a reference to the given string and assigns it to the UsagePercentage field.
func (o *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) SetUsagePercentage(v string) {
	o.UsagePercentage = &v
}

func (o NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mbps) {
		toSerialize["mbps"] = o.Mbps
	}
	if !isNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !isNil(o.UsagePercentage) {
		toSerialize["usagePercentage"] = o.UsagePercentage
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance struct {
	value *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) Get() *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) Set(val *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance(val *NetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) *NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance {
	return &NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessMeshStatusesLatestMeshPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


