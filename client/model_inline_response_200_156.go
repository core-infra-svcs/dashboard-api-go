/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200156 struct for InlineResponse200156
type InlineResponse200156 struct {
	// Management VLAN
	Vlan *int32 `json:"vlan,omitempty"`
	// The use Combined Power as the default behavior of secondary power supplies on supported devices.
	UseCombinedPower *bool `json:"useCombinedPower,omitempty"`
	// Exceptions on a per switch basis to \"useCombinedPower\"
	PowerExceptions []InlineResponse200156PowerExceptions `json:"powerExceptions,omitempty"`
	UplinkClientSampling *InlineResponse200156UplinkClientSampling `json:"uplinkClientSampling,omitempty"`
	MacBlocklist *InlineResponse200156MacBlocklist `json:"macBlocklist,omitempty"`
}

// NewInlineResponse200156 instantiates a new InlineResponse200156 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200156() *InlineResponse200156 {
	this := InlineResponse200156{}
	return &this
}

// NewInlineResponse200156WithDefaults instantiates a new InlineResponse200156 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200156WithDefaults() *InlineResponse200156 {
	this := InlineResponse200156{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse200156) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200156) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse200156) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse200156) SetVlan(v int32) {
	o.Vlan = &v
}

// GetUseCombinedPower returns the UseCombinedPower field value if set, zero value otherwise.
func (o *InlineResponse200156) GetUseCombinedPower() bool {
	if o == nil || isNil(o.UseCombinedPower) {
		var ret bool
		return ret
	}
	return *o.UseCombinedPower
}

// GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200156) GetUseCombinedPowerOk() (*bool, bool) {
	if o == nil || isNil(o.UseCombinedPower) {
    return nil, false
	}
	return o.UseCombinedPower, true
}

// HasUseCombinedPower returns a boolean if a field has been set.
func (o *InlineResponse200156) HasUseCombinedPower() bool {
	if o != nil && !isNil(o.UseCombinedPower) {
		return true
	}

	return false
}

// SetUseCombinedPower gets a reference to the given bool and assigns it to the UseCombinedPower field.
func (o *InlineResponse200156) SetUseCombinedPower(v bool) {
	o.UseCombinedPower = &v
}

// GetPowerExceptions returns the PowerExceptions field value if set, zero value otherwise.
func (o *InlineResponse200156) GetPowerExceptions() []InlineResponse200156PowerExceptions {
	if o == nil || isNil(o.PowerExceptions) {
		var ret []InlineResponse200156PowerExceptions
		return ret
	}
	return o.PowerExceptions
}

// GetPowerExceptionsOk returns a tuple with the PowerExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200156) GetPowerExceptionsOk() ([]InlineResponse200156PowerExceptions, bool) {
	if o == nil || isNil(o.PowerExceptions) {
    return nil, false
	}
	return o.PowerExceptions, true
}

// HasPowerExceptions returns a boolean if a field has been set.
func (o *InlineResponse200156) HasPowerExceptions() bool {
	if o != nil && !isNil(o.PowerExceptions) {
		return true
	}

	return false
}

// SetPowerExceptions gets a reference to the given []InlineResponse200156PowerExceptions and assigns it to the PowerExceptions field.
func (o *InlineResponse200156) SetPowerExceptions(v []InlineResponse200156PowerExceptions) {
	o.PowerExceptions = v
}

// GetUplinkClientSampling returns the UplinkClientSampling field value if set, zero value otherwise.
func (o *InlineResponse200156) GetUplinkClientSampling() InlineResponse200156UplinkClientSampling {
	if o == nil || isNil(o.UplinkClientSampling) {
		var ret InlineResponse200156UplinkClientSampling
		return ret
	}
	return *o.UplinkClientSampling
}

// GetUplinkClientSamplingOk returns a tuple with the UplinkClientSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200156) GetUplinkClientSamplingOk() (*InlineResponse200156UplinkClientSampling, bool) {
	if o == nil || isNil(o.UplinkClientSampling) {
    return nil, false
	}
	return o.UplinkClientSampling, true
}

// HasUplinkClientSampling returns a boolean if a field has been set.
func (o *InlineResponse200156) HasUplinkClientSampling() bool {
	if o != nil && !isNil(o.UplinkClientSampling) {
		return true
	}

	return false
}

// SetUplinkClientSampling gets a reference to the given InlineResponse200156UplinkClientSampling and assigns it to the UplinkClientSampling field.
func (o *InlineResponse200156) SetUplinkClientSampling(v InlineResponse200156UplinkClientSampling) {
	o.UplinkClientSampling = &v
}

// GetMacBlocklist returns the MacBlocklist field value if set, zero value otherwise.
func (o *InlineResponse200156) GetMacBlocklist() InlineResponse200156MacBlocklist {
	if o == nil || isNil(o.MacBlocklist) {
		var ret InlineResponse200156MacBlocklist
		return ret
	}
	return *o.MacBlocklist
}

// GetMacBlocklistOk returns a tuple with the MacBlocklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200156) GetMacBlocklistOk() (*InlineResponse200156MacBlocklist, bool) {
	if o == nil || isNil(o.MacBlocklist) {
    return nil, false
	}
	return o.MacBlocklist, true
}

// HasMacBlocklist returns a boolean if a field has been set.
func (o *InlineResponse200156) HasMacBlocklist() bool {
	if o != nil && !isNil(o.MacBlocklist) {
		return true
	}

	return false
}

// SetMacBlocklist gets a reference to the given InlineResponse200156MacBlocklist and assigns it to the MacBlocklist field.
func (o *InlineResponse200156) SetMacBlocklist(v InlineResponse200156MacBlocklist) {
	o.MacBlocklist = &v
}

func (o InlineResponse200156) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200156 struct {
	value *InlineResponse200156
	isSet bool
}

func (v NullableInlineResponse200156) Get() *InlineResponse200156 {
	return v.value
}

func (v *NullableInlineResponse200156) Set(val *InlineResponse200156) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200156) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200156) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200156(val *InlineResponse200156) *NullableInlineResponse200156 {
	return &NullableInlineResponse200156{value: val, isSet: true}
}

func (v NullableInlineResponse200156) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200156) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


