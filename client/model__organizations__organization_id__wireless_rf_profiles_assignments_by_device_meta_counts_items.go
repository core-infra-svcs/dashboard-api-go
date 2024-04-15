/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems The count metadata.
type OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems struct {
	// The total number of serials.
	Total *int32 `json:"total,omitempty"`
	// The number of serials remaining based on current pagination location within the dataset.
	Remaining *int32 `json:"remaining,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItemsWithDefaults instantiates a new OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItemsWithDefaults() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems {
	this := OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) SetTotal(v int32) {
	o.Total = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) GetRemaining() int32 {
	if o == nil || isNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) GetRemainingOk() (*int32, bool) {
	if o == nil || isNil(o.Remaining) {
    return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) HasRemaining() bool {
	if o != nil && !isNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) SetRemaining(v int32) {
	o.Remaining = &v
}

func (o OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems struct {
	value *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) Get() *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) Set(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems(val *OrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems {
	return &NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessRfProfilesAssignmentsByDeviceMetaCountsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


