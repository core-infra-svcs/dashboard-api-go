/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsStatusesSpanningTree The Spanning Tree Protocol (STP) information of the connected device.
type DevicesSerialSwitchPortsStatusesSpanningTree struct {
	// The current Spanning Tree Protocol statuses of the port.
	Statuses []string `json:"statuses,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesSpanningTree instantiates a new DevicesSerialSwitchPortsStatusesSpanningTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesSpanningTree() *DevicesSerialSwitchPortsStatusesSpanningTree {
	this := DevicesSerialSwitchPortsStatusesSpanningTree{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesSpanningTreeWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesSpanningTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesSpanningTreeWithDefaults() *DevicesSerialSwitchPortsStatusesSpanningTree {
	this := DevicesSerialSwitchPortsStatusesSpanningTree{}
	return &this
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSpanningTree) GetStatuses() []string {
	if o == nil || isNil(o.Statuses) {
		var ret []string
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSpanningTree) GetStatusesOk() ([]string, bool) {
	if o == nil || isNil(o.Statuses) {
    return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSpanningTree) HasStatuses() bool {
	if o != nil && !isNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []string and assigns it to the Statuses field.
func (o *DevicesSerialSwitchPortsStatusesSpanningTree) SetStatuses(v []string) {
	o.Statuses = v
}

func (o DevicesSerialSwitchPortsStatusesSpanningTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesSpanningTree struct {
	value *DevicesSerialSwitchPortsStatusesSpanningTree
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesSpanningTree) Get() *DevicesSerialSwitchPortsStatusesSpanningTree {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesSpanningTree) Set(val *DevicesSerialSwitchPortsStatusesSpanningTree) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesSpanningTree) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesSpanningTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesSpanningTree(val *DevicesSerialSwitchPortsStatusesSpanningTree) *NullableDevicesSerialSwitchPortsStatusesSpanningTree {
	return &NullableDevicesSerialSwitchPortsStatusesSpanningTree{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesSpanningTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesSpanningTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


