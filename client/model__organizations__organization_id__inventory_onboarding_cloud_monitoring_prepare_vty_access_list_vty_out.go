/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut VTY out ACL
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut struct {
	// Name
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOutWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOutWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessListVtyOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


