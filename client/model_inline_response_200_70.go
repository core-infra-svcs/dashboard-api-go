/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20070 struct for InlineResponse20070
type InlineResponse20070 struct {
	// Management VLAN
	Vlan *int32 `json:"vlan,omitempty"`
	// The use Combined Power as the default behavior of secondary power supplies on supported devices.
	UseCombinedPower *bool `json:"useCombinedPower,omitempty"`
	// Exceptions on a per switch basis to \"useCombinedPower\"
	PowerExceptions []InlineResponse20070PowerExceptions `json:"powerExceptions,omitempty"`
}

// NewInlineResponse20070 instantiates a new InlineResponse20070 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070() *InlineResponse20070 {
	this := InlineResponse20070{}
	return &this
}

// NewInlineResponse20070WithDefaults instantiates a new InlineResponse20070 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070WithDefaults() *InlineResponse20070 {
	this := InlineResponse20070{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse20070) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse20070) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse20070) SetVlan(v int32) {
	o.Vlan = &v
}

// GetUseCombinedPower returns the UseCombinedPower field value if set, zero value otherwise.
func (o *InlineResponse20070) GetUseCombinedPower() bool {
	if o == nil || isNil(o.UseCombinedPower) {
		var ret bool
		return ret
	}
	return *o.UseCombinedPower
}

// GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetUseCombinedPowerOk() (*bool, bool) {
	if o == nil || isNil(o.UseCombinedPower) {
    return nil, false
	}
	return o.UseCombinedPower, true
}

// HasUseCombinedPower returns a boolean if a field has been set.
func (o *InlineResponse20070) HasUseCombinedPower() bool {
	if o != nil && !isNil(o.UseCombinedPower) {
		return true
	}

	return false
}

// SetUseCombinedPower gets a reference to the given bool and assigns it to the UseCombinedPower field.
func (o *InlineResponse20070) SetUseCombinedPower(v bool) {
	o.UseCombinedPower = &v
}

// GetPowerExceptions returns the PowerExceptions field value if set, zero value otherwise.
func (o *InlineResponse20070) GetPowerExceptions() []InlineResponse20070PowerExceptions {
	if o == nil || isNil(o.PowerExceptions) {
		var ret []InlineResponse20070PowerExceptions
		return ret
	}
	return o.PowerExceptions
}

// GetPowerExceptionsOk returns a tuple with the PowerExceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetPowerExceptionsOk() ([]InlineResponse20070PowerExceptions, bool) {
	if o == nil || isNil(o.PowerExceptions) {
    return nil, false
	}
	return o.PowerExceptions, true
}

// HasPowerExceptions returns a boolean if a field has been set.
func (o *InlineResponse20070) HasPowerExceptions() bool {
	if o != nil && !isNil(o.PowerExceptions) {
		return true
	}

	return false
}

// SetPowerExceptions gets a reference to the given []InlineResponse20070PowerExceptions and assigns it to the PowerExceptions field.
func (o *InlineResponse20070) SetPowerExceptions(v []InlineResponse20070PowerExceptions) {
	o.PowerExceptions = v
}

func (o InlineResponse20070) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20070 struct {
	value *InlineResponse20070
	isSet bool
}

func (v NullableInlineResponse20070) Get() *InlineResponse20070 {
	return v.value
}

func (v *NullableInlineResponse20070) Set(val *InlineResponse20070) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070(val *InlineResponse20070) *NullableInlineResponse20070 {
	return &NullableInlineResponse20070{value: val, isSet: true}
}

func (v NullableInlineResponse20070) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


