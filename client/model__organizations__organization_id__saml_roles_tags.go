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

// OrganizationsOrganizationIdSamlRolesTags struct for OrganizationsOrganizationIdSamlRolesTags
type OrganizationsOrganizationIdSamlRolesTags struct {
	// The name of the tag
	Tag *string `json:"tag,omitempty"`
	// The privilege of the SAML administrator on the tag
	Access *string `json:"access,omitempty"`
}

// NewOrganizationsOrganizationIdSamlRolesTags instantiates a new OrganizationsOrganizationIdSamlRolesTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSamlRolesTags() *OrganizationsOrganizationIdSamlRolesTags {
	this := OrganizationsOrganizationIdSamlRolesTags{}
	return &this
}

// NewOrganizationsOrganizationIdSamlRolesTagsWithDefaults instantiates a new OrganizationsOrganizationIdSamlRolesTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSamlRolesTagsWithDefaults() *OrganizationsOrganizationIdSamlRolesTags {
	this := OrganizationsOrganizationIdSamlRolesTags{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSamlRolesTags) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesTags) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSamlRolesTags) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *OrganizationsOrganizationIdSamlRolesTags) SetTag(v string) {
	o.Tag = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSamlRolesTags) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesTags) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSamlRolesTags) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *OrganizationsOrganizationIdSamlRolesTags) SetAccess(v string) {
	o.Access = &v
}

func (o OrganizationsOrganizationIdSamlRolesTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSamlRolesTags struct {
	value *OrganizationsOrganizationIdSamlRolesTags
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSamlRolesTags) Get() *OrganizationsOrganizationIdSamlRolesTags {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSamlRolesTags) Set(val *OrganizationsOrganizationIdSamlRolesTags) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSamlRolesTags) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSamlRolesTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSamlRolesTags(val *OrganizationsOrganizationIdSamlRolesTags) *NullableOrganizationsOrganizationIdSamlRolesTags {
	return &NullableOrganizationsOrganizationIdSamlRolesTags{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSamlRolesTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSamlRolesTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


