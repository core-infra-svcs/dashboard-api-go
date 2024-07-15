/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200159 struct for InlineResponse200159
type InlineResponse200159 struct {
	// The spanning tree protocol status in network
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []InlineResponse200159StpBridgePriority `json:"stpBridgePriority,omitempty"`
}

// NewInlineResponse200159 instantiates a new InlineResponse200159 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200159() *InlineResponse200159 {
	this := InlineResponse200159{}
	return &this
}

// NewInlineResponse200159WithDefaults instantiates a new InlineResponse200159 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200159WithDefaults() *InlineResponse200159 {
	this := InlineResponse200159{}
	return &this
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *InlineResponse200159) GetRstpEnabled() bool {
	if o == nil || isNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200159) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RstpEnabled) {
    return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *InlineResponse200159) HasRstpEnabled() bool {
	if o != nil && !isNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *InlineResponse200159) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpBridgePriority returns the StpBridgePriority field value if set, zero value otherwise.
func (o *InlineResponse200159) GetStpBridgePriority() []InlineResponse200159StpBridgePriority {
	if o == nil || isNil(o.StpBridgePriority) {
		var ret []InlineResponse200159StpBridgePriority
		return ret
	}
	return o.StpBridgePriority
}

// GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200159) GetStpBridgePriorityOk() ([]InlineResponse200159StpBridgePriority, bool) {
	if o == nil || isNil(o.StpBridgePriority) {
    return nil, false
	}
	return o.StpBridgePriority, true
}

// HasStpBridgePriority returns a boolean if a field has been set.
func (o *InlineResponse200159) HasStpBridgePriority() bool {
	if o != nil && !isNil(o.StpBridgePriority) {
		return true
	}

	return false
}

// SetStpBridgePriority gets a reference to the given []InlineResponse200159StpBridgePriority and assigns it to the StpBridgePriority field.
func (o *InlineResponse200159) SetStpBridgePriority(v []InlineResponse200159StpBridgePriority) {
	o.StpBridgePriority = v
}

func (o InlineResponse200159) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !isNil(o.StpBridgePriority) {
		toSerialize["stpBridgePriority"] = o.StpBridgePriority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200159 struct {
	value *InlineResponse200159
	isSet bool
}

func (v NullableInlineResponse200159) Get() *InlineResponse200159 {
	return v.value
}

func (v *NullableInlineResponse200159) Set(val *InlineResponse200159) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200159) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200159) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200159(val *InlineResponse200159) *NullableInlineResponse200159 {
	return &NullableInlineResponse200159{value: val, isSet: true}
}

func (v NullableInlineResponse200159) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200159) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


