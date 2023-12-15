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

// NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones The Staged Upgrade Milestones for the specific stage
type NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones struct {
	// The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)
	ScheduledFor string `json:"scheduledFor"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones(scheduledFor string) *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones{}
	this.ScheduledFor = scheduledFor
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedEventsMilestonesWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedEventsMilestonesWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones {
	this := NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) GetScheduledFor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) GetScheduledForOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScheduledFor, true
}

// SetScheduledFor sets field value
func (o *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) SetScheduledFor(v string) {
	o.ScheduledFor = v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scheduledFor"] = o.ScheduledFor
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) Get() *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) Set(val *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones(val *NetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedEventsMilestones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


