/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject101 struct for InlineObject101
type InlineObject101 struct {
	// Email address of the user
	Email string `json:"email"`
	// Name of the user. Only required If the user is not a Dashboard administrator.
	Name *string `json:"name,omitempty"`
	// The password for this user account. Only required If the user is not a Dashboard administrator.
	Password *string `json:"password,omitempty"`
	// Authorization type for user. Can be 'Guest' or '802.1X' for wireless networks, or 'Client VPN' for MX networks. Defaults to '802.1X'.
	AccountType *string `json:"accountType,omitempty"`
	// Whether or not Meraki should email the password to user. Default is false.
	EmailPasswordToUser *bool `json:"emailPasswordToUser,omitempty"`
	// Whether or not the user is a Dashboard administrator.
	IsAdmin *bool `json:"isAdmin,omitempty"`
	// Authorization zones and expiration dates for the user.
	Authorizations []NetworksNetworkIdMerakiAuthUsersAuthorizations1 `json:"authorizations"`
}

// NewInlineObject101 instantiates a new InlineObject101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject101(email string, authorizations []NetworksNetworkIdMerakiAuthUsersAuthorizations1) *InlineObject101 {
	this := InlineObject101{}
	this.Email = email
	var accountType string = "802.1X"
	this.AccountType = &accountType
	this.Authorizations = authorizations
	return &this
}

// NewInlineObject101WithDefaults instantiates a new InlineObject101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject101WithDefaults() *InlineObject101 {
	this := InlineObject101{}
	var accountType string = "802.1X"
	this.AccountType = &accountType
	return &this
}

// GetEmail returns the Email field value
func (o *InlineObject101) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineObject101) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject101) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject101) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject101) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InlineObject101) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineObject101) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineObject101) SetPassword(v string) {
	o.Password = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *InlineObject101) GetAccountType() string {
	if o == nil || isNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetAccountTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccountType) {
    return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *InlineObject101) HasAccountType() bool {
	if o != nil && !isNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *InlineObject101) SetAccountType(v string) {
	o.AccountType = &v
}

// GetEmailPasswordToUser returns the EmailPasswordToUser field value if set, zero value otherwise.
func (o *InlineObject101) GetEmailPasswordToUser() bool {
	if o == nil || isNil(o.EmailPasswordToUser) {
		var ret bool
		return ret
	}
	return *o.EmailPasswordToUser
}

// GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetEmailPasswordToUserOk() (*bool, bool) {
	if o == nil || isNil(o.EmailPasswordToUser) {
    return nil, false
	}
	return o.EmailPasswordToUser, true
}

// HasEmailPasswordToUser returns a boolean if a field has been set.
func (o *InlineObject101) HasEmailPasswordToUser() bool {
	if o != nil && !isNil(o.EmailPasswordToUser) {
		return true
	}

	return false
}

// SetEmailPasswordToUser gets a reference to the given bool and assigns it to the EmailPasswordToUser field.
func (o *InlineObject101) SetEmailPasswordToUser(v bool) {
	o.EmailPasswordToUser = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *InlineObject101) GetIsAdmin() bool {
	if o == nil || isNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetIsAdminOk() (*bool, bool) {
	if o == nil || isNil(o.IsAdmin) {
    return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *InlineObject101) HasIsAdmin() bool {
	if o != nil && !isNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *InlineObject101) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetAuthorizations returns the Authorizations field value
func (o *InlineObject101) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersAuthorizations1 {
	if o == nil {
		var ret []NetworksNetworkIdMerakiAuthUsersAuthorizations1
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetAuthorizationsOk() ([]NetworksNetworkIdMerakiAuthUsersAuthorizations1, bool) {
	if o == nil {
    return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *InlineObject101) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersAuthorizations1) {
	o.Authorizations = v
}

func (o InlineObject101) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !isNil(o.EmailPasswordToUser) {
		toSerialize["emailPasswordToUser"] = o.EmailPasswordToUser
	}
	if !isNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if true {
		toSerialize["authorizations"] = o.Authorizations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject101 struct {
	value *InlineObject101
	isSet bool
}

func (v NullableInlineObject101) Get() *InlineObject101 {
	return v.value
}

func (v *NullableInlineObject101) Set(val *InlineObject101) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject101) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject101) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject101(val *InlineObject101) *NullableInlineObject101 {
	return &NullableInlineObject101{value: val, isSet: true}
}

func (v NullableInlineObject101) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject101) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


