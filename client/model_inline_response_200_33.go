/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20033 struct for InlineResponse20033
type InlineResponse20033 struct {
	// Meraki auth user id
	Id *string `json:"id,omitempty"`
	// Email address of the user
	Email *string `json:"email,omitempty"`
	// Name of the user
	Name *string `json:"name,omitempty"`
	// Creation time of the user
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Authorization type for user.
	AccountType *string `json:"accountType,omitempty"`
	// Whether or not the user is a Dashboard administrator
	IsAdmin *bool `json:"isAdmin,omitempty"`
	// User authorization info
	Authorizations []NetworksNetworkIdMerakiAuthUsersAuthorizations `json:"authorizations,omitempty"`
}

// NewInlineResponse20033 instantiates a new InlineResponse20033 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20033() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20033WithDefaults() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20033) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20033) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20033) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse20033) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse20033) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse20033) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20033) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20033) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20033) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse20033) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse20033) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse20033) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *InlineResponse20033) GetAccountType() string {
	if o == nil || isNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetAccountTypeOk() (*string, bool) {
	if o == nil || isNil(o.AccountType) {
    return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *InlineResponse20033) HasAccountType() bool {
	if o != nil && !isNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *InlineResponse20033) SetAccountType(v string) {
	o.AccountType = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *InlineResponse20033) GetIsAdmin() bool {
	if o == nil || isNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetIsAdminOk() (*bool, bool) {
	if o == nil || isNil(o.IsAdmin) {
    return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *InlineResponse20033) HasIsAdmin() bool {
	if o != nil && !isNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *InlineResponse20033) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetAuthorizations returns the Authorizations field value if set, zero value otherwise.
func (o *InlineResponse20033) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersAuthorizations {
	if o == nil || isNil(o.Authorizations) {
		var ret []NetworksNetworkIdMerakiAuthUsersAuthorizations
		return ret
	}
	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetAuthorizationsOk() ([]NetworksNetworkIdMerakiAuthUsersAuthorizations, bool) {
	if o == nil || isNil(o.Authorizations) {
    return nil, false
	}
	return o.Authorizations, true
}

// HasAuthorizations returns a boolean if a field has been set.
func (o *InlineResponse20033) HasAuthorizations() bool {
	if o != nil && !isNil(o.Authorizations) {
		return true
	}

	return false
}

// SetAuthorizations gets a reference to the given []NetworksNetworkIdMerakiAuthUsersAuthorizations and assigns it to the Authorizations field.
func (o *InlineResponse20033) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersAuthorizations) {
	o.Authorizations = v
}

func (o InlineResponse20033) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !isNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if !isNil(o.Authorizations) {
		toSerialize["authorizations"] = o.Authorizations
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20033 struct {
	value *InlineResponse20033
	isSet bool
}

func (v NullableInlineResponse20033) Get() *InlineResponse20033 {
	return v.value
}

func (v *NullableInlineResponse20033) Set(val *InlineResponse20033) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20033) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20033) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20033(val *InlineResponse20033) *NullableInlineResponse20033 {
	return &NullableInlineResponse20033{value: val, isSet: true}
}

func (v NullableInlineResponse20033) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20033) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


