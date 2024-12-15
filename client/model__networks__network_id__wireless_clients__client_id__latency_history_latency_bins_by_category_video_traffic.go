/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic The time bucket's video traffic latency history
type NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic struct {
	// The latency bucket for video traffic in 0.5 seconds
	Var05 *int32 `json:"0.5,omitempty"`
	// The latency bucket for video traffic in 1.0 seconds
	Var10 *int32 `json:"1.0,omitempty"`
	// The latency bucket for video traffic in 2.0 seconds
	Var20 *int32 `json:"2.0,omitempty"`
	// The latency bucket for video traffic in 4.0 seconds
	Var40 *int32 `json:"4.0,omitempty"`
	// The latency bucket for video traffic in 8.0 seconds
	Var80 *int32 `json:"8.0,omitempty"`
	// The latency bucket for video traffic in 16.0 seconds
	Var160 *int32 `json:"16.0,omitempty"`
	// The latency bucket for video traffic in 32.0 seconds
	Var320 *int32 `json:"32.0,omitempty"`
	// The latency bucket for video traffic in 64.0 seconds
	Var640 *int32 `json:"64.0,omitempty"`
	// The latency bucket for video traffic in 128.0 seconds
	Var1280 *int32 `json:"128.0,omitempty"`
	// The latency bucket for video traffic in 256.0 seconds
	Var2560 *int32 `json:"256.0,omitempty"`
	// The latency bucket for video traffic in 512.0 seconds
	Var5120 *int32 `json:"512.0,omitempty"`
	// The latency bucket for video traffic in 1024.0 seconds
	Var10240 *int32 `json:"1024.0,omitempty"`
	// The latency bucket for video traffic in 2048.0 seconds
	Var20480 *int32 `json:"2048.0,omitempty"`
}

// NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic instantiates a new NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic() *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic {
	this := NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic{}
	return &this
}

// NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTrafficWithDefaults instantiates a new NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTrafficWithDefaults() *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic {
	this := NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic{}
	return &this
}

// GetVar05 returns the Var05 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar05() int32 {
	if o == nil || isNil(o.Var05) {
		var ret int32
		return ret
	}
	return *o.Var05
}

// GetVar05Ok returns a tuple with the Var05 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar05Ok() (*int32, bool) {
	if o == nil || isNil(o.Var05) {
    return nil, false
	}
	return o.Var05, true
}

// HasVar05 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar05() bool {
	if o != nil && !isNil(o.Var05) {
		return true
	}

	return false
}

// SetVar05 gets a reference to the given int32 and assigns it to the Var05 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar05(v int32) {
	o.Var05 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar10() int32 {
	if o == nil || isNil(o.Var10) {
		var ret int32
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar10Ok() (*int32, bool) {
	if o == nil || isNil(o.Var10) {
    return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar10() bool {
	if o != nil && !isNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given int32 and assigns it to the Var10 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar10(v int32) {
	o.Var10 = &v
}

// GetVar20 returns the Var20 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar20() int32 {
	if o == nil || isNil(o.Var20) {
		var ret int32
		return ret
	}
	return *o.Var20
}

// GetVar20Ok returns a tuple with the Var20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar20Ok() (*int32, bool) {
	if o == nil || isNil(o.Var20) {
    return nil, false
	}
	return o.Var20, true
}

// HasVar20 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar20() bool {
	if o != nil && !isNil(o.Var20) {
		return true
	}

	return false
}

// SetVar20 gets a reference to the given int32 and assigns it to the Var20 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar20(v int32) {
	o.Var20 = &v
}

// GetVar40 returns the Var40 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar40() int32 {
	if o == nil || isNil(o.Var40) {
		var ret int32
		return ret
	}
	return *o.Var40
}

// GetVar40Ok returns a tuple with the Var40 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar40Ok() (*int32, bool) {
	if o == nil || isNil(o.Var40) {
    return nil, false
	}
	return o.Var40, true
}

// HasVar40 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar40() bool {
	if o != nil && !isNil(o.Var40) {
		return true
	}

	return false
}

// SetVar40 gets a reference to the given int32 and assigns it to the Var40 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar40(v int32) {
	o.Var40 = &v
}

// GetVar80 returns the Var80 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar80() int32 {
	if o == nil || isNil(o.Var80) {
		var ret int32
		return ret
	}
	return *o.Var80
}

// GetVar80Ok returns a tuple with the Var80 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar80Ok() (*int32, bool) {
	if o == nil || isNil(o.Var80) {
    return nil, false
	}
	return o.Var80, true
}

// HasVar80 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar80() bool {
	if o != nil && !isNil(o.Var80) {
		return true
	}

	return false
}

// SetVar80 gets a reference to the given int32 and assigns it to the Var80 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar80(v int32) {
	o.Var80 = &v
}

// GetVar160 returns the Var160 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar160() int32 {
	if o == nil || isNil(o.Var160) {
		var ret int32
		return ret
	}
	return *o.Var160
}

// GetVar160Ok returns a tuple with the Var160 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar160Ok() (*int32, bool) {
	if o == nil || isNil(o.Var160) {
    return nil, false
	}
	return o.Var160, true
}

// HasVar160 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar160() bool {
	if o != nil && !isNil(o.Var160) {
		return true
	}

	return false
}

// SetVar160 gets a reference to the given int32 and assigns it to the Var160 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar160(v int32) {
	o.Var160 = &v
}

// GetVar320 returns the Var320 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar320() int32 {
	if o == nil || isNil(o.Var320) {
		var ret int32
		return ret
	}
	return *o.Var320
}

