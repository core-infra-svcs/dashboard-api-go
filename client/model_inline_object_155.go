/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject155 struct for InlineObject155
type InlineObject155 struct {
	// The spanning tree protocol status in network
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []NetworksNetworkIdSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"`
}

// NewInlineObject155 instantiates a new InlineObject155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject155() *InlineObject155 {
	this := InlineObject155{}
	return &this
}

// NewInlineObject155WithDefaults instantiates a new InlineObject155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject155WithDefaults() *InlineObject155 {
	this := InlineObject155{}
	return &this
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *InlineObject155) GetRstpEnabled() bool {
	if o == nil || isNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RstpEnabled) {
    return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *InlineObject155) HasRstpEnabled() bool {
	if o != nil && !isNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *InlineObject155) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpBridgePriority returns the StpBridgePriority field value if set, zero value otherwise.
func (o *InlineObject155) GetStpBridgePriority() []NetworksNetworkIdSwitchStpStpBridgePriority {
	if o == nil || isNil(o.StpBridgePriority) {
		var ret []NetworksNetworkIdSwitchStpStpBridgePriority
		return ret
	}
	return o.StpBridgePriority
}

// GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetStpBridgePriorityOk() ([]NetworksNetworkIdSwitchStpStpBridgePriority, bool) {
	if o == nil || isNil(o.StpBridgePriority) {
    return nil, false
	}
	return o.StpBridgePriority, true
}

// HasStpBridgePriority returns a boolean if a field has been set.
func (o *InlineObject155) HasStpBridgePriority() bool {
	if o != nil && !isNil(o.StpBridgePriority) {
		return true
	}

	return false
}

// SetStpBridgePriority gets a reference to the given []NetworksNetworkIdSwitchStpStpBridgePriority and assigns it to the StpBridgePriority field.
func (o *InlineObject155) SetStpBridgePriority(v []NetworksNetworkIdSwitchStpStpBridgePriority) {
	o.StpBridgePriority = v
}

func (o InlineObject155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !isNil(o.StpBridgePriority) {
		toSerialize["stpBridgePriority"] = o.StpBridgePriority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject155 struct {
	value *InlineObject155
	isSet bool
}

func (v NullableInlineObject155) Get() *InlineObject155 {
	return v.value
}

func (v *NullableInlineObject155) Set(val *InlineObject155) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject155) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject155(val *InlineObject155) *NullableInlineObject155 {
	return &NullableInlineObject155{value: val, isSet: true}
}

func (v NullableInlineObject155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


