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

// NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches struct for NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches
type NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches struct {
	// Switch serial number
	Serial string `json:"serial"`
	// Switch alternative management IP. To remove a prior IP setting, provide an empty string
	AlternateManagementIp string `json:"alternateManagementIp"`
	// Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Gateway *string `json:"gateway,omitempty"`
}

// NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches instantiates a new NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches(serial string, alternateManagementIp string) *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches {
	this := NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches{}
	this.Serial = serial
	this.AlternateManagementIp = alternateManagementIp
	return &this
}

// NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitchesWithDefaults instantiates a new NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAlternateManagementInterfaceSwitchesWithDefaults() *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches {
	this := NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches{}
	return &this
}

// GetSerial returns the Serial field value
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSerialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetSerial(v string) {
	o.Serial = v
}

// GetAlternateManagementIp returns the AlternateManagementIp field value
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetAlternateManagementIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateManagementIp
}

// GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetAlternateManagementIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AlternateManagementIp, true
}

// SetAlternateManagementIp sets field value
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetAlternateManagementIp(v string) {
	o.AlternateManagementIp = v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSubnetMask() string {
	if o == nil || o.SubnetMask == nil {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetSubnetMaskOk() (*string, bool) {
	if o == nil || o.SubnetMask == nil {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) HasSubnetMask() bool {
	if o != nil && o.SubnetMask != nil {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) SetGateway(v string) {
	o.Gateway = &v
}

func (o NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["alternateManagementIp"] = o.AlternateManagementIp
	}
	if o.SubnetMask != nil {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches struct {
	value *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) Get() *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) Set(val *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches(val *NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) *NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches {
	return &NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


