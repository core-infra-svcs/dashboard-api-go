/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks struct for OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks
type OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks struct {
	// Network tag
	Tag *string `json:"tag,omitempty"`
	// Network id
	Id *string `json:"id,omitempty"`
	// Permission scope id
	PermissionScopeId string `json:"permissionScopeId"`
}

// NewOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks instantiates a new OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks(permissionScopeId string) *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks {
	this := OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks{}
	this.PermissionScopeId = permissionScopeId
	return &this
}

// NewOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworksWithDefaults instantiates a new OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworksWithDefaults() *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks {
	this := OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) SetTag(v string) {
	o.Tag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) SetId(v string) {
	o.Id = &v
}

// GetPermissionScopeId returns the PermissionScopeId field value
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) GetPermissionScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionScopeId
}

// GetPermissionScopeIdOk returns a tuple with the PermissionScopeId field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) GetPermissionScopeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PermissionScopeId, true
}

// SetPermissionScopeId sets field value
func (o *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) SetPermissionScopeId(v string) {
	o.PermissionScopeId = v
}

func (o OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["permissionScopeId"] = o.PermissionScopeId
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks struct {
	value *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) Get() *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) Set(val *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks(val *OrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) *NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks {
	return &NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraRolesRoleIdAppliedOnNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


