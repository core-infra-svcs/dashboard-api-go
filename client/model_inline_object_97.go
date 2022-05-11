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

// InlineObject97 struct for InlineObject97
type InlineObject97 struct {
	// Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchPorts *[]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts `json:"switchPorts,omitempty"`
	// Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.
	SwitchProfilePorts *[]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts `json:"switchProfilePorts,omitempty"`
}

// NewInlineObject97 instantiates a new InlineObject97 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject97() *InlineObject97 {
	this := InlineObject97{}
	return &this
}

// NewInlineObject97WithDefaults instantiates a new InlineObject97 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject97WithDefaults() *InlineObject97 {
	this := InlineObject97{}
	return &this
}

// GetSwitchPorts returns the SwitchPorts field value if set, zero value otherwise.
func (o *InlineObject97) GetSwitchPorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	if o == nil || o.SwitchPorts == nil {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
		return ret
	}
	return *o.SwitchPorts
}

// GetSwitchPortsOk returns a tuple with the SwitchPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetSwitchPortsOk() (*[]NetworksNetworkIdSwitchLinkAggregationsSwitchPorts, bool) {
	if o == nil || o.SwitchPorts == nil {
		return nil, false
	}
	return o.SwitchPorts, true
}

// HasSwitchPorts returns a boolean if a field has been set.
func (o *InlineObject97) HasSwitchPorts() bool {
	if o != nil && o.SwitchPorts != nil {
		return true
	}

	return false
}

// SetSwitchPorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts and assigns it to the SwitchPorts field.
func (o *InlineObject97) SetSwitchPorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) {
	o.SwitchPorts = &v
}

// GetSwitchProfilePorts returns the SwitchProfilePorts field value if set, zero value otherwise.
func (o *InlineObject97) GetSwitchProfilePorts() []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	if o == nil || o.SwitchProfilePorts == nil {
		var ret []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts
		return ret
	}
	return *o.SwitchProfilePorts
}

// GetSwitchProfilePortsOk returns a tuple with the SwitchProfilePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetSwitchProfilePortsOk() (*[]NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts, bool) {
	if o == nil || o.SwitchProfilePorts == nil {
		return nil, false
	}
	return o.SwitchProfilePorts, true
}

// HasSwitchProfilePorts returns a boolean if a field has been set.
func (o *InlineObject97) HasSwitchProfilePorts() bool {
	if o != nil && o.SwitchProfilePorts != nil {
		return true
	}

	return false
}

// SetSwitchProfilePorts gets a reference to the given []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts and assigns it to the SwitchProfilePorts field.
func (o *InlineObject97) SetSwitchProfilePorts(v []NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) {
	o.SwitchProfilePorts = &v
}

func (o InlineObject97) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SwitchPorts != nil {
		toSerialize["switchPorts"] = o.SwitchPorts
	}
	if o.SwitchProfilePorts != nil {
		toSerialize["switchProfilePorts"] = o.SwitchProfilePorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject97 struct {
	value *InlineObject97
	isSet bool
}

func (v NullableInlineObject97) Get() *InlineObject97 {
	return v.value
}

func (v *NullableInlineObject97) Set(val *InlineObject97) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject97) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject97) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject97(val *InlineObject97) *NullableInlineObject97 {
	return &NullableInlineObject97{value: val, isSet: true}
}

func (v NullableInlineObject97) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject97) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


