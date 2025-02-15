/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchStpStpBridgePriority struct for NetworksNetworkIdSwitchStpStpBridgePriority
type NetworksNetworkIdSwitchStpStpBridgePriority struct {
	// List of switch template IDs
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// List of switch serial numbers
	Switches []string `json:"switches,omitempty"`
	// List of stack IDs
	Stacks []string `json:"stacks,omitempty"`
	// STP priority for switch, stacks, or switch templates
	StpPriority int32 `json:"stpPriority"`
}

// NewNetworksNetworkIdSwitchStpStpBridgePriority instantiates a new NetworksNetworkIdSwitchStpStpBridgePriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchStpStpBridgePriority(stpPriority int32) *NetworksNetworkIdSwitchStpStpBridgePriority {
	this := NetworksNetworkIdSwitchStpStpBridgePriority{}
	this.StpPriority = stpPriority
	return &this
}

// NewNetworksNetworkIdSwitchStpStpBridgePriorityWithDefaults instantiates a new NetworksNetworkIdSwitchStpStpBridgePriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchStpStpBridgePriorityWithDefaults() *NetworksNetworkIdSwitchStpStpBridgePriority {
	this := NetworksNetworkIdSwitchStpStpBridgePriority{}
	return &this
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetStacks() []string {
	if o == nil || isNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetStacksOk() ([]string, bool) {
	if o == nil || isNil(o.Stacks) {
    return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) HasStacks() bool {
	if o != nil && !isNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) SetStacks(v []string) {
	o.Stacks = v
}

// GetStpPriority returns the StpPriority field value
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetStpPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StpPriority
}

// GetStpPriorityOk returns a tuple with the StpPriority field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) GetStpPriorityOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StpPriority, true
}

// SetStpPriority sets field value
func (o *NetworksNetworkIdSwitchStpStpBridgePriority) SetStpPriority(v int32) {
	o.StpPriority = v
}

func (o NetworksNetworkIdSwitchStpStpBridgePriority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !isNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}
	if true {
		toSerialize["stpPriority"] = o.StpPriority
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchStpStpBridgePriority struct {
	value *NetworksNetworkIdSwitchStpStpBridgePriority
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchStpStpBridgePriority) Get() *NetworksNetworkIdSwitchStpStpBridgePriority {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchStpStpBridgePriority) Set(val *NetworksNetworkIdSwitchStpStpBridgePriority) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchStpStpBridgePriority) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchStpStpBridgePriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchStpStpBridgePriority(val *NetworksNetworkIdSwitchStpStpBridgePriority) *NullableNetworksNetworkIdSwitchStpStpBridgePriority {
	return &NullableNetworksNetworkIdSwitchStpStpBridgePriority{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchStpStpBridgePriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchStpStpBridgePriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


