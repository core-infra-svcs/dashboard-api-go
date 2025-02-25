/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject213 struct for InlineObject213
type InlineObject213 struct {
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
	// No longer used as of Cisco SecureX end-of-life. Can be one of 'Email'. The default is Email authentication.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
}

// NewInlineObject213 instantiates a new InlineObject213 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject213(email string, name string, orgAccess string) *InlineObject213 {
	this := InlineObject213{}
	this.Email = email
	this.Name = name
	this.OrgAccess = orgAccess
	return &this
}

// NewInlineObject213WithDefaults instantiates a new InlineObject213 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject213WithDefaults() *InlineObject213 {
	this := InlineObject213{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineObject213) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineObject213) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineObject213) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *InlineObject213) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject213) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject213) SetName(v string) {
	o.Name = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *InlineObject213) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *InlineObject213) GetOrgAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *InlineObject213) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject213) GetTags() []OrganizationsOrganizationIdAdminsTags1 {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdAdminsTags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject213) GetTagsOk() ([]OrganizationsOrganizationIdAdminsTags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject213) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdAdminsTags1 and assigns it to the Tags field.
func (o *InlineObject213) SetTags(v []OrganizationsOrganizationIdAdminsTags1) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineObject213) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks1 {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdAdminsNetworks1
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject213) GetNetworksOk() ([]OrganizationsOrganizationIdAdminsNetworks1, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineObject213) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdAdminsNetworks1 and assigns it to the Networks field.
func (o *InlineObject213) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks1) {
	o.Networks = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *InlineObject213) GetAuthenticationMethod() string {
	if o == nil || isNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject213) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || isNil(o.AuthenticationMethod) {
    return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *InlineObject213) HasAuthenticationMethod() bool {
	if o != nil && !isNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *InlineObject213) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o InlineObject213) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject213 struct {
	value *InlineObject213
	isSet bool
}

func (v NullableInlineObject213) Get() *InlineObject213 {
	return v.value
}

func (v *NullableInlineObject213) Set(val *InlineObject213) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject213) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject213) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject213(val *InlineObject213) *NullableInlineObject213 {
	return &NullableInlineObject213{value: val, isSet: true}
}

func (v NullableInlineObject213) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject213) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


