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

// InlineObject109 struct for InlineObject109
type InlineObject109 struct {
	// Management VLAN
	Vlan *int32 `json:"vlan,omitempty"`
	// The use Combined Power as the default behavior of secondary power supplies on supported devices.
	UseCombinedPower *bool `json:"useCombinedPower,omitempty"`
	// Exceptions on a per switch basis to \"useCombinedPower\"
	PowerExceptions []NetworksNetworkIdSwitchSettingsPowerExceptions `json:"powerExceptions,omitempty"`
}

// NewInlineObject109 instantiates a new InlineObject109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject109() *InlineObject109 {
	this := InlineObject109{}
	return &this
}

// NewInlineObject109WithDefaults instantiates a new InlineObject109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject109WithDefaults() *InlineObject109 {
	this := InlineObject109{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject109) GetVlan() int32 {
	if o == nil || o.Vlan == nil {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject109) GetVlanOk() (*int32, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject109) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject109) SetVlan(v int32) {
	o.Vlan = &v
}

// GetUseCombinedPower returns the UseCombinedPower field value if set, zero value otherwise.
func (o *InlineObject109) GetUseCombinedPower() bool {
	if o == nil || o.UseCombinedPower == nil {
		var ret bool
		return ret
	}
	return *o.UseCombinedPower
}

// GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject109) GetUseCombinedPowerOk() (*bool, bool) {
	if o == nil || o.UseCombinedPower == nil {
		return nil, false
	}
	return o.UseCombinedPower, true
}

// HasUseCombinedPower returns a boolean if a field has been set.
func (o *InlineObject109) HasUseCombinedPower() bool {
	if o != nil && o.UseCombinedPower != nil {
		return true
	}

	return false
}

// SetUseCombinedPower gets a reference to the given bool and assigns it to the UseCombinedPower field.
func (o *InlineObject109) SetUseCombinedPower(v bool) {
	o.UseCombinedPower = &v
}

// GetPowerExceptions returns the PowerExceptions field value if set, zero value otherwise.
func (o *InlineObject109) GetPowerExceptions() []NetworksNetworkIdSwitchSettingsPowerExceptions {
	if o == nil || o.PowerExceptions == nil {
		var ret []NetworksNetworkIdSwitchSettingsPowerExceptions
		return ret
	}
	return o.PowerExceptions
}

// GetPowerExceptionsOk returns a tuple with the PowerExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject109) GetPowerExceptionsOk() ([]NetworksNetworkIdSwitchSettingsPowerExceptions, bool) {
	if o == nil || o.PowerExceptions == nil {
		return nil, false
	}
	return o.PowerExceptions, true
}

// HasPowerExceptions returns a boolean if a field has been set.
func (o *InlineObject109) HasPowerExceptions() bool {
	if o != nil && o.PowerExceptions != nil {
		return true
	}

	return false
}

// SetPowerExceptions gets a reference to the given []NetworksNetworkIdSwitchSettingsPowerExceptions and assigns it to the PowerExceptions field.
func (o *InlineObject109) SetPowerExceptions(v []NetworksNetworkIdSwitchSettingsPowerExceptions) {
	o.PowerExceptions = v
}

func (o InlineObject109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vlan != nil {
		toSerialize["vlan"] = o.Vlan
	}
	if o.UseCombinedPower != nil {
		toSerialize["useCombinedPower"] = o.UseCombinedPower
	}
	if o.PowerExceptions != nil {
		toSerialize["powerExceptions"] = o.PowerExceptions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject109 struct {
	value *InlineObject109
	isSet bool
}

func (v NullableInlineObject109) Get() *InlineObject109 {
	return v.value
}

func (v *NullableInlineObject109) Set(val *InlineObject109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject109(val *InlineObject109) *NullableInlineObject109 {
	return &NullableInlineObject109{value: val, isSet: true}
}

func (v NullableInlineObject109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