// GetVar320Ok returns a tuple with the Var320 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar320Ok() (*int32, bool) {
	if o == nil || isNil(o.Var320) {
    return nil, false
	}
	return o.Var320, true
}

// HasVar320 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar320() bool {
	if o != nil && !isNil(o.Var320) {
		return true
	}

	return false
}

// SetVar320 gets a reference to the given int32 and assigns it to the Var320 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar320(v int32) {
	o.Var320 = &v
}

// GetVar640 returns the Var640 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar640() int32 {
	if o == nil || isNil(o.Var640) {
		var ret int32
		return ret
	}
	return *o.Var640
}

// GetVar640Ok returns a tuple with the Var640 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar640Ok() (*int32, bool) {
	if o == nil || isNil(o.Var640) {
    return nil, false
	}
	return o.Var640, true
}

// HasVar640 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar640() bool {
	if o != nil && !isNil(o.Var640) {
		return true
	}

	return false
}

// SetVar640 gets a reference to the given int32 and assigns it to the Var640 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar640(v int32) {
	o.Var640 = &v
}

// GetVar1280 returns the Var1280 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar1280() int32 {
	if o == nil || isNil(o.Var1280) {
		var ret int32
		return ret
	}
	return *o.Var1280
}

// GetVar1280Ok returns a tuple with the Var1280 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar1280Ok() (*int32, bool) {
	if o == nil || isNil(o.Var1280) {
    return nil, false
	}
	return o.Var1280, true
}

// HasVar1280 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar1280() bool {
	if o != nil && !isNil(o.Var1280) {
		return true
	}

	return false
}

// SetVar1280 gets a reference to the given int32 and assigns it to the Var1280 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar1280(v int32) {
	o.Var1280 = &v
}

// GetVar2560 returns the Var2560 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar2560() int32 {
	if o == nil || isNil(o.Var2560) {
		var ret int32
		return ret
	}
	return *o.Var2560
}

// GetVar2560Ok returns a tuple with the Var2560 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar2560Ok() (*int32, bool) {
	if o == nil || isNil(o.Var2560) {
    return nil, false
	}
	return o.Var2560, true
}

// HasVar2560 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar2560() bool {
	if o != nil && !isNil(o.Var2560) {
		return true
	}

	return false
}

// SetVar2560 gets a reference to the given int32 and assigns it to the Var2560 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar2560(v int32) {
	o.Var2560 = &v
}

// GetVar5120 returns the Var5120 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar5120() int32 {
	if o == nil || isNil(o.Var5120) {
		var ret int32
		return ret
	}
	return *o.Var5120
}

// GetVar5120Ok returns a tuple with the Var5120 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar5120Ok() (*int32, bool) {
	if o == nil || isNil(o.Var5120) {
    return nil, false
	}
	return o.Var5120, true
}

// HasVar5120 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar5120() bool {
	if o != nil && !isNil(o.Var5120) {
		return true
	}

	return false
}

// SetVar5120 gets a reference to the given int32 and assigns it to the Var5120 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar5120(v int32) {
	o.Var5120 = &v
}

// GetVar10240 returns the Var10240 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar10240() int32 {
	if o == nil || isNil(o.Var10240) {
		var ret int32
		return ret
	}
	return *o.Var10240
}

// GetVar10240Ok returns a tuple with the Var10240 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar10240Ok() (*int32, bool) {
	if o == nil || isNil(o.Var10240) {
    return nil, false
	}
	return o.Var10240, true
}

// HasVar10240 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar10240() bool {
	if o != nil && !isNil(o.Var10240) {
		return true
	}

	return false
}

// SetVar10240 gets a reference to the given int32 and assigns it to the Var10240 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar10240(v int32) {
	o.Var10240 = &v
}

// GetVar20480 returns the Var20480 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar20480() int32 {
	if o == nil || isNil(o.Var20480) {
		var ret int32
		return ret
	}
	return *o.Var20480
}

// GetVar20480Ok returns a tuple with the Var20480 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) GetVar20480Ok() (*int32, bool) {
	if o == nil || isNil(o.Var20480) {
    return nil, false
	}
	return o.Var20480, true
}

// HasVar20480 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) HasVar20480() bool {
	if o != nil && !isNil(o.Var20480) {
		return true
	}

	return false
}

// SetVar20480 gets a reference to the given int32 and assigns it to the Var20480 field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) SetVar20480(v int32) {
	o.Var20480 = &v
}

func (o NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Var05) {
		toSerialize["0.5"] = o.Var05
	}
	if !isNil(o.Var10) {
		toSerialize["1.0"] = o.Var10
	}
	if !isNil(o.Var20) {
		toSerialize["2.0"] = o.Var20
	}
	if !isNil(o.Var40) {
		toSerialize["4.0"] = o.Var40
	}
	if !isNil(o.Var80) {
		toSerialize["8.0"] = o.Var80
	}
	if !isNil(o.Var160) {
		toSerialize["16.0"] = o.Var160
	}
	if !isNil(o.Var320) {
		toSerialize["32.0"] = o.Var320
	}
	if !isNil(o.Var640) {
		toSerialize["64.0"] = o.Var640
	}
	if !isNil(o.Var1280) {
		toSerialize["128.0"] = o.Var1280
	}
	if !isNil(o.Var2560) {
		toSerialize["256.0"] = o.Var2560
	}
	if !isNil(o.Var5120) {
		toSerialize["512.0"] = o.Var5120
	}
	if !isNil(o.Var10240) {
		toSerialize["1024.0"] = o.Var10240
	}
	if !isNil(o.Var20480) {
		toSerialize["2048.0"] = o.Var20480
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic struct {
	value *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) Get() *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) Set(val *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic(val *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic {
	return &NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


