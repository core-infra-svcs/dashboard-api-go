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

// NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile The VLAN Profile
type NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile struct {
	// IName of the VLAN Profile
	Iname *string `json:"iname,omitempty"`
}

// NewNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile instantiates a new NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile() *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile {
	this := NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile{}
	return &this
}

// NewNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfileWithDefaults instantiates a new NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfileWithDefaults() *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile {
	this := NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile{}
	return &this
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) GetIname() string {
	if o == nil || isNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) GetInameOk() (*string, bool) {
	if o == nil || isNil(o.Iname) {
    return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) HasIname() bool {
	if o != nil && !isNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) SetIname(v string) {
	o.Iname = &v
}

func (o NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile struct {
	value *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile
	isSet bool
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) Get() *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile {
	return v.value
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) Set(val *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile(val *NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) *NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile {
	return &NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


