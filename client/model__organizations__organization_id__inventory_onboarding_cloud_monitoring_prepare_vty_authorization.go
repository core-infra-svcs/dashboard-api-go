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

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization VTY AAA authorization
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization struct {
	Group *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup `json:"group,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorizationWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorizationWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) GetGroup() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup {
	if o == nil || isNil(o.Group) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) GetGroupOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup and assigns it to the Group field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) SetGroup(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) {
	o.Group = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


