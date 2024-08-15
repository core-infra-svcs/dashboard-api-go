/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraRolesAppliedOrgWide struct for OrganizationsOrganizationIdCameraRolesAppliedOrgWide
type OrganizationsOrganizationIdCameraRolesAppliedOrgWide struct {
	// Permission scope id
	PermissionScopeId string `json:"permissionScopeId"`
}

// NewOrganizationsOrganizationIdCameraRolesAppliedOrgWide instantiates a new OrganizationsOrganizationIdCameraRolesAppliedOrgWide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraRolesAppliedOrgWide(permissionScopeId string) *OrganizationsOrganizationIdCameraRolesAppliedOrgWide {
	this := OrganizationsOrganizationIdCameraRolesAppliedOrgWide{}
	this.PermissionScopeId = permissionScopeId
	return &this
}

// NewOrganizationsOrganizationIdCameraRolesAppliedOrgWideWithDefaults instantiates a new OrganizationsOrganizationIdCameraRolesAppliedOrgWide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraRolesAppliedOrgWideWithDefaults() *OrganizationsOrganizationIdCameraRolesAppliedOrgWide {
	this := OrganizationsOrganizationIdCameraRolesAppliedOrgWide{}
	return &this
}

// GetPermissionScopeId returns the PermissionScopeId field value
func (o *OrganizationsOrganizationIdCameraRolesAppliedOrgWide) GetPermissionScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionScopeId
}

// GetPermissionScopeIdOk returns a tuple with the PermissionScopeId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraRolesAppliedOrgWide) GetPermissionScopeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PermissionScopeId, true
}

// SetPermissionScopeId sets field value
func (o *OrganizationsOrganizationIdCameraRolesAppliedOrgWide) SetPermissionScopeId(v string) {
	o.PermissionScopeId = v
}

func (o OrganizationsOrganizationIdCameraRolesAppliedOrgWide) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionScopeId"] = o.PermissionScopeId
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide struct {
	value *OrganizationsOrganizationIdCameraRolesAppliedOrgWide
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide) Get() *OrganizationsOrganizationIdCameraRolesAppliedOrgWide {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide) Set(val *OrganizationsOrganizationIdCameraRolesAppliedOrgWide) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide(val *OrganizationsOrganizationIdCameraRolesAppliedOrgWide) *NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide {
	return &NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraRolesAppliedOrgWide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


