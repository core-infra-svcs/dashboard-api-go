/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchAccessPoliciesCounts Counts associated with the access policy
type NetworksNetworkIdSwitchAccessPoliciesCounts struct {
	Ports *NetworksNetworkIdSwitchAccessPoliciesCountsPorts `json:"ports,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesCounts instantiates a new NetworksNetworkIdSwitchAccessPoliciesCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesCounts() *NetworksNetworkIdSwitchAccessPoliciesCounts {
	this := NetworksNetworkIdSwitchAccessPoliciesCounts{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesCountsWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesCountsWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesCounts {
	this := NetworksNetworkIdSwitchAccessPoliciesCounts{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesCounts) GetPorts() NetworksNetworkIdSwitchAccessPoliciesCountsPorts {
	if o == nil || isNil(o.Ports) {
		var ret NetworksNetworkIdSwitchAccessPoliciesCountsPorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesCounts) GetPortsOk() (*NetworksNetworkIdSwitchAccessPoliciesCountsPorts, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesCounts) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given NetworksNetworkIdSwitchAccessPoliciesCountsPorts and assigns it to the Ports field.
func (o *NetworksNetworkIdSwitchAccessPoliciesCounts) SetPorts(v NetworksNetworkIdSwitchAccessPoliciesCountsPorts) {
	o.Ports = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesCounts struct {
	value *NetworksNetworkIdSwitchAccessPoliciesCounts
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesCounts) Get() *NetworksNetworkIdSwitchAccessPoliciesCounts {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesCounts) Set(val *NetworksNetworkIdSwitchAccessPoliciesCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesCounts(val *NetworksNetworkIdSwitchAccessPoliciesCounts) *NullableNetworksNetworkIdSwitchAccessPoliciesCounts {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesCounts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


