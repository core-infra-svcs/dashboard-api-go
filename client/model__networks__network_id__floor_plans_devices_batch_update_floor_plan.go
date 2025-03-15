/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan Floorplan to be assigned or unassigned
type NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan struct {
	// The ID of the floor plan to assign the device to, or null to unassign the device from its floor plan
	Id string `json:"id"`
}

// NewNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan instantiates a new NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan(id string) *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan {
	this := NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlanWithDefaults instantiates a new NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlanWithDefaults() *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan {
	this := NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) SetId(v string) {
	o.Id = v
}

func (o NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan struct {
	value *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) Get() *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) Set(val *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan(val *NetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) *NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan {
	return &NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


