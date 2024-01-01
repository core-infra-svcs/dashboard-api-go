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

// NetworksNetworkIdWirelessSsidsNumberGre Ethernet over GRE settings
type NetworksNetworkIdWirelessSsidsNumberGre struct {
	Concentrator *NetworksNetworkIdWirelessSsidsNumberGreConcentrator `json:"concentrator,omitempty"`
	// Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.
	Key *int32 `json:"key,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberGre instantiates a new NetworksNetworkIdWirelessSsidsNumberGre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberGre() *NetworksNetworkIdWirelessSsidsNumberGre {
	this := NetworksNetworkIdWirelessSsidsNumberGre{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberGreWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberGre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberGreWithDefaults() *NetworksNetworkIdWirelessSsidsNumberGre {
	this := NetworksNetworkIdWirelessSsidsNumberGre{}
	return &this
}

// GetConcentrator returns the Concentrator field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) GetConcentrator() NetworksNetworkIdWirelessSsidsNumberGreConcentrator {
	if o == nil || isNil(o.Concentrator) {
		var ret NetworksNetworkIdWirelessSsidsNumberGreConcentrator
		return ret
	}
	return *o.Concentrator
}

// GetConcentratorOk returns a tuple with the Concentrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) GetConcentratorOk() (*NetworksNetworkIdWirelessSsidsNumberGreConcentrator, bool) {
	if o == nil || isNil(o.Concentrator) {
    return nil, false
	}
	return o.Concentrator, true
}

// HasConcentrator returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) HasConcentrator() bool {
	if o != nil && !isNil(o.Concentrator) {
		return true
	}

	return false
}

// SetConcentrator gets a reference to the given NetworksNetworkIdWirelessSsidsNumberGreConcentrator and assigns it to the Concentrator field.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) SetConcentrator(v NetworksNetworkIdWirelessSsidsNumberGreConcentrator) {
	o.Concentrator = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) GetKey() int32 {
	if o == nil || isNil(o.Key) {
		var ret int32
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) GetKeyOk() (*int32, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given int32 and assigns it to the Key field.
func (o *NetworksNetworkIdWirelessSsidsNumberGre) SetKey(v int32) {
	o.Key = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberGre) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentrator) {
		toSerialize["concentrator"] = o.Concentrator
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberGre struct {
	value *NetworksNetworkIdWirelessSsidsNumberGre
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberGre) Get() *NetworksNetworkIdWirelessSsidsNumberGre {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberGre) Set(val *NetworksNetworkIdWirelessSsidsNumberGre) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberGre) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberGre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberGre(val *NetworksNetworkIdWirelessSsidsNumberGre) *NullableNetworksNetworkIdWirelessSsidsNumberGre {
	return &NullableNetworksNetworkIdWirelessSsidsNumberGre{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberGre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberGre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


