/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds struct for NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds
type NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds struct {
	// Seconds since Sunday at midnight when the outage range starts.
	Start int32 `json:"start"`
	// Seconds since Sunday at midnight when that outage range ends.
	End int32 `json:"end"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds instantiates a new NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds(start int32, end int32) *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds {
	this := NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds{}
	this.Start = start
	this.End = end
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSecondsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSecondsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds {
	this := NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds{}
	return &this
}

// GetStart returns the Start field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) GetStartOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) GetEndOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) SetEnd(v int32) {
	o.End = v
}

func (o NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds struct {
	value *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) Get() *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) Set(val *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds(val *NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


