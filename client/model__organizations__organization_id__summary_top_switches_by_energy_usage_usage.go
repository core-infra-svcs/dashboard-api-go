/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage Energy usage of the switch
type OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage struct {
	// Total energy usage of the switch
	Total *float32 `json:"total,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage instantiates a new OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage() *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsageWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsageWithDefaults() *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage {
	this := OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) SetTotal(v float32) {
	o.Total = &v
}

func (o OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage struct {
	value *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) Get() *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) Set(val *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage(val *OrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) *NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage {
	return &NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopSwitchesByEnergyUsageUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


