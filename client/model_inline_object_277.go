/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject277 struct for InlineObject277
type InlineObject277 struct {
	// The role of the SAML administrator
	Role string `json:"role"`
	// The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	OrgAccess string `json:"orgAccess"`
	// The list of tags that the SAML administrator has privileges on
	Tags []OrganizationsOrganizationIdSamlRolesTags1 `json:"tags,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []OrganizationsOrganizationIdSamlRolesNetworks1 `json:"networks,omitempty"`
}

// NewInlineObject277 instantiates a new InlineObject277 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject277(role string, orgAccess string) *InlineObject277 {
	this := InlineObject277{}
	this.Role = role
	this.OrgAccess = orgAccess
	return &this
}

// NewInlineObject277WithDefaults instantiates a new InlineObject277 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject277WithDefaults() *InlineObject277 {
	this := InlineObject277{}
	return &this
}

// GetRole returns the Role field value
func (o *InlineObject277) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InlineObject277) GetRoleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InlineObject277) SetRole(v string) {
	o.Role = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *InlineObject277) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *InlineObject277) GetOrgAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *InlineObject277) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject277) GetTags() []OrganizationsOrganizationIdSamlRolesTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdSamlRolesTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject277) GetTagsOk() ([]OrganizationsOrganizationIdSamlRolesTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject277) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdSamlRolesTags1 and assigns it to the Tags field.
func (o *InlineObject277) SetTags(v []OrganizationsOrganizationIdSamlRolesTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject277) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdSamlRolesNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject277) GetNetworksOk() ([]OrganizationsOrganizationIdSamlRolesNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject277) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdSamlRolesNetworks1 and assigns it to the Networks field.
func (o *InlineObject277) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks1) {
	o.Networks = v
}

func (o InlineObject277) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["orgAccess"] = o.OrgAccess
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject277 struct {
	value *InlineObject277
	isSet bool
}

func (v NullableInlineObject277) Get() *InlineObject277 {
	return v.value
}

func (v *NullableInlineObject277) Set(val *InlineObject277) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject277) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject277) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject277(val *InlineObject277) *NullableInlineObject277 {
	return &NullableInlineObject277{value: val, isSet: true}
}

func (v NullableInlineObject277) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject277) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


