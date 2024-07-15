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

// InlineObject208 struct for InlineObject208
type InlineObject208 struct {
	// The email of the dashboard administrator. This attribute can not be updated.
	Email string `json:"email"`
	// The name of the dashboard administrator
	Name string `json:"name"`
	// The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	OrgAccess string `json:"orgAccess"`
	// The list of tags that the dashboard administrator has privileges on
	Tags []OrganizationsOrganizationIdAdminsTags1 `json:"tags,omitempty"`
	// The list of networks that the dashboard administrator has privileges on
	Networks []OrganizationsOrganizationIdAdminsNetworks1 `json:"networks,omitempty"`
	// The method of authentication the user will use to sign in to the Meraki dashboard. Can be one of 'Email' or 'Cisco SecureX Sign-On'. The default is Email authentication
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
}

// NewInlineObject208 instantiates a new InlineObject208 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject208(email string, name string, orgAccess string) *InlineObject208 {
	this := InlineObject208{}
	this.Email = email
	this.Name = name
	this.OrgAccess = orgAccess
	return &this
}

// NewInlineObject208WithDefaults instantiates a new InlineObject208 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject208WithDefaults() *InlineObject208 {
	this := InlineObject208{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineObject208) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineObject208) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineObject208) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *InlineObject208) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject208) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject208) SetName(v string) {
	o.Name = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *InlineObject208) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *InlineObject208) GetOrgAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *InlineObject208) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject208) GetTags() []OrganizationsOrganizationIdAdminsTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdAdminsTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject208) GetTagsOk() ([]OrganizationsOrganizationIdAdminsTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject208) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdAdminsTags1 and assigns it to the Tags field.
func (o *InlineObject208) SetTags(v []OrganizationsOrganizationIdAdminsTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject208) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdAdminsNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject208) GetNetworksOk() ([]OrganizationsOrganizationIdAdminsNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject208) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdAdminsNetworks1 and assigns it to the Networks field.
func (o *InlineObject208) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks1) {
	o.Networks = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *InlineObject208) GetAuthenticationMethod() string {
	if o == nil || isNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject208) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || isNil(o.AuthenticationMethod) {
    return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *InlineObject208) HasAuthenticationMethod() bool {
	if o != nil && !isNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *InlineObject208) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o InlineObject208) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["name"] = o.Name
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
	if !isNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject208 struct {
	value *InlineObject208
	isSet bool
}

func (v NullableInlineObject208) Get() *InlineObject208 {
	return v.value
}

func (v *NullableInlineObject208) Set(val *InlineObject208) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject208) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject208) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject208(val *InlineObject208) *NullableInlineObject208 {
	return &NullableInlineObject208{value: val, isSet: true}
}

func (v NullableInlineObject208) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject208) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


