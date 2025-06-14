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

// NetworksNetworkIdCampusGatewayClustersUplink Uplink interface details if uplink is reused as tunnel
type NetworksNetworkIdCampusGatewayClustersUplink struct {
	// Uplink interface identifier
	Interface string `json:"interface"`
}

// NewNetworksNetworkIdCampusGatewayClustersUplink instantiates a new NetworksNetworkIdCampusGatewayClustersUplink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCampusGatewayClustersUplink(interface_ string) *NetworksNetworkIdCampusGatewayClustersUplink {
	this := NetworksNetworkIdCampusGatewayClustersUplink{}
	this.Interface = interface_
	return &this
}

// NewNetworksNetworkIdCampusGatewayClustersUplinkWithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersUplink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCampusGatewayClustersUplinkWithDefaults() *NetworksNetworkIdCampusGatewayClustersUplink {
	this := NetworksNetworkIdCampusGatewayClustersUplink{}
	return &this
}

// GetInterface returns the Interface field value
func (o *NetworksNetworkIdCampusGatewayClustersUplink) GetInterface() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCampusGatewayClustersUplink) GetInterfaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *NetworksNetworkIdCampusGatewayClustersUplink) SetInterface(v string) {
	o.Interface = v
}

func (o NetworksNetworkIdCampusGatewayClustersUplink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interface"] = o.Interface
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCampusGatewayClustersUplink struct {
	value *NetworksNetworkIdCampusGatewayClustersUplink
	isSet bool
}

func (v NullableNetworksNetworkIdCampusGatewayClustersUplink) Get() *NetworksNetworkIdCampusGatewayClustersUplink {
	return v.value
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersUplink) Set(val *NetworksNetworkIdCampusGatewayClustersUplink) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCampusGatewayClustersUplink) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersUplink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCampusGatewayClustersUplink(val *NetworksNetworkIdCampusGatewayClustersUplink) *NullableNetworksNetworkIdCampusGatewayClustersUplink {
	return &NullableNetworksNetworkIdCampusGatewayClustersUplink{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCampusGatewayClustersUplink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersUplink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


