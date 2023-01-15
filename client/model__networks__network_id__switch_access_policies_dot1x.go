/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesDot1x 802.1x Settings
type NetworksNetworkIdSwitchAccessPoliciesDot1x struct {
	// Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
	ControlDirection *string `json:"controlDirection,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesDot1x instantiates a new NetworksNetworkIdSwitchAccessPoliciesDot1x object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesDot1x() *NetworksNetworkIdSwitchAccessPoliciesDot1x {
	this := NetworksNetworkIdSwitchAccessPoliciesDot1x{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesDot1xWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesDot1x object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesDot1xWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesDot1x {
	this := NetworksNetworkIdSwitchAccessPoliciesDot1x{}
	return &this
}

// GetControlDirection returns the ControlDirection field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesDot1x) GetControlDirection() string {
	if o == nil || isNil(o.ControlDirection) {
		var ret string
		return ret
	}
	return *o.ControlDirection
}

// GetControlDirectionOk returns a tuple with the ControlDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesDot1x) GetControlDirectionOk() (*string, bool) {
	if o == nil || isNil(o.ControlDirection) {
    return nil, false
	}
	return o.ControlDirection, true
}

// HasControlDirection returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesDot1x) HasControlDirection() bool {
	if o != nil && !isNil(o.ControlDirection) {
		return true
	}

	return false
}

// SetControlDirection gets a reference to the given string and assigns it to the ControlDirection field.
func (o *NetworksNetworkIdSwitchAccessPoliciesDot1x) SetControlDirection(v string) {
	o.ControlDirection = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesDot1x) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ControlDirection) {
		toSerialize["controlDirection"] = o.ControlDirection
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesDot1x struct {
	value *NetworksNetworkIdSwitchAccessPoliciesDot1x
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesDot1x) Get() *NetworksNetworkIdSwitchAccessPoliciesDot1x {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesDot1x) Set(val *NetworksNetworkIdSwitchAccessPoliciesDot1x) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesDot1x) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesDot1x) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesDot1x(val *NetworksNetworkIdSwitchAccessPoliciesDot1x) *NullableNetworksNetworkIdSwitchAccessPoliciesDot1x {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesDot1x{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesDot1x) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesDot1x) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


