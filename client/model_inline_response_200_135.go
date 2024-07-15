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

// InlineResponse200135 struct for InlineResponse200135
type InlineResponse200135 struct {
	// The Meraki managed Id of the user record.
	Id *string `json:"id,omitempty"`
	// User email.
	Email *string `json:"email,omitempty"`
	// User full name.
	FullName *string `json:"fullName,omitempty"`
	// The users username.
	Username *string `json:"username,omitempty"`
	// A boolean denoting if the user has a password associated with the record.
	HasPassword *bool `json:"hasPassword,omitempty"`
	// The set of tags the user is scoped to.
	Tags *string `json:"tags,omitempty"`
	// Active Directory Groups the user belongs to.
	AdGroups []string `json:"adGroups,omitempty"`
	// Azure Active Directory Groups the user belongs to.
	AzureAdGroups []string `json:"azureAdGroups,omitempty"`
	// SAML Groups the user belongs to.
	SamlGroups []string `json:"samlGroups,omitempty"`
	// Apple School Manager Groups the user belongs to.
	AsmGroups []string `json:"asmGroups,omitempty"`
	// Whether the user was created using an external integration, or via the Meraki Dashboard.
	IsExternal *bool `json:"isExternal,omitempty"`
	// The user display name.
	DisplayName *string `json:"displayName,omitempty"`
	// A boolean indicating if the user has an associated identity certificate..
	HasIdentityCertificate *bool `json:"hasIdentityCertificate,omitempty"`
	// The url where the users thumbnail is hosted.
	UserThumbnail *string `json:"userThumbnail,omitempty"`
}

// NewInlineResponse200135 instantiates a new InlineResponse200135 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200135() *InlineResponse200135 {
	this := InlineResponse200135{}
	return &this
}

// NewInlineResponse200135WithDefaults instantiates a new InlineResponse200135 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200135WithDefaults() *InlineResponse200135 {
	this := InlineResponse200135{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200135) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200135) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200135) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200135) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200135) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse200135) SetEmail(v string) {
	o.Email = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *InlineResponse200135) GetFullName() string {
	if o == nil || isNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetFullNameOk() (*string, bool) {
	if o == nil || isNil(o.FullName) {
    return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *InlineResponse200135) HasFullName() bool {
	if o != nil && !isNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *InlineResponse200135) SetFullName(v string) {
	o.FullName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineResponse200135) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineResponse200135) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineResponse200135) SetUsername(v string) {
	o.Username = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *InlineResponse200135) GetHasPassword() bool {
	if o == nil || isNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetHasPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.HasPassword) {
    return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *InlineResponse200135) HasHasPassword() bool {
	if o != nil && !isNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *InlineResponse200135) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200135) GetTags() string {
	if o == nil || isNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetTagsOk() (*string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200135) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *InlineResponse200135) SetTags(v string) {
	o.Tags = &v
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *InlineResponse200135) GetAdGroups() []string {
	if o == nil || isNil(o.AdGroups) {
		var ret []string
		return ret
	}
	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetAdGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.AdGroups) {
    return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *InlineResponse200135) HasAdGroups() bool {
	if o != nil && !isNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given []string and assigns it to the AdGroups field.
func (o *InlineResponse200135) SetAdGroups(v []string) {
	o.AdGroups = v
}

// GetAzureAdGroups returns the AzureAdGroups field value if set, zero value otherwise.
func (o *InlineResponse200135) GetAzureAdGroups() []string {
	if o == nil || isNil(o.AzureAdGroups) {
		var ret []string
		return ret
	}
	return o.AzureAdGroups
}

// GetAzureAdGroupsOk returns a tuple with the AzureAdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetAzureAdGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.AzureAdGroups) {
    return nil, false
	}
	return o.AzureAdGroups, true
}

// HasAzureAdGroups returns a boolean if a field has been set.
func (o *InlineResponse200135) HasAzureAdGroups() bool {
	if o != nil && !isNil(o.AzureAdGroups) {
		return true
	}

	return false
}

// SetAzureAdGroups gets a reference to the given []string and assigns it to the AzureAdGroups field.
func (o *InlineResponse200135) SetAzureAdGroups(v []string) {
	o.AzureAdGroups = v
}

// GetSamlGroups returns the SamlGroups field value if set, zero value otherwise.
func (o *InlineResponse200135) GetSamlGroups() []string {
	if o == nil || isNil(o.SamlGroups) {
		var ret []string
		return ret
	}
	return o.SamlGroups
}

// GetSamlGroupsOk returns a tuple with the SamlGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetSamlGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.SamlGroups) {
    return nil, false
	}
	return o.SamlGroups, true
}

// HasSamlGroups returns a boolean if a field has been set.
func (o *InlineResponse200135) HasSamlGroups() bool {
	if o != nil && !isNil(o.SamlGroups) {
		return true
	}

	return false
}

