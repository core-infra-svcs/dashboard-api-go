/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication VTY AAA authentication
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication struct {
	Group *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup `json:"group,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) GetGroup() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup {
	if o == nil || isNil(o.Group) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) GetGroupOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup and assigns it to the Group field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) SetGroup(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) {
	o.Group = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


