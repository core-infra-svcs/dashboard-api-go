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

// NetworksNetworkIdSwitchMtuOverrides struct for NetworksNetworkIdSwitchMtuOverrides
type NetworksNetworkIdSwitchMtuOverrides struct {
	// List of switch serials. Applicable only for switch network.
	Switches *[]string `json:"switches,omitempty"`
	// List of switch profile IDs. Applicable only for template network.
	SwitchProfiles *[]string `json:"switchProfiles,omitempty"`
	// MTU size for the switches or switch profiles.
	MtuSize int32 `json:"mtuSize"`
}

// NewNetworksNetworkIdSwitchMtuOverrides instantiates a new NetworksNetworkIdSwitchMtuOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchMtuOverrides(mtuSize int32) *NetworksNetworkIdSwitchMtuOverrides {
	this := NetworksNetworkIdSwitchMtuOverrides{}
	this.MtuSize = mtuSize
	return &this
}

// NewNetworksNetworkIdSwitchMtuOverridesWithDefaults instantiates a new NetworksNetworkIdSwitchMtuOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchMtuOverridesWithDefaults() *NetworksNetworkIdSwitchMtuOverrides {
	this := NetworksNetworkIdSwitchMtuOverrides{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitches() []string {
	if o == nil || o.Switches == nil {
		var ret []string
		return ret
	}
	return *o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitchesOk() (*[]string, bool) {
	if o == nil || o.Switches == nil {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchMtuOverrides) HasSwitches() bool {
	if o != nil && o.Switches != nil {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *NetworksNetworkIdSwitchMtuOverrides) SetSwitches(v []string) {
	o.Switches = &v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitchProfiles() []string {
	if o == nil || o.SwitchProfiles == nil {
		var ret []string
		return ret
	}
	return *o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchMtuOverrides) GetSwitchProfilesOk() (*[]string, bool) {
	if o == nil || o.SwitchProfiles == nil {
		return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchMtuOverrides) HasSwitchProfiles() bool {
	if o != nil && o.SwitchProfiles != nil {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *NetworksNetworkIdSwitchMtuOverrides) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = &v
}

// GetMtuSize returns the MtuSize field value
func (o *NetworksNetworkIdSwitchMtuOverrides) GetMtuSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MtuSize
}

// GetMtuSizeOk returns a tuple with the MtuSize field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchMtuOverrides) GetMtuSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MtuSize, true
}

// SetMtuSize sets field value
func (o *NetworksNetworkIdSwitchMtuOverrides) SetMtuSize(v int32) {
	o.MtuSize = v
}

func (o NetworksNetworkIdSwitchMtuOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Switches != nil {
		toSerialize["switches"] = o.Switches
	}
	if o.SwitchProfiles != nil {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if true {
		toSerialize["mtuSize"] = o.MtuSize
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchMtuOverrides struct {
	value *NetworksNetworkIdSwitchMtuOverrides
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchMtuOverrides) Get() *NetworksNetworkIdSwitchMtuOverrides {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchMtuOverrides) Set(val *NetworksNetworkIdSwitchMtuOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchMtuOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchMtuOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchMtuOverrides(val *NetworksNetworkIdSwitchMtuOverrides) *NullableNetworksNetworkIdSwitchMtuOverrides {
	return &NullableNetworksNetworkIdSwitchMtuOverrides{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchMtuOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchMtuOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


