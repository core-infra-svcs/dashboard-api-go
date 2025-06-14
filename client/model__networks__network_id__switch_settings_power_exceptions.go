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

// NetworksNetworkIdSwitchSettingsPowerExceptions struct for NetworksNetworkIdSwitchSettingsPowerExceptions
type NetworksNetworkIdSwitchSettingsPowerExceptions struct {
	// Serial number of the switch
	Serial string `json:"serial"`
	// Per switch exception (combined, redundant, useNetworkSetting)
	PowerType string `json:"powerType"`
}

// NewNetworksNetworkIdSwitchSettingsPowerExceptions instantiates a new NetworksNetworkIdSwitchSettingsPowerExceptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchSettingsPowerExceptions(serial string, powerType string) *NetworksNetworkIdSwitchSettingsPowerExceptions {
	this := NetworksNetworkIdSwitchSettingsPowerExceptions{}
	this.Serial = serial
	this.PowerType = powerType
	return &this
}

// NewNetworksNetworkIdSwitchSettingsPowerExceptionsWithDefaults instantiates a new NetworksNetworkIdSwitchSettingsPowerExceptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchSettingsPowerExceptionsWithDefaults() *NetworksNetworkIdSwitchSettingsPowerExceptions {
	this := NetworksNetworkIdSwitchSettingsPowerExceptions{}
	return &this
}

// GetSerial returns the Serial field value
func (o *NetworksNetworkIdSwitchSettingsPowerExceptions) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchSettingsPowerExceptions) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *NetworksNetworkIdSwitchSettingsPowerExceptions) SetSerial(v string) {
	o.Serial = v
}

// GetPowerType returns the PowerType field value
func (o *NetworksNetworkIdSwitchSettingsPowerExceptions) GetPowerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerType
}

// GetPowerTypeOk returns a tuple with the PowerType field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchSettingsPowerExceptions) GetPowerTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PowerType, true
}

// SetPowerType sets field value
func (o *NetworksNetworkIdSwitchSettingsPowerExceptions) SetPowerType(v string) {
	o.PowerType = v
}

func (o NetworksNetworkIdSwitchSettingsPowerExceptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["powerType"] = o.PowerType
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchSettingsPowerExceptions struct {
	value *NetworksNetworkIdSwitchSettingsPowerExceptions
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchSettingsPowerExceptions) Get() *NetworksNetworkIdSwitchSettingsPowerExceptions {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchSettingsPowerExceptions) Set(val *NetworksNetworkIdSwitchSettingsPowerExceptions) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchSettingsPowerExceptions) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchSettingsPowerExceptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchSettingsPowerExceptions(val *NetworksNetworkIdSwitchSettingsPowerExceptions) *NullableNetworksNetworkIdSwitchSettingsPowerExceptions {
	return &NullableNetworksNetworkIdSwitchSettingsPowerExceptions{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchSettingsPowerExceptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchSettingsPowerExceptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


