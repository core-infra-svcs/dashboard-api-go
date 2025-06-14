/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile The profile the DNS record is associated with
type OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile struct {
	// Profile ID
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile instantiates a new OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile() *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile {
	this := OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceDnsLocalRecordsProfileWithDefaults instantiates a new OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceDnsLocalRecordsProfileWithDefaults() *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile {
	this := OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile struct {
	value *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) Get() *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) Set(val *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile(val *OrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) *NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile {
	return &NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


