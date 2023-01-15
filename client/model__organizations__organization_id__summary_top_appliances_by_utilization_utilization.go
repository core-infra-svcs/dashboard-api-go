/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization Utilization of the appliance
type OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization struct {
	Average *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationAverage `json:"average,omitempty"`
}

// NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization instantiates a new OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization() *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization {
	this := OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization{}
	return &this
}

// NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationWithDefaults instantiates a new OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationWithDefaults() *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization {
	this := OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) GetAverage() OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationAverage {
	if o == nil || isNil(o.Average) {
		var ret OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationAverage
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) GetAverageOk() (*OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationAverage, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationAverage and assigns it to the Average field.
func (o *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) SetAverage(v OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilizationAverage) {
	o.Average = &v
}

func (o OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization struct {
	value *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) Get() *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) Set(val *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization(val *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization {
	return &NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


