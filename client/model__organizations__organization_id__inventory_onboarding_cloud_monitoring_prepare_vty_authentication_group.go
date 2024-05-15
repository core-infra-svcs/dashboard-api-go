/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup Group Details
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup struct {
	// Group Name
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroupWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroupWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthenticationGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


