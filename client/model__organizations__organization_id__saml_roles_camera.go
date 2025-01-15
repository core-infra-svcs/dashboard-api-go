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

// OrganizationsOrganizationIdSamlRolesCamera struct for OrganizationsOrganizationIdSamlRolesCamera
type OrganizationsOrganizationIdSamlRolesCamera struct {
	// Whether or not SAML administrator has org-wide access
	OrgWide *bool `json:"orgWide,omitempty"`
	// Camera access ability
	Access *string `json:"access,omitempty"`
}

// NewOrganizationsOrganizationIdSamlRolesCamera instantiates a new OrganizationsOrganizationIdSamlRolesCamera object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSamlRolesCamera() *OrganizationsOrganizationIdSamlRolesCamera {
	this := OrganizationsOrganizationIdSamlRolesCamera{}
	return &this
}

// NewOrganizationsOrganizationIdSamlRolesCameraWithDefaults instantiates a new OrganizationsOrganizationIdSamlRolesCamera object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSamlRolesCameraWithDefaults() *OrganizationsOrganizationIdSamlRolesCamera {
	this := OrganizationsOrganizationIdSamlRolesCamera{}
	return &this
}

// GetOrgWide returns the OrgWide field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSamlRolesCamera) GetOrgWide() bool {
	if o == nil || isNil(o.OrgWide) {
		var ret bool
		return ret
	}
	return *o.OrgWide
}

// GetOrgWideOk returns a tuple with the OrgWide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesCamera) GetOrgWideOk() (*bool, bool) {
	if o == nil || isNil(o.OrgWide) {
    return nil, false
	}
	return o.OrgWide, true
}

// HasOrgWide returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSamlRolesCamera) HasOrgWide() bool {
	if o != nil && !isNil(o.OrgWide) {
		return true
	}

	return false
}

// SetOrgWide gets a reference to the given bool and assigns it to the OrgWide field.
func (o *OrganizationsOrganizationIdSamlRolesCamera) SetOrgWide(v bool) {
	o.OrgWide = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSamlRolesCamera) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesCamera) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSamlRolesCamera) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *OrganizationsOrganizationIdSamlRolesCamera) SetAccess(v string) {
	o.Access = &v
}

func (o OrganizationsOrganizationIdSamlRolesCamera) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrgWide) {
		toSerialize["orgWide"] = o.OrgWide
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSamlRolesCamera struct {
	value *OrganizationsOrganizationIdSamlRolesCamera
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSamlRolesCamera) Get() *OrganizationsOrganizationIdSamlRolesCamera {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSamlRolesCamera) Set(val *OrganizationsOrganizationIdSamlRolesCamera) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSamlRolesCamera) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSamlRolesCamera) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSamlRolesCamera(val *OrganizationsOrganizationIdSamlRolesCamera) *NullableOrganizationsOrganizationIdSamlRolesCamera {
	return &NullableOrganizationsOrganizationIdSamlRolesCamera{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSamlRolesCamera) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSamlRolesCamera) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


