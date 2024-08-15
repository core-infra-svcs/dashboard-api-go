/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments struct for DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments
type DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments struct {
	// The name of the client which has fixed IP address
	Name string `json:"name"`
	// The MAC address of the client which has fixed IP address
	Mac string `json:"mac"`
	// The IP address of the client which has fixed IP address assigned to it
	Ip string `json:"ip"`
}

// NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments instantiates a new DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments(name string, mac string, ip string) *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments {
	this := DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments{}
	this.Name = name
	this.Mac = mac
	this.Ip = ip
	return &this
}

// NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignmentsWithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignmentsWithDefaults() *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments {
	this := DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments{}
	return &this
}

// GetName returns the Name field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) SetName(v string) {
	o.Name = v
}

// GetMac returns the Mac field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) SetMac(v string) {
	o.Mac = v
}

// GetIp returns the Ip field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) GetIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) SetIp(v string) {
	o.Ip = v
}

func (o DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments struct {
	value *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments
	isSet bool
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) Get() *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments {
	return v.value
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) Set(val *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments(val *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments {
	return &NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpFixedIpAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


