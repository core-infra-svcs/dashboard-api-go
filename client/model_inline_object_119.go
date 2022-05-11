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

// InlineObject119 struct for InlineObject119
type InlineObject119 struct {
	// The spanning tree protocol status in network
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings.
	StpBridgePriority *[]NetworksNetworkIdSwitchStpStpBridgePriority `json:"stpBridgePriority,omitempty"`
}

// NewInlineObject119 instantiates a new InlineObject119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject119() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// NewInlineObject119WithDefaults instantiates a new InlineObject119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject119WithDefaults() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *InlineObject119) GetRstpEnabled() bool {
	if o == nil || o.RstpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || o.RstpEnabled == nil {
		return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *InlineObject119) HasRstpEnabled() bool {
	if o != nil && o.RstpEnabled != nil {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *InlineObject119) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpBridgePriority returns the StpBridgePriority field value if set, zero value otherwise.
func (o *InlineObject119) GetStpBridgePriority() []NetworksNetworkIdSwitchStpStpBridgePriority {
	if o == nil || o.StpBridgePriority == nil {
		var ret []NetworksNetworkIdSwitchStpStpBridgePriority
		return ret
	}
	return *o.StpBridgePriority
}

// GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetStpBridgePriorityOk() (*[]NetworksNetworkIdSwitchStpStpBridgePriority, bool) {
	if o == nil || o.StpBridgePriority == nil {
		return nil, false
	}
	return o.StpBridgePriority, true
}

// HasStpBridgePriority returns a boolean if a field has been set.
func (o *InlineObject119) HasStpBridgePriority() bool {
	if o != nil && o.StpBridgePriority != nil {
		return true
	}

	return false
}

// SetStpBridgePriority gets a reference to the given []NetworksNetworkIdSwitchStpStpBridgePriority and assigns it to the StpBridgePriority field.
func (o *InlineObject119) SetStpBridgePriority(v []NetworksNetworkIdSwitchStpStpBridgePriority) {
	o.StpBridgePriority = &v
}

func (o InlineObject119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RstpEnabled != nil {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if o.StpBridgePriority != nil {
		toSerialize["stpBridgePriority"] = o.StpBridgePriority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject119 struct {
	value *InlineObject119
	isSet bool
}

func (v NullableInlineObject119) Get() *InlineObject119 {
	return v.value
}

func (v *NullableInlineObject119) Set(val *InlineObject119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject119(val *InlineObject119) *NullableInlineObject119 {
	return &NullableInlineObject119{value: val, isSet: true}
}

func (v NullableInlineObject119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


