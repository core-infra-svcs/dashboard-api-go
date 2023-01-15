/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchRoutingMulticastOverrides struct for NetworksNetworkIdSwitchRoutingMulticastOverrides
type NetworksNetworkIdSwitchRoutingMulticastOverrides struct {
	// List of switch profiles ids for template network
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// List of switch serials for non-template network
	Switches []string `json:"switches,omitempty"`
	// List of switch stack ids for non-template network
	Stacks []string `json:"stacks,omitempty"`
	// IGMP snooping setting for switches, switch stacks or switch profiles
	IgmpSnoopingEnabled bool `json:"igmpSnoopingEnabled"`
	// Flood unknown multicast traffic setting for switches, switch stacks or switch profiles
	FloodUnknownMulticastTrafficEnabled bool `json:"floodUnknownMulticastTrafficEnabled"`
}

// NewNetworksNetworkIdSwitchRoutingMulticastOverrides instantiates a new NetworksNetworkIdSwitchRoutingMulticastOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchRoutingMulticastOverrides(igmpSnoopingEnabled bool, floodUnknownMulticastTrafficEnabled bool) *NetworksNetworkIdSwitchRoutingMulticastOverrides {
	this := NetworksNetworkIdSwitchRoutingMulticastOverrides{}
	this.IgmpSnoopingEnabled = igmpSnoopingEnabled
	this.FloodUnknownMulticastTrafficEnabled = floodUnknownMulticastTrafficEnabled
	return &this
}

// NewNetworksNetworkIdSwitchRoutingMulticastOverridesWithDefaults instantiates a new NetworksNetworkIdSwitchRoutingMulticastOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchRoutingMulticastOverridesWithDefaults() *NetworksNetworkIdSwitchRoutingMulticastOverrides {
	this := NetworksNetworkIdSwitchRoutingMulticastOverrides{}
	return &this
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetStacks() []string {
	if o == nil || isNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetStacksOk() ([]string, bool) {
	if o == nil || isNil(o.Stacks) {
    return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) HasStacks() bool {
	if o != nil && !isNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetStacks(v []string) {
	o.Stacks = v
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetIgmpSnoopingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IgmpSnoopingEnabled, true
}

// SetIgmpSnoopingEnabled sets field value
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FloodUnknownMulticastTrafficEnabled, true
}

// SetFloodUnknownMulticastTrafficEnabled sets field value
func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = v
}

func (o NetworksNetworkIdSwitchRoutingMulticastOverrides) MarshalJSON() ([]byte, error) {
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
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if true {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchRoutingMulticastOverrides struct {
	value *NetworksNetworkIdSwitchRoutingMulticastOverrides
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchRoutingMulticastOverrides) Get() *NetworksNetworkIdSwitchRoutingMulticastOverrides {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchRoutingMulticastOverrides) Set(val *NetworksNetworkIdSwitchRoutingMulticastOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchRoutingMulticastOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchRoutingMulticastOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchRoutingMulticastOverrides(val *NetworksNetworkIdSwitchRoutingMulticastOverrides) *NullableNetworksNetworkIdSwitchRoutingMulticastOverrides {
	return &NullableNetworksNetworkIdSwitchRoutingMulticastOverrides{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchRoutingMulticastOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchRoutingMulticastOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


