/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser User credentials used to connect to the device
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser struct {
	// The public key for the registered user
	PublicKey *string `json:"publicKey,omitempty"`
	// The username added to Catalyst device
	Username *string `json:"username,omitempty"`
	Secret *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret `json:"secret,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
    return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) SetUsername(v string) {
	o.Username = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) GetSecret() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret {
	if o == nil || isNil(o.Secret) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) GetSecretOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret and assigns it to the Secret field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) SetSecret(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) {
	o.Secret = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


