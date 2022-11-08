/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdminsTags struct for OrganizationsOrganizationIdAdminsTags
type OrganizationsOrganizationIdAdminsTags struct {
	// The name of the tag
	Tag string `json:"tag"`
	// The privilege of the dashboard administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

// NewOrganizationsOrganizationIdAdminsTags instantiates a new OrganizationsOrganizationIdAdminsTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdminsTags(tag string, access string) *OrganizationsOrganizationIdAdminsTags {
	this := OrganizationsOrganizationIdAdminsTags{}
	this.Tag = tag
	this.Access = access
	return &this
}

// NewOrganizationsOrganizationIdAdminsTagsWithDefaults instantiates a new OrganizationsOrganizationIdAdminsTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdminsTagsWithDefaults() *OrganizationsOrganizationIdAdminsTags {
	this := OrganizationsOrganizationIdAdminsTags{}
	return &this
}

// GetTag returns the Tag field value
func (o *OrganizationsOrganizationIdAdminsTags) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsTags) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *OrganizationsOrganizationIdAdminsTags) SetTag(v string) {
	o.Tag = v
}

// GetAccess returns the Access field value
func (o *OrganizationsOrganizationIdAdminsTags) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsTags) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *OrganizationsOrganizationIdAdminsTags) SetAccess(v string) {
	o.Access = v
}

func (o OrganizationsOrganizationIdAdminsTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAdminsTags struct {
	value *OrganizationsOrganizationIdAdminsTags
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAdminsTags) Get() *OrganizationsOrganizationIdAdminsTags {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAdminsTags) Set(val *OrganizationsOrganizationIdAdminsTags) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAdminsTags) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAdminsTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAdminsTags(val *OrganizationsOrganizationIdAdminsTags) *NullableOrganizationsOrganizationIdAdminsTags {
	return &NullableOrganizationsOrganizationIdAdminsTags{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAdminsTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAdminsTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


