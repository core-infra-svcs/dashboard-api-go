/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts Count metadata related to this result set.
type OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts struct {
	Items *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems `json:"items,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsWithDefaults instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsWithDefaults() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) GetItems() OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) GetItemsOk() (*OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems and assigns it to the Items field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) SetItems(v OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) {
	o.Items = &v
}

func (o OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts struct {
	value *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) Get() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) Set(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts {
	return &NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


