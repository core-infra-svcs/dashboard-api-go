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

// NetworksNetworkIdCampusGatewayClustersUplinks1 struct for NetworksNetworkIdCampusGatewayClustersUplinks1
type NetworksNetworkIdCampusGatewayClustersUplinks1 struct {
	// Uplink interface name
	Interface string `json:"interface"`
	// Uplink IP addresses of the device
	Addresses []NetworksNetworkIdCampusGatewayClustersAddresses2 `json:"addresses"`
}

// NewNetworksNetworkIdCampusGatewayClustersUplinks1 instantiates a new NetworksNetworkIdCampusGatewayClustersUplinks1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCampusGatewayClustersUplinks1(interface_ string, addresses []NetworksNetworkIdCampusGatewayClustersAddresses2) *NetworksNetworkIdCampusGatewayClustersUplinks1 {
	this := NetworksNetworkIdCampusGatewayClustersUplinks1{}
	this.Interface = interface_
	this.Addresses = addresses
	return &this
}

// NewNetworksNetworkIdCampusGatewayClustersUplinks1WithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersUplinks1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCampusGatewayClustersUplinks1WithDefaults() *NetworksNetworkIdCampusGatewayClustersUplinks1 {
	this := NetworksNetworkIdCampusGatewayClustersUplinks1{}
	return &this
}

// GetInterface returns the Interface field value
func (o *NetworksNetworkIdCampusGatewayClustersUplinks1) GetInterface() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCampusGatewayClustersUplinks1) GetInterfaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *NetworksNetworkIdCampusGatewayClustersUplinks1) SetInterface(v string) {
	o.Interface = v
}

// GetAddresses returns the Addresses field value
func (o *NetworksNetworkIdCampusGatewayClustersUplinks1) GetAddresses() []NetworksNetworkIdCampusGatewayClustersAddresses2 {
	if o == nil {
		var ret []NetworksNetworkIdCampusGatewayClustersAddresses2
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCampusGatewayClustersUplinks1) GetAddressesOk() ([]NetworksNetworkIdCampusGatewayClustersAddresses2, bool) {
	if o == nil {
    return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *NetworksNetworkIdCampusGatewayClustersUplinks1) SetAddresses(v []NetworksNetworkIdCampusGatewayClustersAddresses2) {
	o.Addresses = v
}

func (o NetworksNetworkIdCampusGatewayClustersUplinks1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interface"] = o.Interface
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCampusGatewayClustersUplinks1 struct {
	value *NetworksNetworkIdCampusGatewayClustersUplinks1
	isSet bool
}

func (v NullableNetworksNetworkIdCampusGatewayClustersUplinks1) Get() *NetworksNetworkIdCampusGatewayClustersUplinks1 {
	return v.value
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersUplinks1) Set(val *NetworksNetworkIdCampusGatewayClustersUplinks1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCampusGatewayClustersUplinks1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersUplinks1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCampusGatewayClustersUplinks1(val *NetworksNetworkIdCampusGatewayClustersUplinks1) *NullableNetworksNetworkIdCampusGatewayClustersUplinks1 {
	return &NullableNetworksNetworkIdCampusGatewayClustersUplinks1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCampusGatewayClustersUplinks1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersUplinks1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


