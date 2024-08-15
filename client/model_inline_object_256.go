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

// InlineObject256 struct for InlineObject256
type InlineObject256 struct {
	// The role of the SAML administrator
	Role string `json:"role"`
	// The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	OrgAccess string `json:"orgAccess"`
	// The list of tags that the SAML administrator has privileges on
	Tags []OrganizationsOrganizationIdSamlRolesTags1 `json:"tags,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []OrganizationsOrganizationIdSamlRolesNetworks1 `json:"networks,omitempty"`
}

// NewInlineObject256 instantiates a new InlineObject256 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject256(role string, orgAccess string) *InlineObject256 {
	this := InlineObject256{}
	this.Role = role
	this.OrgAccess = orgAccess
	return &this
}

// NewInlineObject256WithDefaults instantiates a new InlineObject256 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject256WithDefaults() *InlineObject256 {
	this := InlineObject256{}
	return &this
}

// GetRole returns the Role field value
func (o *InlineObject256) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InlineObject256) GetRoleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InlineObject256) SetRole(v string) {
	o.Role = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *InlineObject256) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *InlineObject256) GetOrgAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *InlineObject256) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject256) GetTags() []OrganizationsOrganizationIdSamlRolesTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdSamlRolesTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject256) GetTagsOk() ([]OrganizationsOrganizationIdSamlRolesTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject256) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdSamlRolesTags1 and assigns it to the Tags field.
func (o *InlineObject256) SetTags(v []OrganizationsOrganizationIdSamlRolesTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject256) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdSamlRolesNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject256) GetNetworksOk() ([]OrganizationsOrganizationIdSamlRolesNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject256) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdSamlRolesNetworks1 and assigns it to the Networks field.
func (o *InlineObject256) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks1) {
	o.Networks = v
}

func (o InlineObject256) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject256 struct {
	value *InlineObject256
	isSet bool
}

func (v NullableInlineObject256) Get() *InlineObject256 {
	return v.value
}

func (v *NullableInlineObject256) Set(val *InlineObject256) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject256) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject256) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject256(val *InlineObject256) *NullableInlineObject256 {
	return &NullableInlineObject256{value: val, isSet: true}
}

func (v NullableInlineObject256) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject256) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


