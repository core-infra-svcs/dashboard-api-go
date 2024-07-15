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

// NetworksNetworkIdSwitchAccessPoliciesCountsPorts Counts associated with ports
type NetworksNetworkIdSwitchAccessPoliciesCountsPorts struct {
	// Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.
	WithThisPolicy *int32 `json:"withThisPolicy,omitempty"`
}

// NewNetworksNetworkIdSwitchAccessPoliciesCountsPorts instantiates a new NetworksNetworkIdSwitchAccessPoliciesCountsPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchAccessPoliciesCountsPorts() *NetworksNetworkIdSwitchAccessPoliciesCountsPorts {
	this := NetworksNetworkIdSwitchAccessPoliciesCountsPorts{}
	return &this
}

// NewNetworksNetworkIdSwitchAccessPoliciesCountsPortsWithDefaults instantiates a new NetworksNetworkIdSwitchAccessPoliciesCountsPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchAccessPoliciesCountsPortsWithDefaults() *NetworksNetworkIdSwitchAccessPoliciesCountsPorts {
	this := NetworksNetworkIdSwitchAccessPoliciesCountsPorts{}
	return &this
}

// GetWithThisPolicy returns the WithThisPolicy field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchAccessPoliciesCountsPorts) GetWithThisPolicy() int32 {
	if o == nil || isNil(o.WithThisPolicy) {
		var ret int32
		return ret
	}
	return *o.WithThisPolicy
}

// GetWithThisPolicyOk returns a tuple with the WithThisPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesCountsPorts) GetWithThisPolicyOk() (*int32, bool) {
	if o == nil || isNil(o.WithThisPolicy) {
    return nil, false
	}
	return o.WithThisPolicy, true
}

// HasWithThisPolicy returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchAccessPoliciesCountsPorts) HasWithThisPolicy() bool {
	if o != nil && !isNil(o.WithThisPolicy) {
		return true
	}

	return false
}

// SetWithThisPolicy gets a reference to the given int32 and assigns it to the WithThisPolicy field.
func (o *NetworksNetworkIdSwitchAccessPoliciesCountsPorts) SetWithThisPolicy(v int32) {
	o.WithThisPolicy = &v
}

func (o NetworksNetworkIdSwitchAccessPoliciesCountsPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WithThisPolicy) {
		toSerialize["withThisPolicy"] = o.WithThisPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts struct {
	value *NetworksNetworkIdSwitchAccessPoliciesCountsPorts
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts) Get() *NetworksNetworkIdSwitchAccessPoliciesCountsPorts {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts) Set(val *NetworksNetworkIdSwitchAccessPoliciesCountsPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts(val *NetworksNetworkIdSwitchAccessPoliciesCountsPorts) *NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts {
	return &NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchAccessPoliciesCountsPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


