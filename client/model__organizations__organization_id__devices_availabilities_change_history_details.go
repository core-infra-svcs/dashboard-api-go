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

// OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails Details about the status changes
type OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails struct {
	// Details about the old status
	Old []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld `json:"old,omitempty"`
	// Details about the new status
	New []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld `json:"new,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsWithDefaults instantiates a new OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsWithDefaults() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails {
	this := OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails{}
	return &this
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) GetOld() []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld {
	if o == nil || isNil(o.Old) {
		var ret []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld
		return ret
	}
	return o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) GetOldOk() ([]OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld, bool) {
	if o == nil || isNil(o.Old) {
    return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) HasOld() bool {
	if o != nil && !isNil(o.Old) {
		return true
	}

	return false
}

// SetOld gets a reference to the given []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld and assigns it to the Old field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) SetOld(v []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) {
	o.Old = v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) GetNew() []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld {
	if o == nil || isNil(o.New) {
		var ret []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld
		return ret
	}
	return o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) GetNewOk() ([]OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld, bool) {
	if o == nil || isNil(o.New) {
    return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) HasNew() bool {
	if o != nil && !isNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld and assigns it to the New field.
func (o *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) SetNew(v []OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetailsOld) {
	o.New = v
}

func (o OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Old) {
		toSerialize["old"] = o.Old
	}
	if !isNil(o.New) {
		toSerialize["new"] = o.New
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails struct {
	value *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) Get() *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) Set(val *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails(val *OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails {
	return &NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


