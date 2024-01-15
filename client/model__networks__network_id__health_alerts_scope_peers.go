/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdHealthAlertsScopePeers struct for NetworksNetworkIdHealthAlertsScopePeers
type NetworksNetworkIdHealthAlertsScopePeers struct {
	// URL to the peer
	Url *string `json:"url,omitempty"`
	Network *NetworksNetworkIdHealthAlertsScopeNetwork `json:"network,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScopePeers instantiates a new NetworksNetworkIdHealthAlertsScopePeers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScopePeers() *NetworksNetworkIdHealthAlertsScopePeers {
	this := NetworksNetworkIdHealthAlertsScopePeers{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopePeersWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopePeers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopePeersWithDefaults() *NetworksNetworkIdHealthAlertsScopePeers {
	this := NetworksNetworkIdHealthAlertsScopePeers{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopePeers) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopePeers) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopePeers) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworksNetworkIdHealthAlertsScopePeers) SetUrl(v string) {
	o.Url = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopePeers) GetNetwork() NetworksNetworkIdHealthAlertsScopeNetwork {
	if o == nil || isNil(o.Network) {
		var ret NetworksNetworkIdHealthAlertsScopeNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopePeers) GetNetworkOk() (*NetworksNetworkIdHealthAlertsScopeNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopePeers) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NetworksNetworkIdHealthAlertsScopeNetwork and assigns it to the Network field.
func (o *NetworksNetworkIdHealthAlertsScopePeers) SetNetwork(v NetworksNetworkIdHealthAlertsScopeNetwork) {
	o.Network = &v
}

func (o NetworksNetworkIdHealthAlertsScopePeers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScopePeers struct {
	value *NetworksNetworkIdHealthAlertsScopePeers
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScopePeers) Get() *NetworksNetworkIdHealthAlertsScopePeers {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScopePeers) Set(val *NetworksNetworkIdHealthAlertsScopePeers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScopePeers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScopePeers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScopePeers(val *NetworksNetworkIdHealthAlertsScopePeers) *NullableNetworksNetworkIdHealthAlertsScopePeers {
	return &NullableNetworksNetworkIdHealthAlertsScopePeers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScopePeers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScopePeers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


