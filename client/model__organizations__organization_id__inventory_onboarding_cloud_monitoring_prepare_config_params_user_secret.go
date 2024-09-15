/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret Stores the user secret hash
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret struct {
	// The hashed secret
	Hash *string `json:"hash,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecretWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecretWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) GetHash() string {
	if o == nil || isNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) GetHashOk() (*string, bool) {
	if o == nil || isNil(o.Hash) {
    return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) HasHash() bool {
	if o != nil && !isNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) SetHash(v string) {
	o.Hash = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


