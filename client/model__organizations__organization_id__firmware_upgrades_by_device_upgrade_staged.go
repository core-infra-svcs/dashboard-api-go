/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged Staged upgrade
type OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged struct {
	Group *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup `json:"group,omitempty"`
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged{}
	return &this
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedWithDefaults instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedWithDefaults() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) GetGroup() OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup {
	if o == nil || isNil(o.Group) {
		var ret OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) GetGroupOk() (*OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup and assigns it to the Group field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) SetGroup(v OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) {
	o.Group = &v
}

func (o OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged struct {
	value *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) Get() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) Set(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged {
	return &NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStaged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


