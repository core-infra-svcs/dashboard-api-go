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

// OrganizationsOrganizationIdSamlRolesTags1 struct for OrganizationsOrganizationIdSamlRolesTags1
type OrganizationsOrganizationIdSamlRolesTags1 struct {
	// The name of the tag
	Tag string `json:"tag"`
	// The privilege of the SAML administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

// NewOrganizationsOrganizationIdSamlRolesTags1 instantiates a new OrganizationsOrganizationIdSamlRolesTags1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSamlRolesTags1(tag string, access string) *OrganizationsOrganizationIdSamlRolesTags1 {
	this := OrganizationsOrganizationIdSamlRolesTags1{}
	this.Tag = tag
	this.Access = access
	return &this
}

// NewOrganizationsOrganizationIdSamlRolesTags1WithDefaults instantiates a new OrganizationsOrganizationIdSamlRolesTags1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSamlRolesTags1WithDefaults() *OrganizationsOrganizationIdSamlRolesTags1 {
	this := OrganizationsOrganizationIdSamlRolesTags1{}
	return &this
}

// GetTag returns the Tag field value
func (o *OrganizationsOrganizationIdSamlRolesTags1) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesTags1) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *OrganizationsOrganizationIdSamlRolesTags1) SetTag(v string) {
	o.Tag = v
}

// GetAccess returns the Access field value
func (o *OrganizationsOrganizationIdSamlRolesTags1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSamlRolesTags1) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *OrganizationsOrganizationIdSamlRolesTags1) SetAccess(v string) {
	o.Access = v
}

func (o OrganizationsOrganizationIdSamlRolesTags1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSamlRolesTags1 struct {
	value *OrganizationsOrganizationIdSamlRolesTags1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSamlRolesTags1) Get() *OrganizationsOrganizationIdSamlRolesTags1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSamlRolesTags1) Set(val *OrganizationsOrganizationIdSamlRolesTags1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSamlRolesTags1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSamlRolesTags1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSamlRolesTags1(val *OrganizationsOrganizationIdSamlRolesTags1) *NullableOrganizationsOrganizationIdSamlRolesTags1 {
	return &NullableOrganizationsOrganizationIdSamlRolesTags1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSamlRolesTags1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSamlRolesTags1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


