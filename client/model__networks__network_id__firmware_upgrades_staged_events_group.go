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

// NetworksNetworkIdFirmwareUpgradesStagedEventsGroup The Staged Upgrade Group containing the name and ID
type NetworksNetworkIdFirmwareUpgradesStagedEventsGroup struct {
	// ID of the Staged Upgrade Group
	Id string `json:"id"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsGroup instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsGroup(id string) *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsGroup{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsGroupWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsGroupWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) SetId(v string) {
	o.Id = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup(val *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


