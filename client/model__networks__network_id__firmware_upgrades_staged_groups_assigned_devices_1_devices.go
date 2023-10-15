/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices struct for NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices
type NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices struct {
	// Serial of the device
	Serial string `json:"serial"`
	// Name of the device
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices(serial string) *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices{}
	this.Serial = serial
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1DevicesWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1DevicesWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices{}
	return &this
}

// GetSerial returns the Serial field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) SetSerial(v string) {
	o.Serial = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) Get() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) Set(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1Devices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


