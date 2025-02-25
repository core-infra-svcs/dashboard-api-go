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

// InlineResponse200168 struct for InlineResponse200168
type InlineResponse200168 struct {
	// The spanning tree protocol status in network
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []InlineResponse200168StpBridgePriority `json:"stpBridgePriority,omitempty"`
}

// NewInlineResponse200168 instantiates a new InlineResponse200168 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200168() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// NewInlineResponse200168WithDefaults instantiates a new InlineResponse200168 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200168WithDefaults() *InlineResponse200168 {
	this := InlineResponse200168{}
	return &this
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *InlineResponse200168) GetRstpEnabled() bool {
	if o == nil || isNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RstpEnabled) {
    return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *InlineResponse200168) HasRstpEnabled() bool {
	if o != nil && !isNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *InlineResponse200168) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpBridgePriority returns the StpBridgePriority field value if set, zero value otherwise.
func (o *InlineResponse200168) GetStpBridgePriority() []InlineResponse200168StpBridgePriority {
	if o == nil || isNil(o.StpBridgePriority) {
		var ret []InlineResponse200168StpBridgePriority
		return ret
	}
	return o.StpBridgePriority
}

// GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168) GetStpBridgePriorityOk() ([]InlineResponse200168StpBridgePriority, bool) {
	if o == nil || isNil(o.StpBridgePriority) {
    return nil, false
	}
	return o.StpBridgePriority, true
}

// HasStpBridgePriority returns a boolean if a field has been set.
func (o *InlineResponse200168) HasStpBridgePriority() bool {
	if o != nil && !isNil(o.StpBridgePriority) {
		return true
	}

	return false
}

// SetStpBridgePriority gets a reference to the given []InlineResponse200168StpBridgePriority and assigns it to the StpBridgePriority field.
func (o *InlineResponse200168) SetStpBridgePriority(v []InlineResponse200168StpBridgePriority) {
	o.StpBridgePriority = v
}

func (o InlineResponse200168) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !isNil(o.StpBridgePriority) {
		toSerialize["stpBridgePriority"] = o.StpBridgePriority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200168 struct {
	value *InlineResponse200168
	isSet bool
}

func (v NullableInlineResponse200168) Get() *InlineResponse200168 {
	return v.value
}

func (v *NullableInlineResponse200168) Set(val *InlineResponse200168) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200168) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200168) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200168(val *InlineResponse200168) *NullableInlineResponse200168 {
	return &NullableInlineResponse200168{value: val, isSet: true}
}

func (v NullableInlineResponse200168) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200168) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


