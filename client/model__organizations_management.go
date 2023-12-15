/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsManagement Information about the organization's management system
type OrganizationsManagement struct {
	// Details related to organization management, possibly empty. Details may be named 'MSP ID', 'IP restriction mode for API', or 'IP restriction mode for dashboard', if the organization admin has configured any.
	Details []OrganizationsManagementDetails `json:"details,omitempty"`
}

// NewOrganizationsManagement instantiates a new OrganizationsManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsManagement() *OrganizationsManagement {
	this := OrganizationsManagement{}
	return &this
}

// NewOrganizationsManagementWithDefaults instantiates a new OrganizationsManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsManagementWithDefaults() *OrganizationsManagement {
	this := OrganizationsManagement{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OrganizationsManagement) GetDetails() []OrganizationsManagementDetails {
	if o == nil || isNil(o.Details) {
		var ret []OrganizationsManagementDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsManagement) GetDetailsOk() ([]OrganizationsManagementDetails, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OrganizationsManagement) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []OrganizationsManagementDetails and assigns it to the Details field.
func (o *OrganizationsManagement) SetDetails(v []OrganizationsManagementDetails) {
	o.Details = v
}

func (o OrganizationsManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsManagement struct {
	value *OrganizationsManagement
	isSet bool
}

func (v NullableOrganizationsManagement) Get() *OrganizationsManagement {
	return v.value
}

func (v *NullableOrganizationsManagement) Set(val *OrganizationsManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsManagement(val *OrganizationsManagement) *NullableOrganizationsManagement {
	return &NullableOrganizationsManagement{value: val, isSet: true}
}

func (v NullableOrganizationsManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


