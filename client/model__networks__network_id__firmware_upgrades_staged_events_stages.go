/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesStagedEventsStages struct for NetworksNetworkIdFirmwareUpgradesStagedEventsStages
type NetworksNetworkIdFirmwareUpgradesStagedEventsStages struct {
	Group *NetworksNetworkIdFirmwareUpgradesStagedEventsGroup `json:"group,omitempty"`
	Milestones *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones `json:"milestones,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsStages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages() *NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsStages{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsStagesWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsStages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsStagesWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsStages{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) GetGroup() NetworksNetworkIdFirmwareUpgradesStagedEventsGroup {
	if o == nil || isNil(o.Group) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) GetGroupOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsGroup and assigns it to the Group field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) SetGroup(v NetworksNetworkIdFirmwareUpgradesStagedEventsGroup) {
	o.Group = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) GetMilestones() NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones {
	if o == nil || isNil(o.Milestones) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) GetMilestonesOk() (*NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones, bool) {
	if o == nil || isNil(o.Milestones) {
    return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) HasMilestones() bool {
	if o != nil && !isNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones and assigns it to the Milestones field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) SetMilestones(v NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) {
	o.Milestones = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsStages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Milestones) {
		toSerialize["milestones"] = o.Milestones
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsStages
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages(val *NetworksNetworkIdFirmwareUpgradesStagedEventsStages) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsStages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


