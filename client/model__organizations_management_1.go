/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsManagement1 Information about the organization's management system
type OrganizationsManagement1 struct {
	// Details related to organization management, possibly empty
	Details []OrganizationsManagementDetails `json:"details,omitempty"`
}

// NewOrganizationsManagement1 instantiates a new OrganizationsManagement1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsManagement1() *OrganizationsManagement1 {
	this := OrganizationsManagement1{}
	return &this
}

// NewOrganizationsManagement1WithDefaults instantiates a new OrganizationsManagement1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsManagement1WithDefaults() *OrganizationsManagement1 {
	this := OrganizationsManagement1{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OrganizationsManagement1) GetDetails() []OrganizationsManagementDetails {
	if o == nil || isNil(o.Details) {
		var ret []OrganizationsManagementDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsManagement1) GetDetailsOk() ([]OrganizationsManagementDetails, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OrganizationsManagement1) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []OrganizationsManagementDetails and assigns it to the Details field.
func (o *OrganizationsManagement1) SetDetails(v []OrganizationsManagementDetails) {
	o.Details = v
}

func (o OrganizationsManagement1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsManagement1 struct {
	value *OrganizationsManagement1
	isSet bool
}

func (v NullableOrganizationsManagement1) Get() *OrganizationsManagement1 {
	return v.value
}

func (v *NullableOrganizationsManagement1) Set(val *OrganizationsManagement1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsManagement1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsManagement1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsManagement1(val *OrganizationsManagement1) *NullableOrganizationsManagement1 {
	return &NullableOrganizationsManagement1{value: val, isSet: true}
}

func (v NullableOrganizationsManagement1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsManagement1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

