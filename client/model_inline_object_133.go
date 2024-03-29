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

// InlineObject133 struct for InlineObject133
type InlineObject133 struct {
	// Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchPorts []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"`
	// Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts `json:"switchProfilePorts,omitempty"`
}

// NewInlineObject133 instantiates a new InlineObject133 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject133() *InlineObject133 {
	this := InlineObject133{}
	return &this
}

// NewInlineObject133WithDefaults instantiates a new InlineObject133 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject133WithDefaults() *InlineObject133 {
	this := InlineObject133{}
	return &this
}

// GetSwitchPorts returns the SwitchPorts field value if set, zero value otherwise.
func (o *InlineObject133) GetSwitchPorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	if o == nil || isNil(o.SwitchPorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
		return ret
	}
	return o.SwitchPorts
}

// GetSwitchPortsOk returns a tuple with the SwitchPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject133) GetSwitchPortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts, bool) {
	if o == nil || isNil(o.SwitchPorts) {
    return nil, false
	}
	return o.SwitchPorts, true
}

// HasSwitchPorts returns a boolean if a field has been set.
func (o *InlineObject133) HasSwitchPorts() bool {
	if o != nil && !isNil(o.SwitchPorts) {
		return true
	}

	return false
}

// SetSwitchPorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts and assigns it to the SwitchPorts field.
func (o *InlineObject133) SetSwitchPorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) {
	o.SwitchPorts = v
}

// GetSwitchProfilePorts returns the SwitchProfilePorts field value if set, zero value otherwise.
func (o *InlineObject133) GetSwitchProfilePorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	if o == nil || isNil(o.SwitchProfilePorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts
		return ret
	}
	return o.SwitchProfilePorts
}

// GetSwitchProfilePortsOk returns a tuple with the SwitchProfilePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject133) GetSwitchProfilePortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts, bool) {
	if o == nil || isNil(o.SwitchProfilePorts) {
    return nil, false
	}
	return o.SwitchProfilePorts, true
}

// HasSwitchProfilePorts returns a boolean if a field has been set.
func (o *InlineObject133) HasSwitchProfilePorts() bool {
	if o != nil && !isNil(o.SwitchProfilePorts) {
		return true
	}

	return false
}

// SetSwitchProfilePorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts and assigns it to the SwitchProfilePorts field.
func (o *InlineObject133) SetSwitchProfilePorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) {
	o.SwitchProfilePorts = v
}

func (o InlineObject133) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SwitchPorts) {
		toSerialize["switchPorts"] = o.SwitchPorts
	}
	if !isNil(o.SwitchProfilePorts) {
		toSerialize["switchProfilePorts"] = o.SwitchProfilePorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject133 struct {
	value *InlineObject133
	isSet bool
}

func (v NullableInlineObject133) Get() *InlineObject133 {
	return v.value
}

func (v *NullableInlineObject133) Set(val *InlineObject133) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject133) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject133) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject133(val *InlineObject133) *NullableInlineObject133 {
	return &NullableInlineObject133{value: val, isSet: true}
}

func (v NullableInlineObject133) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject133) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


