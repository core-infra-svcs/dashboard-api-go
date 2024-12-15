/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices struct for NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices
type NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices struct {
	// Serial of the device
	Serial *string `json:"serial,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevicesWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevicesWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices {
	this := NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) Get() *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) Set(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices(val *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevicesDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


