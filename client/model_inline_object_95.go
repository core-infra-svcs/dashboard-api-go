/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject95 struct for InlineObject95
type InlineObject95 struct {
	// Name of the user. Only allowed If the user is not Dashboard administrator.
	Name *string `json:"name,omitempty"`
	// The password for this user account. Only allowed If the user is not Dashboard administrator.
	Password *string `json:"password,omitempty"`
	// Whether or not Meraki should email the password to user. Default is false.
	EmailPasswordToUser *bool `json:"emailPasswordToUser,omitempty"`
	// Authorization zones and expiration dates for the user.
	Authorizations []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations `json:"authorizations,omitempty"`
}

// NewInlineObject95 instantiates a new InlineObject95 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject95() *InlineObject95 {
	this := InlineObject95{}
	return &this
}

// NewInlineObject95WithDefaults instantiates a new InlineObject95 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject95WithDefaults() *InlineObject95 {
	this := InlineObject95{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject95) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject95) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject95) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InlineObject95) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineObject95) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineObject95) SetPassword(v string) {
	o.Password = &v
}

// GetEmailPasswordToUser returns the EmailPasswordToUser field value if set, zero value otherwise.
func (o *InlineObject95) GetEmailPasswordToUser() bool {
	if o == nil || isNil(o.EmailPasswordToUser) {
		var ret bool
		return ret
	}
	return *o.EmailPasswordToUser
}

// GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetEmailPasswordToUserOk() (*bool, bool) {
	if o == nil || isNil(o.EmailPasswordToUser) {
    return nil, false
	}
	return o.EmailPasswordToUser, true
}

// HasEmailPasswordToUser returns a boolean if a field has been set.
func (o *InlineObject95) HasEmailPasswordToUser() bool {
	if o != nil && !isNil(o.EmailPasswordToUser) {
		return true
	}

	return false
}

// SetEmailPasswordToUser gets a reference to the given bool and assigns it to the EmailPasswordToUser field.
func (o *InlineObject95) SetEmailPasswordToUser(v bool) {
	o.EmailPasswordToUser = &v
}

// GetAuthorizations returns the Authorizations field value if set, zero value otherwise.
func (o *InlineObject95) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations {
	if o == nil || isNil(o.Authorizations) {
		var ret []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations
		return ret
	}
	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject95) GetAuthorizationsOk() ([]NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations, bool) {
	if o == nil || isNil(o.Authorizations) {
    return nil, false
	}
	return o.Authorizations, true
}

// HasAuthorizations returns a boolean if a field has been set.
func (o *InlineObject95) HasAuthorizations() bool {
	if o != nil && !isNil(o.Authorizations) {
		return true
	}

	return false
}

// SetAuthorizations gets a reference to the given []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations and assigns it to the Authorizations field.
func (o *InlineObject95) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) {
	o.Authorizations = v
}

func (o InlineObject95) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.EmailPasswordToUser) {
		toSerialize["emailPasswordToUser"] = o.EmailPasswordToUser
	}
	if !isNil(o.Authorizations) {
		toSerialize["authorizations"] = o.Authorizations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject95 struct {
	value *InlineObject95
	isSet bool
}

func (v NullableInlineObject95) Get() *InlineObject95 {
	return v.value
}

func (v *NullableInlineObject95) Set(val *InlineObject95) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject95) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject95) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject95(val *InlineObject95) *NullableInlineObject95 {
	return &NullableInlineObject95{value: val, isSet: true}
}

func (v NullableInlineObject95) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject95) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


