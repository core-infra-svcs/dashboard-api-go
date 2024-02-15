/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdminsTags1 struct for OrganizationsOrganizationIdAdminsTags1
type OrganizationsOrganizationIdAdminsTags1 struct {
	// The name of the tag
	Tag string `json:"tag"`
	// The privilege of the dashboard administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

// NewOrganizationsOrganizationIdAdminsTags1 instantiates a new OrganizationsOrganizationIdAdminsTags1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdminsTags1(tag string, access string) *OrganizationsOrganizationIdAdminsTags1 {
	this := OrganizationsOrganizationIdAdminsTags1{}
	this.Tag = tag
	this.Access = access
	return &this
}

// NewOrganizationsOrganizationIdAdminsTags1WithDefaults instantiates a new OrganizationsOrganizationIdAdminsTags1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdminsTags1WithDefaults() *OrganizationsOrganizationIdAdminsTags1 {
	this := OrganizationsOrganizationIdAdminsTags1{}
	return &this
}

// GetTag returns the Tag field value
func (o *OrganizationsOrganizationIdAdminsTags1) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsTags1) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *OrganizationsOrganizationIdAdminsTags1) SetTag(v string) {
	o.Tag = v
}

// GetAccess returns the Access field value
func (o *OrganizationsOrganizationIdAdminsTags1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsTags1) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *OrganizationsOrganizationIdAdminsTags1) SetAccess(v string) {
	o.Access = v
}

func (o OrganizationsOrganizationIdAdminsTags1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdminsTags1 struct {
	value *OrganizationsOrganizationIdAdminsTags1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdminsTags1) Get() *OrganizationsOrganizationIdAdminsTags1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdminsTags1) Set(val *OrganizationsOrganizationIdAdminsTags1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdminsTags1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdminsTags1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdminsTags1(val *OrganizationsOrganizationIdAdminsTags1) *NullableOrganizationsOrganizationIdAdminsTags1 {
	return &NullableOrganizationsOrganizationIdAdminsTags1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdminsTags1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdminsTags1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


