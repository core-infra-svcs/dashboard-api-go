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

// InlineResponse200136 struct for InlineResponse200136
type InlineResponse200136 struct {
	// The percentage of CPU used as a decimal format.
	CpuPercentUsed *float32 `json:"cpuPercentUsed,omitempty"`
	// Memory that is not yet in use by the system.
	MemFree *int32 `json:"memFree,omitempty"`
	// Memory used for core OS functions on the device.
	MemWired *int32 `json:"memWired,omitempty"`
	// The active RAM on the device.
	MemActive *int32 `json:"memActive,omitempty"`
	// The inactive RAM on the device.
	MemInactive *int32 `json:"memInactive,omitempty"`
	// Network bandwith transmitted.
	NetworkSent *int32 `json:"networkSent,omitempty"`
	// Network bandwith received.
	NetworkReceived *int32 `json:"networkReceived,omitempty"`
	// The amount of space being used on the startup disk to swap unused files to and from RAM.
	SwapUsed *int32 `json:"swapUsed,omitempty"`
	DiskUsage *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage `json:"diskUsage,omitempty"`
	// The time at which the performance was measured.
	Ts *string `json:"ts,omitempty"`
}

// NewInlineResponse200136 instantiates a new InlineResponse200136 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200136() *InlineResponse200136 {
	this := InlineResponse200136{}
	return &this
}

// NewInlineResponse200136WithDefaults instantiates a new InlineResponse200136 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200136WithDefaults() *InlineResponse200136 {
	this := InlineResponse200136{}
	return &this
}

// GetCpuPercentUsed returns the CpuPercentUsed field value if set, zero value otherwise.
func (o *InlineResponse200136) GetCpuPercentUsed() float32 {
	if o == nil || isNil(o.CpuPercentUsed) {
		var ret float32
		return ret
	}
	return *o.CpuPercentUsed
}

// GetCpuPercentUsedOk returns a tuple with the CpuPercentUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetCpuPercentUsedOk() (*float32, bool) {
	if o == nil || isNil(o.CpuPercentUsed) {
    return nil, false
	}
	return o.CpuPercentUsed, true
}

// HasCpuPercentUsed returns a boolean if a field has been set.
func (o *InlineResponse200136) HasCpuPercentUsed() bool {
	if o != nil && !isNil(o.CpuPercentUsed) {
		return true
	}

	return false
}

// SetCpuPercentUsed gets a reference to the given float32 and assigns it to the CpuPercentUsed field.
func (o *InlineResponse200136) SetCpuPercentUsed(v float32) {
	o.CpuPercentUsed = &v
}

// GetMemFree returns the MemFree field value if set, zero value otherwise.
func (o *InlineResponse200136) GetMemFree() int32 {
	if o == nil || isNil(o.MemFree) {
		var ret int32
		return ret
	}
	return *o.MemFree
}

// GetMemFreeOk returns a tuple with the MemFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetMemFreeOk() (*int32, bool) {
	if o == nil || isNil(o.MemFree) {
    return nil, false
	}
	return o.MemFree, true
}

// HasMemFree returns a boolean if a field has been set.
func (o *InlineResponse200136) HasMemFree() bool {
	if o != nil && !isNil(o.MemFree) {
		return true
	}

	return false
}

// SetMemFree gets a reference to the given int32 and assigns it to the MemFree field.
func (o *InlineResponse200136) SetMemFree(v int32) {
	o.MemFree = &v
}

// GetMemWired returns the MemWired field value if set, zero value otherwise.
func (o *InlineResponse200136) GetMemWired() int32 {
	if o == nil || isNil(o.MemWired) {
		var ret int32
		return ret
	}
	return *o.MemWired
}

// GetMemWiredOk returns a tuple with the MemWired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetMemWiredOk() (*int32, bool) {
	if o == nil || isNil(o.MemWired) {
    return nil, false
	}
	return o.MemWired, true
}

// HasMemWired returns a boolean if a field has been set.
func (o *InlineResponse200136) HasMemWired() bool {
	if o != nil && !isNil(o.MemWired) {
		return true
	}

	return false
}

// SetMemWired gets a reference to the given int32 and assigns it to the MemWired field.
func (o *InlineResponse200136) SetMemWired(v int32) {
	o.MemWired = &v
}

// GetMemActive returns the MemActive field value if set, zero value otherwise.
func (o *InlineResponse200136) GetMemActive() int32 {
	if o == nil || isNil(o.MemActive) {
		var ret int32
		return ret
	}
	return *o.MemActive
}

// GetMemActiveOk returns a tuple with the MemActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetMemActiveOk() (*int32, bool) {
	if o == nil || isNil(o.MemActive) {
    return nil, false
	}
	return o.MemActive, true
}

// HasMemActive returns a boolean if a field has been set.
func (o *InlineResponse200136) HasMemActive() bool {
	if o != nil && !isNil(o.MemActive) {
		return true
	}

	return false
}

// SetMemActive gets a reference to the given int32 and assigns it to the MemActive field.
func (o *InlineResponse200136) SetMemActive(v int32) {
	o.MemActive = &v
}

// GetMemInactive returns the MemInactive field value if set, zero value otherwise.
func (o *InlineResponse200136) GetMemInactive() int32 {
	if o == nil || isNil(o.MemInactive) {
		var ret int32
		return ret
	}
	return *o.MemInactive
}

// GetMemInactiveOk returns a tuple with the MemInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetMemInactiveOk() (*int32, bool) {
	if o == nil || isNil(o.MemInactive) {
    return nil, false
	}
	return o.MemInactive, true
}

