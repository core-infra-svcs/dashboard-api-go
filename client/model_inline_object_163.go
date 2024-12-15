/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject163 struct for InlineObject163
type InlineObject163 struct {
	// The spanning tree protocol status in network
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []NetworksNetworkIdSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"`
}

// NewInlineObject163 instantiates a new InlineObject163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject163() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// NewInlineObject163WithDefaults instantiates a new InlineObject163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject163WithDefaults() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *InlineObject163) GetRstpEnabled() bool {
	if o == nil || isNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RstpEnabled) {
    return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *InlineObject163) HasRstpEnabled() bool {
	if o != nil && !isNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *InlineObject163) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpBridgePriority returns the StpBridgePriority field value if set, zero value otherwise.
func (o *InlineObject163) GetStpBridgePriority() []NetworksNetworkIdSwitchStpStpBridgePriority {
	if o == nil || isNil(o.StpBridgePriority) {
		var ret []NetworksNetworkIdSwitchStpStpBridgePriority
		return ret
	}
	return o.StpBridgePriority
}

// GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetStpBridgePriorityOk() ([]NetworksNetworkIdSwitchStpStpBridgePriority, bool) {
	if o == nil || isNil(o.StpBridgePriority) {
    return nil, false
	}
	return o.StpBridgePriority, true
}

// HasStpBridgePriority returns a boolean if a field has been set.
func (o *InlineObject163) HasStpBridgePriority() bool {
	if o != nil && !isNil(o.StpBridgePriority) {
		return true
	}

	return false
}

// SetStpBridgePriority gets a reference to the given []NetworksNetworkIdSwitchStpStpBridgePriority and assigns it to the StpBridgePriority field.
func (o *InlineObject163) SetStpBridgePriority(v []NetworksNetworkIdSwitchStpStpBridgePriority) {
	o.StpBridgePriority = v
}

func (o InlineObject163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !isNil(o.StpBridgePriority) {
		toSerialize["stpBridgePriority"] = o.StpBridgePriority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject163 struct {
	value *InlineObject163
	isSet bool
}

func (v NullableInlineObject163) Get() *InlineObject163 {
	return v.value
}

func (v *NullableInlineObject163) Set(val *InlineObject163) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject163) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject163(val *InlineObject163) *NullableInlineObject163 {
	return &NullableInlineObject163{value: val, isSet: true}
}

func (v NullableInlineObject163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


