/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta Other metadata related to this result set.
type OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta struct {
	Counts *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts `json:"counts,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaWithDefaults instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaWithDefaults() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) GetCounts() OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) GetCountsOk() (*OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts and assigns it to the Counts field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) SetCounts(v OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) {
	o.Counts = &v
}

func (o OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta struct {
	value *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) Get() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) Set(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta {
	return &NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


