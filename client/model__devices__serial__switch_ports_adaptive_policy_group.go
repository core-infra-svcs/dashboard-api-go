/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsAdaptivePolicyGroup The adaptive policy group data of the port.
type DevicesSerialSwitchPortsAdaptivePolicyGroup struct {
	// The ID of the adaptive policy group.
	Id *string `json:"id,omitempty"`
	// The name of the adaptive policy group.
	Name *string `json:"name,omitempty"`
}

// NewDevicesSerialSwitchPortsAdaptivePolicyGroup instantiates a new DevicesSerialSwitchPortsAdaptivePolicyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsAdaptivePolicyGroup() *DevicesSerialSwitchPortsAdaptivePolicyGroup {
	this := DevicesSerialSwitchPortsAdaptivePolicyGroup{}
	return &this
}

// NewDevicesSerialSwitchPortsAdaptivePolicyGroupWithDefaults instantiates a new DevicesSerialSwitchPortsAdaptivePolicyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsAdaptivePolicyGroupWithDefaults() *DevicesSerialSwitchPortsAdaptivePolicyGroup {
	this := DevicesSerialSwitchPortsAdaptivePolicyGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevicesSerialSwitchPortsAdaptivePolicyGroup) SetName(v string) {
	o.Name = &v
}

func (o DevicesSerialSwitchPortsAdaptivePolicyGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsAdaptivePolicyGroup struct {
	value *DevicesSerialSwitchPortsAdaptivePolicyGroup
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsAdaptivePolicyGroup) Get() *DevicesSerialSwitchPortsAdaptivePolicyGroup {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsAdaptivePolicyGroup) Set(val *DevicesSerialSwitchPortsAdaptivePolicyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsAdaptivePolicyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsAdaptivePolicyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsAdaptivePolicyGroup(val *DevicesSerialSwitchPortsAdaptivePolicyGroup) *NullableDevicesSerialSwitchPortsAdaptivePolicyGroup {
	return &NullableDevicesSerialSwitchPortsAdaptivePolicyGroup{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsAdaptivePolicyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsAdaptivePolicyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


