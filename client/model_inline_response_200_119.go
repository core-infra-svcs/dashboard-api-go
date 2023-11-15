/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200119 struct for InlineResponse200119
type InlineResponse200119 struct {
	// Admin's ID
	Id *string `json:"id,omitempty"`
	// Admin's username
	Name *string `json:"name,omitempty"`
	// Admin's email address
	Email *string `json:"email,omitempty"`
	// Admin's level of access to the organization
	OrgAccess *string `json:"orgAccess,omitempty"`
	// Status of the admin's account
	AccountStatus *string `json:"accountStatus,omitempty"`
	// Indicates whether two-factor authentication is enabled
	TwoFactorAuthEnabled *bool `json:"twoFactorAuthEnabled,omitempty"`
	// Indicates whether the admin has an API key
	HasApiKey *bool `json:"hasApiKey,omitempty"`
	// Time when the admin was last active
	LastActive *time.Time `json:"lastActive,omitempty"`
	// Admin tag information
	Tags []OrganizationsOrganizationIdAdminsTags `json:"tags,omitempty"`
	// Admin network access information
	Networks []OrganizationsOrganizationIdAdminsNetworks `json:"networks,omitempty"`
	// Admin's authentication method
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
}

// NewInlineResponse200119 instantiates a new InlineResponse200119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200119() *InlineResponse200119 {
	this := InlineResponse200119{}
	return &this
}

// NewInlineResponse200119WithDefaults instantiates a new InlineResponse200119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200119WithDefaults() *InlineResponse200119 {
	this := InlineResponse200119{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200119) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200119) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200119) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200119) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200119) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200119) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200119) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200119) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse200119) SetEmail(v string) {
	o.Email = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *InlineResponse200119) GetOrgAccess() string {
	if o == nil || isNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetOrgAccessOk() (*string, bool) {
	if o == nil || isNil(o.OrgAccess) {
    return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *InlineResponse200119) HasOrgAccess() bool {
	if o != nil && !isNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *InlineResponse200119) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *InlineResponse200119) GetAccountStatus() string {
	if o == nil || isNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetAccountStatusOk() (*string, bool) {
	if o == nil || isNil(o.AccountStatus) {
    return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *InlineResponse200119) HasAccountStatus() bool {
	if o != nil && !isNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *InlineResponse200119) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field value if set, zero value otherwise.
func (o *InlineResponse200119) GetTwoFactorAuthEnabled() bool {
	if o == nil || isNil(o.TwoFactorAuthEnabled) {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthEnabled
}

// GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetTwoFactorAuthEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.TwoFactorAuthEnabled) {
    return nil, false
	}
	return o.TwoFactorAuthEnabled, true
}

// HasTwoFactorAuthEnabled returns a boolean if a field has been set.
func (o *InlineResponse200119) HasTwoFactorAuthEnabled() bool {
	if o != nil && !isNil(o.TwoFactorAuthEnabled) {
		return true
	}

	return false
}

// SetTwoFactorAuthEnabled gets a reference to the given bool and assigns it to the TwoFactorAuthEnabled field.
func (o *InlineResponse200119) SetTwoFactorAuthEnabled(v bool) {
	o.TwoFactorAuthEnabled = &v
}

// GetHasApiKey returns the HasApiKey field value if set, zero value otherwise.
func (o *InlineResponse200119) GetHasApiKey() bool {
	if o == nil || isNil(o.HasApiKey) {
		var ret bool
		return ret
	}
	return *o.HasApiKey
}

// GetHasApiKeyOk returns a tuple with the HasApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetHasApiKeyOk() (*bool, bool) {
	if o == nil || isNil(o.HasApiKey) {
    return nil, false
	}
	return o.HasApiKey, true
}

// HasHasApiKey returns a boolean if a field has been set.
func (o *InlineResponse200119) HasHasApiKey() bool {
	if o != nil && !isNil(o.HasApiKey) {
		return true
	}

	return false
}

// SetHasApiKey gets a reference to the given bool and assigns it to the HasApiKey field.
func (o *InlineResponse200119) SetHasApiKey(v bool) {
	o.HasApiKey = &v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise.
func (o *InlineResponse200119) GetLastActive() time.Time {
	if o == nil || isNil(o.LastActive) {
		var ret time.Time
		return ret
	}
	return *o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetLastActiveOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastActive) {
    return nil, false
	}
	return o.LastActive, true
}

// HasLastActive returns a boolean if a field has been set.
func (o *InlineResponse200119) HasLastActive() bool {
	if o != nil && !isNil(o.LastActive) {
		return true
	}

	return false
}

// SetLastActive gets a reference to the given time.Time and assigns it to the LastActive field.
func (o *InlineResponse200119) SetLastActive(v time.Time) {
	o.LastActive = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200119) GetTags() []OrganizationsOrganizationIdAdminsTags {
	if o == nil || isNil(o.Tags) {
		var ret []OrganizationsOrganizationIdAdminsTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetTagsOk() ([]OrganizationsOrganizationIdAdminsTags, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200119) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []OrganizationsOrganizationIdAdminsTags and assigns it to the Tags field.
func (o *InlineResponse200119) SetTags(v []OrganizationsOrganizationIdAdminsTags) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineResponse200119) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks {
	if o == nil || isNil(o.Networks) {
		var ret []OrganizationsOrganizationIdAdminsNetworks
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetNetworksOk() ([]OrganizationsOrganizationIdAdminsNetworks, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineResponse200119) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []OrganizationsOrganizationIdAdminsNetworks and assigns it to the Networks field.
func (o *InlineResponse200119) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks) {
	o.Networks = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *InlineResponse200119) GetAuthenticationMethod() string {
	if o == nil || isNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200119) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || isNil(o.AuthenticationMethod) {
    return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *InlineResponse200119) HasAuthenticationMethod() bool {
	if o != nil && !isNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *InlineResponse200119) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o InlineResponse200119) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.OrgAccess) {
		toSerialize["orgAccess"] = o.OrgAccess
	}
	if !isNil(o.AccountStatus) {
		toSerialize["accountStatus"] = o.AccountStatus
	}
	if !isNil(o.TwoFactorAuthEnabled) {
		toSerialize["twoFactorAuthEnabled"] = o.TwoFactorAuthEnabled
	}
	if !isNil(o.HasApiKey) {
		toSerialize["hasApiKey"] = o.HasApiKey
	}
	if !isNil(o.LastActive) {
		toSerialize["lastActive"] = o.LastActive
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

type NullableInlineResponse200119 struct {
	value *InlineResponse200119
	isSet bool
}

func (v NullableInlineResponse200119) Get() *InlineResponse200119 {
	return v.value
}

func (v *NullableInlineResponse200119) Set(val *InlineResponse200119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200119(val *InlineResponse200119) *NullableInlineResponse200119 {
	return &NullableInlineResponse200119{value: val, isSet: true}
}

func (v NullableInlineResponse200119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


