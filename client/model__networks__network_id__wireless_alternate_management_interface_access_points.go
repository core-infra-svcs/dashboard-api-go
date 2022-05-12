/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints struct for NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints
type NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints struct {
	// Serial number of access point to be configured with alternate management IP
	Serial string `json:"serial"`
	// Wireless alternate management interface device IP. Provide an empty string to remove alternate management IP assignment
	AlternateManagementIp string `json:"alternateManagementIp"`
	// Subnet mask must be in IP format
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Gateway must be in IP format
	Gateway *string `json:"gateway,omitempty"`
	// Primary DNS must be in IP format
	Dns1 *string `json:"dns1,omitempty"`
	// Optional secondary DNS must be in IP format
	Dns2 *string `json:"dns2,omitempty"`
}

// NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints instantiates a new NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints(serial string, alternateManagementIp string) *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints {
	this := NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints{}
	this.Serial = serial
	this.AlternateManagementIp = alternateManagementIp
	return &this
}

// NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPointsWithDefaults instantiates a new NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPointsWithDefaults() *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints {
	this := NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints{}
	return &this
}

// GetSerial returns the Serial field value
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSerialOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetSerial(v string) {
	o.Serial = v
}

// GetAlternateManagementIp returns the AlternateManagementIp field value
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetAlternateManagementIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateManagementIp
}

// GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetAlternateManagementIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AlternateManagementIp, true
}

// SetAlternateManagementIp sets field value
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetAlternateManagementIp(v string) {
	o.AlternateManagementIp = v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSubnetMask() string {
	if o == nil || o.SubnetMask == nil {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetSubnetMaskOk() (*string, bool) {
	if o == nil || o.SubnetMask == nil {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasSubnetMask() bool {
	if o != nil && o.SubnetMask != nil {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetGateway(v string) {
	o.Gateway = &v
}

// GetDns1 returns the Dns1 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns1() string {
	if o == nil || o.Dns1 == nil {
		var ret string
		return ret
	}
	return *o.Dns1
}

// GetDns1Ok returns a tuple with the Dns1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns1Ok() (*string, bool) {
	if o == nil || o.Dns1 == nil {
		return nil, false
	}
	return o.Dns1, true
}

// HasDns1 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasDns1() bool {
	if o != nil && o.Dns1 != nil {
		return true
	}

	return false
}

// SetDns1 gets a reference to the given string and assigns it to the Dns1 field.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetDns1(v string) {
	o.Dns1 = &v
}

// GetDns2 returns the Dns2 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns2() string {
	if o == nil || o.Dns2 == nil {
		var ret string
		return ret
	}
	return *o.Dns2
}

// GetDns2Ok returns a tuple with the Dns2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) GetDns2Ok() (*string, bool) {
	if o == nil || o.Dns2 == nil {
		return nil, false
	}
	return o.Dns2, true
}

// HasDns2 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) HasDns2() bool {
	if o != nil && o.Dns2 != nil {
		return true
	}

	return false
}

// SetDns2 gets a reference to the given string and assigns it to the Dns2 field.
func (o *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) SetDns2(v string) {
	o.Dns2 = &v
}

func (o NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) MarshalJSON() ([]byte, error) {
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
	if o.Dns1 != nil {
		toSerialize["dns1"] = o.Dns1
	}
	if o.Dns2 != nil {
		toSerialize["dns2"] = o.Dns2
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints struct {
	value *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) Get() *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) Set(val *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints(val *NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) *NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints {
	return &NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

