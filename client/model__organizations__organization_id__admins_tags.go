/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAdminsTags struct for OrganizationsOrganizationIdAdminsTags
type OrganizationsOrganizationIdAdminsTags struct {
	// Tag value
	Tag *string `json:"tag,omitempty"`
	// Access level for the tag
	Access *string `json:"access,omitempty"`
}

// NewOrganizationsOrganizationIdAdminsTags instantiates a new OrganizationsOrganizationIdAdminsTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAdminsTags() *OrganizationsOrganizationIdAdminsTags {
	this := OrganizationsOrganizationIdAdminsTags{}
	return &this
}

// NewOrganizationsOrganizationIdAdminsTagsWithDefaults instantiates a new OrganizationsOrganizationIdAdminsTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAdminsTagsWithDefaults() *OrganizationsOrganizationIdAdminsTags {
	this := OrganizationsOrganizationIdAdminsTags{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdminsTags) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsTags) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdminsTags) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *OrganizationsOrganizationIdAdminsTags) SetTag(v string) {
	o.Tag = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAdminsTags) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAdminsTags) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAdminsTags) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *OrganizationsOrganizationIdAdminsTags) SetAccess(v string) {
	o.Access = &v
}

func (o OrganizationsOrganizationIdAdminsTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Access) {
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


