/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 struct for NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1
type NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 struct {
	// Public IP address of the RADIUS accounting server
	Host string `json:"host"`
	// UDP port that the RADIUS Accounting server listens on for access requests
	Port int32 `json:"port"`
	// RADIUS client shared secret
	Secret string `json:"secret"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1(host string, port int32, secret string) *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1{}
	this.Host = host
	this.Port = port
	this.Secret = secret
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1WithDefaults() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	this := NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetPort(v int32) {
	o.Port = v
}

// GetSecret returns the Secret field value
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) GetSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) SetSecret(v string) {
	o.Secret = v
}

func (o NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 struct {
	value *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) Get() *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) Set(val *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1(val *NetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1 {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesRadiusAccountingServers1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


