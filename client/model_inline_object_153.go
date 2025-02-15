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

// InlineObject153 struct for InlineObject153
type InlineObject153 struct {
	// Management VLAN
	Vlan *int32 `json:"vlan,omitempty"`
	// The use Combined Power as the default behavior of secondary power supplies on supported devices.
	UseCombinedPower *bool `json:"useCombinedPower,omitempty"`
	// Exceptions on a per switch basis to \"useCombinedPower\"
	PowerExceptions []NetworksNetworkIdSwitchSettingsPowerExceptions `json:"powerExceptions,omitempty"`
	UplinkClientSampling *NetworksNetworkIdSwitchSettingsUplinkClientSampling `json:"uplinkClientSampling,omitempty"`
	MacBlocklist *NetworksNetworkIdSwitchSettingsMacBlocklist `json:"macBlocklist,omitempty"`
}

// NewInlineObject153 instantiates a new InlineObject153 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject153() *InlineObject153 {
	this := InlineObject153{}
	return &this
}

// NewInlineObject153WithDefaults instantiates a new InlineObject153 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject153WithDefaults() *InlineObject153 {
	this := InlineObject153{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject153) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject153) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject153) SetVlan(v int32) {
	o.Vlan = &v
}

// GetUseCombinedPower returns the UseCombinedPower field value if set, zero value otherwise.
func (o *InlineObject153) GetUseCombinedPower() bool {
	if o == nil || isNil(o.UseCombinedPower) {
		var ret bool
		return ret
	}
	return *o.UseCombinedPower
}

// GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetUseCombinedPowerOk() (*bool, bool) {
	if o == nil || isNil(o.UseCombinedPower) {
    return nil, false
	}
	return o.UseCombinedPower, true
}

// HasUseCombinedPower returns a boolean if a field has been set.
func (o *InlineObject153) HasUseCombinedPower() bool {
	if o != nil && !isNil(o.UseCombinedPower) {
		return true
	}

	return false
}

// SetUseCombinedPower gets a reference to the given bool and assigns it to the UseCombinedPower field.
func (o *InlineObject153) SetUseCombinedPower(v bool) {
	o.UseCombinedPower = &v
}

// GetPowerExceptions returns the PowerExceptions field value if set, zero value otherwise.
func (o *InlineObject153) GetPowerExceptions() []NetworksNetworkIdSwitchSettingsPowerExceptions {
	if o == nil || isNil(o.PowerExceptions) {
		var ret []NetworksNetworkIdSwitchSettingsPowerExceptions
		return ret
	}
	return o.PowerExceptions
}

// GetPowerExceptionsOk returns a tuple with the PowerExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetPowerExceptionsOk() ([]NetworksNetworkIdSwitchSettingsPowerExceptions, bool) {
	if o == nil || isNil(o.PowerExceptions) {
    return nil, false
	}
	return o.PowerExceptions, true
}

// HasPowerExceptions returns a boolean if a field has been set.
func (o *InlineObject153) HasPowerExceptions() bool {
	if o != nil && !isNil(o.PowerExceptions) {
		return true
	}

	return false
}

// SetPowerExceptions gets a reference to the given []NetworksNetworkIdSwitchSettingsPowerExceptions and assigns it to the PowerExceptions field.
func (o *InlineObject153) SetPowerExceptions(v []NetworksNetworkIdSwitchSettingsPowerExceptions) {
	o.PowerExceptions = v
}

// GetUplinkClientSampling returns the UplinkClientSampling field value if set, zero value otherwise.
func (o *InlineObject153) GetUplinkClientSampling() NetworksNetworkIdSwitchSettingsUplinkClientSampling {
	if o == nil || isNil(o.UplinkClientSampling) {
		var ret NetworksNetworkIdSwitchSettingsUplinkClientSampling
		return ret
	}
	return *o.UplinkClientSampling
}

// GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetUplinkClientSamplingOk() (*NetworksNetworkIdSwitchSettingsUplinkClientSampling, bool) {
	if o == nil || isNil(o.UplinkClientSampling) {
    return nil, false
	}
	return o.UplinkClientSampling, true
}

// HasUplinkClientSampling returns a boolean if a field has been set.
func (o *InlineObject153) HasUplinkClientSampling() bool {
	if o != nil && !isNil(o.UplinkClientSampling) {
		return true
	}

	return false
}

// SetUplinkClientSampling gets a reference to the given NetworksNetworkIdSwitchSettingsUplinkClientSampling and assigns it to the UplinkClientSampling field.
func (o *InlineObject153) SetUplinkClientSampling(v NetworksNetworkIdSwitchSettingsUplinkClientSampling) {
	o.UplinkClientSampling = &v
}

// GetMacBlocklist returns the MacBlocklist field value if set, zero value otherwise.
func (o *InlineObject153) GetMacBlocklist() NetworksNetworkIdSwitchSettingsMacBlocklist {
	if o == nil || isNil(o.MacBlocklist) {
		var ret NetworksNetworkIdSwitchSettingsMacBlocklist
		return ret
	}
	return *o.MacBlocklist
}

// GetMacBlocklistOk returns a tuple with the MacBlocklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject153) GetMacBlocklistOk() (*NetworksNetworkIdSwitchSettingsMacBlocklist, bool) {
	if o == nil || isNil(o.MacBlocklist) {
    return nil, false
	}
	return o.MacBlocklist, true
}

// HasMacBlocklist returns a boolean if a field has been set.
func (o *InlineObject153) HasMacBlocklist() bool {
	if o != nil && !isNil(o.MacBlocklist) {
		return true
	}

	return false
}

// SetMacBlocklist gets a reference to the given NetworksNetworkIdSwitchSettingsMacBlocklist and assigns it to the MacBlocklist field.
func (o *InlineObject153) SetMacBlocklist(v NetworksNetworkIdSwitchSettingsMacBlocklist) {
	o.MacBlocklist = &v
}

func (o InlineObject153) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.UseCombinedPower) {
		toSerialize["useCombinedPower"] = o.UseCombinedPower
	}
	if !isNil(o.PowerExceptions) {
		toSerialize["powerExceptions"] = o.PowerExceptions
	}
	if !isNil(o.UplinkClientSampling) {
		toSerialize["uplinkClientSampling"] = o.UplinkClientSampling
	}
	if !isNil(o.MacBlocklist) {
		toSerialize["macBlocklist"] = o.MacBlocklist
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject153 struct {
	value *InlineObject153
	isSet bool
}

func (v NullableInlineObject153) Get() *InlineObject153 {
	return v.value
}

func (v *NullableInlineObject153) Set(val *InlineObject153) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject153) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject153) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject153(val *InlineObject153) *NullableInlineObject153 {
	return &NullableInlineObject153{value: val, isSet: true}
}

func (v NullableInlineObject153) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject153) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


