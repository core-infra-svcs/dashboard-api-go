/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSmVppAccountsParsedToken The parsed VPP service token
type OrganizationsOrganizationIdSmVppAccountsParsedToken struct {
	// The organization name
	OrgName *string `json:"orgName,omitempty"`
	// The hashed token
	HashedToken *string `json:"hashedToken,omitempty"`
	// The expiration time of the token
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewOrganizationsOrganizationIdSmVppAccountsParsedToken instantiates a new OrganizationsOrganizationIdSmVppAccountsParsedToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSmVppAccountsParsedToken() *OrganizationsOrganizationIdSmVppAccountsParsedToken {
	this := OrganizationsOrganizationIdSmVppAccountsParsedToken{}
	return &this
}

// NewOrganizationsOrganizationIdSmVppAccountsParsedTokenWithDefaults instantiates a new OrganizationsOrganizationIdSmVppAccountsParsedToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSmVppAccountsParsedTokenWithDefaults() *OrganizationsOrganizationIdSmVppAccountsParsedToken {
	this := OrganizationsOrganizationIdSmVppAccountsParsedToken{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) GetOrgName() string {
	if o == nil || isNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) GetOrgNameOk() (*string, bool) {
	if o == nil || isNil(o.OrgName) {
    return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) HasOrgName() bool {
	if o != nil && !isNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) SetOrgName(v string) {
	o.OrgName = &v
}

// GetHashedToken returns the HashedToken field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) GetHashedToken() string {
	if o == nil || isNil(o.HashedToken) {
		var ret string
		return ret
	}
	return *o.HashedToken
}

// GetHashedTokenOk returns a tuple with the HashedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) GetHashedTokenOk() (*string, bool) {
	if o == nil || isNil(o.HashedToken) {
    return nil, false
	}
	return o.HashedToken, true
}

// HasHashedToken returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) HasHashedToken() bool {
	if o != nil && !isNil(o.HashedToken) {
		return true
	}

	return false
}

// SetHashedToken gets a reference to the given string and assigns it to the HashedToken field.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) SetHashedToken(v string) {
	o.HashedToken = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) GetExpiresAt() string {
	if o == nil || isNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) GetExpiresAtOk() (*string, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *OrganizationsOrganizationIdSmVppAccountsParsedToken) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o OrganizationsOrganizationIdSmVppAccountsParsedToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !isNil(o.HashedToken) {
		toSerialize["hashedToken"] = o.HashedToken
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSmVppAccountsParsedToken struct {
	value *OrganizationsOrganizationIdSmVppAccountsParsedToken
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSmVppAccountsParsedToken) Get() *OrganizationsOrganizationIdSmVppAccountsParsedToken {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSmVppAccountsParsedToken) Set(val *OrganizationsOrganizationIdSmVppAccountsParsedToken) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSmVppAccountsParsedToken) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSmVppAccountsParsedToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSmVppAccountsParsedToken(val *OrganizationsOrganizationIdSmVppAccountsParsedToken) *NullableOrganizationsOrganizationIdSmVppAccountsParsedToken {
	return &NullableOrganizationsOrganizationIdSmVppAccountsParsedToken{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSmVppAccountsParsedToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSmVppAccountsParsedToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


