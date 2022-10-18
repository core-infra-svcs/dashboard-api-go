/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberGreConcentrator The EoGRE concentrator's settings
type NetworksNetworkIdWirelessSsidsNumberGreConcentrator struct {
	// The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is 'Ethernet over GRE'.
	Host string `json:"host"`
}

// NewNetworksNetworkIdWirelessSsidsNumberGreConcentrator instantiates a new NetworksNetworkIdWirelessSsidsNumberGreConcentrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberGreConcentrator(host string) *NetworksNetworkIdWirelessSsidsNumberGreConcentrator {
	this := NetworksNetworkIdWirelessSsidsNumberGreConcentrator{}
	this.Host = host
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberGreConcentratorWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberGreConcentrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberGreConcentratorWithDefaults() *NetworksNetworkIdWirelessSsidsNumberGreConcentrator {
	this := NetworksNetworkIdWirelessSsidsNumberGreConcentrator{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworksNetworkIdWirelessSsidsNumberGreConcentrator) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberGreConcentrator) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberGreConcentrator) SetHost(v string) {
	o.Host = v
}

func (o NetworksNetworkIdWirelessSsidsNumberGreConcentrator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator struct {
	value *NetworksNetworkIdWirelessSsidsNumberGreConcentrator
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator) Get() *NetworksNetworkIdWirelessSsidsNumberGreConcentrator {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator) Set(val *NetworksNetworkIdWirelessSsidsNumberGreConcentrator) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator(val *NetworksNetworkIdWirelessSsidsNumberGreConcentrator) *NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator {
	return &NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberGreConcentrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


