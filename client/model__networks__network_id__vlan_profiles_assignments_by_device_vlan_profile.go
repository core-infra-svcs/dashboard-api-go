/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile The VLAN Profile
type NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile struct {
	// IName of the VLAN Profile
	Iname *string `json:"iname,omitempty"`
	// Name of the VLAN Profile
	Name *string `json:"name,omitempty"`
	// Is this VLAN profile the default for the network?
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile instantiates a new NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile() *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile {
	this := NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile{}
	return &this
}

// NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfileWithDefaults instantiates a new NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfileWithDefaults() *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile {
	this := NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile{}
	return &this
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) GetIname() string {
	if o == nil || isNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) GetInameOk() (*string, bool) {
	if o == nil || isNil(o.Iname) {
    return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) HasIname() bool {
	if o != nil && !isNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) SetIname(v string) {
	o.Iname = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) SetName(v string) {
	o.Name = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile struct {
	value *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile
	isSet bool
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) Get() *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile {
	return v.value
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) Set(val *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile(val *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile {
	return &NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


