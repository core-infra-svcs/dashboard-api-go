/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject97 struct for InlineObject97
type InlineObject97 struct {
	// Name of the Staged Upgrade Group. Length must be 1 to 255 characters
	Name string `json:"name"`
	// Description of the Staged Upgrade Group. Length must be 1 to 255 characters
	Description *string `json:"description,omitempty"`
	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault bool `json:"isDefault"`
	AssignedDevices *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 `json:"assignedDevices,omitempty"`
}

// NewInlineObject97 instantiates a new InlineObject97 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject97(name string, isDefault bool) *InlineObject97 {
	this := InlineObject97{}
	this.Name = name
	this.IsDefault = isDefault
	return &this
}

// NewInlineObject97WithDefaults instantiates a new InlineObject97 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject97WithDefaults() *InlineObject97 {
	this := InlineObject97{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject97) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject97) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject97) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject97) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject97) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value
func (o *InlineObject97) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *InlineObject97) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetAssignedDevices returns the AssignedDevices field value if set, zero value otherwise.
func (o *InlineObject97) GetAssignedDevices() NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 {
	if o == nil || isNil(o.AssignedDevices) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1
		return ret
	}
	return *o.AssignedDevices
}

// GetAssignedDevicesOk returns a tuple with the AssignedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject97) GetAssignedDevicesOk() (*NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1, bool) {
	if o == nil || isNil(o.AssignedDevices) {
    return nil, false
	}
	return o.AssignedDevices, true
}

// HasAssignedDevices returns a boolean if a field has been set.
func (o *InlineObject97) HasAssignedDevices() bool {
	if o != nil && !isNil(o.AssignedDevices) {
		return true
	}

	return false
}

// SetAssignedDevices gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1 and assigns it to the AssignedDevices field.
func (o *InlineObject97) SetAssignedDevices(v NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices1) {
	o.AssignedDevices = &v
}

func (o InlineObject97) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.AssignedDevices) {
		toSerialize["assignedDevices"] = o.AssignedDevices
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


