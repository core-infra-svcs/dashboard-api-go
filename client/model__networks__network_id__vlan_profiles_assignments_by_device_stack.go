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

// NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack The Switch Stack the device belongs to
type NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack struct {
	// ID of the Switch Stack
	Id *string `json:"id,omitempty"`
}

// NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack instantiates a new NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack() *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack {
	this := NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack{}
	return &this
}

// NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceStackWithDefaults instantiates a new NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdVlanProfilesAssignmentsByDeviceStackWithDefaults() *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack {
	this := NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) SetId(v string) {
	o.Id = &v
}

func (o NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack struct {
	value *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack
	isSet bool
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) Get() *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack {
	return v.value
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) Set(val *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack(val *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack {
	return &NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


