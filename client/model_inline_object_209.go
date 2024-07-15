/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject209 struct for InlineObject209
type InlineObject209 struct {
	// The name of the dashboard administrator
	Name *string `json:"name,omitempty"`
	// The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	OrgAccess *string `json:"orgAccess,omitempty"`
	// The list of tags that the dashboard administrator has privileges on
	Tags []OrganizationsOrganizationIdAdminsTags1 `json:"tags,omitempty"`
	// The list of networks that the dashboard administrator has privileges on
	Networks []OrganizationsOrganizationIdAdminsNetworks1 `json:"networks,omitempty"`
}

// NewInlineObject209 instantiates a new InlineObject209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject209() *InlineObject209 {
	this := InlineObject209{}
	return &this
}

// NewInlineObject209WithDefaults instantiates a new InlineObject209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject209WithDefaults() *InlineObject209 {
	this := InlineObject209{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject209) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject209) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject209) SetName(v string) {
	o.Name = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *InlineObject209) GetOrgAccess() string {
	if o == nil || isNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetOrgAccessOk() (*string, bool) {
	if o == nil || isNil(o.OrgAccess) {
    return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *InlineObject209) HasOrgAccess() bool {
	if o != nil && !isNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *InlineObject209) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject209) GetTags() []OrganizationsOrganizationIdAdminsTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdAdminsTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetTagsOk() ([]OrganizationsOrganizationIdAdminsTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject209) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdAdminsTags1 and assigns it to the Tags field.
func (o *InlineObject209) SetTags(v []OrganizationsOrganizationIdAdminsTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject209) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdAdminsNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetNetworksOk() ([]OrganizationsOrganizationIdAdminsNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject209) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdAdminsNetworks1 and assigns it to the Networks field.
func (o *InlineObject209) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks1) {
	o.Networks = v
}

func (o InlineObject209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
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

type NullableInlineObject209 struct {
	value *InlineObject209
	isSet bool
}

func (v NullableInlineObject209) Get() *InlineObject209 {
	return v.value
}

func (v *NullableInlineObject209) Set(val *InlineObject209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject209(val *InlineObject209) *NullableInlineObject209 {
	return &NullableInlineObject209{value: val, isSet: true}
}

func (v NullableInlineObject209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


