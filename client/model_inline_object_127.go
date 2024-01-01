/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject127 struct for InlineObject127
type InlineObject127 struct {
	// Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchPorts []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"`
	// Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts `json:"switchProfilePorts,omitempty"`
}

// NewInlineObject127 instantiates a new InlineObject127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject127() *InlineObject127 {
	this := InlineObject127{}
	return &this
}

// NewInlineObject127WithDefaults instantiates a new InlineObject127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject127WithDefaults() *InlineObject127 {
	this := InlineObject127{}
	return &this
}

// GetSwitchPorts returns the SwitchPorts field value if set, zero value otherwise.
func (o *InlineObject127) GetSwitchPorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	if o == nil || isNil(o.SwitchPorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
		return ret
	}
	return o.SwitchPorts
}

// GetSwitchPortsOk returns a tuple with the SwitchPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetSwitchPortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts, bool) {
	if o == nil || isNil(o.SwitchPorts) {
    return nil, false
	}
	return o.SwitchPorts, true
}

// HasSwitchPorts returns a boolean if a field has been set.
func (o *InlineObject127) HasSwitchPorts() bool {
	if o != nil && !isNil(o.SwitchPorts) {
		return true
	}

	return false
}

// SetSwitchPorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts and assigns it to the SwitchPorts field.
func (o *InlineObject127) SetSwitchPorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) {
	o.SwitchPorts = v
}

// GetSwitchProfilePorts returns the SwitchProfilePorts field value if set, zero value otherwise.
func (o *InlineObject127) GetSwitchProfilePorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	if o == nil || isNil(o.SwitchProfilePorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts
		return ret
	}
	return o.SwitchProfilePorts
}

// GetSwitchProfilePortsOk returns a tuple with the SwitchProfilePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetSwitchProfilePortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts, bool) {
	if o == nil || isNil(o.SwitchProfilePorts) {
    return nil, false
	}
	return o.SwitchProfilePorts, true
}

// HasSwitchProfilePorts returns a boolean if a field has been set.
func (o *InlineObject127) HasSwitchProfilePorts() bool {
	if o != nil && !isNil(o.SwitchProfilePorts) {
		return true
	}

	return false
}

// SetSwitchProfilePorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts and assigns it to the SwitchProfilePorts field.
func (o *InlineObject127) SetSwitchProfilePorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) {
	o.SwitchProfilePorts = v
}

func (o InlineObject127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SwitchPorts) {
		toSerialize["switchPorts"] = o.SwitchPorts
	}
	if !isNil(o.SwitchProfilePorts) {
		toSerialize["switchProfilePorts"] = o.SwitchProfilePorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject127 struct {
	value *InlineObject127
	isSet bool
}

func (v NullableInlineObject127) Get() *InlineObject127 {
	return v.value
}

func (v *NullableInlineObject127) Set(val *InlineObject127) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject127) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject127(val *InlineObject127) *NullableInlineObject127 {
	return &NullableInlineObject127{value: val, isSet: true}
}

func (v NullableInlineObject127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