// HasMemInactive returns a boolean if a field has been set.
func (o *InlineResponse200136) HasMemInactive() bool {
	if o != nil && !isNil(o.MemInactive) {
		return true
	}

	return false
}

// SetMemInactive gets a reference to the given int32 and assigns it to the MemInactive field.
func (o *InlineResponse200136) SetMemInactive(v int32) {
	o.MemInactive = &v
}

// GetNetworkSent returns the NetworkSent field value if set, zero value otherwise.
func (o *InlineResponse200136) GetNetworkSent() int32 {
	if o == nil || isNil(o.NetworkSent) {
		var ret int32
		return ret
	}
	return *o.NetworkSent
}

// GetNetworkSentOk returns a tuple with the NetworkSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetNetworkSentOk() (*int32, bool) {
	if o == nil || isNil(o.NetworkSent) {
    return nil, false
	}
	return o.NetworkSent, true
}

// HasNetworkSent returns a boolean if a field has been set.
func (o *InlineResponse200136) HasNetworkSent() bool {
	if o != nil && !isNil(o.NetworkSent) {
		return true
	}

	return false
}

// SetNetworkSent gets a reference to the given int32 and assigns it to the NetworkSent field.
func (o *InlineResponse200136) SetNetworkSent(v int32) {
	o.NetworkSent = &v
}

// GetNetworkReceived returns the NetworkReceived field value if set, zero value otherwise.
func (o *InlineResponse200136) GetNetworkReceived() int32 {
	if o == nil || isNil(o.NetworkReceived) {
		var ret int32
		return ret
	}
	return *o.NetworkReceived
}

// GetNetworkReceivedOk returns a tuple with the NetworkReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetNetworkReceivedOk() (*int32, bool) {
	if o == nil || isNil(o.NetworkReceived) {
    return nil, false
	}
	return o.NetworkReceived, true
}

// HasNetworkReceived returns a boolean if a field has been set.
func (o *InlineResponse200136) HasNetworkReceived() bool {
	if o != nil && !isNil(o.NetworkReceived) {
		return true
	}

	return false
}

// SetNetworkReceived gets a reference to the given int32 and assigns it to the NetworkReceived field.
func (o *InlineResponse200136) SetNetworkReceived(v int32) {
	o.NetworkReceived = &v
}

// GetSwapUsed returns the SwapUsed field value if set, zero value otherwise.
func (o *InlineResponse200136) GetSwapUsed() int32 {
	if o == nil || isNil(o.SwapUsed) {
		var ret int32
		return ret
	}
	return *o.SwapUsed
}

// GetSwapUsedOk returns a tuple with the SwapUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetSwapUsedOk() (*int32, bool) {
	if o == nil || isNil(o.SwapUsed) {
    return nil, false
	}
	return o.SwapUsed, true
}

// HasSwapUsed returns a boolean if a field has been set.
func (o *InlineResponse200136) HasSwapUsed() bool {
	if o != nil && !isNil(o.SwapUsed) {
		return true
	}

	return false
}

// SetSwapUsed gets a reference to the given int32 and assigns it to the SwapUsed field.
func (o *InlineResponse200136) SetSwapUsed(v int32) {
	o.SwapUsed = &v
}

// GetDiskUsage returns the DiskUsage field value if set, zero value otherwise.
func (o *InlineResponse200136) GetDiskUsage() NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage {
	if o == nil || isNil(o.DiskUsage) {
		var ret NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage
		return ret
	}
	return *o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetDiskUsageOk() (*NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage, bool) {
	if o == nil || isNil(o.DiskUsage) {
    return nil, false
	}
	return o.DiskUsage, true
}

// HasDiskUsage returns a boolean if a field has been set.
func (o *InlineResponse200136) HasDiskUsage() bool {
	if o != nil && !isNil(o.DiskUsage) {
		return true
	}

	return false
}

// SetDiskUsage gets a reference to the given NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage and assigns it to the DiskUsage field.
func (o *InlineResponse200136) SetDiskUsage(v NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsage) {
	o.DiskUsage = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200136) GetTs() string {
	if o == nil || isNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetTsOk() (*string, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200136) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *InlineResponse200136) SetTs(v string) {
	o.Ts = &v
}

func (o InlineResponse200136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CpuPercentUsed) {
		toSerialize["cpuPercentUsed"] = o.CpuPercentUsed
	}
	if !isNil(o.MemFree) {
		toSerialize["memFree"] = o.MemFree
	}
	if !isNil(o.MemWired) {
		toSerialize["memWired"] = o.MemWired
	}
	if !isNil(o.MemActive) {
		toSerialize["memActive"] = o.MemActive
	}
	if !isNil(o.MemInactive) {
		toSerialize["memInactive"] = o.MemInactive
	}
	if !isNil(o.NetworkSent) {
		toSerialize["networkSent"] = o.NetworkSent
	}
	if !isNil(o.NetworkReceived) {
		toSerialize["networkReceived"] = o.NetworkReceived
	}
	if !isNil(o.SwapUsed) {
		toSerialize["swapUsed"] = o.SwapUsed
	}
	if !isNil(o.DiskUsage) {
		toSerialize["diskUsage"] = o.DiskUsage
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200136 struct {
	value *InlineResponse200136
	isSet bool
}

func (v NullableInlineResponse200136) Get() *InlineResponse200136 {
	return v.value
}

func (v *NullableInlineResponse200136) Set(val *InlineResponse200136) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200136) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200136) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200136(val *InlineResponse200136) *NullableInlineResponse200136 {
	return &NullableInlineResponse200136{value: val, isSet: true}
}

func (v NullableInlineResponse200136) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200136) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


