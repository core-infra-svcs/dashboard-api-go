/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject224 struct for InlineObject224
type InlineObject224 struct {
	// The role of the SAML administrator
	Role *string `json:"role,omitempty"`
	// The privilege of the SAML administrator on the organization. Can be one of 'none', 'read-only', 'full' or 'enterprise'
	OrgAccess *string `json:"orgAccess,omitempty"`
	// The list of tags that the SAML administrator has privileges on
	Tags []OrganizationsOrganizationIdSamlRolesTags1 `json:"tags,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []OrganizationsOrganizationIdSamlRolesNetworks1 `json:"networks,omitempty"`
}

// NewInlineObject224 instantiates a new InlineObject224 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject224() *InlineObject224 {
	this := InlineObject224{}
	return &this
}

// NewInlineObject224WithDefaults instantiates a new InlineObject224 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject224WithDefaults() *InlineObject224 {
	this := InlineObject224{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InlineObject224) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InlineObject224) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *InlineObject224) SetRole(v string) {
	o.Role = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *InlineObject224) GetOrgAccess() string {
	if o == nil || isNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetOrgAccessOk() (*string, bool) {
	if o == nil || isNil(o.OrgAccess) {
    return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *InlineObject224) HasOrgAccess() bool {
	if o != nil && !isNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *InlineObject224) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject224) GetTags() []OrganizationsOrganizationIdSamlRolesTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdSamlRolesTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetTagsOk() ([]OrganizationsOrganizationIdSamlRolesTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject224) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdSamlRolesTags1 and assigns it to the Tags field.
func (o *InlineObject224) SetTags(v []OrganizationsOrganizationIdSamlRolesTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject224) GetNetworks() []OrganizationsOrganizationIdSamlRolesNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdSamlRolesNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject224) GetNetworksOk() ([]OrganizationsOrganizationIdSamlRolesNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject224) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdSamlRolesNetworks1 and assigns it to the Networks field.
func (o *InlineObject224) SetNetworks(v []OrganizationsOrganizationIdSamlRolesNetworks1) {
	o.Networks = v
}

func (o InlineObject224) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject224 struct {
	value *InlineObject224
	isSet bool
}

func (v NullableInlineObject224) Get() *InlineObject224 {
	return v.value
}

func (v *NullableInlineObject224) Set(val *InlineObject224) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject224) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject224) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject224(val *InlineObject224) *NullableInlineObject224 {
	return &NullableInlineObject224{value: val, isSet: true}
}

func (v NullableInlineObject224) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject224) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


