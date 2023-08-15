/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVlansDhcpOptions struct for NetworksNetworkIdApplianceVlansDhcpOptions
type NetworksNetworkIdApplianceVlansDhcpOptions struct {
	// The code for the DHCP option. This should be an integer between 2 and 254.
	Code string `json:"code"`
	// The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	Type string `json:"type"`
	// The value for the DHCP option
	Value string `json:"value"`
}

// NewNetworksNetworkIdApplianceVlansDhcpOptions instantiates a new NetworksNetworkIdApplianceVlansDhcpOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVlansDhcpOptions(code string, type_ string, value string) *NetworksNetworkIdApplianceVlansDhcpOptions {
	this := NetworksNetworkIdApplianceVlansDhcpOptions{}
	this.Code = code
	this.Type = type_
	this.Value = value
	return &this
}

// NewNetworksNetworkIdApplianceVlansDhcpOptionsWithDefaults instantiates a new NetworksNetworkIdApplianceVlansDhcpOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVlansDhcpOptionsWithDefaults() *NetworksNetworkIdApplianceVlansDhcpOptions {
	this := NetworksNetworkIdApplianceVlansDhcpOptions{}
	return &this
}

// GetCode returns the Code field value
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) GetCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NetworksNetworkIdApplianceVlansDhcpOptions) SetValue(v string) {
	o.Value = v
}

func (o NetworksNetworkIdApplianceVlansDhcpOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVlansDhcpOptions struct {
	value *NetworksNetworkIdApplianceVlansDhcpOptions
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVlansDhcpOptions) Get() *NetworksNetworkIdApplianceVlansDhcpOptions {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVlansDhcpOptions) Set(val *NetworksNetworkIdApplianceVlansDhcpOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVlansDhcpOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVlansDhcpOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVlansDhcpOptions(val *NetworksNetworkIdApplianceVlansDhcpOptions) *NullableNetworksNetworkIdApplianceVlansDhcpOptions {
	return &NullableNetworksNetworkIdApplianceVlansDhcpOptions{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVlansDhcpOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVlansDhcpOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


