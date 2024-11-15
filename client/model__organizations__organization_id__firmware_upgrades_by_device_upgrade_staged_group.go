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

// OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup The staged upgrade group
type OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup struct {
	// Id of the staged upgrade group
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup{}
	return &this
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroupWithDefaults instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroupWithDefaults() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup struct {
	value *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) Get() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) Set(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup {
	return &NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeStagedGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