// SetSamlGroups gets a reference to the given []string and assigns it to the SamlGroups field.
func (o *InlineResponse200135) SetSamlGroups(v []string) {
	o.SamlGroups = v
}

// GetAsmGroups returns the AsmGroups field value if set, zero value otherwise.
func (o *InlineResponse200135) GetAsmGroups() []string {
	if o == nil || isNil(o.AsmGroups) {
		var ret []string
		return ret
	}
	return o.AsmGroups
}

// GetAsmGroupsOk returns a tuple with the AsmGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetAsmGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.AsmGroups) {
    return nil, false
	}
	return o.AsmGroups, true
}

// HasAsmGroups returns a boolean if a field has been set.
func (o *InlineResponse200135) HasAsmGroups() bool {
	if o != nil && !isNil(o.AsmGroups) {
		return true
	}

	return false
}

// SetAsmGroups gets a reference to the given []string and assigns it to the AsmGroups field.
func (o *InlineResponse200135) SetAsmGroups(v []string) {
	o.AsmGroups = v
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise.
func (o *InlineResponse200135) GetIsExternal() bool {
	if o == nil || isNil(o.IsExternal) {
		var ret bool
		return ret
	}
	return *o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetIsExternalOk() (*bool, bool) {
	if o == nil || isNil(o.IsExternal) {
    return nil, false
	}
	return o.IsExternal, true
}

// HasIsExternal returns a boolean if a field has been set.
func (o *InlineResponse200135) HasIsExternal() bool {
	if o != nil && !isNil(o.IsExternal) {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given bool and assigns it to the IsExternal field.
func (o *InlineResponse200135) SetIsExternal(v bool) {
	o.IsExternal = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InlineResponse200135) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InlineResponse200135) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InlineResponse200135) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHasIdentityCertificate returns the HasIdentityCertificate field value if set, zero value otherwise.
func (o *InlineResponse200135) GetHasIdentityCertificate() bool {
	if o == nil || isNil(o.HasIdentityCertificate) {
		var ret bool
		return ret
	}
	return *o.HasIdentityCertificate
}

// GetHasIdentityCertificateOk returns a tuple with the HasIdentityCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetHasIdentityCertificateOk() (*bool, bool) {
	if o == nil || isNil(o.HasIdentityCertificate) {
    return nil, false
	}
	return o.HasIdentityCertificate, true
}

// HasHasIdentityCertificate returns a boolean if a field has been set.
func (o *InlineResponse200135) HasHasIdentityCertificate() bool {
	if o != nil && !isNil(o.HasIdentityCertificate) {
		return true
	}

	return false
}

// SetHasIdentityCertificate gets a reference to the given bool and assigns it to the HasIdentityCertificate field.
func (o *InlineResponse200135) SetHasIdentityCertificate(v bool) {
	o.HasIdentityCertificate = &v
}

// GetUserThumbnail returns the UserThumbnail field value if set, zero value otherwise.
func (o *InlineResponse200135) GetUserThumbnail() string {
	if o == nil || isNil(o.UserThumbnail) {
		var ret string
		return ret
	}
	return *o.UserThumbnail
}

// GetUserThumbnailOk returns a tuple with the UserThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200135) GetUserThumbnailOk() (*string, bool) {
	if o == nil || isNil(o.UserThumbnail) {
    return nil, false
	}
	return o.UserThumbnail, true
}

// HasUserThumbnail returns a boolean if a field has been set.
func (o *InlineResponse200135) HasUserThumbnail() bool {
	if o != nil && !isNil(o.UserThumbnail) {
		return true
	}

	return false
}

// SetUserThumbnail gets a reference to the given string and assigns it to the UserThumbnail field.
func (o *InlineResponse200135) SetUserThumbnail(v string) {
	o.UserThumbnail = &v
}

func (o InlineResponse200135) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	if !isNil(o.AzureAdGroups) {
		toSerialize["azureAdGroups"] = o.AzureAdGroups
	}
	if !isNil(o.SamlGroups) {
		toSerialize["samlGroups"] = o.SamlGroups
	}
	if !isNil(o.AsmGroups) {
		toSerialize["asmGroups"] = o.AsmGroups
	}
	if !isNil(o.IsExternal) {
		toSerialize["isExternal"] = o.IsExternal
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.HasIdentityCertificate) {
		toSerialize["hasIdentityCertificate"] = o.HasIdentityCertificate
	}
	if !isNil(o.UserThumbnail) {
		toSerialize["userThumbnail"] = o.UserThumbnail
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200135 struct {
	value *InlineResponse200135
	isSet bool
}

func (v NullableInlineResponse200135) Get() *InlineResponse200135 {
	return v.value
}

func (v *NullableInlineResponse200135) Set(val *InlineResponse200135) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200135) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200135) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200135(val *InlineResponse200135) *NullableInlineResponse200135 {
	return &NullableInlineResponse200135{value: val, isSet: true}
}

func (v NullableInlineResponse200135) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200135) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


