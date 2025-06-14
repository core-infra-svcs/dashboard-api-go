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

// NetworksNetworkIdCampusGatewayClustersAddresses1 struct for NetworksNetworkIdCampusGatewayClustersAddresses1
type NetworksNetworkIdCampusGatewayClustersAddresses1 struct {
	// Protocol of the interface
	Protocol string `json:"protocol"`
	// Gateway of the interface
	Gateway string `json:"gateway"`
	// Subnet mask of the interface
	SubnetMask string `json:"subnetMask"`
}

// NewNetworksNetworkIdCampusGatewayClustersAddresses1 instantiates a new NetworksNetworkIdCampusGatewayClustersAddresses1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCampusGatewayClustersAddresses1(protocol string, gateway string, subnetMask string) *NetworksNetworkIdCampusGatewayClustersAddresses1 {
	this := NetworksNetworkIdCampusGatewayClustersAddresses1{}
	this.Protocol = protocol
	this.Gateway = gateway
	this.SubnetMask = subnetMask
	return &this
}

// NewNetworksNetworkIdCampusGatewayClustersAddresses1WithDefaults instantiates a new NetworksNetworkIdCampusGatewayClustersAddresses1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCampusGatewayClustersAddresses1WithDefaults() *NetworksNetworkIdCampusGatewayClustersAddresses1 {
	this := NetworksNetworkIdCampusGatewayClustersAddresses1{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) SetProtocol(v string) {
	o.Protocol = v
}

// GetGateway returns the Gateway field value
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) GetGatewayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) SetGateway(v string) {
	o.Gateway = v
}

// GetSubnetMask returns the SubnetMask field value
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) GetSubnetMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) GetSubnetMaskOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubnetMask, true
}

// SetSubnetMask sets field value
func (o *NetworksNetworkIdCampusGatewayClustersAddresses1) SetSubnetMask(v string) {
	o.SubnetMask = v
}

func (o NetworksNetworkIdCampusGatewayClustersAddresses1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["gateway"] = o.Gateway
	}
	if true {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCampusGatewayClustersAddresses1 struct {
	value *NetworksNetworkIdCampusGatewayClustersAddresses1
	isSet bool
}

func (v NullableNetworksNetworkIdCampusGatewayClustersAddresses1) Get() *NetworksNetworkIdCampusGatewayClustersAddresses1 {
	return v.value
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersAddresses1) Set(val *NetworksNetworkIdCampusGatewayClustersAddresses1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCampusGatewayClustersAddresses1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersAddresses1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCampusGatewayClustersAddresses1(val *NetworksNetworkIdCampusGatewayClustersAddresses1) *NullableNetworksNetworkIdCampusGatewayClustersAddresses1 {
	return &NullableNetworksNetworkIdCampusGatewayClustersAddresses1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCampusGatewayClustersAddresses1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCampusGatewayClustersAddresses1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


