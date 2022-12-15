/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20022 struct for InlineResponse20022
type InlineResponse20022 struct {
	// Id of staged upgrade group
	GroupId *string `json:"groupId,omitempty"`
	// Name of the Staged Upgrade Group
	Name *string `json:"name,omitempty"`
	// Description of the Staged Upgrade Group
	Description *string `json:"description,omitempty"`
	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault *bool `json:"isDefault,omitempty"`
	AssignedDevices *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices `json:"assignedDevices,omitempty"`
}

// NewInlineResponse20022 instantiates a new InlineResponse20022 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022() *InlineResponse20022 {
	this := InlineResponse20022{}
	return &this
}

// NewInlineResponse20022WithDefaults instantiates a new InlineResponse20022 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022WithDefaults() *InlineResponse20022 {
	this := InlineResponse20022{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *InlineResponse20022) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineResponse20022) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *InlineResponse20022) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20022) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20022) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20022) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20022) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20022) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20022) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *InlineResponse20022) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *InlineResponse20022) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *InlineResponse20022) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetAssignedDevices returns the AssignedDevices field value if set, zero value otherwise.
func (o *InlineResponse20022) GetAssignedDevices() NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices {
	if o == nil || isNil(o.AssignedDevices) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices
		return ret
	}
	return *o.AssignedDevices
}

// GetAssignedDevicesOk returns a tuple with the AssignedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetAssignedDevicesOk() (*NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices, bool) {
	if o == nil || isNil(o.AssignedDevices) {
    return nil, false
	}
	return o.AssignedDevices, true
}

// HasAssignedDevices returns a boolean if a field has been set.
func (o *InlineResponse20022) HasAssignedDevices() bool {
	if o != nil && !isNil(o.AssignedDevices) {
		return true
	}

	return false
}

// SetAssignedDevices gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices and assigns it to the AssignedDevices field.
func (o *InlineResponse20022) SetAssignedDevices(v NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) {
	o.AssignedDevices = &v
}

func (o InlineResponse20022) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.AssignedDevices) {
		toSerialize["assignedDevices"] = o.AssignedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022 struct {
	value *InlineResponse20022
	isSet bool
}

func (v NullableInlineResponse20022) Get() *InlineResponse20022 {
	return v.value
}

func (v *NullableInlineResponse20022) Set(val *InlineResponse20022) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022(val *InlineResponse20022) *NullableInlineResponse20022 {
	return &NullableInlineResponse20022{value: val, isSet: true}
}

func (v NullableInlineResponse20022) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


