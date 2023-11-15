/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject237 struct for InlineObject237
type InlineObject237 struct {
	// The role of the SAML administrator
	Role *string `json:"role,omitempty"`
	// The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	OrgAccess *string `json:"orgAccess,omitempty"`
	// The list of tags that the SAML administrator has privileges on
	Tags []OrganizationsOrganizationIdSamlRolesTags1 `json:"tags,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []OrganizationsOrganizationIdSamlRolesNetworks1 `json:"networks,omitempty"`
}

// NewInlineObject237 instantiates a new InlineObject237 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject237() *InlineObject237 {
	this := InlineObject237{}
	return &this
}

// NewInlineObject237WithDefaults instantiates a new InlineObject237 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject237WithDefaults() *InlineObject237 {
	this := InlineObject237{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InlineObject237) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject237) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InlineObject237) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *InlineObject237) SetRole(v string) {
	o.Role = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *InlineObject237) GetOrgAccess() string {
	if o == nil || isNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject237) GetOrgAccessOk() (*string, bool) {
	if o == nil || isNil(o.OrgAccess) {
    return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *InlineObject237) HasOrgAccess() bool {
	if o != nil && !isNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *InlineObject237) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject237) GetTags() []OrganizationsOrganizationIdSamlRolesTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdSamlRolesTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject237) GetTagsOk() ([]OrganizationsOrganizationIdSamlRolesTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject237) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdSamlRolesTags1 and assigns it to the Tags field.
func (o *InlineObject237) SetTags(v []OrganizationsOrganizationIdSamlRolesTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject237) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdSamlRolesNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject237) GetNetworksOk() ([]OrganizationsOrganizationIdSamlRolesNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject237) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdSamlRolesNetworks1 and assigns it to the Networks field.
func (o *InlineObject237) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks1) {
	o.Networks = v
}

func (o InlineObject237) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.OrgAccess) {
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

type NullableInlineObject237 struct {
	value *InlineObject237
	isSet bool
}

func (v NullableInlineObject237) Get() *InlineObject237 {
	return v.value
}

func (v *NullableInlineObject237) Set(val *InlineObject237) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject237) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject237) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject237(val *InlineObject237) *NullableInlineObject237 {
	return &NullableInlineObject237{value: val, isSet: true}
}

func (v NullableInlineObject237) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject237) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


