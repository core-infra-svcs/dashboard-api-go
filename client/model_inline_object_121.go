/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject121 struct for InlineObject121
type InlineObject121 struct {
	// Array of switch or stack ports for updating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchPorts []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"`
	// Array of switch profile ports for updating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts `json:"switchProfilePorts,omitempty"`
}

// NewInlineObject121 instantiates a new InlineObject121 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject121() *InlineObject121 {
	this := InlineObject121{}
	return &this
}

// NewInlineObject121WithDefaults instantiates a new InlineObject121 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject121WithDefaults() *InlineObject121 {
	this := InlineObject121{}
	return &this
}

// GetSwitchPorts returns the SwitchPorts field value if set, zero value otherwise.
func (o *InlineObject121) GetSwitchPorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	if o == nil || isNil(o.SwitchPorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
		return ret
	}
	return o.SwitchPorts
}

// GetSwitchPortsOk returns a tuple with the SwitchPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject121) GetSwitchPortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts, bool) {
	if o == nil || isNil(o.SwitchPorts) {
    return nil, false
	}
	return o.SwitchPorts, true
}

// HasSwitchPorts returns a boolean if a field has been set.
func (o *InlineObject121) HasSwitchPorts() bool {
	if o != nil && !isNil(o.SwitchPorts) {
		return true
	}

	return false
}

// SetSwitchPorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts and assigns it to the SwitchPorts field.
func (o *InlineObject121) SetSwitchPorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) {
	o.SwitchPorts = v
}

// GetSwitchProfilePorts returns the SwitchProfilePorts field value if set, zero value otherwise.
func (o *InlineObject121) GetSwitchProfilePorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	if o == nil || isNil(o.SwitchProfilePorts) {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts
		return ret
	}
	return o.SwitchProfilePorts
}

// GetSwitchProfilePortsOk returns a tuple with the SwitchProfilePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject121) GetSwitchProfilePortsOk() ([]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts, bool) {
	if o == nil || isNil(o.SwitchProfilePorts) {
    return nil, false
	}
	return o.SwitchProfilePorts, true
}

// HasSwitchProfilePorts returns a boolean if a field has been set.
func (o *InlineObject121) HasSwitchProfilePorts() bool {
	if o != nil && !isNil(o.SwitchProfilePorts) {
		return true
	}

	return false
}

// SetSwitchProfilePorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts and assigns it to the SwitchProfilePorts field.
func (o *InlineObject121) SetSwitchProfilePorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) {
	o.SwitchProfilePorts = v
}

func (o InlineObject121) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SwitchPorts) {
		toSerialize["switchPorts"] = o.SwitchPorts
	}
	if !isNil(o.SwitchProfilePorts) {
		toSerialize["switchProfilePorts"] = o.SwitchProfilePorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject121 struct {
	value *InlineObject121
	isSet bool
}

func (v NullableInlineObject121) Get() *InlineObject121 {
	return v.value
}

func (v *NullableInlineObject121) Set(val *InlineObject121) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject121) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject121) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject121(val *InlineObject121) *NullableInlineObject121 {
	return &NullableInlineObject121{value: val, isSet: true}
}

func (v NullableInlineObject121) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject121) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


